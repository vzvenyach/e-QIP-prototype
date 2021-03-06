package twofactor

import (
	"bytes"
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"errors"
	"net/url"
	"text/template"

	"github.com/dgryski/dgoogauth"
	"github.com/keighl/mandrill"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	qr "github.com/skip2/go-qrcode"
	"github.com/truetandem/e-QIP-prototype/api/cf"
)

const (
	auth string = "totp"
)

var (
	templateEmail = template.Must(template.New("email").Parse(`# Passcode\n\n{{ . }}`))
)

// Secret creates a random secret and then base32 encodes it.
func Secret() string {
	secret := make([]byte, 6)
	_, err := rand.Read(secret)
	if err != nil {
		return ""
	}

	secret32 := base32.StdEncoding.EncodeToString(secret)
	return secret32
}

// Generate will create a QR code in PNG format which will then
// be base64 encoded so it can traverse the wire to the front end.
func Generate(account, secret string) (string, error) {
	u, err := url.Parse("otpauth://" + auth)
	if err != nil {
		return "", err
	}

	u.Path += "/eqip:" + account
	params := &url.Values{}
	params.Add("secret", secret)
	params.Add("issuer", "eqip")
	u.RawQuery = params.Encode()

	png, err := qr.Encode(u.String(), qr.Medium, 256)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(png), nil
}

// Authenticate validates the initial token generated when configuring two-factor
// authentication for the first time.
func Authenticate(token, secret string) (ok bool, err error) {
	otpc := &dgoogauth.OTPConfig{
		Secret:      secret,
		WindowSize:  3,
		HotpCounter: 0,
	}
	return otpc.Authenticate(token)
}

// Email delivers code to the specified address.
func Email(address, secret string) error {
	// Get a valid token for two-factor authentication
	code := dgoogauth.ComputeCode(secret, 0)
	if code == -1 {
		return errors.New("Failed to generate code")
	}

	// Pull the API key for the mail service
	key := cf.UserService("eqip-smtp", "api_key")
	if key == "" {
		return errors.New("Could not retrieve API key")
	}

	// Form the mail service
	client := mandrill.ClientWithKey(key)
	message := &mandrill.Message{
		FromEmail: "noreply@mail.gov",
		FromName:  "E-QIP",
		Subject:   "E-QIP Passcode",
	}
	message.AddRecipient(address, address, "to")

	// The template is stored in markdown format to easily
	// transform something human readable to HTML as well
	var plain bytes.Buffer
	err := templateEmail.Execute(&plain, code)
	if err != nil {
		return err
	}
	message.Text = plain.String()
	unsafe := blackfriday.MarkdownCommon(plain.Bytes())
	message.HTML = string(bluemonday.UGCPolicy().SanitizeBytes(unsafe))

	// Send the message and return any errors from the service
	_, err = client.MessagesSend(message)
	return err
}

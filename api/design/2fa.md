FORMAT: 1A

# E-QIP API

# Two-factor Authentication [/2fa]

Two-factor authentication endpoints

## Retrieve a QR code [GET]

+ Request
    + Headers

        Authorization: Bearer
        Accept: text/plain

+ Response 200 (text/plain)
    + Headers

        X-Eqip-Media-Type: eqip.v1

    + Body

        image/png;base64

+ Response 500

# 2FA verify endpoint [/2fa/verify]

## Verify two-factor token [POST]

+ Request
    + Headers

        Authorization: Bearer
        Accept: text/plain

+ Response 200 (text/plain)
    + Headers

        X-Eqip-Media-Type: eqip.v1

+ Response 500

+ Response 401

# 2FA email endpoint [/2fa/email]

## Request two-factor token by e-mail [POST]

+ Request
    + Headers

        Authorization: Bearer
        Accept: text/plain

+ Response 200 (text/plain)
    + Headers

        X-Eqip-Media-Type: eqip.v1

+ Response 500

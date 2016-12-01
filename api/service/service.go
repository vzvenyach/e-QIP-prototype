package service

var (
	Account AccountService
	Token   TokenService
)

func init() {
	Account = NewAccountPgService()
	Token = NewJwtTokenService()
}

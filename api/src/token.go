package kafebar

type TokenService interface {
	GetToken() (string, error)
}

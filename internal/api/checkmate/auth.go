package checkmate

type Authenticator struct {
	key       string
	headerKey string
	header    string
}

func NewBearerAuthenticator(credentials string) *Authenticator {
	return &Authenticator{
		key:       credentials,
		headerKey: "Authorization",
		header:    "Bearer ",
	}
}

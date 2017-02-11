package parameters

type TokenAuthentication struct {
	Token string `json:"token" form:"token"`
}

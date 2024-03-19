package token

type Maker interface {
	CreateToken(params CreateTokenParams) (string, *Payload, error)

	VerifyToken(token string) (*Payload, error)
}

type CreateTokenParams struct {
	// UID is used to identify the user. Can be username, user id, email or any other user identifier
	UID string
	// Roles is used for rbac
	Roles []string
}

package token_manager

type InvalidTokenError struct{}

func (e InvalidTokenError) Error() string {
	return "invalid token"
}

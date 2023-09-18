package auth

import "context"

type Authentication struct {
	User     string
	Password string
}

// RequireTransportSecurity indicates whether the credentials requires
// transport security.

func (a *Authentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
	return true
}

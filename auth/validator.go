package auth

import "context"

type (
	AuthorizationRequestValidator func(ctx context.Context) (*AuthorizationRequest, error)
)

func DefaultAuthorizationRequestValidator(ctx context.Context) (*AuthorizationRequest, error) {
	return nil, nil
}

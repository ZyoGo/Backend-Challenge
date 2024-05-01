package core

import "context"

type Business interface {
	LoginUser(ctx context.Context, dto LoginUserParams) (Auth, error)
}

package core

import "context"

type Business interface {
	CreateOrder(ctx context.Context, params CreateOrderDTO) error
}

package service

import (
	"context"
	"ordersystem/middleware"
)

var (
	logger = middleware.Logger()
	ctx    = context.Background()
)

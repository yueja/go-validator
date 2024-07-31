package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
)

// ValidatorMiddlewareKratos is a validator middleware.
func ValidatorMiddlewareKratos() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if err = validateV1(req); err != nil {
				return nil, err
			}
			return handler(ctx, req)
		}
	}
}

// https://blog.csdn.net/NiDeHaoPengYou/article/details/138793196

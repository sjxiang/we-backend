package data

import (
	"context"
	"we-backend/internal/biz"
)

type otpRepo struct {
	cache     OtpCache 
}

func NewOtpRepo(cache OtpCache) biz.OtpRepo {
	return &otpRepo{
		cache: cache,
	}
}

func (impl *otpRepo) Insert(ctx context.Context, biz, phoneNumber, code string) error {
	return impl.cache.Set(ctx, biz, phoneNumber, code)
}

func (impl *otpRepo) FindOne(ctx context.Context, biz, phoneNumber, code string) (bool, error) {
	return impl.cache.Verify(ctx, biz, phoneNumber, code)
}

package context

import (
	"context"
	"errors"

	"github.com/nghiatk54/go_ecommerce_api/internal/constant"
	"github.com/nghiatk54/go_ecommerce_api/internal/util/cache"
)

type InfoUserUuid struct {
	UserId      uint64
	UserAccount string
}

func GetSubjectUuid(ctx context.Context) (string, error) {
	sUuid, ok := ctx.Value(constant.SUBJECT_UUID).(string)
	if !ok {
		return "", errors.New("subject uuid not found")
	}
	return sUuid, nil
}

func GetUserIdFromUuid(ctx context.Context) (uint64, error) {
	sUuid, err := GetSubjectUuid(ctx)
	if err != nil {
		return 0, err
	}
	var infoUser InfoUserUuid
	if err := cache.GetCache(ctx, sUuid, &infoUser); err != nil {
		return 0, err
	}
	return infoUser.UserId, nil
}

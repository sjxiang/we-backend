package utils

import (
	"context"
	"errors"
)


var (
	ErrNoUserIDInContext  = errors.New("no user id in context")
)

type contextKey string

var (
	contextAuthIDKey contextKey = "currentUserId"
)

func GetUserIDFromContextx(ctx context.Context) (string, error) {
	if ctx.Value(contextAuthIDKey) == nil {
		return "", ErrNoUserIDInContext
	}

	userID, ok := ctx.Value(contextAuthIDKey).(string)
	if !ok {
		return "", ErrNoUserIDInContext
	}

	return userID, nil
}

func PutUserIDIntoContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, contextAuthIDKey, id)
}

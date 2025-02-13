package utils

import "context"

func GetUserIdFromCtx (ctx context.Context) int32{
	userId := ctx.Value(SessionsUserId)
	if userId == nil{
          return 0
	}
    return userId.(int32)
}
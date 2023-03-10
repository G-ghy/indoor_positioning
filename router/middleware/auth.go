package middleware

import (
	"indoor_positioning/handler"
	"indoor_positioning/pkg/errno"
	"indoor_positioning/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(ctx); err != nil {
			handler.SendResponse(ctx, errno.ErrorTokenInvalid, nil)
			ctx.Abort() // 终止当前中间件以后的中间件执行，但是会执行当前中间件的后续逻辑
			return
		}

		ctx.Next()
	}
}

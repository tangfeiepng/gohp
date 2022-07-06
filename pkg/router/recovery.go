package router

import (
	"Walker/pkg/contract"
	"Walker/pkg/exception"
	"github.com/gin-gonic/gin"
)

func WithWriter(logger contract.Logger, httpHandler func(exception2 contract.Exception, context *gin.Context)) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				exceptions := exception.WithError(err)
				httpHandler(exceptions, context)
				logger.Error(exceptions.GetMessage(), contract.LoggerField{
					Key: "stack",
					Val: exceptions.GetStack(),
				})
			}
		}()
		context.Next()
	}
}

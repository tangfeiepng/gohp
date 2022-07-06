package request

import (
	"Walker/pkg/contract"
	"github.com/gin-gonic/gin"
)

type Request struct {
	context *gin.Context
}

func (r *Request) CreateRequest(context *gin.Context) contract.Request {
	return &Request{context: context}
}

func (r *Request) Param(key string) string {
	return r.context.Param(key)
}
func (r *Request) Query(key string) string {
	return r.context.Query(key)
}
func (r *Request) Next() {
	r.context.Next()
}
func (r *Request) Context() interface{} {
	return r.context
}

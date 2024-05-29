package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Controller[R, S any] interface {
	HandlerFunc(ctx *gin.Context)
}

type controller[R, S any] struct {
	serviceFunc func(ctx context.Context, request R) (response S, err error)
}

func NewController[R, S any](fn func(ctx context.Context, request R) (S, error)) Controller[R, S] {
	return controller[R, S]{serviceFunc: fn}
}

func (ctrl controller[R, S]) HandlerFunc(ctx *gin.Context) {

	var request R

	if bindError := ctx.ShouldBindBodyWith(&request, binding.JSON); bindError != nil {
		log.Printf("Failed to bind request. Error: %v", bindError)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	resp, err := ctrl.serviceFunc(context.TODO(), request)
	if err != nil {
		log.Printf("Service call failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
	log.Default().Println("call complted successfully")
	ctx.JSON(http.StatusOK, resp)
}

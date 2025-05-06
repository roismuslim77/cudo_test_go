package transaction

import (
	"context"
	"github.com/gin-gonic/gin"
	"simple-go/pkg/response"
)

type Service interface {
	GetFraudDetection(ctx context.Context) ([]ResultDetectionData, response.ErrorResponse)
}

type handler struct {
	service Service
}

func NewHandler(svc Service) handler {
	return handler{
		service: svc,
	}
}

func (h handler) GetFraudDetection(ctx *gin.Context) {

	result, err := h.service.GetFraudDetection(ctx)
	if !err.IsNoError {
		resp := response.Error(err.Code).WithError(err.Message).WithStatusCode(err.StatusCode)
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	resp := response.Success("22152").WithData(result)
	ctx.JSON(resp.StatusCode, resp)
}

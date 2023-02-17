package middleware

import (
	"errors"

	"go_blockChain_server/customerror"
	"go_blockChain_server/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CheckBodyBinding(req interface{}, ctx *gin.Context) []models.ErrorMsg {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// bind 체크를 위한 코드
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]models.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = models.ErrorMsg{fe.Field(), customerror.GetErrorMsg(fe)}
			}
			return out
		}
	}

	return nil
}

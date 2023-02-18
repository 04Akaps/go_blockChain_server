package middleware

import (
	"errors"

	"go_blockChain_server/customerror"
	"go_blockChain_server/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// interface를 사용하는 이유는 어떤 model타입이 들어 올 지 모르기 떄문에
func CheckBodyBinding(req interface{}, ctx *gin.Context) []models.ErrorMsg {
	// interface의 default값은 nil 이기 떄문에
	// 사실 어차피 error가 있다는 뜻이기 떄문에 의미는 없는 코드가 된다.

	if req == nil {
		return []models.ErrorMsg{{
			Field:   "req is nil",
			Message: "req is nil",
		}}
	}

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

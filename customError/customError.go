package customerror

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required" + fe.Param()
	case "min":
		return "Check min number" + fe.Param()
	}

	return "Unknown error"
}

func ErrorMsg(err error) gin.H {
	return gin.H{"error Message ": err}
}

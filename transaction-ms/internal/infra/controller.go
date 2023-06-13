package infra

import (
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type StatusController struct {
}

func NewStatusController() StatusController {
	return StatusController{}
}

func (c StatusController) HandleStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "time": time.Now().Format("2006-01-02 15:04:05")})
}

type CreditTransactionsController struct {
}

func NewCreditTransactionsController() CreditTransactionsController {
	return CreditTransactionsController{}
}

func RegisterCustomValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

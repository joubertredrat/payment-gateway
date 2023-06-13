package infra

import (
	"joubertredrat/transaction-ms/internal/application"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type ApiBaseController struct {
}

func NewApiBaseController() ApiBaseController {
	return ApiBaseController{}
}

func (c ApiBaseController) HandleStatus(ctx *gin.Context) {
	t := time.Now()
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"time":   GetDatetimeCanonical(&t),
	})
}

func (c ApiBaseController) HandleNotFound(ctx *gin.Context) {
	t := time.Now()
	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "404 page not found",
		"time":  GetDatetimeCanonical(&t),
	})
}

type CreditTransactionsController struct {
}

func NewCreditTransactionsController() CreditTransactionsController {
	return CreditTransactionsController{}
}

func (c CreditTransactionsController) HandleList(usecase application.UsecaseListCreditCardTransaction) gin.HandlerFunc {
	t := time.Now()
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"time":      GetDatetimeCanonical(&t),
			"operation": "list",
		})
	}
}

func (c CreditTransactionsController) HandleCreate(usecase application.UsecaseCreateCreditCardTransaction) gin.HandlerFunc {
	t := time.Now()
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"time":      GetDatetimeCanonical(&t),
			"operation": "create",
		})
	}
}

func (c CreditTransactionsController) HandleGet(usecase application.UsecaseGetCreditCardTransaction) gin.HandlerFunc {
	t := time.Now()
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"time":      GetDatetimeCanonical(&t),
			"operation": "get",
		})
	}
}

func (c CreditTransactionsController) HandleEdit(usecase application.UsecaseEditCreditCardTransaction) gin.HandlerFunc {
	t := time.Now()
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"time":      GetDatetimeCanonical(&t),
			"operation": "edit",
		})
	}
}

func (c CreditTransactionsController) HandleDelete(usecase application.UsecaseDeleteCreditCardTransaction) gin.HandlerFunc {
	t := time.Now()
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"time":      GetDatetimeCanonical(&t),
			"operation": "delete",
		})
	}
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

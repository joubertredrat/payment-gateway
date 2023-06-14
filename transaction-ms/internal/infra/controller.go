package infra

import (
	"errors"
	"fmt"
	"joubertredrat/transaction-ms/internal/application"
	"joubertredrat/transaction-ms/internal/domain"
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
		"time":   DatetimeCanonical(&t),
	})
}

func (c ApiBaseController) HandleNotFound(ctx *gin.Context) {
	t := time.Now()
	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "page not found",
		"time":  DatetimeCanonical(&t),
	})
}

type CreditTransactionsController struct {
}

func NewCreditTransactionsController() CreditTransactionsController {
	return CreditTransactionsController{}
}

func (c CreditTransactionsController) HandleList(usecase application.UsecaseListCreditCardTransaction) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		l, _ := usecase.Execute(application.ListCreditCardTransactionInput{
			Page:         1,
			ItemsPerPage: 10,
		})
		response := NewListResponse(1, 10, CreateCreditCardTransactionListResponseFromUsecase(l))
		ctx.JSON(http.StatusOK, response)
	}
}

func (c CreditTransactionsController) HandleCreate(usecase application.UsecaseCreateCreditCardTransaction) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request CreateCreditCardTransactionRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			fmt.Println(err)
			responseWithError(ctx, err)
			return
		}
		expireDate, err := CardExpireTime(request.ExpireYear, request.ExpireMonth)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid credit card expire year or month",
			})
			return
		}

		t, _ := usecase.Execute(application.CreateCreditCardTransactionInput{
			HolderName:   request.HolderName,
			CardNumber:   request.CardNumber,
			CVV:          request.CVV,
			ExpireDate:   expireDate,
			Amount:       request.Amount,
			Installments: request.Installments,
			Description:  request.Description,
		})
		response := CreateCreditCardTransactionResponseFromUsecase(t)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c CreditTransactionsController) HandleGet(usecase application.UsecaseGetCreditCardTransaction) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		transactionID := ctx.Param("transactionid")
		if transactionID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "transaction id required",
			})
			return
		}

		t, err := usecase.Execute(application.GetCreditCardTransactionInput{
			TransactionID: transactionID,
		})
		if err != nil {
			switch err.(type) {
			case domain.ErrCreditCardTransactionNotFound:
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
			default:
				t := time.Now()
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "internal server error",
					"time":  DatetimeCanonical(&t),
				})
			}
			return
		}

		response := CreateCreditCardTransactionResponseFromUsecase(t)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c CreditTransactionsController) HandleEdit(usecase application.UsecaseEditCreditCardTransaction) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		transactionID := ctx.Param("transactionid")
		if transactionID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "transaction id required",
			})
			return
		}

		var request EditCreditCardTransactionRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			fmt.Println(err)
			responseWithError(ctx, err)
			return
		}

		t, _ := usecase.Execute(application.EditCreditCardTransactionInput{
			TransactionID: transactionID,
			Description:   request.Description,
		})

		response := CreateCreditCardTransactionResponseFromUsecase(t)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c CreditTransactionsController) HandleDelete(usecase application.UsecaseDeleteCreditCardTransaction) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		transactionID := ctx.Param("transactionid")
		if transactionID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "transaction id required",
			})
			return
		}

		usecase.Execute(application.DeleteCreditCardTransactionInput{
			TransactionID: transactionID,
		})

		ctx.JSON(http.StatusNoContent, gin.H{})
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

func responseWithError(c *gin.Context, err error) {
	var verr validator.ValidationErrors
	if errors.As(err, &verr) {
		c.JSON(http.StatusBadRequest, gin.H{"errors": getValidatorErrors(verr)})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func getValidatorErrors(verr validator.ValidationErrors) []RequestValidationError {
	var errs []RequestValidationError

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}

		errs = append(errs, RequestValidationError{Field: f.Field(), Reason: err})
	}

	return errs
}

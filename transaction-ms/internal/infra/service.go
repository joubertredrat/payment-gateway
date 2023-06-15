package infra

import (
	"context"
	"joubertredrat/transaction-ms/internal/application"
	"joubertredrat/transaction-ms/internal/domain"
	"joubertredrat/transaction-ms/internal/infra/authorization"
)

type AuthorizationServiceMicroservice struct {
	logger  application.Logger
	service authorization.AuthorizationClient
}

func NewAuthorizationServiceMicroservice(
	logger application.Logger,
	service authorization.AuthorizationClient,
) AuthorizationServiceMicroservice {
	return AuthorizationServiceMicroservice{
		logger:  logger,
		service: service,
	}
}

func (a AuthorizationServiceMicroservice) Handle(request domain.AuthorizationRequest) (domain.AuthorizationResponse, error) {
	response, err := a.service.Authorize(context.Background(), &authorization.CreditCardRequest{
		HolderName:   request.HolderName,
		CardNumber:   request.CardNumber,
		Cvv:          request.CVV,
		ExpireDate:   request.ExpireDate.Unix(),
		Amount:       uint64(request.Amount),
		Installments: uint32(request.Installments),
	})
	if err != nil {
		a.logger.Error(err)
		return domain.AuthorizationResponse{}, domain.ErrAuthorizationServiceHouston
	}

	return domain.NewAuthorizationResponse(a.convertStatus(response.Status))
}

func (a AuthorizationServiceMicroservice) convertStatus(s authorization.AuthorizationStatus) string {
	status := a.correlatedStatus()
	return status[s]
}

func (a AuthorizationServiceMicroservice) correlatedStatus() map[authorization.AuthorizationStatus]string {
	return map[authorization.AuthorizationStatus]string{
		authorization.AuthorizationStatus_AUTHORIZED:                  domain.AUTHORIZATION_STATUS_AUTHORIZED,
		authorization.AuthorizationStatus_DECLINED_UNKNOWN:            domain.AUTHORIZATION_STATUS_DECLINED,
		authorization.AuthorizationStatus_DECLINED_INVALID_DATA:       domain.AUTHORIZATION_STATUS_DECLINED,
		authorization.AuthorizationStatus_DECLINED_EXPIRED_CARD:       domain.AUTHORIZATION_STATUS_DECLINED,
		authorization.AuthorizationStatus_DECLINED_INSUFFICIENT_FUNDS: domain.AUTHORIZATION_STATUS_DECLINED,
		authorization.AuthorizationStatus_DECLINED_SUSPECT_FRAUD:      domain.AUTHORIZATION_STATUS_DECLINED,
	}
}

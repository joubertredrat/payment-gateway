package infra

import "joubertredrat/transaction-ms/internal/domain"

type AuthorizationServiceMicroservice struct {
}

func NewAuthorizationServiceMicroservice() AuthorizationServiceMicroservice {
	return AuthorizationServiceMicroservice{}
}

func (a AuthorizationServiceMicroservice) Handle(request domain.AuthorizationRequest) (domain.AuthorizationResponse, error) {
	return domain.NewAuthorizationResponse(domain.AUTHORIZATION_STATUS_AUTHORIZED)
}

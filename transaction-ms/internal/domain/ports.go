package domain

type AuthorizationService interface {
	Handle(request AuthorizationRequest) (AuthorizationResponse, error)
}

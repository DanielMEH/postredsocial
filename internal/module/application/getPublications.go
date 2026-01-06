package application

import (
	"github.com/postredsocial/internal/infrastructure/config"
	"github.com/postredsocial/internal/infrastructure/utils"
	"github.com/postredsocial/internal/module/domains/entities"
	"github.com/postredsocial/internal/module/domains/ports"
)

type GetPublicationUseCase struct {
	portsPublications ports.PortsRepositoryPosts
}

func NewGetPublicationUseCase(portsPublications ports.PortsRepositoryPosts) *GetPublicationUseCase {

	return &GetPublicationUseCase{portsPublications}
}

func (uc *GetPublicationUseCase) ExecuteGetPublicationUseCase(tokenString string) (entities.EntityPublicationsResponse, error) {

	_, errToken := utils.ValidateJWT(tokenString)

	if errToken != nil {
		return entities.EntityPublicationsResponse{}, config.NewStatusUnauthorized(errToken)
	}
	responseUsc, err := uc.portsPublications.GetPublications()

	if err != nil {
		return entities.EntityPublicationsResponse{}, err
	}

	return responseUsc, nil
}

package application

import (
	"github.com/postredsocial/internal/infrastructure/config"
	"github.com/postredsocial/internal/infrastructure/utils"
	"github.com/postredsocial/internal/module/domains/entities"
	"github.com/postredsocial/internal/module/domains/ports"
)

type GetNewPublicationUseCase struct {
	portsPublications ports.PortsRepositoryPosts
}

func NewPublicationAccountUseCase(portsPublications ports.PortsRepositoryPosts) *GetNewPublicationUseCase {

	return &GetNewPublicationUseCase{portsPublications}
}

func (uc *GetNewPublicationUseCase) ExecutePostPublicationeCase(message string, tokenString string) (entities.EntityNewPostResponse, error) {

	accountId, errToken := utils.ValidateJWT(tokenString)

	if errToken != nil {
		return entities.EntityNewPostResponse{}, config.NewStatusUnauthorized(errToken)
	}
	responseUsc, err := uc.portsPublications.NewPublication(message, accountId)

	if err != nil {
		return entities.EntityNewPostResponse{}, err
	}

	return responseUsc, nil
}

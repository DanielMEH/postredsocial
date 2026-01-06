package application

import (
	"github.com/google/uuid"
	"github.com/postredsocial/internal/infrastructure/config"
	"github.com/postredsocial/internal/infrastructure/utils"
	"github.com/postredsocial/internal/module/domains/entities"
	"github.com/postredsocial/internal/module/domains/ports"
)

type LikePublicationUseCase struct {
	portsPublications ports.PortsRepositoryPosts
}

func NewLikePublicationUseCase(portsPublications ports.PortsRepositoryPosts) *LikePublicationUseCase {

	return &LikePublicationUseCase{portsPublications}
}

func (uc *LikePublicationUseCase) ExecuteLikePublicationeCase(PostId uuid.UUID, userId uuid.UUID, tokenString string) (entities.EntityLikePostResponse, error) {

	_, errToken := utils.ValidateJWT(tokenString)

	if errToken != nil {
		return entities.EntityLikePostResponse{}, config.NewStatusUnauthorized(errToken)
	}
	responseUsc, err := uc.portsPublications.LikePublication(PostId, userId)

	if err != nil {
		return entities.EntityLikePostResponse{}, err
	}

	return responseUsc, nil
}

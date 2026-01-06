package ports

import (
	"github.com/google/uuid"
	"github.com/postredsocial/internal/module/domains/entities"
)

type PortsRepositoryPosts interface {
	NewPublication(message string, accountId string) (entities.EntityNewPostResponse, error)
	LikePublication(postId uuid.UUID, UserId uuid.UUID) (entities.EntityLikePostResponse, error)
}

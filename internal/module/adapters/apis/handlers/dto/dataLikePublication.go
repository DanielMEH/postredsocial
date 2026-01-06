package dto

import "github.com/google/uuid"

type DtoLikePublication struct {
	UserId uuid.UUID `json:"user_id"`
	PostId uuid.UUID `json:"post_id"`
}

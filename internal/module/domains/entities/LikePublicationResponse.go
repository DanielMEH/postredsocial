package entities

type EntityLikePostResponse struct {
	Message string `json:"message"`
	Details struct {
		Like_id string `json:"like_id"`
	} `json:"details"`
}

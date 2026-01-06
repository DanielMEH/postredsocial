package entities

type EntityNewPostResponse struct {
	Message string `json:"message"`
	Details struct {
		PublicationId string `json:"publication_id"`
	} `json:"details"`
}

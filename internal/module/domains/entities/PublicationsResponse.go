package entities

type EntityPublicationsResponse struct {
	Message string `json:"message"`
	Details struct {
		Data []map[string]interface{} `json:"data"`
	} `json:"details"`
}

package api

type GetIndexResponse struct {
	Index        int    `json:"index"`
	Value        int    `json:"value"`
	ErrorMessage string `json:"error_message,omitempty"`
}

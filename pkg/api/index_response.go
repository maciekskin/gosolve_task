package api

type GetIndexResponse struct {
	Index        int    `json:"index"`
	ErrorMessage string `json:"error_message,omitempty"`
}

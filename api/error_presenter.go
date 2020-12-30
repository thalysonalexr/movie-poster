package api

// ErrorPresenter error presenter
type ErrorPresenter struct {
	Type    string `json:"type"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

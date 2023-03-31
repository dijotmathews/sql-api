package dto

type Identity struct {
}

type QueryRequestBody struct {
	Query string `json:"query"`
}

type QueryResponseBody struct {
	Status  int    `json:"status"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

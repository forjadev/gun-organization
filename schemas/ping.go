package schemas

type PingResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

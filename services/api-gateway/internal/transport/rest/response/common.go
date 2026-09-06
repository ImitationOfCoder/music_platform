package response

type BaseResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

type StatusResponse struct {
	Success bool `json:"success"`
}

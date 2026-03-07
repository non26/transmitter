package push

type PushResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewPushResponse() *PushResponse {
	return &PushResponse{}
}

func (p *PushResponse) Success() *PushResponse {
	return &PushResponse{
		Code:    "0000",
		Message: "success",
	}
}

func (p *PushResponse) Error(message string) *PushResponse {
	return &PushResponse{
		Code:    "9999",
		Message: message,
	}
}

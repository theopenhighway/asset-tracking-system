package dto

type ResponseMessage struct {
	Msg  string `json:"message"`
	Data any    `json:"data"`
}

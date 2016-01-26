package api

type Response struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Text string      `json:"text,omitempty"`
}

func WithID(ID string) Response {
	return Response{
		Code: 10001,
		Data: ID,
		Text: "id",
	}
}

func WithData(data string) Response {
	return Response{
		Code: 10002,
		Data: data,
		Text: "data",
	}
}

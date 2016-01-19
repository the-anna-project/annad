package api

type Response struct {
	Code int         `json:"code,omitempty"`
	Text string      `json:"text,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func Success() Response {
	return Response{
		Code: 10001,
		Text: "success",
	}
}

func WithData(v interface{}) Response {
	return Response{
		Code: 10002,
		Text: "data",
		Data: v,
	}
}

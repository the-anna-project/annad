package api

type Response struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
	ID   string      `json:"id,omitempty"`
	Text string      `json:"text,omitempty"`
}

func WithID(ID string) Response {
	return Response{
		Code: 10001,
		ID:   ID,
		Text: "success",
	}
}

func WithData(v interface{}) Response {
	return Response{
		Code: 10002,
		Data: v,
		Text: "data",
	}
}

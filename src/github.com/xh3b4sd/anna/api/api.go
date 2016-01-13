package api

type Response struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func Success() Response {
	return Response{
		Code: 10001,
		Text: "success",
	}
}

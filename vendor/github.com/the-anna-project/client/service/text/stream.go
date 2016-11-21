package text

import (
	"io"

	"google.golang.org/grpc"

	textoutputobject "github.com/the-anna-project/output/object/text"
	apispec "github.com/the-anna-project/spec/api"
	objectspec "github.com/the-anna-project/spec/object"
)

func (s *service) decodeResponse(streamTextResponse *StreamTextResponse) (objectspec.TextOutput, error) {
	if streamTextResponse.Code != apispec.CodeData {
		return nil, maskAnyf(invalidAPIResponseError, "API response code must be %d", apispec.CodeData)
	}

	textOutputObject := textoutputobject.New()
	textOutputObject.SetOutput(streamTextResponse.Data.Output)

	return textOutputObject, nil
}

func (s *service) encodeRequest(textInput objectspec.TextInput) *StreamTextRequest {
	streamTextRequest := &StreamTextRequest{
		Echo:      textInput.Echo(),
		Input:     textInput.Input(),
		SessionID: textInput.SessionID(),
	}

	return streamTextRequest
}

func (s *service) streamText() error {
	done := make(chan struct{}, 1)
	fail := make(chan error, 1)

	conn, err := grpc.Dial(s.address, grpc.WithInsecure())
	if err != nil {
		return maskAny(err)
	}
	defer conn.Close()

	client := NewTextEndpointClient(conn)
	stream, err := client.StreamText(s.context)
	if err != nil {
		return maskAny(err)
	}

	// Listen on the outout of the text endpoint stream send it back to the
	// client.
	go func() {
		for {
			streamTextResponse, err := stream.Recv()
			if err == io.EOF {
				// The stream ended. We broadcast to all goroutines by closing the done
				// channel.
				close(done)
				return
			} else if err != nil {
				fail <- maskAny(err)
				return
			}

			textResponse, err := s.decodeResponse(streamTextResponse)
			if err != nil {
				fail <- maskAny(err)
				return
			}
			s.Service().Output().Text().Channel() <- textResponse
		}
	}()

	// Listen on the client input channel and forward it to the server stream.
	go func() {
		for {
			select {
			case <-done:
				return
			case textInput := <-s.Service().Input().Text().Channel():
				streamTextRequest := s.encodeRequest(textInput)
				err := stream.Send(streamTextRequest)
				if err != nil {
					fail <- maskAny(err)
					return
				}
			}
		}
	}()

	for {
		select {
		case <-stream.Context().Done():
			close(done)
			return maskAny(stream.Context().Err())
		case <-done:
			return nil
		case err := <-fail:
			return maskAny(err)
		}
	}
}

package text

import (
	"io"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	apispec "github.com/the-anna-project/spec/api"
	systemspec "github.com/the-anna-project/spec/legacy"
	objectspec "github.com/the-anna-project/spec/object"
	servicespec "github.com/the-anna-project/spec/service"
	"github.com/xh3b4sd/anna/object/textoutput"
)

// New creates a new text interface service.
func New() systemspec.TextInterfaceClient {
	return &client{}
}

type client struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

	// Settings.

	// gRPCAddr is the host:port representation based on the golang convention
	// for net.Listen to serve gRPC traffic.
	gRPCAddr string
	metadata map[string]string
}

func (c *client) Boot() {
	id, err := c.Service().ID().New()
	if err != nil {
		panic(err)
	}
	c.metadata = map[string]string{
		"id":   id,
		"name": "text-interface",
		"type": "service",
	}
}

func (c *client) DecodeResponse(streamTextResponse *StreamTextResponse) (objectspec.TextOutput, error) {
	if streamTextResponse.Code != apispec.CodeData {
		return nil, maskAnyf(invalidAPIResponseError, "API response code must be %d", apispec.CodeData)
	}

	textOutputConfig := textoutput.DefaultConfig()
	textOutputConfig.Output = streamTextResponse.Data.Output
	textResponse, err := textoutput.New(textOutputConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return textResponse, nil
}

func (c *client) EncodeRequest(textInput objectspec.TextInput) *StreamTextRequest {
	streamTextRequest := &StreamTextRequest{
		Echo:      textInput.GetEcho(),
		Input:     textInput.GetInput(),
		SessionID: textInput.GetSessionID(),
	}

	return streamTextRequest
}

func (c *client) Service() servicespec.ServiceCollection {
	return c.serviceCollection
}

func (c *client) SetGRPCAddress(gRPCAddr string) {
	c.gRPCAddr = gRPCAddr
}

func (c *client) SetServiceCollection(sc servicespec.ServiceCollection) {
	c.serviceCollection = sc
}

func (c *client) StreamText(ctx context.Context) error {
	done := make(chan struct{}, 1)
	fail := make(chan error, 1)

	conn, err := grpc.Dial(c.gRPCAddr, grpc.WithInsecure())
	if err != nil {
		return maskAny(err)
	}
	defer conn.Close()

	client := NewTextEndpointClient(conn)
	stream, err := client.StreamText(ctx)
	if err != nil {
		return maskAny(err)
	}

	// Listen on the outout of the text interface stream send it back to the
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

			textResponse, err := c.DecodeResponse(streamTextResponse)
			if err != nil {
				fail <- maskAny(err)
				return
			}
			c.Service().TextOutput().Channel() <- textResponse
		}
	}()

	// Listen on the client input channel and forward it to the server stream.
	go func() {
		for {
			select {
			case <-done:
				return
			case textInput := <-c.Service().Input().Text().Channel():
				streamTextRequest := c.EncodeRequest(textInput)
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

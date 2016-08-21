package text

import (
	"io"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
)

// InterfaceConfig represents the configuration used to create a new text interface
// object.
type InterfaceConfig struct {
	// Settings.

	// GRPCAddr is the host:port representation based on the golang convention
	// for net.Listen to serve gRPC traffic.
	GRPCAddr string
}

// DefaultInterfaceConfig provides a default configuration to create a new text
// interface object by best effort.
func DefaultInterfaceConfig() InterfaceConfig {
	newConfig := InterfaceConfig{
		// Settings.
		GRPCAddr: "127.0.0.1:9119",
	}

	return newConfig
}

// NewInterface creates a new configured text interface object.
func NewInterface(config InterfaceConfig) (spec.TextInterface, error) {
	newInterface := &tinterface{
		InterfaceConfig: config,
	}

	// Settings.

	if newInterface.GRPCAddr == "" {
		return nil, maskAnyf(invalidConfigError, "gRPC address must not be empty")
	}

	return newInterface, nil
}

type tinterface struct {
	InterfaceConfig
}

func (i tinterface) StreamText(ctx context.Context, in chan spec.TextRequest, out chan spec.TextResponse) error {
	done := make(chan struct{}, 1)
	fail := make(chan error, 1)

	conn, err := grpc.Dial(i.GRPCAddr, grpc.WithInsecure())
	if err != nil {
		return maskAny(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			fail <- maskAny(err)
		}
	}()

	client := api.NewTextInterfaceClient(conn)
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

			textResponseConfig := api.DefaultTextResponseConfig()
			textResponseConfig.Output = streamTextResponse.Output
			textResponse, err := api.NewTextResponse(textResponseConfig)
			if err != nil {
				fail <- maskAny(err)
				return
			}

			out <- textResponse
		}
	}()

	// Listen on the client input channel and forward it to the server stream.
	go func() {
		for {
			select {
			case <-done:
				return
			case textRequest := <-in:
				streamTextRequest := &api.StreamTextRequest{
					Input: textRequest.GetInput(),
				}
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
			return maskAny(stream.Context().Err())
		case <-done:
			return nil
		case err := <-fail:
			if err != nil {
				close(done)
				return maskAny(err)
			}
		}
	}
}

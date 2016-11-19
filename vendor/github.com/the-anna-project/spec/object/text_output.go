package spec

// TextOutput represents a streamed response being send to the client. This
// is basically good for responding calculated output of the neural network.
type TextOutput interface {
	// Output returns the output of the current text response.
	Output() string
	SetOutput(output string)
}

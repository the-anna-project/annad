// Package id provides a simple ID generating service using pseudo random
// strings.
package id

// MustNew returns a new string of type Hex128. In case of any error
// this method panics.
func MustNew() string {
	newID, err := MustNewService().WithType(Hex128)
	if err != nil {
		panic(err)
	}

	return newID
}

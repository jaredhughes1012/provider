package provider

import (
	"context"
)

// Function that is configured to provide a value of a given type
type Provider[T any] func(ctx context.Context) (T, error)

// Creates a new provider. This function does nothing but return the function that is
// provider, this is a helper function that is used to easily encapsulate scoped values
func New[T any](provider func(ctx context.Context) (T, error)) Provider[T] {
	return provider
}

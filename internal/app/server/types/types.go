package types

type Result[T any] struct {
	Code int   `json:"code"`
	Data T     `json:"data,omitempty"`
	Err  error `json:"error,omitempty"`
}

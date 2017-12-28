package main

// Equaler ...
type Equaler interface {
	Equal(Equaler) bool
}

// T ...
type T int

// Equal ...
func (t T) Equal(u T) bool { return t == u }

// var _ Equaler = T(0) // error

// T2 ...
type T2 int

// Equal ...
func (t T2) Equal(u Equaler) bool { return t == u.(T2) }

var _ Equaler = T2(0)

func main() {
}

package model

type Args struct {
	A, B int
}

type Reply struct {
	Sum int
}

const (
	TCP = iota
	HTTP
)

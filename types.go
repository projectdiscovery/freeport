package freeport

type Protocol uint8

const (
	TCP Protocol = iota
	UDP
)

type Port struct {
	Address  string
	Port     int
	Protocol Protocol
}

package freeport

import (
	"errors"
	"fmt"
	"net"
)

func GetFreePort(address string, protocol Protocol) (*Port, error) {
	switch protocol {
	case UDP:
		return GetFreeUDPPort(address)
	default:
		return GetFreeTCPPort(address)
	}
}

func GetFreePorts(address string, protocol Protocol, count int) ([]*Port, error) {
	ports := make([]*Port, count)
	for i := 0; i < count; i++ {
		port, err := GetFreePort(address, protocol)
		if err != nil {
			return nil, err
		}
		ports[i] = port
	}
	return ports, nil
}

func GetFreePortInRange(address string, protocol Protocol, minPort, maxPort int) (*Port, error) {
	if minPort > maxPort {
		return nil, errors.New("invalid interval")
	}
	for port := minPort; port <= maxPort; port++ {
		if port, err := GetPort(protocol, address, port); err == nil {
			return port, nil
		}
	}
	return nil, fmt.Errorf("couldn't find free ports between %d and %d", minPort, maxPort)
}

func GetFreeTCPPort(address string) (*Port, error) {
	addr, err := net.ResolveTCPAddr("tcp", address+":0")
	if err != nil {
		return nil, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return nil, err
	}
	if err := l.Close(); err != nil {
		return nil, err
	}
	var port int
	if tcpAddr, ok := l.Addr().(*net.TCPAddr); ok {
		port = tcpAddr.Port
	}

	return &Port{Address: l.Addr().String(), Port: port, Protocol: TCP}, nil
}

func GetPort(protocol Protocol, address string, port int) (*Port, error) {
	hostport := net.JoinHostPort(address, fmt.Sprint(port))
	switch protocol {
	case UDP:
		addr, err := net.ResolveUDPAddr("udp", hostport)
		if err != nil {
			return nil, err
		}
		l, err := net.ListenUDP("udp", addr)
		if err != nil {
			return nil, err
		}
		if err := l.Close(); err != nil {
			return nil, err
		}
		return &Port{Address: address, Port: port, Protocol: UDP}, nil
	default:
		l, err := net.Listen("tcp", hostport)
		if err != nil {
			return nil, err
		}
		if err := l.Close(); err != nil {
			return nil, err
		}
		return &Port{Address: address, Port: port, Protocol: TCP}, nil
	}
}

func GetFreeUDPPort(address string) (*Port, error) {
	addr, err := net.ResolveUDPAddr("udp", address+":0")
	if err != nil {
		return nil, err
	}

	l, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}
	if err := l.Close(); err != nil {
		return nil, err
	}
	var port int
	if udpAddr, ok := l.LocalAddr().(*net.UDPAddr); ok {
		port = udpAddr.Port
	}

	return &Port{Address: l.LocalAddr().String(), Port: port, Protocol: UDP}, nil
}

func MustGetFreeTCPPort(address string) *Port {
	port, err := GetFreeTCPPort(address)
	if err != nil {
		panic(err)
	}
	return port
}

func MustGetFreeUDPPort(address string) *Port {
	port, err := GetFreeUDPPort(address)
	if err != nil {
		panic(err)
	}
	return port
}

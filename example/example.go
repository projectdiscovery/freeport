package main

import (
	"fmt"

	"github.com/projectdiscovery/freeport"
)

func main() {

	fmt.Println("Single port:")
	// UDP port
	udpPort, err := freeport.GetFreePort("127.0.0.1", freeport.UDP)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("UDP port: %+v\n", udpPort)
	fmt.Println("UDP port: ", udpPort.Port)

	// TCP port
	tcpPort, err := freeport.GetFreePort("127.0.0.1", freeport.TCP)
	if err != nil {
		panic(err)
	}
	fmt.Println("TCP port: ", tcpPort.Port)

	fmt.Println("Multiple ports:")

	// get multiple ports
	// test with TCP
	tcpPorts, err := freeport.GetFreePorts("127.0.0.1", freeport.TCP, 5)
	if err != nil {
		panic(err)
	}

	for _, port := range tcpPorts {
		fmt.Println("TCP port: ", port.Port)
	}

	fmt.Println("Multiple ports in range:")

	tcpPortsInRange, err := freeport.GetFreePortInRange("127.0.0.1", freeport.TCP, 10000, 20000)
	if err != nil {
		panic(err)
	}

	fmt.Println("TCP port in range 10000-20000: ", tcpPortsInRange.Port)

}

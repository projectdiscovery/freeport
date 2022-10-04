package freeport

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetFreePort(t *testing.T) {
	t.Run("UDP Port", func(t *testing.T) {
		port, err := GetFreePort("127.0.0.1", UDP)
		require.Nil(t, err)
		require.NotNil(t, port)
	})
	t.Run("TCP Port", func(t *testing.T) {
		port, err := GetFreePort("127.0.0.1", TCP)
		require.Nil(t, err)
		require.NotNil(t, port)
	})
	t.Run("TCP Port on unknown address", func(t *testing.T) {
		port, err := GetFreePort("1.2.3.4", TCP)
		require.NotNil(t, err)
		require.Nil(t, port)
	})
}

func TestGetFreePorts(t *testing.T) {
	wanted := 10
	t.Run("UDP Ports", func(t *testing.T) {
		ports, err := GetFreePorts("127.0.0.1", UDP, wanted)
		require.Nil(t, err)
		require.Len(t, ports, wanted)
	})
	t.Run("TCP Ports", func(t *testing.T) {
		ports, err := GetFreePorts("127.0.0.1", TCP, wanted)
		require.Nil(t, err)
		require.Len(t, ports, wanted)
	})
	t.Run("Error on too many ports", func(t *testing.T) {
		ports, err := GetFreePorts("1.2.3.4", TCP, 70000)
		require.NotNil(t, err)
		require.Nil(t, ports)
	})
}

func TestGetFreePortInRange(t *testing.T) {
	t.Run("UDP in range", func(t *testing.T) {
		min, max := 10000, 20000
		port, err := GetFreePortInRange("127.0.0.1", UDP, min, max)
		require.Nil(t, err)
		require.NotNil(t, port)
		require.True(t, min <= port.Port && port.Port <= max)
	})
	t.Run("TCP in range", func(t *testing.T) {
		min, max := 10000, 20000
		port, err := GetFreePortInRange("127.0.0.1", TCP, min, max)
		require.Nil(t, err)
		require.NotNil(t, port)
		require.True(t, min <= port.Port && port.Port <= max)
	})
	t.Run("Invalid Interval", func(t *testing.T) {
		min, max := 20000, 10000
		port, err := GetFreePortInRange("127.0.0.1", TCP, min, max)
		require.NotNil(t, err)
		require.Nil(t, port)
	})
}

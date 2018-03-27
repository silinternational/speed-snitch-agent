package httputils

import (
	"net"
	"time"
)

// Established connection with local address and timeout support
func DialTimeout(network string, laddr *net.TCPAddr, raddr *net.TCPAddr, timeoutSeconds int) (net.Conn, error) {

	timeout := time.Duration(timeoutSeconds) * time.Second

	dialer := &net.Dialer{
		Timeout:   timeout,
		LocalAddr: laddr,
	}

	conn, err := dialer.Dial(network, raddr.String())
	return conn, err
}

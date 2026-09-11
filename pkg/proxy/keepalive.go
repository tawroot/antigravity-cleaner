package proxy

import (
	"context"
	"net"
	"time"
)

// NewResilientDialer creates a net.Dialer configured with aggressive keep-alives
// to keep SSE and gRPC streams alive through Iranian DPI middleboxes.
func NewResilientDialer() *net.Dialer {
	return &net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 15 * time.Second, // Aggressive 15s ping to prevent firewall idle drop
		DualStack: true,
	}
}

// DialWithKeepAlive dials a target network address with stream keepalive enabled
func DialWithKeepAlive(ctx context.Context, network, addr string) (net.Conn, error) {
	dialer := NewResilientDialer()
	conn, err := dialer.DialContext(ctx, network, addr)
	if err != nil {
		return nil, err
	}

	if tcpConn, ok := conn.(*net.TCPConn); ok {
		_ = tcpConn.SetKeepAlive(true)
		_ = tcpConn.SetKeepAlivePeriod(15 * time.Second)
		_ = tcpConn.SetNoDelay(true)
	}

	return conn, nil
}

package quick

import (
	"context"
	"net"
	"time"

	quic "github.com/lucas-clemente/quic-go"
)

var _ net.Conn = &Conn{}

// Conn is a generic quic connection implements net.Conn.
type Conn struct {
	conn    *net.UDPConn
	session quic.Session

	stream quic.Stream
}

func newConn(sess quic.Session, conn *net.UDPConn) (*Conn, error) {
	stream, err := sess.OpenStreamSync(context.Background())
	if err != nil {
		return nil, err
	}
	return &Conn{
		conn:    conn,
		session: sess,
		stream:  stream,
	}, nil
}

// Read implements the Conn Read method.
func (c *Conn) Read(b []byte) (int, error) {
	return c.stream.Read(b)
}

// Write implements the Conn Write method.
func (c *Conn) Write(b []byte) (int, error) {
	return c.stream.Write(b)
}

// LocalAddr returns the local network address.
func (c *Conn) LocalAddr() net.Addr {
	return c.session.LocalAddr()
}

// RemoteAddr returns the remote network address.
func (c *Conn) RemoteAddr() net.Addr {
	return c.session.RemoteAddr()
}

// Close closes the connection.
func (c *Conn) Close() error {
	if c.stream != nil {
		c.stream.Close()
	}
	if c.conn != nil {
		return c.conn.Close()
	}

	return nil
}

// SetDeadline sets the deadline associated with the listener. A zero time value disables the deadline.
func (c *Conn) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}

// SetReadDeadline implements the Conn SetReadDeadline method.
func (c *Conn) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

// SetWriteDeadline implements the Conn SetWriteDeadline method.
func (c *Conn) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}

// SetReadBuffer sets the size of the operating system's receive buffer associated with the connection.
func (c *Conn) SetReadBuffer(bytes int) error {
	return c.conn.SetReadBuffer(bytes)
}

// SetWriteBuffer sets the size of the operating system's transmit buffer associated with the connection.
func (c *Conn) SetWriteBuffer(bytes int) error {
	return c.conn.SetWriteBuffer(bytes)
}

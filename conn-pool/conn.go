package conn_pool

import (
	"net"
	"log"
)

type customConn struct {
	net.Conn
	pool *poolImpl
}

func(c *customConn) Close() error {
	if c.pool == nil {
		return ErrNilPool
	}

	err := c.pool.putConnToPool(c.Conn)
	if err != nil {
		log.Printf("Failed to put connection back to pool.")
		return err
	}

	return nil
}

func(p *poolImpl) wrapWithCustomConn(conn net.Conn) *customConn {
	customConn := &customConn{
		pool: p,
	}
	customConn.Conn = conn

	return customConn
}
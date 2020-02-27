package conn_pool

import (
	_ "github.com/go-sql-driver/mysql"
	"net"
	"log"
	"errors"
	"sync"
)

type Pool interface {
	GetConn() (conn net.Conn, err error)
	Close() error
	GetIdleConnsLength() int
	GetOpenConnsLength() int
}

type PoolConfig struct {
	InitialConns int
	MaxOpenConns int
	AllowNewConnOverMax bool
}

type poolImpl struct {
	conns chan net.Conn
	mu   sync.RWMutex
	config *PoolConfig
	connFactory factory
	openConns int
}

type factory func() (net.Conn, error)

var (
	ErrInvalidConfig = errors.New("pool config is not valid")
	ErrFactoryFailed =  errors.New("factory method failed to create a connection")
	ErrNilPool = errors.New("connection pool is nil")
	ErrNoIdleConn = errors.New("no idle connection is available")
	ErrMaxConnReached = errors.New("max connection limit reached")
)

func (p *poolImpl) GetConn() (net.Conn, error) {
	if p.conns == nil {
		log.Println("Connection channel is nil.")
		return nil, ErrNilPool
	}

	select {
	case conn := <- p.conns:
		return p.wrapWithCustomConn(conn), nil
	default:
		conn, err := p.createNewConn()
		if err != nil {
			return nil, err
		}

		return p.wrapWithCustomConn(conn), nil
	}
}

func (p *poolImpl) createNewConn() (net.Conn, error) {
	if p.hasMaxConns() && !p.config.AllowNewConnOverMax {
		return nil, ErrMaxConnReached
	}

	conn, err := p.connFactory()
	if err != nil {
		log.Println("Factory method failed to create a connection.")
		return nil, ErrFactoryFailed
	}

	p.openConns += 1
	return conn, nil
}

func (p *poolImpl) hasMaxConns() bool {
	if p.conns == nil {
		log.Println("Connection channel is nil.")
		return false
	}

	return p.openConns >= p.config.MaxOpenConns
}

func (p *poolImpl) putConnToPool(conn net.Conn) error {
	if p.conns == nil {
		log.Println("Connection channel is nil.")
		return ErrNilPool
	}

	p.conns <- conn
	return nil
}

func (p *poolImpl) Close() error {
	conns := p.conns

	close(p.conns)
	for conn := range conns {
		conn.Close()
		p.openConns -= 1
	}

	return nil
}


func (p *poolImpl) GetIdleConnsLength() int {
	if p.conns == nil {
		return 0
	}

	return len(p.conns)
}

func (p *poolImpl) GetOpenConnsLength() int {
	return p.openConns
}

func NewPool(config *PoolConfig, connFactory factory) (Pool, error) {
	if err := validateConfig(*config); err != nil {
		log.Println("Invalid pool config")
		return nil, err
	}

	conns := make(chan net.Conn, config.MaxOpenConns)
	pool := &poolImpl{
		conns: conns,
		config: config,
		connFactory: connFactory,
	}

	for i := 0; i < config.InitialConns; i++ {
		conn, err := pool.createNewConn()
		if err != nil {
			return nil, err
		}

		pool.conns <- conn
	}

	return pool, nil
}

func validateConfig(config PoolConfig) error {
	if config.InitialConns < 0 || config.MaxOpenConns <= 0 {
		return ErrInvalidConfig
	}
	return nil
}
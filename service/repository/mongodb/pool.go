package mongodb

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"time"
)

type driverConn struct {
	db        *DB
	createdAt time.Time

	sync.Mutex // guards following
	closed     bool
	cli        *mongo.Client

	// guarded by db.mu
	inUse bool
}

type DB struct {
	mu       sync.Mutex
	freeConn []*driverConn
	numOpen  int // number of opened and pending open connections
	// connRequests map[uint64]chan connRequest
	// nextRequest  uint64 // Next key to use in connRequests.

	maxOpen     int
	minIdle     int
	maxIdle     int
	maxLifetime time.Duration
	closed      bool

	fnNew   func() (*mongo.Client, error)
	fnClose func(*mongo.Client) error
	fnPing  func(*mongo.Client) error
}

func NewPool() *DB {

	d := &DB{
		maxOpen:     20,
		minIdle:     1,
		maxIdle:     5,
		maxLifetime: 1 * time.Hour,
	}
	return d
}

func (db *DB) SetMaxOpen(n int) {
	db.mu.Lock()
	db.maxOpen = n
	db.mu.Unlock()
}

func (db *DB) SetMaxIdle(n int) {
	db.mu.Lock()
	db.maxIdle = n
	db.mu.Unlock()
}

func (db *DB) SetFunNew(f func() (*mongo.Client, error)) {
	db.fnNew = f
}

func (db *DB) SetFunClose(f func(*mongo.Client) error) {
	db.fnClose = f
}

func (db *DB) SetFunPing(f func(*mongo.Client) error) {
	db.fnPing = f
}

func (db *DB) Get() (*driverConn, error) {

	// 1.
	db.mu.Lock()
	numFree := len(db.freeConn)
	if numFree > 0 {
		conn := db.freeConn[0]
		copy(db.freeConn, db.freeConn[1:])
		db.freeConn = db.freeConn[:numFree-1]
		conn.inUse = true

		if conn.expired(db.maxLifetime) {
			db.mu.Unlock()
			conn.Close()
			return nil, fmt.Errorf("out of lifetime")
		}

		db.mu.Unlock()
		return conn, nil
	}

	// 2. todo 排队等待
	if db.maxOpen > 0 && db.numOpen >= db.maxOpen {
		db.mu.Unlock()
		return nil, fmt.Errorf("it`s reached the maxOpen")
	}

	// 3. new conn
	db.numOpen++
	db.mu.Unlock()
	cli, err := db.fnNew()
	if err != nil {
		db.mu.Lock()
		db.numOpen--
		db.mu.Unlock()
		return nil, err
	}

	db.mu.Lock()
	dc := &driverConn{
		db:        db,
		createdAt: time.Now(),
		inUse:     true,
		cli:       cli,
	}
	db.mu.Unlock()
	return dc, nil
}

func (db *DB) Put(conn *driverConn) bool {

	if db.maxOpen > 0 && db.numOpen > db.maxOpen {
		return false
	}

	db.mu.Lock()
	conn.inUse = false
	db.freeConn = append(db.freeConn, conn)
	db.mu.Unlock()
	return true
}

func (dc *driverConn) Close() error {

	dc.Lock()
	if dc.closed {
		dc.Unlock()
		return fmt.Errorf("duplicate driverConn close")
	}
	dc.closed = true
	dc.Unlock()

	return dc.db.fnClose(dc.GetCli())
}

func (dc *driverConn) GetCli() *mongo.Client {
	return dc.cli
}

func (dc *driverConn) expired(timeout time.Duration) bool {
	if timeout <= 0 {
		return false
	}
	return dc.createdAt.Add(timeout).Before(time.Now())
}

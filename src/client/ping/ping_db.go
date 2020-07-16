package ping

import (
	"database/sql"
	"errors"
	"flag"
	"time"
)

var (
	databaseTimeout = flag.Int64("database-timeout-ms", 5000, "")
)

func WaitForDB(conn *sql.DB) error {
	ready := make(chan struct{})
	go func() {
		for {
			if err := conn.Ping(); err == nil {
				// if err := conn.Ping(); err == nil {
				close(ready)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
	select {
	case <-ready:
		return nil
	case <-time.After(time.Duration(*databaseTimeout) * time.Millisecond):
		return errors.New("Database Not ready")
	}
}

package postgres

import "github.com/jackc/pgx/v4/pgxpool"

type Conn struct {
	*pgxpool.Conn
}

func (c *Conn) Close() {
	c.Conn.Release()
}

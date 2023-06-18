// db/db.go -- db submodule for demo
//
// Keeping it simple here!  The point of this is that the tests for this
// will deal with a live database.

package db

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DefaultTable = "foo"

// Client is a database client with a Pool of connections.
type Client struct {
	Pool  *pgxpool.Pool
	Table string
}

// New creates a Client with a Pool connected to the Postgres database
// described in the environment variable DATABASE_URL.
func New() (*Client, error) {

	dburl := os.Getenv("DATABASE_URL")
	if dburl == "" {
		return nil, errors.New("DATABASE_URL not defined in env")
	}

	pool, err := pgxpool.New(context.Background(), dburl)
	if err != nil {
		return nil, err
	}

	client := &Client{
		Pool:  pool,
		Table: DefaultTable,
	}
	return client, nil
}

// InsertFoo inserts bar into the foo table.
func (c *Client) InsertFoo(bar string) error {
	sql := fmt.Sprintf("INSERT INTO %s (bar) VALUES ($1);", c.Table)
	_, err := c.Pool.Exec(context.Background(), sql, bar)
	return err
}

// DeleteFoo deletes the foo for bar.
func (c *Client) DeleteFoo(bar string) error {
	sql := fmt.Sprintf("DELETE FROM %s WHERE bar = $1", c.Table)
	_, err := c.Pool.Exec(context.Background(), sql, bar)
	return err
}

// CountFoo counts the distinct bar columns in foo, which should be identical
// to all rows in the database, foo.bar being a primary key.
func (c *Client) CountFoo() (int, error) {
	sql := fmt.Sprintf("SELECT COUNT(DISTINCT bar) FROM %s;", c.Table)
	count := -1
	row := c.Pool.QueryRow(context.Background(), sql)
	err := row.Scan(&count)
	return count, err
}

// demo.go

package demo

import (
	"example.com/demo/db"
)

// NewClient is a thin wrapper for db.New.
func NewClient() (*db.Client, error) {

	return db.New()

}

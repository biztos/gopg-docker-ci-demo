// demo_test.go

package demo_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"example.com/demo"
)

func TestNewClientErrorNoUrl(t *testing.T) {

	assert := assert.New(t)

	orig := os.Getenv("DATABASE_URL")
	defer os.Setenv("DATABASE_URL", orig)
	os.Setenv("DATABASE_URL", "")

	_, err := demo.NewClient()

	assert.ErrorContains(err, "DATABASE_URL")
}

func TestDeliberateFail(t *testing.T) {

	assert := assert.New(t)

	assert.True(false, "this will fail")
}

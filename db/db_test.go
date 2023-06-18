// db_test.go -- tests for the db submodule

package db_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/suite"

	"example.com/demo/db"
)

func (suite *DbTestSuite) TestNewOK() {

	require := suite.Require()

	client, err := db.New()
	require.NoError(err)
	require.NotNil(client)
	require.NotNil(client.Pool)

}

func (suite *DbTestSuite) TestNewErrorNoUrl() {

	require := suite.Require()

	orig := os.Getenv("DATABASE_URL")
	defer os.Setenv("DATABASE_URL", orig)
	os.Setenv("DATABASE_URL", "")

	client, err := db.New()
	require.ErrorContains(err, "DATABASE_URL")
	require.Nil(client)

}

func (suite *DbTestSuite) TestNewErrorBadUrl() {

	require := suite.Require()

	orig := os.Getenv("DATABASE_URL")
	defer os.Setenv("DATABASE_URL", orig)
	os.Setenv("DATABASE_URL", "not gonna work")

	client, err := db.New()
	require.ErrorContains(err, "DSN")
	require.Nil(client)

}

func (suite *DbTestSuite) TestInsertFoo() {

	require := suite.Require()

	// IRL we would check the value; here we just care it didn't error out.
	require.NoError(suite.Client.InsertFoo("hello"))

}

func (suite *DbTestSuite) TestDeleteFoo() {

	require := suite.Require()

	require.NoError(suite.Client.DeleteFoo("hello"))

}

func (suite *DbTestSuite) TestCountFoo() {

	require := suite.Require()

	// could be... more than zero!
	count, err := suite.Client.CountFoo()
	require.NoError(err)
	require.Less(count, 1)
	require.Greater(count, -1)

}

// ---------------------------------------------------------------------------
// SUITE RIGGING FOLLOWS BELOW:
// ---------------------------------------------------------------------------

type DbTestSuite struct {
	suite.Suite
	Client *db.Client
}

func (suite *DbTestSuite) SetupSuite() {

	require := suite.Require()

	// A new table for every run!
	table := fmt.Sprintf("test_foo_%s", ulid.Make())
	db.DefaultTable = table

	client, err := db.New()
	require.NoError(err)
	suite.Client = client

	sql := fmt.Sprintf("CREATE TABLE %s (bar TEXT PRIMARY KEY);", table)
	_, execerr := client.Pool.Exec(context.Background(), sql)
	require.NoError(execerr)

}

func (suite *DbTestSuite) TeardownSuite() {

	require := suite.Require()

	// Thou Shalt Not Leave Turds!
	if suite.Client != nil {
		sql := fmt.Sprintf("DROP TABLE IF EXISTS %s;", suite.Client.Table)
		_, err := suite.Client.Pool.Exec(context.Background(), sql)
		require.NoError(err)

	}

}

func (suite *DbTestSuite) SetupTest() {

	require := suite.Require()

	// Zero out the table for every test (but don't create/drop it).
	sql := fmt.Sprintf("TRUNCATE TABLE %s;", suite.Client.Table)
	_, err := suite.Client.Pool.Exec(context.Background(), sql)
	require.NoError(err)

}

// The actual runner func:
func TestDbTestSuite(t *testing.T) {
	suite.Run(t, new(DbTestSuite))
}

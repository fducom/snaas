// +build integration

package user

import (
	"flag"
	"fmt"
	"os/user"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var pgTestURL string

func TestPostgresCount(t *testing.T) {
	testServiceCount(t, preparePostgres)
}

func TestPostgresPut(t *testing.T) {
	testServicePut(t, preparePostgres)
}

func TestPostgresPutEmailUnique(t *testing.T) {
	testServicePutEmailUnique(t, preparePostgres)
}

func TestPostgresPutLastRead(t *testing.T) {
	testServicePutLastRead(t, preparePostgres)
}

func TestPostgresPutUsernameUnique(t *testing.T) {
	testServicePutUsernameUnique(t, preparePostgres)
}

func TestPostgresQuery(t *testing.T) {
	testServiceQuery(t, preparePostgres)
}

func TestPostgresSearch(t *testing.T) {
	testServiceSearch(t, preparePostgres)
}

func preparePostgres(t *testing.T, namespace string) Service {
	db, err := sqlx.Connect("postgres", pgTestURL)
	if err != nil {
		t.Fatal(err)
	}

	s := PostgresService(db)

	if err := s.Teardown(namespace); err != nil {
		t.Fatal(err)
	}

	return s
}

func init() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	d := fmt.Sprintf(
		"postgres://%s@127.0.0.1:5432/tapglue_test?sslmode=disable&connect_timeout=5",
		user.Username,
	)

	url := flag.String("postgres.url", d, "Postgres connection URL")
	flag.Parse()

	pgTestURL = *url
}

package gomodtest

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"testing"
)

func OpenDB(host, port string) (*sql.DB, error) {
	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "4000"
	}
	return sql.Open("mysql", fmt.Sprintf("root@tcp(%s:%s)/test?charset=utf8mb4,utf8", host, port))
}

func TestRandom(t *testing.T) {
	src := rand.NewSource(1024)
	r1 := rand.New(src)
	fmt.Println(r1.Uint32())
}

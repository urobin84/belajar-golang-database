package belajar_golang_database

import (
	"database/sql"
	//_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "user:ACT2019DBMaster@tcp(dev-act-2019.cluster-clgqjwhmhdp5.ap-southeast-1.rds.amazonaws.com:3306)/dataact")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}

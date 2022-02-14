package db

import (
	_ "embed" // support embedding files in variables
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/jmoiron/sqlx"

	"github.com/ebobo/investment_calulator_record/pkg/model"
)

//go:embed schema.sql
var schema string

var mutex sync.RWMutex

func CreateDataBase() error {
	os.Remove("ic-database.db")

	// make data dir if it is not exit
	err := makeDirIfNotExists("../data")
	if err != nil {
		log.Fatal(err.Error())
	}
	// Create SQLite file
	file, err := os.Create("../data/ic-database.db")
	if err != nil {
		return err
	}
	file.Close()

	log.Println("ic-database.db created")

	return nil
}

func CreateSchema(db *sqlx.DB) error {
	for n, statement := range strings.Split(schema, ";") {
		_, err := db.Exec(statement)
		if err != nil {
			return fmt.Errorf("statement %d failed: \"%s\" : %w", n+1, statement, err)
		}
	}
	return nil
}

func Addrecord(db *sqlx.DB, record *model.Report) error {
	mutex.Lock()
	defer mutex.Unlock()

	log.Println(record)
	_, err := db.NamedExec(
		`INSERT INTO records (client, total_interest, periodic_payment, total_payment)
         VALUES (:client, :total_interest, :periodic_payment, :total_payment)`,
		record,
	)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

func PrintRecords(db *sqlx.DB) {
	row, err := db.Query("SELECT * FROM records ORDER BY client")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int32
		var client string
		var interest float32
		var payment float32
		var t_payment float32
		row.Scan(&id, &client, &interest, &payment, &t_payment)
	}
}

func GetRecordsByClient(db *sqlx.DB, client string) ([]model.Report, error) {
	var reports []model.Report
	return reports, db.Select(&reports, "SELECT * FROM records WHERE client = ? ORDER BY id", client)
}

// sqlite
// func GetRecordsByClient(db *sqlx.DB, client string) ([]model.Report, error) {
// 	mutex.RLock()
// 	defer mutex.RUnlock()

// 	var reports []model.Report
// 	return reports, db.Select(&reports, "SELECT * FROM records WHERE client = ? ORDER BY id", client)
// }

func makeDirIfNotExists(dirpath string) error {
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		err := os.Mkdir(dirpath, os.ModeDir|os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

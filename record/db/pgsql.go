package db

import (
	_ "embed" // support embedding files in variables
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/ebobo/investment_calulator_record/pkg/model"
)

//go:embed schema_pg.sql
var schema_pg string

func CreateSchemaPG(db *sqlx.DB) error {
	for n, statement := range strings.Split(schema_pg, ";") {
		_, err := db.Exec(statement)
		if err != nil {
			return fmt.Errorf("statement %d failed: \"%s\" : %w", n+1, statement, err)
		}
	}
	return nil
}

func AddrecordToTable(db *sqlx.DB, record *model.Report) error {
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

func PrintAllRecords(db *sqlx.DB) {
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

func GetRecordsByClientName(db *sqlx.DB, client string) ([]model.Report, error) {
	var reports []model.Report
	return reports, db.Select(&reports, "SELECT * FROM records WHERE client = $1 ORDER BY id", client)
}

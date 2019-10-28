package account

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	"github.com/jinzhu/gorm"
	"github.com/lancatlin/library/pkg/model"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/joho/godotenv/autoload"
)

var DB = os.Getenv("ACCOUNT_DB")

func LoadAllAccounts(db *gorm.DB) (err error) {
	return loadFromDB(db)
}

func loadFromDB(db *gorm.DB) (err error) {
	acctDB, err := getAccountDB()
	if err != nil {
		return
	}
	return saveToDB(readRaws(acctDB), db)
}

func getAccountDB() (acctDB *sql.DB, err error) {
	log.Println(DB)
	acctDB, err = sql.Open("sqlite3", DB)
	if err != nil {
		return nil, fmt.Errorf("getAccountDB fatal: %w", err)
	}
	if err = acctDB.Ping(); err != nil {
		return nil, fmt.Errorf("getAccountDB fatal: %w", err)
	}
	return
}

func readRaws(acctDB *sql.DB) (raws *sql.Rows) {
	query := `select id, name, phone, role from accounts;`
	raws, err := acctDB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func saveToDB(rows *sql.Rows, db *gorm.DB) (err error) {
	for rows.Next() {
		var acct model.Account
		if err = rows.Scan(&acct.ID, &acct.Name, &acct.Phone, &acct.Role); err != nil {
			return fmt.Errorf("Scan fatal: %w", err)
		}
		if db.NewRecord(acct) {
			if err = db.Create(&acct).Error; err != nil {
				return fmt.Errorf("Write to database fatal: %w on %v", err, acct)
			}
		} else {
			if err = db.Save(&acct).Error; err != nil {
				return fmt.Errorf("Update to database fatal: %w on %v", err, acct)
			}
		}
	}
	return nil
}

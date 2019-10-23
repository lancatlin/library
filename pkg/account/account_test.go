package account

import (
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lancatlin/library/pkg/model"
)

func TestLoadFromDB(t *testing.T) {
	db, err := gorm.Open("sqlite3", time.Now().Format("testDB/01-02_03-04-05.sqlite"))
	if err != nil {
		t.Fatal(err)
	}
	model.SetDB(db)
	err = loadFromDB(db)
	if err != nil {
		t.Fatal(err)
	}
}

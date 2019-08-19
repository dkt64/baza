package baza

import (
	"fmt"

	"github.com/jinzhu/gorm"

	// required
	_ "github.com/mattn/go-sqlite3"
)

var bazaFilename string

// ErrCheck - obsługa błedów
// ========================================================
func errCheck(errNr error) {
	if errNr != nil {
		fmt.Println(errNr)
	}
}

// Init - inicjacja bazy
// ========================================================
func Init(datatype interface{}, filename string) {

	db, err := gorm.Open("sqlite3", filename)
	errCheck(err)
	defer db.Close()

	bazaFilename = filename

	// Jeżeli jeszcze nie utworzona to utworzyć
	if !db.HasTable(datatype) {
		db.CreateTable(datatype)
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(datatype)
	}
}

// GetAllRecords - odczytanie bazy
// ========================================================
func GetAllRecords(data interface{}) {

	db, err := gorm.Open("sqlite3", bazaFilename)
	errCheck(err)
	defer db.Close()

	db.Find(data)
}

// GetOneRecord - odczytanie rekordu z bazy
// ========================================================
func GetOneRecord(data interface{}, nr int) {

	db, err := gorm.Open("sqlite3", bazaFilename)
	errCheck(err)
	defer db.Close()

	db.First(data, nr)
}

// AddNewRecord - dodanie rekordu
// ========================================================
func AddNewRecord(data interface{}) {

	db, err := gorm.Open("sqlite3", bazaFilename)
	errCheck(err)
	defer db.Close()

	db.Create(data)
}

// UpdateRecord - zmiana rekordu
// ========================================================
func UpdateRecord(data interface{}) {

	db, err := gorm.Open("sqlite3", bazaFilename)
	errCheck(err)
	defer db.Close()

	db.Save(data)
}

// DeleteRecord - usunięcie rekordu z bazy
// ========================================================
func DeleteRecord(data interface{}) {

	db, err := gorm.Open("sqlite3", bazaFilename)
	errCheck(err)
	defer db.Close()

	db.Delete(data)
}

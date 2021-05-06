package database

import (
	"fmt"
	"github.com/abhiongithub/university-erp/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "erp.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	DBConn.AutoMigrate(&models.Course{})
	fmt.Println("Database Migrated")
}
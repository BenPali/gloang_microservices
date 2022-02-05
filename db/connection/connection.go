package connection

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// connect opens a connection to the database.
func connect(host, port, username, password, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

var db_host = os.Getenv("DB_HOST")
var db_port = os.Getenv("DB_PORT")
var db_user = os.Getenv("DB_USER")
var db_password = os.Getenv("MARIADB_ROOT_PASSWORD")
var db_database = os.Getenv("DB_DATABASE")
var DB *gorm.DB

//InitDb initialize a connection to the database, stored in the DB variable.
// returns error if gorm could not connect to the database
func InitDB() error {
	var err error

	DB, err = connect(db_host, db_port, db_user, db_password, db_database)

	if err != nil {
		return err
	}

	return nil
}

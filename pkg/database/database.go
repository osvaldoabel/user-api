package database

import (
	"fmt"

	"gorm.io/driver/postgres"

	"github.com/joho/godotenv"
	"github.com/osvaldoabel/user-api/configs"
	"github.com/osvaldoabel/user-api/internal/entity"
	"gorm.io/gorm"
)

type DBConfig struct {
	DB            *gorm.DB
	DBUser        string
	DBSSLmode     string
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

type Database struct {
	DB            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func init() {
	godotenv.Load(".env")

}

func FakeDB() *Database {
	return &Database{}
}

func newDb() *Database {
	return &Database{}
}

// InitDB
func InitDB(conf configs.Conf) (*gorm.DB, error) {
	db := newDb()

	db.Dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		conf.DBhost,
		conf.DBUser,
		conf.DBPassword,
		conf.DBName,
		conf.DBPort,
		conf.DBSSLmode,
		conf.DBtimezone)
	db.DbType = conf.DBDriver
	db.AutoMigrateDb = conf.DBAutoMigrate

	conn, err := db.Connect()
	if err != nil {
		panic(err)
	}

	return conn, nil
}

func (d *Database) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(d.Dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if d.AutoMigrateDb {
		db.AutoMigrate(&entity.User{})
	}

	return db, nil
}

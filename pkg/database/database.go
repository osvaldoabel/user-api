package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"

	"github.com/joho/godotenv"
	"github.com/osvaldoabel/user-api/configs"
	"github.com/osvaldoabel/user-api/internal/entity"
	"gorm.io/gorm"
)

const (
	IN_MEMORY_DSN = "file::memory:"
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

func NewInmemoryDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(IN_MEMORY_DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func NewEmptyDB() *Database {
	return &Database{
		DB: NewInmemoryDB(),
	}
}

func NewDbTest() *Database {
	db := NewInmemoryDB()

	if err := db.Migrator().HasTable(&entity.User{}); !err {
		db.AutoMigrate(&entity.User{})
	}

	return &Database{DB: db}
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

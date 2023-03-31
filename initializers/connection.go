package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

// DB ...
var DB *gorm.DB

// Init ...
func Init() *gorm.DB {

	// 	DB_HOST=los-db
	// # DB_HOST=127.0.0.1                             # when running the app without docker
	// DB_DRIVER=postgres                         # Used for creating a JWT. Can be anything
	// DB_USER=uatpostgres
	// DB_PASSWORD=12345
	// DB_NAME=los
	// DB_PORT=5432

	h := os.Getenv("DB_HOST")
	u := os.Getenv("DB_USER")
	p := os.Getenv("DB_PASSWORD")
	dbn := os.Getenv("DB_NAME")
	// port := os.Getenv("DB_PORT")
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//   db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:13306)/%s?charset=utf8mb4&parseTime=True&loc=Local", u, p, dbn)
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", h, u, p, dbn, port)

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// fmt.Println(dsn)
	if err != nil {
		log.Fatal("DB connection failed!")

	}
	fmt.Printf("Connected to PostgreSQL on %s as user %s \n", h, u)
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}

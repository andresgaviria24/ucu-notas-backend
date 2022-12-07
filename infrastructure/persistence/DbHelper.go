package persistence

import (
	"fmt"
	"time"
	"university/infrastructure/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbHelper struct {
	StudentsRepository repository.StudentsRepository
	ordersDb           *gorm.DB
}

func InitOrdersHelper() (*DbHelper, error) {

	/*user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)*/
	config := fmt.Sprintf("root@tcp(127.0.0.1:3306)/sys?charset=utf8mb4&parseTime=True&loc=Local")

	ordersDb, err := gorm.Open(mysql.Open(config),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info), SkipDefaultTransaction: true})

	if err != nil {
		panic(err)
	}

	sqlDB, err := ordersDb.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(600)
	sqlDB.SetMaxOpenConns(0)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)
	ordersDb.AutoMigrate()
	return &DbHelper{
		StudentsRepository: &StudentsRepositoryImpl{ordersDb},
		ordersDb:           ordersDb,
	}, nil
}

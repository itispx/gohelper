package gorm

import (
	"fmt"
	"os"

	"github.com/itispx/eruka/aws/secretsmanager"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Host   = os.Getenv("DB_HOST")
	Logger logger.LogLevel
)

func ConnectPostgres() (*gorm.DB, error) {
	if os.Getenv("ENV") != "dev" {
		dbSecretArn := os.Getenv("DB_SECRET_ARN")

		return ConnPostgresWithAWSSM(dbSecretArn)
	}

	return ConnPostgresWithEnvironment()
}

func ConnPostgresWithAWSSM(dbSecretArn string) (*gorm.DB, error) {
	credentials, err := secretsmanager.GetDBCredentials(&dbSecretArn)
	if err != nil {
		return nil, err
	}

	dbName := credentials.DBName
	port := fmt.Sprintf(`%v`, credentials.Port)
	user := credentials.Username
	password := credentials.Password
	sslmode := "require"

	dsn := fmt.Sprintf(`
	host=%s 
	user=%s 
	password=%s 
	dbname=%s 
	port=%v 
	sslmode=%s`,
		Host,
		user,
		password,
		dbName,
		port,
		sslmode,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(Logger),
	})
}

func ConnPostgresWithEnvironment() (*gorm.DB, error) {
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	sslmode := "disable"

	dsn := fmt.Sprintf(`
	host=%s 
	user=%s 
	password=%s 
	dbname=%s 
	port=%v 
	sslmode=%s`,
		host,
		user,
		password,
		dbName,
		port,
		sslmode,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

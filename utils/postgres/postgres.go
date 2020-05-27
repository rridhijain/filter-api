package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	_ "github.com/lib/pq"
	"github.com/rridhijain/filter-api/utils/viper"
)

type PostgresDatabase struct {
	*sql.DB
}

func getConnectionString() string {
	configuration := viper.Setup()
	return fmt.Sprintf("host=%s dbname=%s user=%s "+
		"password=%s sslmode=disable",
		configuration.Database.Host, configuration.Database.DBName,
		configuration.Database.User, configuration.Database.Password)
}

func OpenConnection() *PostgresDatabase {

	db, err := sql.Open("postgres", getConnectionString())
	// db, err := sql.Open("cloudsqlpostgres", getConnectionString())
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return &PostgresDatabase{db}
}

func (db *PostgresDatabase) CloseConnection() {
	defer db.Close()
	fmt.Println("Successfully Disconnected!")
}

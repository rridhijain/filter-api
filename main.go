package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/rridhijain/filter-api/api/filters/salesDashboard/programTypeTimeSlotFilter"
	"github.com/rridhijain/filter-api/utils/postgres"
	"github.com/rridhijain/filter-api/utils/viper"
)

func main() {
	db := postgres.OpenConnection()
	defer db.CloseConnection()
	fmt.Println("Application Running")
	http.HandleFunc("/graphql", programTypeTimeSlotFilter.GetResponse(db))

	configuration := viper.Setup()
	http.ListenAndServe(":"+strconv.Itoa(configuration.Server.Port), nil)
}

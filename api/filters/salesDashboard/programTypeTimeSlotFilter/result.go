package programTypeTimeSlotFilter

import (
	"encoding/json"
	"fmt"

	// "fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/rridhijain/filter-api/api/filters/salesDashboard/programTypeTimeSlotFilter/schemas"
	"github.com/rridhijain/filter-api/utils/postgres"
	"github.com/rridhijain/filter-api/utils/services"
)

// This is the trial documentation
func GetResponse(db *postgres.PostgresDatabase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		resolver := &Resolver{db}

		var queryType = graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"filters": &graphql.Field{
						Type: schemas.DashboardFields,
						Args: graphql.FieldConfigArgument{
							"period_dates": &graphql.ArgumentConfig{
								Type: graphql.NewList(schemas.PeriodType),
							},
							"deviation_period": &graphql.ArgumentConfig{
								Type: graphql.NewList(schemas.PeriodType),
							},
						},
						Resolve: resolver.DashboardResolver,
					},
				},
			})
		var schema, _ = graphql.NewSchema(
			graphql.SchemaConfig{
				Query: queryType,
			},
		)
		result := services.ExecuteQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	}
}

// GetFilters api is used to get cyclic dashboard filters
func GetFilters(db *postgres.PostgresDatabase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		resolver := &Resolver{db}

		decoder := json.NewDecoder(r.Body)
		var t schemas.InputPeriod
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		result, err := resolver.DashboardFilterResolver(t)
		fmt.Println(result)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		//w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(result); err != nil {
			panic(err)
		}
	}
}

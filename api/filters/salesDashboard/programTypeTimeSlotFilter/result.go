package programTypeTimeSlotFilter

import (
	"encoding/json"
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

package programTypeTimeSlotFilter

import (
	"github.com/rridhijain/filter-api/utils/postgres"
	"github.com/graphql-go/graphql"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	*postgres.PostgresDatabase
}

// UserResolver resolves our user query through a db call to GetUserByName
func (r *Resolver) DashboardResolver(p graphql.ResolveParams) (interface{}, error) {
	period_dates, period_dates_present := p.Args["period_dates"].([]interface{})
	deviation_period, deviation_period_present := p.Args["deviation_period"].([]interface{})
	
	startDates, endDates := getDatesArr(period_dates)
	deviation_period_start_dates, deviation_period_end_dates := getDatesArr(deviation_period)
	startDates = append(startDates, deviation_period_start_dates...)
	endDates = append(endDates, deviation_period_end_dates...)
	
	if (period_dates_present || deviation_period_present){
		result := GetDashboardFilters(startDates, endDates, r.PostgresDatabase)
		return result, nil
	}
	return nil, nil
}


func getDatesArr(periods []interface{}) ([]string, []string){
	startDates := make([]string, 0)
	endDates := make([]string, 0)
	for _, value := range periods {
		startDates = append(startDates, value.(map[string]interface{})["start_date"].(string))
		endDates = append(endDates, value.(map[string]interface{})["end_date"].(string))
	}
	return startDates, endDates
}

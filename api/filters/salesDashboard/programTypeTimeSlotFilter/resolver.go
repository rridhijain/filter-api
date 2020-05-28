package programTypeTimeSlotFilter

import (
	"github.com/graphql-go/graphql"
	"github.com/rridhijain/filter-api/utils/postgres"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	*postgres.PostgresDatabase
}

// DashboardResolver resolves our user query through a db call to GetUserByName
func (r *Resolver) DashboardResolver(p graphql.ResolveParams) (interface{}, error) {
	periodDates, periodDatesPresent := p.Args["period_dates"].([]interface{})
	deviationPeriod, deviationPeriodPresent := p.Args["deviation_period"].([]interface{})

	startDates, endDates := getDatesArr(periodDates)
	deviationPeriodStartDates, deviationPeriodEndDates := getDatesArr(deviationPeriod)
	startDates = append(startDates, deviationPeriodStartDates...)
	endDates = append(endDates, deviationPeriodEndDates...)

	if periodDatesPresent || deviationPeriodPresent {
		result := GetDashboardFilters(startDates, endDates, r.PostgresDatabase)
		return result, nil
	}
	return nil, nil
}

func getDatesArr(periods []interface{}) ([]string, []string) {
	startDates := make([]string, 0)
	endDates := make([]string, 0)
	for _, value := range periods {
		startDates = append(startDates, value.(map[string]interface{})["start_date"].(string))
		endDates = append(endDates, value.(map[string]interface{})["end_date"].(string))
	}
	return startDates, endDates
}

package programTypeTimeSlotFilter

import (
	"github.com/graphql-go/graphql"
	"github.com/rridhijain/filter-api/api/filters/salesDashboard/programTypeTimeSlotFilter/schemas"
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

func (r *Resolver) DashboardFilterResolver(p schemas.InputPeriod) (schemas.ProgramTypeAndTimeSlotUpdated1, error) {
	periodDates := p.PeriodDates
	deviationPeriod := p.DeviationPeriod

	startDates, endDates := getDatesArrUp(periodDates)
	deviationPeriodStartDates, deviationPeriodEndDates := getDatesArrUp(deviationPeriod)
	startDates = append(startDates, deviationPeriodStartDates...)
	endDates = append(endDates, deviationPeriodEndDates...)

	result := GetDashboardFiltersUpdate(startDates, endDates, r.PostgresDatabase)
	return result, nil
}

func getDatesArrUp(periods []schemas.Period) ([]string, []string) {
	startDates := make([]string, 0)
	endDates := make([]string, 0)
	for _, value := range periods {
		startDates = append(startDates, value.StartDate)
		endDates = append(endDates, value.EndDate)
	}
	return startDates, endDates
}

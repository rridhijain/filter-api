package programTypeTimeSlotFilter

import (
	"database/sql"
	"fmt"

	"github.com/rridhijain/filter-api/utils/postgres"
	"github.com/rridhijain/filter-api/api/filters/salesDashboard/programTypeTimeSlotFilter/schemas"
)

// Export method to get dashboard filters
func GetDashboardFilters(startDates []string, endDates []string, db *postgres.PostgresDatabase) []schemas.ProgramTypeAndTimeSlot {

	sqlStatement := "SELECT program_type, time_slot, date, advertiser, channel, region FROM program_info " + whereCondition(startDates, endDates)
	fmt.Println(sqlStatement)
	rows, err := db.Query(sqlStatement)
	filters := make([]schemas.ProgramTypeAndTimeSlot, 0)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(filters)
	default:
		panic(err)
	}
	for rows.Next() {
		var filter schemas.ProgramTypeAndTimeSlot 
		rows.Scan(&filter.ProgramType, 
			&filter.TimeSlot, &filter.Date, &filter.Advertiser, &filter.Channel, &filter.Region)
		filters = append(filters, filter)
	}
	return filters
}

func whereCondition(startDates []string, endDates []string) string{
	whereCon := ""
	for index, date := range startDates{
		condition := fmt.Sprintf("date between '%s' and '%s'", date, endDates[index])
		if(len(whereCon) == 0){
			whereCon += fmt.Sprintf("where ( %s )", condition)
		}else{
			whereCon += fmt.Sprintf(" OR ( %s )", condition)
		}
	}
	return whereCon
}

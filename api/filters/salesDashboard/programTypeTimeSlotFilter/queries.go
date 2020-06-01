package programTypeTimeSlotFilter

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/rridhijain/filter-api/api/filters/salesDashboard/programTypeTimeSlotFilter/schemas"
	"github.com/rridhijain/filter-api/utils/postgres"
)

// GetDashboardFilters Export method to get dashboard filters
func GetDashboardFilters(startDates []string, endDates []string, db *postgres.PostgresDatabase) schemas.ProgramTypeAndTimeSlotUpdated {
	sqlStatement := `
		SELECT 
			program_type
			, time_slot, advertiser_group, channel_name, region 
		FROM ent_fact_revenue_advertiser_mappings` +
		whereCondition(startDates, endDates) +
		` and time_slot is not null
			and region is not null
		GROUP BY program_type
		, time_slot, advertiser_group, channel_name, region`

	fmt.Println(sqlStatement)
	rows, err := db.Query(sqlStatement)
	var filters schemas.ProgramTypeAndTimeSlotUpdated
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(filters)
	default:
		panic(err)
	}
	programTypeArr := make([]schemas.ProgramTypeStruct, 0)
	timeSlotArr := make([]schemas.TimeSlotStruct, 0)

	index := 0
	for rows.Next() {
		var programTypeAndTimeSlot schemas.ProgramTypeAndTimeSlot

		rows.Scan(&programTypeAndTimeSlot.ProgramType, &programTypeAndTimeSlot.TimeSlot, &programTypeAndTimeSlot.AdvertiserGroup, &programTypeAndTimeSlot.ChannelName, &programTypeAndTimeSlot.Region)

		advObj := make([]string, 0)
		advObj = append(advObj, programTypeAndTimeSlot.AdvertiserGroup)

		if index == 0 {
			regionObj := getRegionObj(advObj, programTypeAndTimeSlot.Region.String)
			regionObjList := make([]schemas.Region, 0)
			regionObjList = append(regionObjList, regionObj)
			channelNameObj := getChannelObj(regionObjList, programTypeAndTimeSlot.ChannelName.String)
			channelNameObjList := make([]schemas.ChannelName, 0)
			channelNameObjList = append(channelNameObjList, channelNameObj)
			programType := getProgramTypeObj(channelNameObjList, programTypeAndTimeSlot.ProgramType.String)
			programTypeArr = append(programTypeArr, programType)
			timeSlot := getTimeSlotObj(channelNameObjList, programTypeAndTimeSlot.TimeSlot.String)
			timeSlotArr = append(timeSlotArr, timeSlot)
		} else {
			programTypeArr = append(programTypeArr, formatProgramTypeResponse(programTypeAndTimeSlot, advObj)...)
			timeSlotArr = append(timeSlotArr, formatTimeSlotResponse(programTypeAndTimeSlot, advObj)...)
		}
		index++
	}
	filters.ProgramType = programTypeArr
	filters.TimeSlot = timeSlotArr
	return filters
}

func GetDashboardFiltersUpdate(startDates []string, endDates []string, db *postgres.PostgresDatabase) schemas.ProgramTypeAndTimeSlotUpdated1 {
	sqlStatement := `
		SELECT 
			program_type
			, time_slot, advertiser_group, channel_name, region 
		FROM ent_fact_revenue_advertiser_mappings` +
		whereCondition(startDates, endDates) +
		` and time_slot is not null
			and region is not null
		GROUP BY program_type
		, time_slot, advertiser_group, channel_name, region`

	fmt.Println(sqlStatement)
	rows, err := db.Query(sqlStatement)
	var filters schemas.ProgramTypeAndTimeSlotUpdated1
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(filters)
	default:
		panic(err)
	}
	timeSlotMap := make(map[string]map[string]map[string][]string)
	progamTypeMap := make(map[string]map[string]map[string][]string)

	index := 0
	for rows.Next() {
		var programTypeAndTimeSlot schemas.ProgramTypeAndTimeSlot

		rows.Scan(&programTypeAndTimeSlot.ProgramType, &programTypeAndTimeSlot.TimeSlot, &programTypeAndTimeSlot.AdvertiserGroup, &programTypeAndTimeSlot.ChannelName, &programTypeAndTimeSlot.Region)

		advObj := make([]string, 0)
		advObj = append(advObj, programTypeAndTimeSlot.AdvertiserGroup)

		if index == 0 {
			regionObj := make(map[string][]string, 0)
			regionObj[strings.ToUpper(programTypeAndTimeSlot.Region.String)] = advObj
			channelNameObj := make(map[string]map[string][]string, 0)
			channelNameObj[programTypeAndTimeSlot.ChannelName.String] = regionObj
			progamTypeMap[strings.Title(programTypeAndTimeSlot.ProgramType.String)] = channelNameObj
			timeSlotMap[strings.ToUpper(programTypeAndTimeSlot.TimeSlot.String)] = channelNameObj
		} else {
			programType := progamTypeMap[strings.Title(programTypeAndTimeSlot.ProgramType.String)]
			timeSlot := timeSlotMap[strings.ToUpper(programTypeAndTimeSlot.TimeSlot.String)]
			if programType != nil {
				channelType := programType[programTypeAndTimeSlot.ChannelName.String]
				if channelType != nil {
					region := channelType[strings.ToUpper(programTypeAndTimeSlot.Region.String)]
					if region != nil {
						insertValue := true
						for _, values := range region {
							if values == advObj[0] {
								insertValue = false
							}
						}
						if insertValue == true {
							region = append(region, advObj...)
						}
					} else {
						region = advObj
					}
					channelType[strings.ToUpper(programTypeAndTimeSlot.Region.String)] = region
				} else {
					regionObj := make(map[string][]string, 0)
					regionObj[strings.ToUpper(programTypeAndTimeSlot.Region.String)] = advObj
					programType[programTypeAndTimeSlot.ChannelName.String] = regionObj
				}
			} else {
				regionObj := make(map[string][]string, 0)
				regionObj[strings.ToUpper(programTypeAndTimeSlot.Region.String)] = advObj
				channelNameObj := make(map[string]map[string][]string, 0)
				channelNameObj[programTypeAndTimeSlot.ChannelName.String] = regionObj
				progamTypeMap[strings.Title(programTypeAndTimeSlot.ProgramType.String)] = channelNameObj
			}
			if timeSlot != nil {
				channelType := timeSlot[programTypeAndTimeSlot.ChannelName.String]
				if channelType != nil {
					region := channelType[strings.ToUpper(programTypeAndTimeSlot.Region.String)]
					if region != nil {
						insertValue := true
						for _, values := range region {
							if values == advObj[0] {
								insertValue = false
							}
						}
						if insertValue == true {
							region = append(region, advObj...)
						}
					} else {
						region = advObj
					}
					channelType[strings.ToUpper(programTypeAndTimeSlot.Region.String)] = region
				} else {
					regionObj := make(map[string][]string, 0)
					regionObj[strings.ToUpper(programTypeAndTimeSlot.Region.String)] = advObj
					timeSlot[programTypeAndTimeSlot.ChannelName.String] = regionObj
				}
			} else {
				regionObj := make(map[string][]string, 0)
				regionObj[strings.ToUpper(programTypeAndTimeSlot.Region.String)] = advObj
				channelNameObj := make(map[string]map[string][]string, 0)
				channelNameObj[programTypeAndTimeSlot.ChannelName.String] = regionObj
				timeSlotMap[strings.ToUpper(programTypeAndTimeSlot.TimeSlot.String)] = channelNameObj
			}
		}
		index++
	}
	filters.ProgramType = progamTypeMap
	filters.TimeSlot = timeSlotMap
	return filters
}

func whereCondition(startDates []string, endDates []string) string {
	whereCon := ""
	for index, date := range startDates {
		condition := fmt.Sprintf("date between '%s' and '%s'", date, endDates[index])
		if len(whereCon) == 0 {
			whereCon += fmt.Sprintf(" where ( ( %s )", condition)
		} else {
			whereCon += fmt.Sprintf(" OR ( %s )", condition)
		}
	}
	if len(whereCon) != 0 {
		whereCon += " ) "
	}
	return whereCon
}

func getRegionObj(advObj []string, region string) schemas.Region {
	var regionObj schemas.Region
	regionObj.Label = region
	regionObj.Advertisers = advObj
	return regionObj
}

func getChannelObj(regions []schemas.Region, channel string) schemas.ChannelName {
	var channelObj schemas.ChannelName
	channelObj.Label = channel
	channelObj.Regions = regions
	return channelObj
}

func getProgramTypeObj(channels []schemas.ChannelName, programType string) schemas.ProgramTypeStruct {
	var programTypeObj schemas.ProgramTypeStruct
	programTypeObj.Label = programType
	programTypeObj.Channels = channels
	return programTypeObj
}
func getTimeSlotObj(channels []schemas.ChannelName, timeSlot string) schemas.TimeSlotStruct {
	var timeSlotObj schemas.TimeSlotStruct
	timeSlotObj.Label = timeSlot
	timeSlotObj.Channels = channels
	return timeSlotObj
}

func formatProgramTypeResponse(programTypeAndTimeSlot schemas.ProgramTypeAndTimeSlot, advObj []string) []schemas.ProgramTypeStruct {
	programTypeArr := make([]schemas.ProgramTypeStruct, 0)
	programTypePresent := false
	for k, program := range programTypeArr {
		if program.Label == programTypeAndTimeSlot.ProgramType.String {
			channelPresent := false
			for j, key := range program.Channels {
				if key.Label == programTypeAndTimeSlot.ChannelName.String {
					present := false
					for i, regionKey := range key.Regions {
						if regionKey.Label == programTypeAndTimeSlot.Region.String {
							regionKey.Advertisers = append(regionKey.Advertisers, programTypeAndTimeSlot.AdvertiserGroup)
							key.Regions[i] = key.Regions[len(key.Regions)-1]
							key.Regions[len(key.Regions)-1] = schemas.Region{}
							key.Regions = key.Regions[:len(key.Regions)-1]
							regionObjListLocal := make([]schemas.Region, 1)
							regionObjListLocal[0] = regionKey
							regionObjListLocal = append(regionObjListLocal, key.Regions...)
							key.Regions = regionObjListLocal
							present = true
						}
					}
					if present == false {
						regionObj := getRegionObj(advObj, programTypeAndTimeSlot.Region.String)
						regionObjListLocal := make([]schemas.Region, 1)
						regionObjListLocal[0] = regionObj
						key.Regions = append(key.Regions, regionObj)
					}
					program.Channels[j] = program.Channels[len(program.Channels)-1]
					program.Channels[len(program.Channels)-1] = schemas.ChannelName{}
					program.Channels = program.Channels[:len(program.Channels)-1]
					channelNameObjListLocal := make([]schemas.ChannelName, 1)
					channelNameObjListLocal[0] = key
					channelNameObjListLocal = append(channelNameObjListLocal, program.Channels...)
					program.Channels = channelNameObjListLocal
					channelPresent = true
				}
			}
			if channelPresent == false {
				regionObj := getRegionObj(advObj, programTypeAndTimeSlot.Region.String)
				regionObjListLocal := make([]schemas.Region, 1)
				regionObjListLocal[0] = regionObj
				channelNameObj := getChannelObj(regionObjListLocal, programTypeAndTimeSlot.ChannelName.String)
				channelNameObjList := make([]schemas.ChannelName, 0)
				channelNameObjList = append(channelNameObjList, channelNameObj)
				program.Channels = channelNameObjList
			}

			programTypeArr[k] = programTypeArr[len(programTypeArr)-1]
			programTypeArr[len(programTypeArr)-1] = schemas.ProgramTypeStruct{}
			programTypeArr = programTypeArr[:len(programTypeArr)-1]
			programTypeArrLocal := make([]schemas.ProgramTypeStruct, 1)
			programTypeArrLocal[0] = program
			programTypeArrLocal = append(programTypeArrLocal, programTypeArr...)
			programTypeArr = programTypeArrLocal
			programTypePresent = true
		}
	}
	if programTypePresent == false {
		regionObj := getRegionObj(advObj, programTypeAndTimeSlot.Region.String)
		regionObjListLocal := make([]schemas.Region, 1)
		regionObjListLocal[0] = regionObj
		channelNameObj := getChannelObj(regionObjListLocal, programTypeAndTimeSlot.ChannelName.String)
		channelNameObjList := make([]schemas.ChannelName, 0)
		channelNameObjList = append(channelNameObjList, channelNameObj)
		programType := getProgramTypeObj(channelNameObjList, programTypeAndTimeSlot.ProgramType.String)
		programTypeArr = append(programTypeArr, programType)
	}
	return programTypeArr
}

func formatTimeSlotResponse(programTypeAndTimeSlot schemas.ProgramTypeAndTimeSlot, advObj []string) []schemas.TimeSlotStruct {
	timeSlotPresent := false
	timeSlotArr := make([]schemas.TimeSlotStruct, 0)
	for k, timeSlot := range timeSlotArr {
		if timeSlot.Label == programTypeAndTimeSlot.TimeSlot.String {
			fmt.Println("in if timeslot")
			channelPresent := false
			for j, key := range timeSlot.Channels {
				if key.Label == programTypeAndTimeSlot.ChannelName.String {
					present := false
					for i, regionKey := range key.Regions {
						if regionKey.Label == programTypeAndTimeSlot.Region.String {
							regionKey.Advertisers = append(regionKey.Advertisers, programTypeAndTimeSlot.AdvertiserGroup)
							key.Regions[i] = key.Regions[len(key.Regions)-1]
							key.Regions[len(key.Regions)-1] = schemas.Region{}
							key.Regions = key.Regions[:len(key.Regions)-1]
							regionObjListLocal := make([]schemas.Region, 1)
							regionObjListLocal[0] = regionKey
							regionObjListLocal = append(regionObjListLocal, key.Regions...)
							key.Regions = regionObjListLocal
							present = true
						}
					}
					if present == false {
						regionObj := getRegionObj(advObj, programTypeAndTimeSlot.Region.String)
						regionObjListLocal := make([]schemas.Region, 1)
						regionObjListLocal[0] = regionObj
						key.Regions = append(key.Regions, regionObj)
					}
					timeSlot.Channels[j] = timeSlot.Channels[len(timeSlot.Channels)-1]
					timeSlot.Channels[len(timeSlot.Channels)-1] = schemas.ChannelName{}
					timeSlot.Channels = timeSlot.Channels[:len(timeSlot.Channels)-1]
					channelNameObjListLocal := make([]schemas.ChannelName, 1)
					channelNameObjListLocal[0] = key
					channelNameObjListLocal = append(channelNameObjListLocal, timeSlot.Channels...)
					timeSlot.Channels = channelNameObjListLocal
					channelPresent = true
				}
			}
			if channelPresent == false {
				regionObj := getRegionObj(advObj, programTypeAndTimeSlot.Region.String)
				regionObjListLocal := make([]schemas.Region, 1)
				regionObjListLocal[0] = regionObj
				channelNameObj := getChannelObj(regionObjListLocal, programTypeAndTimeSlot.ChannelName.String)
				channelNameObjList := make([]schemas.ChannelName, 0)
				channelNameObjList = append(channelNameObjList, channelNameObj)
				timeSlot.Channels = channelNameObjList
			}

			timeSlotArr[k] = timeSlotArr[len(timeSlotArr)-1]
			timeSlotArr[len(timeSlotArr)-1] = schemas.TimeSlotStruct{}
			timeSlotArr = timeSlotArr[:len(timeSlotArr)-1]
			timeSlotArrLocal := make([]schemas.TimeSlotStruct, 1)
			timeSlotArrLocal[0] = timeSlot
			timeSlotArrLocal = append(timeSlotArrLocal, timeSlotArr...)
			timeSlotArr = timeSlotArrLocal
			timeSlotPresent = true
		}
	}

	if timeSlotPresent == false {
		regionObj := getRegionObj(advObj, programTypeAndTimeSlot.Region.String)
		regionObjListLocal := make([]schemas.Region, 1)
		regionObjListLocal[0] = regionObj
		channelNameObj := getChannelObj(regionObjListLocal, programTypeAndTimeSlot.ChannelName.String)
		channelNameObjList := make([]schemas.ChannelName, 0)
		channelNameObjList = append(channelNameObjList, channelNameObj)
		timeSlot := getTimeSlotObj(channelNameObjList, programTypeAndTimeSlot.TimeSlot.String)
		timeSlotArr = append(timeSlotArr, timeSlot)
	}
	return timeSlotArr
}

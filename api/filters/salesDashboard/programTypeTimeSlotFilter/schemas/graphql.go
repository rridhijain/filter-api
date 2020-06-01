package schemas

import "github.com/graphql-go/graphql"

//Dashboard Fields Params
var DashboardFields = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Dashboard_Fields",
		Fields: graphql.Fields{
			"program_type": &graphql.Field{
				Type: graphql.NewList(programTypeFields),
			},
			"time_slot": &graphql.Field{
				Type: graphql.NewList(timeSlotFields),
			},
		},
	},
)

var programTypeFields = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Program_Type_Fields",
		Fields: graphql.Fields{
			"label": &graphql.Field{
				Type: graphql.String,
			},
			"channels": &graphql.Field{
				Type: graphql.NewList(channelFields),
			},
		},
	},
)

var timeSlotFields = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Time_Slot_Fields",
		Fields: graphql.Fields{
			"label": &graphql.Field{
				Type: graphql.String,
			},
			"channels": &graphql.Field{
				Type: graphql.NewList(channelFields),
			},
		},
	},
)

var channelFields = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Channel_Fields",
		Fields: graphql.Fields{
			"label": &graphql.Field{
				Type: graphql.String,
			},
			"regions": &graphql.Field{
				Type: graphql.NewList(regionFields),
			},
		},
	},
)

var regionFields = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Region_Fields",
		Fields: graphql.Fields{
			"label": &graphql.Field{
				Type: graphql.String,
			},
			"advertisers": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

var PeriodType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "Dates",
		Fields: graphql.InputObjectConfigFieldMap{
			"start_date": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"end_date": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

package schemas

import "github.com/graphql-go/graphql"

//Dashboard Fields Params
var DashboardFields = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Dashboard_Fields",
		Fields: graphql.Fields{
			"program_type": &graphql.Field{
				Type: graphql.String,
			},
			"time_slot": &graphql.Field{
				Type: graphql.String,
			},
			"date": &graphql.Field{
				Type: graphql.String,
			},
			"advertiser_group": &graphql.Field{
				Type: graphql.String,
			},
			"channel_name": &graphql.Field{
				Type: graphql.String,
			},
			"region": &graphql.Field{
				Type: graphql.String,
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

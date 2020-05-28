package schemas

import "database/sql"

type ProgramTypeAndTimeSlot struct {
	ProgramType     sql.NullString `json:"program_type"`
	TimeSlot        sql.NullString `json:"time_slot"`
	AdvertiserGroup string         `json:"advertiser_group"`
	ChannelName     sql.NullString `json:"channel_name"`
	Region          sql.NullString `json:"region"`
}

type ProgramTypeAndTimeSlotUpdated struct {
	ProgramType []ProgramTypeStruct `json:"program_type"`
	TimeSlot    []TimeSlotStruct    `json:"time_slot"`
}

type ProgramTypeStruct struct {
	Label    string        `json:"program_type"`
	Channels []ChannelName `json:"channels"`
}

type TimeSlotStruct struct {
	Label    string        `json:"time_slot"`
	Channels []ChannelName `json:"channels"`
}

type ChannelName struct {
	Label   string   `json:"label"`
	Regions []Region `json:"regions"`
}

type Region struct {
	Label       string   `json:"label"`
	Advertisers []string `json:"advertisers"`
}

type Period struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

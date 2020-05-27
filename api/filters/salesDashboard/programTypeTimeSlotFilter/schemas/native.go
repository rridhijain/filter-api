package schemas

type ProgramTypeAndTimeSlot struct {
	ProgramType     string `json:"program_type"`
	TimeSlot        string `json:"time_slot"`
	Date            string `json:"date"`
	AdvertiserGroup string `json:"advertiser_group"`
	ChannelName     string `json:"channel_name"`
	Region          string `json:"region"`
}

type ProgramTypeAndTimeSlotUpdated struct {
	ProgramType []ProgramTypeStruct `json:"program_type"`
	TimeSlot    []TimeSlotStruct    `json:"time_slot"`
}

type ProgramTypeStruct struct {
	Label  string        `json:"label"`
	Values []ChannelName `json:"values"`
}

type TimeSlotStruct struct {
	Label  string        `json:"label"`
	Values []ChannelName `json:"values"`
}

type ChannelName struct {
	Label  string   `json:"label"`
	Values []Region `json:"values"`
}

type Region struct {
	Label  string   `json:"label"`
	Values []string `json:"values"`
}

type Period struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

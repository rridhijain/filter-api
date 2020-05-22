package schemas

type ProgramTypeAndTimeSlot struct {
	ProgramType string `json:"program_type"`
	TimeSlot    string `json:"time_slot"`
	Date        string `json:"date"`
	Advertiser  string `json:"advertiser"`
	Channel     string `json:"channel"`
	Region      string `json:"region"`
}

type Period struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

package structures

type Data struct {
	ShortName    string  `json:"SHORT_NAME"`
	CategoryName string  `json:"CATEGORY_NAME"`
	Events       []Event `json:"EVENTS"`
}

type Event struct {
	TimeStamp  int      `json:"START_TIME"`
	HomeName   string   `json:"HOME_PARTICIPANT_NAME_ONE"`
	HomeImages []string `json:"HOME_IMAGES"`
	AwayName   string   `json:"AWAY_PARTICIPANT_NAME_ONE"`
	AwayImages []string `json:"AWAY_IMAGES"`
	Winner     int      `json:"WINNER"`
}

package structures

type HockeyMatch struct {
	Date   string  `json:"date"`
	Timer  string  `json:"timer"`
	Status *Status `json:"status"`
	League *League `json:"league"`
	Teams  struct {
		Home *Team `json:"home"`
		Away *Team `json:"away"`
	}
	Scores *Score `json:"scores"`
}

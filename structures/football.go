package structures

type FootballMatch struct {
	Fixture struct {
		Status struct {
			MinuteGame int `json:"elapsed"`
		} `json:"fixture"`
		Date string `json:"date"`
	} `json:"fixture"`
	League *League `json:"league"`
	Teams  struct {
		Home *Team `json:"home"`
		Away *Team `json:"away"`
	} `json:"teams"`
	Score  *Score `json:"goals"`
	Events []Even `json:"events"`
}
type Even struct {
	Type string `json:"type"`
	Time struct {
		Minute int `json:"elapsed"`
	} `json:"time"`
	Team struct {
		Name string `json:"name"`
	} `json:"team"`
	Player struct {
		Name string `json:"name"`
	} `json:"player"`
}

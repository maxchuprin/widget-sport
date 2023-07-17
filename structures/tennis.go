package structures

type TennisMatch struct {
	Date      string       `json:"start_at"`
	Stage     Stage        `json:"round_info,omitempty"`
	League    *League      `json:"league"`
	Home      *Team        `json:"home_team"`
	Away      *Team        `json:"away_team"`
	HomeScore *TennisScore `json:"home_score"`
	AwayScore *TennisScore `json:"away_score"`
}

type TennisScore struct {
	Current  int    `json:"current"`
	Display  int    `json:"display"`
	Period_1 int    `json:"period_1"`
	Period_2 int    `json:"period_2"`
	Period_3 int    `json:"period_3"`
	Period_4 int    `json:"period_4"`
	Period_5 int    `json:"period_5"`
	Point    string `json:"point"`
}

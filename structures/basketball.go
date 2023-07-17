package structures

type BasketballMatch struct {
	Date   string  `json:"date"`
	Status *Status `json:"status"`
	League *League `json:"league"`
	Teams  struct {
		Home *Team `json:"home"`
		Away *Team `json:"away"`
	} `json:"teams"`
	Scores struct {
		Home *Quarter `json:"home"`
		Away *Quarter `json:"away"`
	} `json:"scores"`
}
type Quarter struct {
	Quarter1 int `json:"quarter_1"`
	Quarter2 int `json:"quarter_2"`
	Quarter3 int `json:"quarter_3"`
	Quarter4 int `json:"quarter_4"`
	Overtime int `json:"over_time"`
	Total    int `json:"total"`
}

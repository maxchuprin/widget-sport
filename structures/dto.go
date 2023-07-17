package structures

type League struct {
	Name  string `json:"name"`
	Logo  string `json:"logo"`
	Round string `json:"round"`
}

type Team struct {
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type Status struct {
	Period string `json:"short"`
	Timer  string `json:"timer"`
}

type Score struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type Stage struct {
	Name string `json:"name,omitempty"`
}

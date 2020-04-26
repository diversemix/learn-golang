package domain

type Farm struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Animal struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Location string `json:"location"`
	FarmID   string `json:"farmID"`
}

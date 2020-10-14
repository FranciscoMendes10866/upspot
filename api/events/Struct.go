package events

type Event struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	City          string `json:"city"`
	Img           string `json:"img"`
	Date          string `json:"date"`
	Type          string `json:"type"`
	Max           int    `json:"max"`
	AttendsNumber int    `json:"attends_number"`
	HostName      string `json:"hostName"`
	HostURL       string `json:"hostURL"`
	Link          string `json:"link"`
	UserID        string `json:"userId"`
}

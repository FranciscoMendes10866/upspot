package events

type Event struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Img      string `json:"img"`
	Date     string `json:"date"`
	Type     string `json:"type"`
	Max      int    `json:"max"`
	HostName string `json:"hostName"`
	HostURL  string `json:"hostURL"`
	Link     string `json:"link"`
	UserID   string `json:"userId"`
}

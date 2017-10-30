package common

// Data retrieved from the JSON API
type API struct {
	Requesting  bool   `json:"requesting"`
	Listeners   int    `json:"listeners"`
	CurrentTime int64  `json:"current"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	NowPlaying  string `json:"np"`
	Thread      string `json:"thread"`
	DJ          struct {
		ID    int    `json:"id"`
		Image string `json:"djimage"`
		Name  string `json:"djname"`
	} `json:"dj"`
	Queue      []Song `json:"queue"`
	LastPlayed []Song `json:"lp"`
}

// Song-related data
type Song struct {
	Timestamp int64  `json:"timestamp"`
	Meta      string `json:"meta"`
}

package common

// Data retrieved from the JSON API
type API struct {
	Requesting bool   `json:"requesting"`
	Listeners  uint   `json:"listeners"`
	StartTime  uint   `json:"start_time"`
	EndTime    uint   `json:"end_time"`
	NowPlaying string `json:"np"`
	Thread     string `json:"thread"`
	DJ         struct {
		ID      uint   `json:"id"`
		DJImage uint   `json:"djimage"`
		DJName  string `json:"djname"`
	} `json:"dj"`
	Queue      []Song `json:"queue"`
	LastPlayed []Song `json:"lp"`
}

// Song-related data
type Song struct {
	Timestamp int64  `json:"timestamp"`
	Meta      string `json:"meta"`
}

package common

import (
	"math"
	"time"
)

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

// Song returned by Elastic Search
type SearchSong struct {
	ID            int    `json:"id"`
	Requests      int    `json:"requests"`
	LastRequested int64  `json:"lastrequested"`
	LastPlayed    int64  `json:"lastplayed"`
	RequestDelay  int64  `json:"-"`
	Artist        string `json:"artist"`
	Title         string `json:"title"`
}

// Calculate request delay and requestability for a song
func (s *SearchSong) CalculateRequestDelay() {
	if s.Requests > 30 {
		s.Requests = 30
	}

	var dur float64
	if s.Requests >= 0 && s.Requests <= 7 {
		dur = -11057*math.Pow(float64(s.Requests), 2) +
			172954*float64(s.Requests) + 81720
	} else {
		dur = 599955*math.Exp(0.0372*float64(s.Requests)) + 0.5
	}
	s.RequestDelay = int64(dur)
}

// Return, if song can be requested
func (s *SearchSong) CanRequest() bool {
	s.CalculateRequestDelay()
	return time.Now().Unix()-s.LastPlayed > s.RequestDelay
}

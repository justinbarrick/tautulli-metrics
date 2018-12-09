package tautulli

import (
	"encoding/json"
	"time"
)

type StreamMetadata struct {
	Started    time.Time
	Stopped    time.Time
	Transcoded bool
	Title      string
	User       string
	UserId     int
	MediaType  string
	Player     string
	Platform   string
}

func (s *StreamMetadata) UnmarshalJSON(data []byte) error {
	type streamMetadata struct {
		Started           int64  `json:"started"`
		Stopped           int64  `json:"stopped"`
		TranscodeDecision string `json:"transcode_decision"`
		FullTitle         string `json:"full_title"`
		User              string `json:"user"`
		UserId            int    `json:"user_id"`
		MediaType         string `json:"media_type"`
		Player            string `json:"player"`
		Platform          string `json:"platform"`
	}

	stream := streamMetadata{}

	if err := json.Unmarshal(data, &stream); err != nil {
		return err
	}

	s.Started = time.Unix(stream.Started, 0)
	s.Stopped = time.Unix(stream.Stopped, 0)
	s.Transcoded = stream.TranscodeDecision == "transcode"
	s.Title = stream.FullTitle
	s.User = stream.User
	s.UserId = stream.UserId
	s.MediaType = stream.MediaType
	s.Player = stream.Player
	s.Platform = stream.Platform
	return nil
}

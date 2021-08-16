package infrastructure

import "gopkg.in/segmentio/analytics-go.v3"

type Segment struct {
	Client analytics.Client
}

func NewLogger(client analytics.Client) *Segment {
	return &Segment{
		Client: client,
	}
}

func (s *Segment) Log(user string, data string) error {
	s.Client.Enqueue(analytics.Track{
		UserId: user,
		Event:  data,
	})

	return nil
}

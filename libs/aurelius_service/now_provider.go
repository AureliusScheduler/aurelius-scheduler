package aurelius_service

import "time"

type NowProvider interface {
	Now() time.Time
}

type RealNowProvider struct{}

func (RealNowProvider) Now() time.Time {
	return time.Now()
}

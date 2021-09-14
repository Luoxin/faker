package faker

import (
	"math/rand"
	"time"
)

func (f *Faker) Time() time.Time {
	return time.Unix(rand.Int63n(time.Now().Unix()-94608000)+94608000, 0)
}

func (f *Faker) TimeAfter(time2 time.Time) time.Time {
	return time2.Add(time.Duration(rand.Uint32()) * f.Choice([]time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		time.Minute,
		time.Hour,
	}).(time.Duration))
}

func (f *Faker) Timestamp() uint32 {
	return uint32(f.Time().Unix())
}

func (f *Faker) TimestampAfter(time2 time.Time) uint32 {
	return uint32(f.TimeAfter(time2).Unix())
}

func (f *Faker) TimestampAfterNow() uint32 {
	return f.TimestampAfter(time.Now())
}

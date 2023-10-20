package countdown

import "time"

type Sleeper interface {
	Sleep()
}

// type DefaultSleeper struct{}
//
// func (d *DefaultSleeper) Sleep() {
// 	time.Sleep(1 * time.Second)
// }

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func NewConfigurableSleeper(
	duration time.Duration,
	sleep func(time.Duration),
) *ConfigurableSleeper {
	return &ConfigurableSleeper{duration, sleep}
}

func (c ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

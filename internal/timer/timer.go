package timer

import "time"

// Timer is a simple timer that can be used to measure elapsed time and
// check if a certain duration has passed. It can be set to loop, meaning
// it will restart automatically after finishing.
type Timer struct {
	start    time.Time
	duration time.Duration
	loop     bool
}

// New creates a new Timer with the specified duration and loop setting.
func New(duration time.Duration, loop bool) *Timer {
	return &Timer{
		duration: duration,
		loop:     loop,
	}
}

// Start initializes the timer by setting the start time to the current time.
func (t *Timer) Start() *Timer {
	t.start = time.Now()
	return t
}

// IsStarted checks if the timer has been started.
func (t *Timer) IsStarted() bool {
	return !t.start.IsZero()
}

// IsFinished checks if the timer has finished its duration.
func (t *Timer) IsFinished() bool {
	return time.Since(t.start) >= t.duration
}

// SetDuration sets the duration of the timer.
func (t *Timer) SetDuration(duration time.Duration) {
	t.duration = duration
}

// Reset restarts the timer.
func (t *Timer) Reset() {
	t.start = time.Now()
}

// RestartLoop restarts the timer if it has finished and is set to loop.
// It returns true if the timer was restarted.
func (t *Timer) RestartLoop() bool {
	if !t.IsStarted() || !t.IsFinished() || !t.loop {
		return false
	}

	t.Reset()

	return true
}

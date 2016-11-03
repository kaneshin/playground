package trick

import "time"

func BoolTimer(d time.Duration) func() bool {
	var enable bool

	t := time.AfterFunc(d, func() {
		enable = false
	})

	return func() bool {
		if enable {
			return true
		}
		defer func() { enable = true }()

		t.Reset(d)
		return enable
	}
}

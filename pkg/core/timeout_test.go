package core

import (
	"testing"
	"time"
)

func TestNewTimeout(t *testing.T) {
	for i := 0; i < 100; i++ {
		timeout := NewTimeout()
		if timeout > Timeout(300*time.Millisecond) {
			t.Fail()
		}
		if timeout < Timeout(150*time.Millisecond) {
			t.Fail()
		}
	}
}

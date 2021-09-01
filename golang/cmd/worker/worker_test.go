package main

import (
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHandleExit(t *testing.T) {

	w, err := New()
	assert.NoError(t, err)

	// need something to read exitChan
	go func() {
		code := <-w.exitChan
		assert.Equal(t, 0, code)
	}()

	// must complete this section in 3 seconds maximum
	timeoutTest(t, func() {
		count := 0
		for i := 0; i < 10; i++ {
			// 10 goroutunes
			w.doAsync(func() error {
				<-time.After(time.Second * 1)
				count++
				return nil
			})
		}

		w.handleExit(syscall.SIGHUP)
		assert.Equal(t, 10, count)
	})
}

// helper to prevent tests from hanging indefinitely
func timeoutTest(t *testing.T, fn func()) {
	done := make(chan bool)
	go func() {
		fn()
		done <- true
	}()

	select {
	case <-time.After(3 * time.Second):
		t.Fatal("Test didn't finish in time")
	case <-done:
	}
}

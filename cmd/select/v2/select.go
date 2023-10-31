package selectLesson

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, error error) {
	return RacerWithTimeout(a, b, 10*time.Second)
}

func RacerWithTimeout(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	// as var := <-channel is a blocking call the case will be false until we get a value in the channel
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(link string) chan struct{} {
	// chan struct{} is the smallest data type available from memory perspective
	ch := make(chan struct{})
	go func() {
		http.Get(link)
		close(ch)
	}()
	return ch
}

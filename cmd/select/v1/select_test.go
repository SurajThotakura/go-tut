package selectLesson

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := serverWithDelay(20 * time.Millisecond)
	defer slowServer.Close()
	fastServer := serverWithDelay(0 * time.Millisecond)
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}

func serverWithDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader((http.StatusOK))
	}))
}

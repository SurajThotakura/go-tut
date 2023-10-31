package selectLesson

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare ping time and get the fastest", func(t *testing.T) {
		slowServer := serverWithDelay(20 * time.Millisecond)
		defer slowServer.Close()
		fastServer := serverWithDelay(0 * time.Millisecond)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("want %q got %q", want, got)
		}
	})

	t.Run("when a server is taking longer than 10s return error", func(t *testing.T) {
		slowServer := serverWithDelay(10200 * time.Millisecond)
		defer slowServer.Close()
		fastServer := serverWithDelay(10100 * time.Millisecond)
		defer fastServer.Close()

		_, err := Racer(slowServer.URL, fastServer.URL)

		if err == nil {
			t.Error("Expected error got none")
		}
	})
}

func serverWithDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader((http.StatusOK))
	}))
}

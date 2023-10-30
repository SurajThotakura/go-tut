package concurrency

import (
	"net/http"
	"reflect"
	"testing"
)

func mockWebsiteChecker(link string) bool {
	response, err := http.Head(link)
	if err != nil {
		return false
	}
	return response.StatusCode == http.StatusOK
}

func TestWebsitesChecker(t *testing.T) {
	websitesList := []string{
		"http://www.google.com",
		"http://www.surajthotakura.com",
		"http://wrongwebsitefalsename.inon",
	}

	want := map[string]bool{
		"http://www.google.com":             true,
		"http://www.surajthotakura.com":     true,
		"http://wrongwebsitefalsename.inon": false,
	}

	got := WebsitesChecker(mockWebsiteChecker, websitesList)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkWebsitesChecker(b *testing.B) {
	links := make([]string, 1000)
	for i := 0; i < len(links); i++ {
		links[i] = "mock link"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WebsitesChecker(mockWebsiteChecker, links)
	}
}

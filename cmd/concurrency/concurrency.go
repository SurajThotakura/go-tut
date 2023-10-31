package concurrency

type WebsiteCheckerFunc func(string) bool

type result struct {
	string
	bool
}

func WebsitesChecker(wc WebsiteCheckerFunc, websitesList []string) map[string]bool {
	results := map[string]bool{}
	resultsChannels := make(chan result)

	captureResult := func(l string) {
		// Send statement
		resultsChannels <- result{l, wc(l)}
	}

	for _, link := range websitesList {
		go captureResult(link)
	}

	for range websitesList {
		// Receive statement
		r := <-resultsChannels
		results[r.string] = r.bool
	}

	return results
}

package concurrency

type WebsiteCheckerFunc func(string) bool

type result struct {
	string
	bool
}

func WebsitesChecker(wc WebsiteCheckerFunc, websitesList []string) map[string]bool {
	results := map[string]bool{}
	resultsChannels := make(chan result)

	for _, link := range websitesList {
		go func(l string) {
			resultsChannels <- result{l, wc(l)}
		}(link)
	}

	for range websitesList {
		r := <-resultsChannels
		results[r.string] = r.bool
	}

	return results
}

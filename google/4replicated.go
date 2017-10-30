package google

import (
	"errors"
	"time"
)

func SearchReplicated(query string, timeout time.Duration) ([]Result, error) {
	// TODO same as you did in timeout, but this time use the replicated*FakeSearchFunctions
	// TODO also go complete the FirstFetchAmong function in searchUtils
	timer := time.After(timeout)
	c := make(chan Result, 3)
	go func() { c <- replicatedWebFakeSearch(query) }()
	go func() { c <- replicatedImageFakeSearch(query) }()
	go func() { c <- replicatedVideoFakeSearch(query) }()

	var results []Result
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timer:
			return results, errors.New("timed out")
		}
	}
	return results, nil
}

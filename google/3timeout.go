package google

import (
	"errors"
	"time"
)

func SearchTimeout(query string, timeout time.Duration) ([]Result, error) {
	// should return the same result in the same way as the SearchParallel function, but only if a timeout doesn't occur
	// TODO use the timer and the select statement inside the result array to either receive the search results from
	// the channel you created before, or return with what you have when the timer expires
	// HINT you can think of the timeout as a message on a whatever channel, also return a proper error string
	timer := time.After(timeout)
	c := make(chan Result, 3)

	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

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
package google

func SearchParallel(query string) ([]Result, error) {
	// should return the same result as the SearchSerial function, but the search should be parallelized
	// TODO create a channel of type Result and use Goroutines to parallelize the work you did in SearchSerial
	// HINT send each search result on the channel and receive those results inside the array of results
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	return []Result{<-c, <-c, <-c}, nil
}
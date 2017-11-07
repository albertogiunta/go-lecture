package google

import (
	"fmt"
	"time"
)

// function that will be passed around to be executed in all the different ways (serial, parallel etc.)
type SearchFunc func(query string) Result

// the resulting data type of a Search
type Result struct {
	Title, URL string
}

type Response struct {
	Results []Result
	Elapsed time.Duration
}

// function that will be passed around to be executed in all the different ways (serial, parallel, with timeout, distributed)
func FakeSearch(kind, title, url string) SearchFunc {
	return func(query string) Result {
		time.Sleep(time.Duration(100 * time.Millisecond)) // artificial delay for instructional reasons
		return Result{
			Title: fmt.Sprintf("Query: %q | Source: %s -> Result: %s", query, kind, title),
			URL:   url,
		}
	}
}

func FirstFetchAmong(replicas ...SearchFunc) SearchFunc {
	return func(query string) Result {
		// TODO create a buffered channel of type Result and as long as the number of replicas that were passed in input
		// TODO create an inner function that given a index runs the replicated function at that index of replicas
		// TODO launch this function in a new goroutines for every index of replicas
		// TODO return the very first value received on the channel
		c := make(chan Result, len(replicas))
		searchReplica := func(i int) {
			c <- replicas[i](query)
		}
		for i := range replicas {
			go searchReplica(i)
		}
		return <-c
	}
}

// String returns the result's title, followed by a newline.
func (r Result) String() string { return r.Title + "\n" }

// serial & parallel
var (
	Web   = FakeSearch("web", "Distributed Systems are the best!", "http://apice.unibo.it/xwiki/bin/view/Courses/Sd1718")
	Image = FakeSearch("image", "But what's really a distributed system?", "https://www.unibo.it/uniboweb/utils/UserImage.aspx?IdAnagrafica=341468&IdFoto=23233366")
	Video = FakeSearch("video", "Cool video about distributed systems", "https://www.youtube.com/watch?v=dQw4w9WgXcQ")
)

// distributed
var (
	replicatedWebFakeSearch   = FirstFetchAmong(Web1, Web2)
	replicatedImageFakeSearch = FirstFetchAmong(Image1, Image2)
	replicatedVideoFakeSearch = FirstFetchAmong(Video1, Video2)

	Web1   = FakeSearch("web1", "Distributed Systems are the best!", "http://apice.unibo.it/xwiki/bin/view/Courses/Sd1718")
	Web2   = FakeSearch("web2", "Distributed Systems are the best!", "http://apice.unibo.it/xwiki/bin/view/Courses/Sd1718")
	Image1 = FakeSearch("image2", "But what's really a distributed system?", "https://www.unibo.it/uniboweb/utils/UserImage.aspx?IdAnagrafica=341468&IdFoto=23233366")
	Image2 = FakeSearch("image2", "But what's really a distributed system?", "https://www.unibo.it/uniboweb/utils/UserImage.aspx?IdAnagrafica=341468&IdFoto=23233366")
	Video1 = FakeSearch("video1", "Cool video about distributed systems", "https://www.youtube.com/watch?v=dQw4w9WgXcQ")
	Video2 = FakeSearch("video2", "Cool video about distributed systems", "https://www.youtube.com/watch?v=dQw4w9WgXcQ")
)
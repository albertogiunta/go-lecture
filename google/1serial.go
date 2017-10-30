package google

func SearchSerial(query string) ([]Result, error) {
	// TODO should return an array of results made of every possible result from Web, Image, Video for a specific query
	results := []Result{
		Web(query),
		Image(query),
		Video(query),
	}
	return results, nil
}

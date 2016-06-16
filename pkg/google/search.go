package google

type Result struct {
	Title, URL string
}

func Search(query string) ([]Result, error) {
	results := []Result{
		Web(query),
		Image(query),
		Video(query),
	}
	return results, nil
}
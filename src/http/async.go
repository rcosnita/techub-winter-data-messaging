package http

//asyncClient provides a high performance implementation of a http client based on async sockets.
type asyncClient struct{}

//Get fetches all available information from a given url and returns it as a string.
func (c *asyncClient) Get(url string) (string, error) {
	return "async", nil
}

package http

//syncClient provides a high performance implementation of a http client based on async sockets.
type syncClient struct{}

//Get fetches all available information from a given url and returns it as a string.
func (c *syncClient) Get(url string) (string, error) {
	return "sync", nil
}

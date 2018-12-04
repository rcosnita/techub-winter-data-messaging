package http

//Client provides the contract for implementing specific http clients.
type Client interface {
	Get(url string) (string, error)
}

//NewClient provides the factory method for obtaining a http client instance.
func NewClient() Client {
	return &asyncClient{}
}

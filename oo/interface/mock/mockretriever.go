package mock

type Retriever struct {
	Contents string
}

// Implement, provider Get() method
func (r Retriever) Get(url string) string {
	return r.Contents
}

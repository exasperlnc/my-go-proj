package maps

type Dictionary map[string]string

// Search takes a dictionary and a search key and returns the value associated with that key.
func (d Dictionary) Search(word string) string {
	return d[word]
}
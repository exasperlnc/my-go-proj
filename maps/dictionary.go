package maps


// Search takes a dictionary and a search key and returns the value associated with that key.
func Search (dictionary map[string]string, word string) string {
	return dictionary[word]
}
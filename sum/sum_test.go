package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("array of 5 numbers", func(t *testing.T) {
			
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertCorrectMessage(t, got, want)
	})
}


func assertCorrectMessage(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
package laststoneweight

import "testing"

func TestLastStoneWeight(t *testing.T) {
	var stones = []int{2, 7, 4, 1, 8, 1}

	if actualValue := lastStoneWeight(stones); actualValue != 1 {
		t.Errorf("Got %v expected 1", actualValue)
	}
	stones = []int{}
	if actualValue := lastStoneWeight(stones); actualValue != 0 {
		t.Errorf("Got %v expected 0", actualValue)
	}
}

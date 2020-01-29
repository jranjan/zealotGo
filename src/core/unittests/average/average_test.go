package average_test

import (
    "testing"
    "patterns/unittests/average"
)

var averageTests = []struct {
  name string
  nums []int // input
  expected int // expected result
}{
  { "TestCase 1", []int{2,}, 2},
  { "TestCase 1", []int{2, 4, 21,}, 9},
}

func TestAverage(t *testing.T) {
  for _, tt := range averageTests {
    actual := average.ComputeAverage(tt.nums)
    if actual != tt.expected {
      t.Errorf("Average(%d): expected %d, actual %d", tt.nums, tt.expected, actual)
    }
  }
}
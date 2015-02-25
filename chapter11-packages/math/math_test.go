package math

import "testing"
import "errors"

type testpair struct {
  values []float64
  average float64
}

var tests = []testpair {
  { []float64{1, 2}, 1.5 },
  { []float64{1, 1, 1, 1, 1, 1}, 1 },
  { []float64{1, -1}, 0 },
}

func TestAverage(t *testing.T) {
  var v float64
  v = Average([]float64{1, 2})
  if v != 1.5 {
    t.Error("Expected 1.5, got ", v)
  }
}

func TestAverage2(t *testing.T) {
  for _, pair := range tests {
    v := Average(pair.values)
    if v != pair.average {
      t.Error("For", pair.values, "expected", pair.average, "got", v)
    }
  }
}

func TestAverageWithEmptySlice(t *testing.T) {
  v := Average([]float64{})
  if v != 0.0 {
    t.Error("Average of an empty list should be 0.0, got", v)
  }
}

type testpair_min_max struct {
  values []float64
  expected float64
  err error
}

var minim_inputs = []testpair_min_max {
  { []float64 {1, 2, 3, 4, 5} , 1, nil },
  { []float64 {5, 4, 3, 2, 1}, 1, nil },
  { []float64 {-1, -2, -3, 10}, -3, nil },
  { []float64 {1}, 1, nil },
  { []float64 {}, 0, errors.New("InvalidArgumentException") },
}

func TestMin(t *testing.T) {
  for _, pair := range minim_inputs {
      v, err := Min(pair.values)
      if v != pair.expected && err != pair.err {
        t.Error("The minimum of ", pair.values, "should be", pair.expected)
      }
  }
}

var maximum_inputs = []testpair_min_max {
  { []float64 {1, 2, 3, 4, 5} , 5, nil },
  { []float64 {5, 4, 3, 2, 1}, 5, nil },
  { []float64 {-1, -2, -3, 10}, 10, nil },
  { []float64 {0}, 0, nil },
  { []float64 {}, 0, errors.New("InvalidArgumentException") },
}

func TestMax(t *testing.T) {
  for _, pair := range maximum_inputs {
      v, err := Max(pair.values)
      if v != pair.expected && err != pair.err {
        t.Error("The maximum of ", pair.values, "should be", pair.expected)
      }
  }
}
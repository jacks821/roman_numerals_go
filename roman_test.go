package main

import "testing"


type testpair struct {
  value int
  roman_value string
}

var tests = []testpair {
  { 1, "I"},
  { 5, "V"},
  { 10, "X"},
  { 50, "L"},
  { 100, "C" },
  { 500, "D"},
  { 1000, "M" },
  { 2, "II"},
  { 6, "VI"},
  { 11, "XI"},
  { 3999, "MMMCMXCIX"},
  { 40, "XL" },
  { 445, "CDXLV"},
}

func TestRoman(t *testing.T) {
  for _, pair := range tests {
    v := Roman(pair.value)
      if v != pair.roman_value {
        t.Error(
            "For", pair.value,
            "expected", pair.roman_value,
            "got", v,
          )
      }
  }
}

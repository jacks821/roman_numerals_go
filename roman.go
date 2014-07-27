package main
import "strings"


var romans = map[int]map[int]string{
  0 : map[int]string{
    1000 : "M",
  },
  1 : map[int]string{
     900 : "CM",
  },
  2 : map[int]string{
     500 : "D",
  },
  3 : map[int]string{
    400 : "CD",
  },
  4 : map[int]string{
    100 : "C",
  },
  5 : map[int]string{
    90 : "XC",
  },
  6 : map[int]string{
     50 : "L",
  },
  7 : map[int]string{
     10 : "X",
  },
  8 : map[int]string{
     9 : "IX",
  },
  9 : map[int]string{
     5 : "V",
  },
  10 : map[int]string{
    4 : "IV",
  },
  11 : map[int]string{
     1 : "I",
  },
}


func Roman(arabic int) string {
  var result string
  var division_result string
      for n:= 0; n <= len(romans); n++ {
          for key, value := range romans[n] {
          division_result = strings.Repeat(value, arabic/key)
          result = result + division_result
          arabic = arabic % key
        }

      }
    return result
}

// +build !example

package raindrops
import "strconv"
const TestVersion = 1

func Convert(x int) (s string) {
  if x%3 == 0 {
    s += "Pling"
  }
  if x%5 == 0 {
    s += "Plang"
  }
  if x%7 == 0{
   s += "Plong"
  }
  if len(s) == 0 {
		s = strconv.Itoa(x)
	}
return

}


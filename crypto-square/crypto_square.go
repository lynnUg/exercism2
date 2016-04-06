// +build !example

package cryptosquare
import "math"
import "regexp"
import "strings"
const TestVersion = 1



func Encode(sentence string) (output string) {
  r := regexp.MustCompile("[^a-zA-Z1-9]")
  temp:=r.ReplaceAllString(sentence, "")
  numCols:=int(math.Ceil(math.Sqrt(float64(len(temp)))))
  col:=make([]string,numCols)
  z:=0
  for _,j:= range temp{
    col[z]+=strings.ToLower(string(j))
    z++
    if z>=numCols{
      z=0
    }
  }
  return strings.Join(col," ")
}


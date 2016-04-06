package etl
import "strings"

func Transform(n map[int][]string) (out map[string]int) {

  out=make(map[string]int)
  for key,value := range n {
    for _,i:= range value{
      out[strings.ToLower(i)]=key
    }
  }
  return out
}

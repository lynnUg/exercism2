package slice
import "strings"

func All(n int , s string) (r []string ){
  for i:=0 ;n<=len(s); i++ {
    r=append(r,s[i,n])
    n++
  }
  return

}

func Frist(n int , s string) string {
  return s[:n]
}
func First(n int , s string) (first string , ok bool) {
   if len(s) >= n {
    return s[:n] , true
  }
  return s, false

}

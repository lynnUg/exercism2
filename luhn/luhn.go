package luhn
import  "strconv"
import "regexp"
//import "fmt"
func Valid( s string) bool {

  r:=regexp.MustCompile("[^0-9]")
  s=r.ReplaceAllString(s,"")
  sum:=0
  
  for i,j:= len(s)-1,1; i>=0 ; i,j=i-1,j+1{
    temp , _:=strconv.Atoi(string(s[i]))
    if j%2==0{
      temp*=2
      if temp >=10{
        temp-=9
      }
    }

    sum+=temp
  }

  if sum%10!=0 || sum==0{
    return false
  }
  return true

}

func AddCheck(s string) string {
  for i:=0;i<10;i++{
    t:=strconv.Itoa(i)

    if Valid(s+t) {
      return s+t
    }
  }
 return ""
}

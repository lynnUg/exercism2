// +build !example

package hamming
import "fmt"
const TestVersion = 2

func Distance(x,y string) ( z int , err error){
  if len(x) !=len(y)  {
    return 0 ,fmt.Errorf("wrong length")
  }

  for i:=0;i<len(x);i++{
    if x[i]!=y[i]{
      z++
    }
  }
  return
}

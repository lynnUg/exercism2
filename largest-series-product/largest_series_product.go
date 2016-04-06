package lsproduct
import  (
  "strconv"
 // "fmt"
)

const TestVersion = 2
func LargestSeriesProduct(input string , span int) (product int , err error) {
  product= -1
  temp:= 1
  for i,_:= range input {
    if (i+(span-1))<len(input) {
      for z:=i;z<i+span;z++{
        //fmt.Println(z)
        x,_:=strconv.Atoi(string(input[z]))
        temp*=x
      }
      //fmt.Println("\n")
      if temp >product{
        product=temp
      }
      temp=1
    }

   

  }

  return product,nil
}

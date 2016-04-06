package summultiples
import "fmt"

func MultipleSummer (divisors... int) func(int) int {
  return func (limit int) (y int){
    fmt.Println(divisors)
    for i:=1 ;i<limit;i++{

      for _,z:= range divisors {
        if i%z==0 {
          y+=i
          break
        }
      }

    }
    return

  }
}

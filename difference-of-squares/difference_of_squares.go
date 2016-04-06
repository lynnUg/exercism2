package diffsquares

import (
"math"
//"fmt"
)

func SumOfSquares(N int) int {
c:=make([]int,N+1)
sum :=0
for i:= range c {
sum+=int(math.Pow(float64(i),2))
i+=1

}

return sum
}

func SquareOfSums(N int) int {
c:=make([]int, N+1)
sum := 0
for i:= range c{
sum+=i
i+=1
}

return int(math.Pow(float64(sum),2))
}

func Difference(N int) int {
diff:= math.Abs(float64(SquareOfSums(N))-float64(SumOfSquares(N)))
return int( diff )
}

package pythagorean
type Triplet [3]int
func isPyth(a,b,c int ) bool {
  return (a*a + b*b) == c*c
}
func Range(min, max int) []Triplet {
  h:= []Triplet{}
  for a:=min; a<=max; a++ {
    for b:=a+1 ;b<=max; b++ {
      for c:=b+1; c<= max ; c++{
        if isPyth(a,b,c){
           h = append(h,Triplet{a,b,c})
        }
      }
    }
  }
  return h
}
func Sum(p int) []Triplet{
  h:= []Triplet{}
  for a:=1; a<p; a++ {
    for b:=a+1 ;b<p; b++ {
      c:=p-(a+b)
      if isPyth(a,b,c){
        h = append(h,Triplet{a,b,c})
      }
    }
  }
  return h
}

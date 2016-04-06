package grains

import "fmt"

func Square (n int) (x uint64, err error) {
  if n<1 || n>64 {
    
    return 0 ,fmt.Errorf("invalid value")

  }

  x=1<<(uint16(n)-1)
  err=nil
  return
}
func Total () uint64 {
  return (1<<64)-1
}

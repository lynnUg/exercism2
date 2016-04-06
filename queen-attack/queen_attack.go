package queenattack
import (
  "fmt"
)
func CanQueenAttack(x, y  string) (x1 bool, err error){
  if err=ValidPos(x); err!=nil {
    return false , err
  }
  if err=ValidPos(y); err!=nil {
    return false , err
  }
  if x==y {
    return false,fmt.Errorf("queens on same square")
  }
  if x[0]==y[0] || x[1]==y[1] {
    return true , nil
  }
  m := y[0] - x[0]
  z :=  y[1] - x[1]

  return m==z || m+z==0,nil
}
func ValidPos(pos string) error {

  if len(pos)<2 || pos[0]< 'a' || pos[0]> 'f' || pos[1] <'1' || pos[1] >'8' {
      return fmt.Errorf("invalid value %q",pos)
    }
    return nil

}

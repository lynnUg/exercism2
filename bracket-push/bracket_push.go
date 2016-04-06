package brackets
import (
//  "fmt"
)

var the_brackets map[string]string =  map[string]string{
    "{" :"}",
    "[" :"]",
    "(" :")",

 }
func Bracket (input string) (bool ,error) {
 output :="_"
  for i,_ := range input {
   val,ok :=the_brackets[output[len(output)-1:]]
   output+=input[i:i+1]
   if ok {
    if input[i:i+1]== val {
     output=output[:len(output)-2]
    
    } 
     
   }
  }

  if len(output)==1 {
   return true ,nil
  }
  return false,nil

}

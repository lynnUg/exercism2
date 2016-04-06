package palindrome

import (
  //"fmt"
  "errors"
  "strconv"
  "sort"
)
type Product struct {
  Product int
  Factorizations [][2]int
 }

func isPalindrome (p string) bool {
  for i:=0 ; i <len(p)/2 ; i++ {
    if p[i]!=p[len(p)-1-i]{
      return false
    }
  }
  return true
}
func convertToString(p int) (s string){
  if p< 0 {
         s= strconv.Itoa(p*-1)

       } else {
         s = strconv.Itoa(p)
       }
  return
}
func compareNumbers (i,j int) (f1 ,f2 int) {
        if i<j {
          f1=i
          f2=j
        } else {
          f1=j
          f2=i
        }
  return
}

func Products(fmin, fmax int) (pmin Product, pmax Product, err error) {
  if fmin > fmax {
    return Product{},Product{},errors.New("fmin > fmax")
  }
  dict:= map[int][][2]int{}
  keys:=[]int {}
  for i:=fmax ; i>=fmin ; i-- {
    for j:=i ;j>=fmin ; j-- {
       p := i*j
       s:=convertToString(p)
      if isPalindrome(s) {
        low ,high :=compareNumbers(i,j)
        dict[p]=append(dict[p],[2]int{low,high})
        keys = append(keys,p)
      }

    }
  }

  sort.Ints(keys)
  if len(dict) >=2 {
    pmin.Product=keys[0]
    pmin.Factorizations=append(pmin.Factorizations,dict[keys[0]]...)

    pmax.Product=keys[len(keys)-1]
    pmax.Factorizations=append(pmax.Factorizations,dict[keys[len(keys)-1]]...)
  } else if  len(dict) ==1 {

    pmax.Product=keys[len(keys)-1]
    pmax.Factorizations=append(pmax.Factorizations,dict[keys[len(keys)-1]]...)
  } else {
    err=errors.New("No palindromes")
  }
  return
}


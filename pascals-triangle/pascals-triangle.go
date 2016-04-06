package pascal

func PascalInt(col , row int) int {

  if col == 0 || col==row {
    return 1
  }
  return PascalInt(col ,row-1) +PascalInt(col-1, row -1)
}

func Triangle (n int)( c [][]int ){

  c = make([][]int , n)
  for i:=0 ;i<n ;i++ {
    c[i]=make([]int,i+1)
    for j:=0 ; j<i+1; j++{
      c[i][j]=PascalInt(j,i)
    }
  }
return 
}

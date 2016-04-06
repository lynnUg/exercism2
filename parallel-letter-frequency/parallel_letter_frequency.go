package letter

func ConcurrentFrequency(n []string) (m FreqMap){
  switch len(n)
  {
  case 0 :
    return FreqMap{}
  case 1 :
    return Frequency(n[0])
  }
  f:= func(n []string) {
    ch<-ConcurrentFrequency(n)
  }
  half = len(n)/2
  go f(n[:half])
  go f[n[half:])
  m:= <-ch
  for x,y := range <-ch {
  m[x]+=y
       }
  return m


}

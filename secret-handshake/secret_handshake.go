
package secret
var commands = []string {"wink", "double blink","close your eyes", "jump"}
func Handshake(n int) (h []string) {
  if n<1 {
    return []string{}
  }
  for i:=0;i<len(commands);i++{
    if n%2==1{
        h=append(h,commands[i])
    }
      n=n/2
  }

  if n%2==1{
    for i,j:=0, len(h)-1;i<j;i,j=i+1,j-1{
        h[i],h[j]=h[j],h[i]
    }
  }

  return
}

// Clock stub file


package clock
import "fmt"
// The value of TestVersion here must match `testVersion` in the file
// clock_test.go.
const TestVersion = 2




type Clock int


func Time(h, m int) Clock {

  c := Clock((h*60 + m) % 1440)

  if c<0 {
    c+=1440
  }

  return c



}

func (c Clock) String() string {

  return fmt.Sprintf("%02d:%02d",c/60,c%60)

}

func (c Clock) Add(m int) Clock {

  c=(c+Clock(m)) % 1440
  if c<0 {
    c+= 1440
  }
  return c
}



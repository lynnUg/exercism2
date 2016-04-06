// API:
//
// Open(initalDeposit int64) *Account
// (Account) Close() (payout int64, ok bool)
// (Account) Balance() (balance int64, ok bool)
// (Account) Deposit(amount uint64) (newBalance int64, ok bool)
package account
import "sync"
type Account struct {
  balance int64
  active bool
  mux sync.Mutex
}

func Open(amount int64) *Account{
  if amount >=0 {
    return &Account{balance:amount,active:true}
  }
  return nil
}
func (a *Account) Close() (payout int64, ok bool){
  a.mux.Lock()
  defer a.mux.Unlock()
  if !a.active {
    return 0,false
  }
  payout= a.balance
  ok=true
  a.balance=0
  a.active=false

  return

}
func (a *Account) Balance() (balance int64, ok bool){
  return a.balance,a.active
}
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool){
  a.mux.Lock()
  defer a.mux.Unlock()
  if amount < 0{
    temp:=-1*amount
    if temp>a.balance{

      return a.balance,false
    }
  }
  a.balance+=amount
  return a.balance,a.active
}

package bank

type RequestBank struct {
	amount int
	result chan<- bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan *RequestBank)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdraws <- &RequestBank{amount,ch}
	return <-ch
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case req := <-withdraws:
			amount := req.amount
			if balance >= amount {
				balance -= amount
				req.result <- true
			} else {
				req.result <- false
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
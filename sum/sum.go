package sum

import (
	//"sync"

	"github.com/Payrav-1997/wallet/pkg/types"
	// "log"
)

type Service struct {
	nextAccountID int64
	accounts      []*types.Account
	payments      []*types.Payment
	favorites     []*types.Favorite
}

// func Regular() int64 {
// sum := int64(0)
// for i := 0; i < 200; i++ {
// 	sum++
// }
// return sum
// }

// func Concurrently() int64 {
// wg := sync.WaitGroup{}
// wg.Add(2)
// mu := sync.Mutex{}
// sum := int64(0)
// go func() {
// 	defer wg.Done()
// 	val := int64(0)
// 	for i := 0; i < 100; i++ {
// 		val++
// 	}
// 	mu.Lock()
// 	defer mu.Unlock()
// 	sum += val

// }()
// go func() {
// 	defer wg.Done()
// 	val := int64(0)
// 	for i := 0; i < 100; i++ {
// 		val++
// 	}
// 	mu.Lock()
// 	defer mu.Unlock()
// 	sum += val

// }()

// wg.Wait()
// return sum
// }

// func (s *Service)SumPayments(goroutines int) types.Money {
	
// 	wg := sync.WaitGroup{}
// 	wg.Add(goroutines) 
// 	mu := sync.Mutex{} 
// 	var sum types.Money

// 	go func(){
// 		wg.Done()
// 		for _, payment := range s.payments {
// 			pay := payment
// 			sum += pay.Amount	
// 		}
// 		mu.Lock()
// 		defer mu.Unlock()
// 	}()
// 	go func(){
// 	    wg.Done() 
// 		for _, payment := range s.payments {
// 			pay := payment
// 			sum += pay.Amount	
// 		}
// 		mu.Lock()
// 		defer mu.Unlock()
// 	}()
// 	wg.Wait()
// 	return sum
// }
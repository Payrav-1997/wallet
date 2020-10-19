package sum

import (
	"sync"

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

func (s *Service) SumPayments(gouroutines int) types.Money {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	sum := int64(0)
	sum1 := len(s.payments) / gouroutines
	if gouroutines == 0 {
		sum1 = len(s.payments)
	}
	for i := 0; i < gouroutines-1; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			val := int64(0)
			payment := s.payments[:(i+1)*sum1]
			for _, pay := range payment {
				val += int64(pay.Amount)
			}
			mu.Lock()
			defer mu.Unlock()
			sum += val
		}(i)
		num := 0
		wg.Add(1)
		go func() {
			defer wg.Done()
			val := int64(0)
			payment := s.payments[num*sum1:]
			for _, pay := range payment {
				val += int64(pay.Amount)
			}
			mu.Lock()
			defer mu.Unlock()
			sum += val
		}()
	}
	wg.Wait()
	return types.Money(sum)
}

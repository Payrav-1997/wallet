package sum

import (
	
	"testing"

	"github.com/Payrav-1997/wallet/pkg/types"
)

// func BenchmarkReagular(b *testing.B) {
// 	want:= int64(200)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		result:= Regular()
// 		b.StopTimer()
// 		if result!=want{
// 			b.Fatalf("Invalid result got %v,want %v",result,want)
// 		}
// 		b.StartTimer()
// 	}
// }

// func BenchmarkConcurrently(b *testing.B) {
// 	want:= int64(200)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		result:= Concurrently()
// 		b.StopTimer()
// 		if result!=want{
// 			b.Fatalf("Invalid result got %v,want %v",result,want)
// 		}
// 		b.StartTimer()
// 	}
// }

func BenchmarkSumPayments(b *testing.B){
	want := types.Money(0)
	b.ResetTimer()
	var service Service
	for i := 0; i < b.N; i++ {
		result := service.SumPayments(2)
		b.StopTimer()
		if result != types.Money(want) {
			b.Fatalf("Invalid result got %v,want %v", result, want)
		}
		b.StartTimer()
	}
}



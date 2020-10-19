package sum

import "testing"

func BenchmarkReagular(b *testing.B) {
	want:= int64(200)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result:= Regular()
		b.StopTimer()
		if result!=want{
			b.Fatalf("Invalid result got %v,want %v",result,want)
		}
		b.StartTimer()
	}
}

func BenchmarkConcurrently(b *testing.B) {
	want:= int64(200)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result:= Concurrently()
		b.StopTimer()
		if result!=want{
			b.Fatalf("Invalid result got %v,want %v",result,want)
		}
		b.StartTimer()
	}
}
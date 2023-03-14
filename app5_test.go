package unit_testing

import (
	"testing"
)

//Methode Benchmark
func BenchmarkData1(b *testing.B){
	for i := 0; i < b.N; i++ {
		ModulSay("Rizki")
	}
}

//Methode Benchmark SubTest
func BenchmarkSub(b *testing.B){
	b.Run("Iki", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ModulSay("Iki")
		}
	})
	b.Run("Rizki", func(b *testing.B){
		for i := 0; i < b.N; i++ {
			ModulSay("Rizki")
		}
	})
}

//Methode Benchmark Table
func BenchmarkTable(b *testing.B){
	tableBench := []struct{
		Id string
		request string
	}{
		{
			Id : "2200",
			request : "2200",
		},
		{
			Id : "2400",
			request : "2400",
		},
	}

	for _, benchmark := range tableBench{
		b.Run(benchmark.Id, func (b *testing.B){
			ModulSay(benchmark.request)
		})
	}
}
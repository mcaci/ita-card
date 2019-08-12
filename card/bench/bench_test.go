package bench

import (
	"testing"
)

type cardCreator func(string, string) (id, error)

func bench(b *testing.B, createFrom cardCreator, n, s string) {
	for i := 0; i < b.N; i++ {
		createFrom(n, s)
	}
}

func BenchmarkSuccessOrigCreate(b *testing.B) {
	bench(b, origCreate, "1", "Coin")
}

func BenchmarkOrigInvalidNCreate(b *testing.B) {
	bench(b, origCreate, "25", "Coin")
}

func BenchmarkOrigInvalidSCreate(b *testing.B) {
	bench(b, origCreate, "1", "House")
}

func BenchmarkOrigInvalidCreate(b *testing.B) {
	bench(b, origCreate, "25", "House")
}

func BenchmarkSuccessOldCreate(b *testing.B) {
	bench(b, oldCreate, "1", "Coin")
}

func BenchmarkOldInvalidNCreate(b *testing.B) {
	bench(b, oldCreate, "25", "Coin")
}

func BenchmarkOldInvalidSCreate(b *testing.B) {
	bench(b, oldCreate, "1", "House")
}

func BenchmarkOldInvalidCreate(b *testing.B) {
	bench(b, oldCreate, "25", "House")
}

func BenchmarkSuccessSeqCreate(b *testing.B) {
	bench(b, seqCreate, "1", "Coin")
}

func BenchmarkSeqInvalidNCreate(b *testing.B) {
	bench(b, seqCreate, "25", "Coin")
}

func BenchmarkSeqInvalidSCreate(b *testing.B) {
	bench(b, seqCreate, "1", "House")
}

func BenchmarkSeqInvalidCreate(b *testing.B) {
	bench(b, seqCreate, "25", "House")
}

func BenchmarkSuccessConcurCreate(b *testing.B) {
	bench(b, concurCreate, "1", "Coin")
}

func BenchmarkConcurInvalidNCreate(b *testing.B) {
	bench(b, concurCreate, "25", "Coin")
}

func BenchmarkConcurInvalidSCreate(b *testing.B) {
	bench(b, concurCreate, "1", "House")
}

func BenchmarkConcurInvalidCreate(b *testing.B) {
	bench(b, concurCreate, "25", "House")
}

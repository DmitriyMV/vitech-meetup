package main

import (
	"iter"
	"maps"
	"slices"
)

func main() {
	slc := []string{"Vi", "Tech", "Meetup", "Moscow"}
	slc2 := []string{"Dmitry", "Kirill", "Edgar", "Daniil"}

	var m map[string]string = maps.Collect(
		Zip(
			TakeSeq2(slices.All(slc)),
			TakeSeq2(slices.All(slc2)),
		),
	)

	for k, v := range MapSeq2(maps.All(m), func(k, v string) (string, string, bool) {
		return "key:" + k, "value:" + v, true
	}) {
		println(k, v)
	}
}

func TakeSeq2[T1, T2 any](it iter.Seq2[T1, T2]) iter.Seq[T2] {
	return func(yield func(T2) bool) {
		for _, v := range it {
			if !yield(v) {
				break
			}
		}
	}
}

func MapSeq2[T1, T2, R1, R2 any](it iter.Seq2[T1, T2], fn func(T1, T2) (R1, R2, bool)) iter.Seq2[R1, R2] {
	return func(yield func(R1, R2) bool) {
		for v1, v2 := range it {
			r1, r2, ok := fn(v1, v2)
			if !ok {
				break
			}

			if !yield(r1, r2) {
				break
			}
		}
	}
}

func Zip[T, V any](it iter.Seq[T], it2 iter.Seq[V]) iter.Seq2[T, V] {
	return func(yield func(T, V) bool) {
		next, stop := iter.Pull(it2)
		defer stop()

		for v := range it {
			v2, ok := next()
			if !ok {
				break
			}

			if !yield(v, v2) {
				break
			}
		}
	}
}

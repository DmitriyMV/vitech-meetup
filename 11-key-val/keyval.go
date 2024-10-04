package main

import (
	"iter"
	"sync"
)

type KeyVal[K comparable, V any] struct {
	mx sync.RWMutex
	m  map[K]V
}

func NewKeyVal[K comparable, V any]() *KeyVal[K, V] {
	return &KeyVal[K, V]{m: map[K]V{}}
}

func (kv *KeyVal[K, V]) Set(k K, v V) {
	kv.mx.Lock()
	defer kv.mx.Unlock()

	kv.m[k] = v
}

func (kv *KeyVal[K, V]) Get(k K) (V, bool) {
	kv.mx.RLock()
	defer kv.mx.RUnlock()

	v, ok := kv.m[k]

	return v, ok
}

func (kv *KeyVal[K, V]) Delete(k K) {
	kv.mx.Lock()
	defer kv.mx.Unlock()

	delete(kv.m, k)
}

func (kv *KeyVal[K, V]) Len() int {
	kv.mx.RLock()
	defer kv.mx.RUnlock()

	return len(kv.m)
}

func (kv *KeyVal[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		kv.mx.RLock()

		for k, v := range kv.m {
			kv.mx.RUnlock()

			if !yield(k, v) {
				break
			}

			kv.mx.RLock()
		}

		kv.mx.RUnlock()
	}
}

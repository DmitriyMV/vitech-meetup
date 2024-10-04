package main

func main() {
	kv := NewKeyVal[string, int]()
	kv.Set("one", 1)
	kv.Set("two", 2)
	kv.Set("three", 3)

	for k, v := range kv.All() {
		println(k, v)
	}

	for k, _ := range kv.All() {
		kv.Delete(k)
	}

	println(kv.Len())
}

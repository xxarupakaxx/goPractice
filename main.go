package main

import (
	"fmt"
	"sync"
)

//KeyValueのための型内部にMaｐを所持
type KeyValue struct {
	store map[string]string //keyValue を格納
	mu    sync.RWMutex      //排他制御のためのmutex
}

func NewKeyValue() *KeyValue {
	return &KeyValue{store: make(map[string]string)}
}
func (kv *KeyValue) Set(key, val string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.store[key] = val
}
func (kv *KeyValue) Get(key string) (string, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	val, ok := kv.store[key]
	return val, ok
}

func main() {
	kv := NewKeyValue()
	kv.Set("key", "value")
	value, ok := kv.Get("key")
	if ok {
		fmt.Println(value)
	}
}

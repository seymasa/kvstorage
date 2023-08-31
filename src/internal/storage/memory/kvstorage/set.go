package kvstorage

func (ms *memoryStorage) Set(key string, value any) any {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if _, err := ms.Get(key); err == nil {
		// add set func
	}
	ms.db[key] = value
	return value
}

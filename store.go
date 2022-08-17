package tbkv

type KVStore struct {
	data   map[string]string
	save   chan ItemRequest
	fetch  chan Request
	delete chan Request
}

func (s *KVStore) Get(key string) (string, error) {
	request := Request{
		Key:    key,
		Result: make(chan Result, 1),
	}
	s.fetch <- request
	result := <-request.Result
	if !result.Ok {
		return "", ErrNotFound
	}
	return result.Value, nil
}

func (s *KVStore) Set(key, value string) {
	setItemRequest := ItemRequest{
		Item: Item{
			Key:   key,
			Value: value,
		},
		Ok: make(chan bool, 1),
	}
	s.save <- setItemRequest
	<-setItemRequest.Ok
}

func (s *KVStore) Delete(key string) {
	request := Request{
		Key:    key,
		Result: make(chan Result, 1),
	}
	s.delete <- request
	<-request.Result
}

// NewStore creates a new KVStore.
// requestBufferSize is the size of the request buffer.
// If requestBufferSize is 0, it will use a default of 20.
func NewStore(requestBufferSize uint8) KVStore {
	var buffer uint8
	if requestBufferSize == 0 {
		buffer = 20
	} else {
		buffer = requestBufferSize
	}
	kvs := KVStore{}
	kvs.data = make(map[string]string)
	kvs.save = make(chan ItemRequest, buffer)
	kvs.fetch = make(chan Request, buffer)
	kvs.delete = make(chan Request, buffer)
	go kvs.loop() // starts listening for requests
	return kvs
}

// NewDefaultStore is a convenience function for creating a KVStore with sensible defaults.
// It will create a KVStore with a request buffer size of 20.
func NewDefaultStore() KVStore {
	return NewStore(0)
}

package tbkv

func (s *KVStore) loop() {
	for {
		select {

		// ---
		// SET
		// ---

		case setRequest := <-s.save:
			s.data[setRequest.Item.Key] = setRequest.Item.Value
			setRequest.Ok <- true // cannot fail

		// ---
		// GET
		// ---

		case getRequest := <-s.fetch:
			val, ok := s.data[getRequest.Key]
			getRequest.Result <- Result{
				Value: val,
				Ok:    ok,
			}

		// ------
		// Delete
		// ------

		case delRequest := <-s.delete:
			delete(s.data, delRequest.Key)
			delRequest.Result <- Result{
				Value: delRequest.Key,
				Ok:    true, // cannot fail
			}
		}
	}
}

package ordone

// OrDone returns a channel that closes when the done channel closes
// it transfers all values from the dataChannel to the returned channel
func OrDone[T any](done <-chan struct{}, dataChannel <-chan T) <-chan T {
	valStream := make(chan T)

	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-dataChannel:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <-done:
					return
				}
			}
		}
	}()
	return valStream
}

// OrDone returns a channel that closes when the done channel closes
// it transfers all values from the dataChannel to the returned channel
func OrDone_(done <-chan struct{}, dataChannel <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})

	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-dataChannel:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <-done:
					return
				}
			}
		}
	}()
	return valStream
}

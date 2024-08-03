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
				println("orDone goroutine stopped, done channel closed")
				return
			case v, ok := <-dataChannel:
				if !ok {
					println("orDone goroutine stopped: dataChannel closed!")
					return
				}
				select {
				case valStream <- v:
				case <-done:
					println("orDone goroutine stopped, done channel closed")
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
				println("orDone stopped")
				return
			case v, ok := <-dataChannel:
				if !ok {
					println("orDone stopped: dataChannel closed!")
					return
				}
				select {
				case valStream <- v:
				case <-done:
					println("orDone stopped")
					return
				}
			}
		}
	}()
	return valStream
}

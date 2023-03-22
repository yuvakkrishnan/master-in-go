// In Go, an empty struct is a struct type that has no fields. An empty struct takes up zero bytes of memory and is used in a variety of ways, including:

// As a marker type: An empty struct can be used as a marker type to indicate the absence of data. For example, it can be used to represent an event or a signal without carrying any data payload.

// As a placeholder type: An empty struct can be used as a placeholder type when a struct is required but no data needs to be stored. For example, it can be used as a key type in a map where the keys are used only for their identity.

// As a compact representation: An empty struct can be used as a compact representation of a type, which can be useful when memory usage is a concern.

// Here's an example of using an empty struct as a marker type:
type Done struct{}

func worker(done chan Done) {
	// Do work here
	done <- Done{}
}

func main() {
	done := make(chan Done)
	go worker(done)
	<-done
}

//In this example, the Done struct is used as a marker type to indicate when the worker Goroutine has completed its work.
//The worker Goroutine sends a value of type Done on the done channel when it has finished its work, and the main Goroutine receives the value to indicate that the worker has completed.
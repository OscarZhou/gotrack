package track

// Config is used to support custom configuration parameters
type Config struct {
	// Debug Determine if the log is printed out
	Debug bool
	// Debug Determine if there is asynchronous log
	AsynLog bool
	// AsynLogInterval The interval of the asynchronous log
	AsynLogInterval int
}

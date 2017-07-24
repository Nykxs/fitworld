package fitworld

// Server defines the behaviour that should be implemented by any object that want to listen events doesnt matter the protocol
type Server interface {
	Setup() error
	Start() error
	Stop() error
}

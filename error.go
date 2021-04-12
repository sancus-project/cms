package cms

type Error interface {
	Error() string
	Status() int
	Unwrap() error
}

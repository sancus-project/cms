package cms

type UserCapability int

const (
	CanRead UserCapability = iota
	CanWrite
	CanPurge
)

type User interface {
	Username() string
	CanDo(c UserCapability, r Resource) bool
	CanDoPath(c UserCapability, path string) bool
}

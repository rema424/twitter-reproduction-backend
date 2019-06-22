package dbutil

// DB ...
type DB interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

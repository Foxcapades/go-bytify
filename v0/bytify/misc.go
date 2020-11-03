package bytify

// Writer is an alias of io.Writer.
//
// This is here simply to avoid pulling in any external packages, use io.Writer
// if you are using io.Writer.
type Writer interface {
	Write(p []byte) (n int, err error)
}

//go:build !snet_trace
// +build !snet_trace

package snack

func (l *Listener) trace(format string, args ...interface{}) {
}

func (c *Conn) trace(format string, args ...interface{}) {
}

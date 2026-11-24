package ui

import "strings"

type CollectorWriter struct {
	Output string
}

func (c *CollectorWriter) Write(p []byte) (n int, err error) {
	c.Output += strings.Clone(string(p))
	return len(p), nil
}

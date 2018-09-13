package counter

import "io"

type ByteCounter struct {
	Writer    io.Writer
	ByteCount int64
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	counter := ByteCounter{w, 0}
	return &counter, &counter.ByteCount
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	n, err := c.Writer.Write(p)
	c.ByteCount += int64(n)
	return n, err
}

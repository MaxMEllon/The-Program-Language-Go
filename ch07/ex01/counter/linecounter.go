package counter

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	for _, itr := range p {
		if itr == '\n' {
			*c++
		}
	}
	return len(p), nil
}

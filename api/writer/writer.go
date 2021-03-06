package writer

// Writer writes bytes to output
type Writer struct {
	output *[]byte
}

func (w Writer) Write(p []byte) (n int, err error) {
	*w.output = append(*w.output, p...)
	return len(p), nil
}

// NewWriter creates a Writer with a given byte output
func NewWriter(o *[]byte) Writer {
	return Writer{
		output: o,
	}
}

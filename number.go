package pjn

var (
	Zero Value = produceNumberZero
)

func produceNumberZero(buf *Buffer) (err error) {
	buf.AppendByte(valueNumberZero)
	return nil
}

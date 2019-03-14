package pjn

var (
	Null Value = produceNull
)

func produceNull(buf *Buffer) (err error) {
	buf.AppendBytes(valueNull)
	return nil
}

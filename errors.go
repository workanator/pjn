package pjn

const (
	whereObject = "object"
	whereArray  = "array"
	whereValue  = "value"
	whereField  = "field"
)

var (
	ErrNilReference = errNilReference{}
)

type ErrProduceFailed struct {
	Where  string
	Reason error
}

func (e ErrProduceFailed) Error() (s string) {
	s = string(e.Where) + " produce failed"
	if e.Reason != nil {
		s += ": " + e.Reason.Error()
	}
	return s
}

func ObjectProduceFailed(reason error) ErrProduceFailed {
	return ErrProduceFailed{
		Where:  whereObject,
		Reason: reason,
	}
}

func ArrayProduceFailed(reason error) ErrProduceFailed {
	return ErrProduceFailed{
		Where:  whereArray,
		Reason: reason,
	}
}

func ValueProduceFailed(reason error) ErrProduceFailed {
	return ErrProduceFailed{
		Where:  whereValue,
		Reason: reason,
	}
}

func MemberProduceFailed(reason error) ErrProduceFailed {
	return ErrProduceFailed{
		Where:  whereField,
		Reason: reason,
	}
}

type errNilReference struct{}

func (errNilReference) Error() string { return "nil reference" }

func produceError(err error) Value {
	return func(*Buffer) error {
		return err
	}
}

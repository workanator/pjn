package pjn

func Bool(value bool) Produce {
	if value {
		return produceBoolTrue
	} else {
		return produceBoolFalse
	}
}

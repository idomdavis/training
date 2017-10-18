package intconv

type Converter func(int) string

type Result struct {
	In  int
	Out string
	Converter
}

func (c Converter) Convert(i int) Result {
	r := Result{In: i, Converter: c}
	r.Out = c(i)
	return r
}

func Convert(c Converter, i int) Result {
	r := Result{In: i, Converter: c}
	r.Out = c(i)
	return r
}

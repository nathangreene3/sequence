package sequence

// format ...
type format []baseFmt

func newFormat(baseFmts ...baseFmt) format {
	f := make(format, len(baseFmts))
	copy(f, baseFmts)
	return f
}

func (f format) copy() format {
	cpy := make(format, len(f), cap(f))
	copy(cpy, f)
	return cpy
}

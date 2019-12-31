package sequence

// CharacterType ...
type CharacterType byte

const (
	// Numeric ...
	Numeric = CharacterType(0)
	// ASCIIAlpha ...
	ASCIIAlpha = CharacterType('@')
	// ASCIINumeric ...
	ASCIINumeric = CharacterType('#')
	// ASCIIAlphanumeric ...
	ASCIIAlphanumeric = CharacterType('!')
)

// Bytes implements the iterator interface.
type Bytes struct {
	overflowed bool
	dims       int
	start      Start
	current    Current
	end        End
	format     Format
	incOrder   IncrementOrder
	skip       SkipFnc
}

// Start ...
type Start Field

// Current ...
type Current Field

// End ...
type End Field

// Exceptions ...
type Exceptions map[byte]struct{}

// Format ...
type Format []CharFmt

// NewFormat ...
func NewFormat(cfs ...CharFmt) Format {
	f := make(Format, 0, len(cfs))
	for _, cf := range cfs {
		f = append(f, cf)
	}

	return f
}

// CharFmt ...
type CharFmt struct {
	charType CharacterType
	min, max byte
}

// NewCharFmt ...
func NewCharFmt(min, max byte, incOrder int) CharFmt {
	if max < min {
		panic("invalid range")
	}

	cf := CharFmt{
		min: min,
		max: max,
	}

	switch {
	case 0 <= min && min <= max && max <= 9:
		cf.charType = Numeric
	case '0' <= min && min <= max && max <= '9':
		cf.charType = ASCIINumeric
	case 'A' <= min && min <= max && max <= 'Z':
		cf.charType = ASCIIAlpha
	case '0' <= min && min <= '9' && 'A' <= max && max <= 'Z':
		cf.charType = ASCIIAlphanumeric
	default:
		panic("invalid character type detected")
	}

	return cf
}

// IncrementOrder ... The ith value is the next index to increment.
type IncrementOrder []int

// MaxChars ...
type MaxChars Field

// MinChars ...
type MinChars Field

// SkipFnc ...
type SkipFnc func(c Current) bool

// NewBytes ...
func NewBytes(opts ...interface{}) Bytes {
	bts := Bytes{
		skip: func(c Current) bool { return false },
	}

	for _, opt := range opts {
		switch t := opt.(type) {
		case Start:
			bts.start = t.Copy()
		case Current:
			bts.current = t.Copy()
		case End:
			bts.end = t.Copy()
		case Format:
			bts.dims = len(t)
			bts.format = t.Copy()
		case IncrementOrder:
			bts.incOrder = t.Copy()
		case SkipFnc:
			bts.skip = t
		default:
			panic("invalid option type")
		}
	}

	if bts.format == nil {
		panic("format is required")
	}

	if bts.start == nil {
		bts.start = bts.format.Start()
		// TODO: Increment for any exception characters
	}

	if bts.current == nil {
		bts.current = Current(bts.start.Copy())
	}

	// TODO: compare start to current to end
	if bts.end == nil {
		bts.end = bts.format.End()
		// TODO: Decrement for any exception characters
	}

	bts.validate()
	return bts
}

func incChar(char byte, cf CharFmt) (byte, byte) {
	switch char {
	case cf.max:
		return cf.min, 1
	default:
		switch cf.charType {
		case ASCIIAlphanumeric:
			if char == '9' {
				return 'A', 0
			}

			return char + 1, 0
		default:
			return char + 1, 0
		}
	}
}

func (bts *Bytes) increment() {
	var carry byte
	for _, index := range bts.incOrder {
		var (
			cf = bts.format[index]
		)

		bts.current[index], carry = incChar(bts.current[index], cf)
		if carry != 0 {

		}
	}
}

// TODO
func (bts *Bytes) validate() {}

// Compare ...
func (bts Bytes) Compare(b Bytes) int {
	if bts.dims != b.dims {
		panic("invalid dimension")
	}

	for _, index := range bts.incOrder {
		if bts.format[index].charType != b.format[index].charType {
			panic("invalid character format")
		}

		x, y := bts.current[index], b.current[index]
		switch {
		case x < y:
			return -1
		case y < x:
			return 1
		}
	}

	return 0
}

// Copy ...
func (s Start) Copy() Start {
	cpy := make(Start, len(s), cap(s))
	copy(cpy, s)
	return cpy
}

// Copy ...
func (c Current) Copy() Current {
	cpy := make(Current, len(c), cap(c))
	copy(cpy, c)
	return cpy
}

// Copy ...
func (e End) Copy() End {
	cpy := make(End, len(e), cap(e))
	copy(cpy, e)
	return cpy
}

// Copy ...
func (e Exceptions) Copy() Exceptions {
	cpy := make(Exceptions)
	for k, v := range e {
		cpy[k] = v
	}

	return cpy
}

// Copy ...
func (f Format) Copy() Format {
	cpy := make(Format, 0, cap(f))
	for _, cf := range f {
		cpy = append(cpy, cf.Copy())
	}

	return cpy
}

// End ...
func (f Format) End() End {
	e := make(End, 0, len(f))
	for _, cf := range f {
		e = append(e, cf.max)
	}

	return e
}

// Start ...
func (f Format) Start() Start {
	s := make(Start, 0, len(f))
	for _, cf := range f {
		s = append(s, cf.min)
	}

	return s
}

// Copy ...
func (io IncrementOrder) Copy() IncrementOrder {
	cpy := make(IncrementOrder, len(io), cap(io))
	copy(cpy, io)
	return cpy
}

// Copy ...
func (mc MaxChars) Copy() MaxChars {
	cpy := make(MaxChars, len(mc), cap(mc))
	copy(cpy, mc)
	return cpy
}

// Copy ...
func (mc MinChars) Copy() MinChars {
	cpy := make(MinChars, len(mc), cap(mc))
	copy(cpy, mc)
	return cpy
}

// Copy ...
func (cf CharFmt) Copy() CharFmt {
	return CharFmt{
		charType: cf.charType,
		min:      cf.min,
		max:      cf.max,
	}
}

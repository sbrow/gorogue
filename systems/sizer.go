package systems

// Sizer returns how many entities a system is aware of.
type Sizer interface {
	Size() int
}

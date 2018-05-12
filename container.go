package gorogue

type Container interface {
	Object
	Contents() []Object
	Add(o Object)
	Remove(i int) *Object
}

package foo

type Foo struct {
	Name string
}

func (Foo) Bar() {}

type Interface interface {
	Interface()
}

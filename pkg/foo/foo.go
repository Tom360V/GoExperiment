package foo

import "log"

// Foo object
type Foo struct {
	name string
}

// NewFoo will return a Foo object
func NewFoo(name string) *Foo {
	f := Foo{
		name: name,
	}

	return &f
}

// Show will show content of the Foo
func (f *Foo) Show() {
	log.Println("Foo name: " + f.name)
}

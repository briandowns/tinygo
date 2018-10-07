package main

func main() {
	thing := &Thing{"foo"}
	println("thing:", thing.String())
	thing.Print()
	printItf(5)
	printItf(byte('x'))
	printItf("foo")
	printItf(*thing)
	printItf(thing)
	printItf(Stringer(thing))
	printItf(Number(3))
	array := Array([4]uint32{1, 7, 11, 13})
	printItf(array)
	s := Stringer(thing)
	println("Stringer.String():", s.String())
	var itf interface{} = s
	println("Stringer.(*Thing).String():", itf.(Stringer).String())
}

func printItf(val interface{}) {
	switch val := val.(type) {
	case Doubler:
		println("is Doubler:", val.Double())
	case Tuple:
		println("is Tuple:", val.Nth(0), val.Nth(1), val.Nth(2), val.Nth(3))
		val.Print()
	case int:
		println("is int:", val)
	case byte:
		println("is byte:", val)
	case string:
		println("is string:", val)
	case Thing:
		println("is Thing:", val.String())
	case *Thing:
		println("is *Thing:", val.String())
	case Foo:
		println("is Foo:", val)
	default:
		println("is ?")
	}
}

type Thing struct {
	name string
}

func (t Thing) String() string {
	return t.name
}

func (t Thing) Print() {
	println("Thing.Print:", t.name)
}

type Stringer interface {
	String() string
}

type Foo int

type Number int

func (n Number) Double() int {
	return int(n) * 2
}

type Doubler interface {
	Double() int
}

type Tuple interface {
	Nth(int) uint32
	Print()
}

type Array [4]uint32

func (a Array) Nth(n int) uint32 {
	return a[n]
}

func (a Array) Print() {
	println("Array len:", len(a))
}

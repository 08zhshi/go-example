package main

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

type my_type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      uint8
	align      uint8
	fieldAlign uint8
	kind       uint8
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata    *byte
	str       int32
	ptrToThis int32
}
type myeface struct {
	_type my_type
	data  unsafe.Pointer
}

type ITest interface {
	TestPrint()
}
type STest struct {
	name string
	age  int
}

func NewSTest(s string) ITest {
	return &STest{name: s, age: 18}
}
func (t *STest) TestPrint() {
	fmt.Println("name=", t.name)
}

type geomtry interface {
	area() float64
	perim() float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geomtry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main() {
	x := 10
	var i interface{} = x
	fmt.Println(i)
	st := NewSTest("hua")
	st.TestPrint()
	myInfer := *(*myeface)(unsafe.Pointer(&x))
	fmt.Println(myInfer.data)

	v := reflect.ValueOf(st)
	fmt.Println(v.Kind())
	t := reflect.TypeOf(st)
	fmt.Println(v.Type(), t, t.String())
	n := v.Elem().FieldByName("name")

	fmt.Println(n)
	fmt.Println(v.Elem().Interface().(STest))

	tt, _ := t.Elem().FieldByName("age")
	p := *(*int32)(unsafe.Pointer(v.Pointer() + tt.Offset))
	fmt.Println(p)

	fmt.Println(reflect.PtrTo(t))

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

}

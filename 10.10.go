//package main
////
////import "fmt"
////
////type Base struct {
////	 id int
////}
////
////func (b *Base)Id() int {
////	return b.id
////}
////
////func (b *Base)SetId(id int)  {
////	b.id = id
////}
////
////type Person struct {
////	Base
////	FirstName string
////	LastName string
////}
////
////type Employee struct {
////	Person
////	salary int
////}
////
////func (e *Employee)CalcSalary() int  {
////	return e.salary
////}
////
////func main() {
////	e := &Employee{Person{Base{id:12}, "dasdsa", "dasdasdsa" }, 12121}
////
////	fmt.Println(e.Id())
////}


package main

import (
	"fmt"
	"runtime"
)

//import (
//	"fmt"
//)
//
//type Base struct{}
//
//func (Base) Magic() {
//	fmt.Println("base magic")
//}
//
//func (self Base) MoreMagic() {
//	self.Magic()
//	self.Magic()
//}
//
//type Voodoo struct {
//	Base
//}
//
//func (Voodoo) Magic() {
//	fmt.Println("voodoo magic")
//}
//
//func main() {
//	v := new(Voodoo)
//	v.Magic()
//	v.MoreMagic()
//}

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)
}



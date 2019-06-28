package main
//
//import "fmt"
//
//type rect struct {
//	width, height int
//}
//
//func (r *rect) area() int  {
//	return r.height*r.width
//}
//
//func (r rect) perim() int {
//	return r.width*2 + r.height*2
//}
//
//func (r rect) setWidth(w int)  {
//	r.width = w
//}
//
//func (r *rect) setWidthPtr(w int)  {
//	r.width = w
//}
//
//func main() {
//
//	r:=rect{3,5}
//
//	fmt.Println(r.area())
//
//	fmt.Println(r.perim())
//
//	rptr := &r
//
//	fmt.Println(rptr.area())
//
//	fmt.Println(rptr.perim())
//
//	r.setWidth(100)
//
//	fmt.Println(r.width)
//
//	rptr.setWidthPtr(200)
//
//	fmt.Println("rptr.width", rptr.width)
//
//	fmt.Println("r.width", r.width)
//
//
//
//}
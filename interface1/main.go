package main

type area interface {
	FindArea() int
}
type square struct{
	edges int 
}
type rectangle struct{
	length int 
	width int
}

func (mySquare square) FindArea() int{
	return mySquare.edges * mySquare.edges
}
func (myRectangle rectangle) FindArea() int{
	return myRectangle.length*myRectangle.width
}
func main()  {
	var s square=square{edges:10,}
	result_area_square:=s.FindArea()
	print("area of square: ",result_area_square)
	print("\n==================\n")

	var r rectangle=rectangle{10,7}
	result_area_rectangle:=r.FindArea()
	print("area of rectangle: ",result_area_rectangle)

}
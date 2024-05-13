package main

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Length float64
}

type Circle struct {
	Radius float64
}

// func Perimeter(width,length int) (perimeter int) {
// 	return 2*(width+length)
// }

func Perimeter(r Rectangle) (perimeter float64) {
	return 2 * (r.Width + r.Length)
}

// func Area(width,length int)(area int){
// 	return width*length
// }

func (r Rectangle) Area() (area float64) {
	return r.Width * r.Length
}

func (c Circle) Area() (area float64) {
	return math.Pi * c.Radius * c.Radius
}

func main(){
	fmt.Println(" interface polymorp interface test")
}
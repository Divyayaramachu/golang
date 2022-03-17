package main

import (
	"fmt"
	"math"
)


func Square(){
var l int = 5
 var sq int
 sq = l*l
fmt.Println("Area of Square:", sq)


}
func Rect(){
var l int = 7
var b int = 8
var rect = 2*(l+b)
fmt.Println("Area of Rectangle:", rect)

}

func Cir(){
var r float64 = 5
var area float64
area = math.Pi*(float64(r)*float64(r))
fmt.Println(float64(area))
}

func main(){
Square()
Rect()
Cir()
}
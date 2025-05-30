package samples

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Check number is even, odd, perfect square, prime
Input params: (num int) number
*/
func CheckNumber() {
	var num = 10
	println("Check even number: ", gouse.IsEven(num))
	println("Check odd number: ", gouse.IsOdd(num))
	println("Check perfect square number: ", gouse.IsPerfectSquare(num))
	println("Check prime number: ", gouse.IsPrime(num))
}

/*
Description: Calculate logarithm (base 2, 10) of number (F: float)
Input params: (num, base int)
*/
func Log() {
	println("Logarithm of integer number: ", gouse.Log(16, 2))
	println("Logarithm of float number: ", fmt.Sprintf("%f", gouse.LogF(20.0, 2.0)))

	println("Logarithm of integer number (base 2): ", gouse.Log2(16))
	println("Logarithm of float number (base 2): ", fmt.Sprintf("%f", gouse.Log2F(20.0)))

	println("Logarithm of integer number (base 10): ", gouse.Log10(100))
	println("Logarithm of float number (base 10): ", fmt.Sprintf("%f", gouse.Log10F(20.0)))
}

/*
Description: Calculate Pytago (F: float)
Input params: (num1, num2 int)
*/
func Pytago() {
	fmt.Printf("Pytago of number (integer): %d", gouse.Pytago(3, 4))
	fmt.Printf("Pytago of number (float): %f", gouse.Pytago(3.0, 4.0))
}

/*
Description: Calculate transition (speed, distance, time)
Input params: Speed(distance, time float64), Distance(speed, time float64), Time(distance, speed float64)
*/
func Transition() {
	var distance, speed, time float64 = 100, 10, 10
	fmt.Printf("Speed: %f", gouse.Speed(distance, time))
	fmt.Printf("Distance: %f", gouse.Distance(speed, time))
	fmt.Printf("Time: %f", gouse.Time(distance, speed))
}

/*
Description: Calculate trigonometry (sine, cosine, tangent) (F: float)
Input params: (angle float64)
*/
func Trigonometry() {
	println("Sine of integer number: ", gouse.Sin(90))
	println("Sine of float number: ", fmt.Sprintf("%f", gouse.SinF(90.0)))
	println("Cosine of integer number: ", gouse.Cos(90))
	println("Cosine of float number: ", fmt.Sprintf("%f", gouse.CosF(90.0)))
	println("Tangent of integer number: ", gouse.Tan(90))
	println("Tangent of float number: ", fmt.Sprintf("%f", gouse.TanF(90.0)))
}

/*
Description: Calculate area, perimeter, volume of circle (F: float)
Input params: (radius float64)
*/
func Circle() {
	fmt.Printf("Area of circle (integer): %f", gouse.AreaCir(10))
	fmt.Printf("Area of circle (float): %f", gouse.AreaCir(10.0))
	fmt.Printf("Perimeter of circle (integer): %f", gouse.PeriCir(10))
	fmt.Printf("Perimeter of circle (float): %f", gouse.PeriCir(10.0))
}

/*
Description: Calculate area, volume of cone (F: float)
Input params: (radius, height float64)
*/
func Cone() {
	fmt.Printf("Area of cone (integer): %f", gouse.AreaCone(10, 20))
	fmt.Printf("Area of cone (float): %f", gouse.AreaCone(10.0, 20.0))
	fmt.Printf("Volume of cone (integer): %f", gouse.VolCone(10, 20))
	fmt.Printf("Volume of cone (float): %f", gouse.VolCone(10.0, 20.0))
}

/*
Description: Calculate area, perimeter, volume of cube (F: float)
Input params: (side float64)
*/
func Cube() {
	println("Area of cube (integer): ", gouse.AreaCube(10))
	println("Area of cube (float): ", fmt.Sprintf("%f", gouse.AreaCube(10.0)))
	println("Perimeter of cube (integer): ", gouse.PeriCube(10))
	println("Perimeter of cube (float): ", fmt.Sprintf("%f", gouse.PeriCube(10.0)))
	println("Volume of cube (integer): ", gouse.VolCube(10))
	println("Volume of cube (float): ", fmt.Sprintf("%f", gouse.VolCubeF(10.0)))
}

/*
Description: Calculate area, volume of cylinder (F: float)
Input params: (radius, height float64)
*/
func Cylinder() {
	fmt.Printf("Area of cylinder (integer): %f", gouse.AreaCyl(10, 20))
	fmt.Printf("Area of cylinder (float): %f", gouse.AreaCyl(10.0, 20.0))
	fmt.Printf("Volume of cylinder (integer): %f", gouse.VolCyl(10, 20))
	fmt.Printf("Volume of cylinder (float): %f", gouse.VolCyl(10.0, 20.0))
}

/*
Description: Calculate area, perimeter, volume of ellipse (F: float)
Input params: (major, minor float64)
*/
func Ellipse() {
	fmt.Printf("Area of ellipse (integer): %f", gouse.AreaEllipse(10, 20))
	fmt.Printf("Area of ellipse (float): %f", gouse.AreaEllipse(10.0, 20.0))
}

/*
Description: Calculate area, perimeter, volume of parallelogram (F: float)
Input params: (base, height float64)
*/
func Parallelogram() {
	println("Area of parallelogram (integer): ", gouse.AreaParallelogram(10, 20))
	println("Area of parallelogram (float): ", fmt.Sprintf("%f", gouse.AreaParallelogram(10.0, 20.0)))
}

/*
Description: Calculate area of pentagon (F: float)
Input params: (side, apothem float64)
*/
func Polygon() {
	fmt.Printf("Area of pentagol by number of sides (integer): %f", gouse.AreaPolygon(10, 6))
}

/*
Description: Calculate area, perimeter, volume of rectangle (F: float)
Input params: (length, width float64)
*/
func Rect() {
	println("Area of rectangle: ", gouse.AreaRect(10, 20))
	println("Perimeter of rectangle (integer): ", gouse.PeriRect(10, 20))
	println("Perimeter of rectangle (float): ", fmt.Sprintf("%f", gouse.PeriRect(10.0, 20.0)))
	println("Diagonal of rectangle (integer): ", fmt.Sprintf("%d", gouse.DiagRect(10, 20)))
	println("Diagonal of rectangle (float): ", fmt.Sprintf("%f", gouse.DiagRect(10.0, 20.0)))
	println("Volume of rectangular (integer): ", gouse.VolRect(10, 20, 30))
	println("Volume of rectangular (float): ", fmt.Sprintf("%f", gouse.VolRect(10.0, 20.0, 30.0)))
}

/*
Description: Calculate area, perimeter, volume of rhombus (F: float)
Input params: (diagonal1, diagonal2 float64)
*/
func Rhombus() {
	println("Area of rhombus (integer): ", gouse.AreaRhombus(10, 20))
	println("Area of rhombus (float): ", fmt.Sprintf("%f", gouse.AreaRhombus(10.0, 20.0)))
}

/*
Description: Calculate area, volume of sphere (F: float)
Input params: (radius float64)
*/
func Sphere() {
	fmt.Printf("Area of sphere (integer): %f", gouse.AreaSphere(10))
	fmt.Printf("Area of sphere (float): %f", gouse.AreaSphereF(10.0))
	fmt.Printf("Volume of sphere (integer): %f", gouse.VolSphere(10))
	fmt.Printf("Volume of sphere (float): %f", gouse.VolSphereF(10.0))
}

/*
Description: Calculate area, perimeter, volume of square (F: float)
Input params: (side float64)
*/
func Square() {
	println("Area of square (integer): ", gouse.AreaSquare(10))
	println("Area of square (float): ", fmt.Sprintf("%f", gouse.AreaSquare(10.0)))
	println("Perimeter of square (integer): ", gouse.PeriSquare(10))
	println("Perimeter of square (float): ", fmt.Sprintf("%f", gouse.PeriSquare(10.0)))
}

/*
Description: Calculate area, perimeter, volume of trapezoid (F: float)
Input params: (base1, base2, height float64)
*/
func Trapezoid() {
	fmt.Printf("Area of trapezoid (integer): %f", gouse.AreaTrapezoid(10, 20, 30))
	fmt.Printf("Area of trapezoid (float): %f", gouse.AreaTrapezoid(10.0, 20.0, 30.0))
}

/*
Description: Calculate area, perimeter of triangle (F: float)
Input params: (base, height, side float64)
*/
func Triangle() {
	println("Area of triangle (integer): ", gouse.AreaTri(10, 20))
	println("Area of triangle (float): ", fmt.Sprintf("%f", gouse.AreaTri(10.0, 20.0)))
	println("Perimeter of triangle (integer): ", gouse.PeriTri(10, 20, 30))
	println("Perimeter of triangle (float): ", fmt.Sprintf("%f", gouse.PeriTri(10.0, 20.0, 30.0)))
}

/*
Description: Calculate absolute number
Input params: (num int)
*/
func Abs() {
	var num = -10
	println("Absolute number: ", gouse.Abs(num))
}

/*
Description: Add two numbers
Input params: (num1, num2 int)
*/
func Add() {
	var num1, num2 = 10, -2
	println("Add numbers: ", gouse.Add(num1, num2))
}

/*
Description: Calculate Ceil number
Input params: (num float64)
*/
func Ceil() {
	var num = 10.49
	println("Ceil number: ", gouse.Ceil(num))
}

/*
Description: Divide two numbers
Input params: (num1, num2 int)
*/
func Operators() {
	var num1, num2 = 10, -2
	println("Quotient of numbers: ", gouse.Div(num1, num2))
}

/*
Description: Calculate factorial number
Input params: (num int)
*/
func Factorial() {
	var num = 5
	println("Factorial of number: ", gouse.Factorial(num))
}

/*
Description: Calculate Floor number
Input params: (num float64)
*/
func Floor() {
	var num = 10.49
	println("Floor number: ", gouse.Floor(num))
}

/*
Description: Find Min, Max, Mean of numbers
Input params: (struct interface{}, fieldName string, value interface{})
*/
func MinMaxMean() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Min number: ", gouse.Min(num1, num2, num3, num4))
	println("Max number: ", gouse.Max(num1, num2, num3, num4))
	println("Average/Mean of numbers: ", gouse.Mean(num1, num2, num3, num4))
}

/*
Description: Multiply multiple numbers
Input params: (num ...int)
*/
func Multi() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Multiply of numbers: ", gouse.Multi(num1, num2, num3, num4))
}

/*
Description: Calculate power of number (F: float)
Input params: Powe2(num int), Pow(num, power int), PoweF(num, power float64)
*/
func Power() {
	println("Power square of number: ", gouse.Pow2(2))
	println("Power of integer numbers: ", gouse.Pow(2, 3))
	println("Power of float numbers: ", gouse.Pow(2.0, 3.0))
}

/*
Description: Calculate remainder of two numbers
Input params: (num1, num2 int)
*/
func Remainder() {
	var num1, num2 = 10, -2
	println("Remainder of numbers: ", gouse.Mod(num1, num2))
}

/*
Description: Calculate square-root, cube-root, nth-root of number (F: float)
Input params: Sqrt(num int), SqrtF(num float64), Cbrt(num int), CbrtF(num float64), Root(num, power int), RootF(num, power float64)
*/
func Root() {
	println("Square-Root of integer number: ", gouse.Sqrt(16))
	println("Square-Root of float number: ", fmt.Sprintf("%f", gouse.Sqrt(20.0)))

	println("Cube-Root of integer number: ", gouse.Cbrt(27))
	println("Cube-Root of float number: ", fmt.Sprintf("%f", gouse.Cbrt(20.0)))

	println("Nth-Root of integer number: ", gouse.Root(4, 2))
	println("Nth-Root of float number: ", fmt.Sprintf("%f", gouse.Root(20.0, 3.0)))
}

/*
Description: Calculate round number
Input params: (num float64)
*/
func Round() {
	var num1, num2 = 10.49, 10.51
	println("Round number: ", gouse.Round(num1), gouse.Round(num2))
}

/*
Description: Subtract two numbers
Input params: (num1, num2 int)
*/
func Sub() {
	var num1, num2 = 10, -2
	println("Subtract of numbers: ", gouse.Sub(num1, num2))
}

/*
Description: Sum multiple numbers
Input params: (num ...int)
*/
func Sum() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Sum of numbers: ", gouse.Sum(num1, num2, num3, num4))
}


# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Math' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Check number

Description: Check number is even, odd, perfect square, prime<br>Input params: (num int) number<br>

```go
func CheckNumber() {
	var num = 10
	println("Check even number: ", gouse.IsEven(num))
	println("Check odd number: ", gouse.IsOdd(num))
	println("Check perfect square number: ", gouse.IsPerfectSquare(num))
	println("Check prime number: ", gouse.IsPrime(num))
}
```

## 2. Log

Description: Calculate logarithm (base 2, 10) of number (F: float)<br>Input params: (num, base int)<br>

```go
func Log() {
	println("Logarithm of integer number: ", gouse.Log(16, 2))
	println("Logarithm of float number: ", fmt.Sprintf("%f", gouse.LogF(20.0, 2.0)))

	println("Logarithm of integer number (base 2): ", gouse.Log2(16))
	println("Logarithm of float number (base 2): ", fmt.Sprintf("%f", gouse.Log2F(20.0)))

	println("Logarithm of integer number (base 10): ", gouse.Log10(100))
	println("Logarithm of float number (base 10): ", fmt.Sprintf("%f", gouse.Log10F(20.0)))
}
```

## 3. Pytago

Description: Calculate Pytago (F: float)<br>Input params: (num1, num2 int)<br>

```go
func Pytago() {
	fmt.Printf("Pytago of number (integer): %d", gouse.Pytago(3, 4))
	fmt.Printf("Pytago of number (float): %f", gouse.Pytago(3.0, 4.0))
}
```

## 4. Transition

Description: Calculate transition (speed, distance, time)<br>Input params: Speed(distance, time float64), Distance(speed, time float64), Time(distance, speed float64)<br>

```go
func Transition() {
	var distance, speed, time float64 = 100, 10, 10
	fmt.Printf("Speed: %f", gouse.Speed(distance, time))
	fmt.Printf("Distance: %f", gouse.Distance(speed, time))
	fmt.Printf("Time: %f", gouse.Time(distance, speed))
}
```

## 5. Trigonometry

Description: Calculate trigonometry (sine, cosine, tangent) (F: float)<br>Input params: (angle float64)<br>

```go
func Trigonometry() {
	println("Sine of integer number: ", gouse.Sin(90))
	println("Sine of float number: ", fmt.Sprintf("%f", gouse.SinF(90.0)))
	println("Cosine of integer number: ", gouse.Cos(90))
	println("Cosine of float number: ", fmt.Sprintf("%f", gouse.CosF(90.0)))
	println("Tangent of integer number: ", gouse.Tan(90))
	println("Tangent of float number: ", fmt.Sprintf("%f", gouse.TanF(90.0)))
}
```

## 6. Circle

Description: Calculate area, perimeter, volume of circle (F: float)<br>Input params: (radius float64)<br>

```go
func Circle() {
	fmt.Printf("Area of circle (integer): %f", gouse.AreaCir(10))
	fmt.Printf("Area of circle (float): %f", gouse.AreaCir(10.0))
	fmt.Printf("Perimeter of circle (integer): %f", gouse.PeriCir(10))
	fmt.Printf("Perimeter of circle (float): %f", gouse.PeriCir(10.0))
}
```

## 7. Cone

Description: Calculate area, volume of cone (F: float)<br>Input params: (radius, height float64)<br>

```go
func Cone() {
	fmt.Printf("Area of cone (integer): %f", gouse.AreaCone(10, 20))
	fmt.Printf("Area of cone (float): %f", gouse.AreaCone(10.0, 20.0))
	fmt.Printf("Volume of cone (integer): %f", gouse.VolCone(10, 20))
	fmt.Printf("Volume of cone (float): %f", gouse.VolCone(10.0, 20.0))
}
```

## 8. Cube

Description: Calculate area, perimeter, volume of cube (F: float)<br>Input params: (side float64)<br>

```go
func Cube() {
	println("Area of cube (integer): ", gouse.AreaCube(10))
	println("Area of cube (float): ", fmt.Sprintf("%f", gouse.AreaCube(10.0)))
	println("Perimeter of cube (integer): ", gouse.PeriCube(10))
	println("Perimeter of cube (float): ", fmt.Sprintf("%f", gouse.PeriCube(10.0)))
	println("Volume of cube (integer): ", gouse.VolCube(10))
	println("Volume of cube (float): ", fmt.Sprintf("%f", gouse.VolCubeF(10.0)))
}
```

## 9. Cylinder

Description: Calculate area, volume of cylinder (F: float)<br>Input params: (radius, height float64)<br>

```go
func Cylinder() {
	fmt.Printf("Area of cylinder (integer): %f", gouse.AreaCyl(10, 20))
	fmt.Printf("Area of cylinder (float): %f", gouse.AreaCyl(10.0, 20.0))
	fmt.Printf("Volume of cylinder (integer): %f", gouse.VolCyl(10, 20))
	fmt.Printf("Volume of cylinder (float): %f", gouse.VolCyl(10.0, 20.0))
}
```

## 10. Ellipse

Description: Calculate area, perimeter, volume of ellipse (F: float)<br>Input params: (major, minor float64)<br>

```go
func Ellipse() {
	fmt.Printf("Area of ellipse (integer): %f", gouse.AreaEllipse(10, 20))
	fmt.Printf("Area of ellipse (float): %f", gouse.AreaEllipse(10.0, 20.0))
}
```

## 11. Parallelogram

Description: Calculate area, perimeter, volume of parallelogram (F: float)<br>Input params: (base, height float64)<br>

```go
func Parallelogram() {
	println("Area of parallelogram (integer): ", gouse.AreaParallelogram(10, 20))
	println("Area of parallelogram (float): ", fmt.Sprintf("%f", gouse.AreaParallelogram(10.0, 20.0)))
}
```

## 12. Polygon

Description: Calculate area of pentagon (F: float)<br>Input params: (side, apothem float64)<br>

```go
func Polygon() {
	fmt.Printf("Area of pentagol by number of sides (integer): %f", gouse.AreaPolygon(10, 6))
}
```

## 13. Rect

Description: Calculate area, perimeter, volume of rectangle (F: float)<br>Input params: (length, width float64)<br>

```go
func Rect() {
	println("Area of rectangle: ", gouse.AreaRect(10, 20))
	println("Perimeter of rectangle (integer): ", gouse.PeriRect(10, 20))
	println("Perimeter of rectangle (float): ", fmt.Sprintf("%f", gouse.PeriRect(10.0, 20.0)))
	println("Diagonal of rectangle (integer): ", fmt.Sprintf("%d", gouse.DiagRect(10, 20)))
	println("Diagonal of rectangle (float): ", fmt.Sprintf("%f", gouse.DiagRect(10.0, 20.0)))
	println("Volume of rectangular (integer): ", gouse.VolRect(10, 20, 30))
	println("Volume of rectangular (float): ", fmt.Sprintf("%f", gouse.VolRect(10.0, 20.0, 30.0)))
}
```

## 14. Rhombus

Description: Calculate area, perimeter, volume of rhombus (F: float)<br>Input params: (diagonal1, diagonal2 float64)<br>

```go
func Rhombus() {
	println("Area of rhombus (integer): ", gouse.AreaRhombus(10, 20))
	println("Area of rhombus (float): ", fmt.Sprintf("%f", gouse.AreaRhombus(10.0, 20.0)))
}
```

## 15. Sphere

Description: Calculate area, volume of sphere (F: float)<br>Input params: (radius float64)<br>

```go
func Sphere() {
	fmt.Printf("Area of sphere (integer): %f", gouse.AreaSphere(10))
	fmt.Printf("Area of sphere (float): %f", gouse.AreaSphereF(10.0))
	fmt.Printf("Volume of sphere (integer): %f", gouse.VolSphere(10))
	fmt.Printf("Volume of sphere (float): %f", gouse.VolSphereF(10.0))
}
```

## 16. Square

Description: Calculate area, perimeter, volume of square (F: float)<br>Input params: (side float64)<br>

```go
func Square() {
	println("Area of square (integer): ", gouse.AreaSquare(10))
	println("Area of square (float): ", fmt.Sprintf("%f", gouse.AreaSquare(10.0)))
	println("Perimeter of square (integer): ", gouse.PeriSquare(10))
	println("Perimeter of square (float): ", fmt.Sprintf("%f", gouse.PeriSquare(10.0)))
}
```

## 17. Trapezoid

Description: Calculate area, perimeter, volume of trapezoid (F: float)<br>Input params: (base1, base2, height float64)<br>

```go
func Trapezoid() {
	fmt.Printf("Area of trapezoid (integer): %f", gouse.AreaTrapezoid(10, 20, 30))
	fmt.Printf("Area of trapezoid (float): %f", gouse.AreaTrapezoid(10.0, 20.0, 30.0))
}
```

## 18. Triangle

Description: Calculate area, perimeter of triangle (F: float)<br>Input params: (base, height, side float64)<br>

```go
func Triangle() {
	println("Area of triangle (integer): ", gouse.AreaTri(10, 20))
	println("Area of triangle (float): ", fmt.Sprintf("%f", gouse.AreaTri(10.0, 20.0)))
	println("Perimeter of triangle (integer): ", gouse.PeriTri(10, 20, 30))
	println("Perimeter of triangle (float): ", fmt.Sprintf("%f", gouse.PeriTri(10.0, 20.0, 30.0)))
}
```

## 19. Abs

Description: Calculate absolute number<br>Input params: (num int)<br>

```go
func Abs() {
	var num = -10
	println("Absolute number: ", gouse.Abs(num))
}
```

## 20. Add

Description: Add two numbers<br>Input params: (num1, num2 int)<br>

```go
func Add() {
	var num1, num2 = 10, -2
	println("Add numbers: ", gouse.Add(num1, num2))
}
```

## 21. Ceil

Description: Calculate Ceil number<br>Input params: (num float64)<br>

```go
func Ceil() {
	var num = 10.49
	println("Ceil number: ", gouse.Ceil(num))
}
```

## 22. Operators

Description: Divide two numbers<br>Input params: (num1, num2 int)<br>

```go
func Operators() {
	var num1, num2 = 10, -2
	println("Quotient of numbers: ", gouse.Div(num1, num2))
}
```

## 23. Factorial

Description: Calculate factorial number<br>Input params: (num int)<br>

```go
func Factorial() {
	var num = 5
	println("Factorial of number: ", gouse.Factorial(num))
}
```

## 24. Floor

Description: Calculate Floor number<br>Input params: (num float64)<br>

```go
func Floor() {
	var num = 10.49
	println("Floor number: ", gouse.Floor(num))
}
```

## 25. Min max mean

Description: Find Min, Max, Mean of numbers<br>Input params: (struct interface{}, fieldName string, value interface{})<br>

```go
func MinMaxMean() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Min number: ", gouse.Min(num1, num2, num3, num4))
	println("Max number: ", gouse.Max(num1, num2, num3, num4))
	println("Average/Mean of numbers: ", gouse.Mean(num1, num2, num3, num4))
}
```

## 26. Multi

Description: Multiply multiple numbers<br>Input params: (num ...int)<br>

```go
func Multi() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Multiply of numbers: ", gouse.Multi(num1, num2, num3, num4))
}
```

## 27. Power

Description: Calculate power of number (F: float)<br>Input params: Powe2(num int), Pow(num, power int), PoweF(num, power float64)<br>

```go
func Power() {
	println("Power square of number: ", gouse.Pow2(2))
	println("Power of integer numbers: ", gouse.Pow(2, 3))
	println("Power of float numbers: ", gouse.Pow(2.0, 3.0))
}
```

## 28. Remainder

Description: Calculate remainder of two numbers<br>Input params: (num1, num2 int)<br>

```go
func Remainder() {
	var num1, num2 = 10, -2
	println("Remainder of numbers: ", gouse.Mod(num1, num2))
}
```

## 29. Root

Description: Calculate square-root, cube-root, nth-root of number (F: float)<br>Input params: Sqrt(num int), SqrtF(num float64), Cbrt(num int), CbrtF(num float64), Root(num, power int), RootF(num, power float64)<br>

```go
func Root() {
	println("Square-Root of integer number: ", gouse.Sqrt(16))
	println("Square-Root of float number: ", fmt.Sprintf("%f", gouse.Sqrt(20.0)))

	println("Cube-Root of integer number: ", gouse.Cbrt(27))
	println("Cube-Root of float number: ", fmt.Sprintf("%f", gouse.Cbrt(20.0)))

	println("Nth-Root of integer number: ", gouse.Root(4, 2))
	println("Nth-Root of float number: ", fmt.Sprintf("%f", gouse.Root(20.0, 3.0)))
}
```

## 30. Round

Description: Calculate round number<br>Input params: (num float64)<br>

```go
func Round() {
	var num1, num2 = 10.49, 10.51
	println("Round number: ", gouse.Round(num1), gouse.Round(num2))
}
```

## 31. Sub

Description: Subtract two numbers<br>Input params: (num1, num2 int)<br>

```go
func Sub() {
	var num1, num2 = 10, -2
	println("Subtract of numbers: ", gouse.Sub(num1, num2))
}
```

## 32. Sum

Description: Sum multiple numbers<br>Input params: (num ...int)<br>

```go
func Sum() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Sum of numbers: ", gouse.Sum(num1, num2, num3, num4))
}
```

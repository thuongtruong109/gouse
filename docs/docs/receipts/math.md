
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Math' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

### <Badge style='font-size: 1.1rem;' type='tip' text='1. math is even' />



```go
func MathIsEven() {
	var num = 10
	println("Check even number: ", gouse.IsEven(num))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='2. math is odd' />



```go
func MathIsOdd() {
	var num = 10
	println("Check odd number: ", gouse.IsOdd(num))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='3. math is perfect square' />



```go
func MathIsPerfectSquare() {
	var num = 10
	println("Check perfect square number: ", gouse.IsPerfectSquare(num))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='4. math is prime' />



```go
func MathIsPrime() {
	var num = 10
	println("Check prime number: ", gouse.IsPrime(num))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='5. math log' />



```go
func MathLog() {
	println("Logarithm of integer number: ", gouse.Log(16, 2))
	println("Logarithm of float number: ", fmt.Sprintf("%f", gouse.LogF(20.0, 2.0)))

	println("Logarithm of integer number (base 2): ", gouse.Log2(16))
	println("Logarithm of float number (base 2): ", fmt.Sprintf("%f", gouse.Log2F(20.0)))

	println("Logarithm of integer number (base 10): ", gouse.Log10(100))
	println("Logarithm of float number (base 10): ", fmt.Sprintf("%f", gouse.Log10F(20.0)))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='6. math pytago' />



```go
func MathPytago() {
	println("Pytago of number (integer): ", fmt.Sprintf("%f", gouse.Pytago(3, 4)))
	println("Pytago of number (float): ", fmt.Sprintf("%f", gouse.PytagoF(3.0, 4.0)))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='7. math transition' />



```go
func MathTransition() {
	var distance, speed, time float64 = 100, 10, 10
	println("Speed: ", fmt.Sprintf("%f", gouse.Speed(distance, time)))
	println("Distance: ", fmt.Sprintf("%f", gouse.Distance(speed, time)))
	println("Time: ", fmt.Sprintf("%f", gouse.Time(distance, speed)))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='8. math trigonometry' />



```go
func MathTrigonometry() {
	println("Sine of integer number: ", gouse.Sin(90))
	println("Sine of float number: ", fmt.Sprintf("%f", gouse.SinF(90.0)))
	println("Cosine of integer number: ", gouse.Cos(90))
	println("Cosine of float number: ", fmt.Sprintf("%f", gouse.CosF(90.0)))
	println("Tangent of integer number: ", gouse.Tan(90))
	println("Tangent of float number: ", fmt.Sprintf("%f", gouse.TanF(90.0)))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='9. math circle' />



```go
func MathCircle() {
	println("Area of circle (integer): ", fmt.Sprintf("%f", gouse.AreaCircle(10)))
	println("Area of circle (float): ", fmt.Sprintf("%f", gouse.AreaCircleF(10.0)))
	println("Perimeter of circle (integer): ", fmt.Sprintf("%f", gouse.PeriCircle(10)))
	println("Perimeter of circle (float): ", fmt.Sprintf("%f", gouse.PeriCircleF(10.0)))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='10. math cone' />



```go
func MathCone() {
	println("Area of cone (integer): ", fmt.Sprintf("%f", gouse.AreaCone(10, 20)))
	println("Area of cone (float): ", fmt.Sprintf("%f", gouse.AreaConeF(10.0, 20.0)))
	println("Volume of cone (integer): ", fmt.Sprintf("%f", gouse.VolCone(10, 20)))
	println("Volume of cone (float): ", fmt.Sprintf("%f", gouse.VolConeF(10.0, 20.0)))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='11. math cube' />



```go
func MathCube() {
	println("Area of cube (integer): ", gouse.AreaCube(10))
	println("Area of cube (float): ", fmt.Sprintf("%f", gouse.AreaCubeF(10.0)))
	println("Perimeter of cube (integer): ", gouse.PeriCube(10))
	println("Perimeter of cube (float): ", fmt.Sprintf("%f", gouse.PeriCubeF(10.0)))
	println("Volume of cube (integer): ", gouse.VolCube(10))
	println("Volume of cube (float): ", fmt.Sprintf("%f", gouse.VolCubeF(10.0)))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='12. math cylinder' />



```go
func MathCylinder() {
	println("Area of cylinder (integer): ", fmt.Sprintf("%f", gouse.AreaCylinder(10, 20)))
	println("Area of cylinder (float): ", fmt.Sprintf("%f", gouse.AreaCylinderF(10.0, 20.0)))
	println("Volume of cylinder (integer): ", fmt.Sprintf("%f", gouse.VolCylinder(10, 20)))
	println("Volume of cylinder (float): ", fmt.Sprintf("%f", gouse.VolCylinderF(10.0, 20.0)))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='13. math ellipse' />



```go
func MathEllipse() {
	println("Area of ellipse (integer): ", fmt.Sprintf("%f", gouse.AreaEllipse(10, 20)))
	println("Area of ellipse (float): ", fmt.Sprintf("%f", gouse.AreaEllipseF(10.0, 20.0)))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='14. math parallelogram' />



```go
func MathParallelogram() {
	println("Area of parallelogram (integer): ", gouse.AreaParallelogram(10, 20))
	println("Area of parallelogram (float): ", fmt.Sprintf("%f", gouse.AreaParallelogramF(10.0, 20.0)))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='15. math polygon' />



```go
func MathPolygon() {
	println("Area of pentagol by number of sides (integer): ", fmt.Sprintf("%f", gouse.AreaPolygon(10, 6)))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='16. math rect' />



```go
func MathRect() {
	println("Area of rectangle: ", gouse.AreaRect(10, 20))
	println("Perimeter of rectangle (integer): ", gouse.PeriRect(10, 20))
	println("Perimeter of rectangle (float): ", fmt.Sprintf("%f", gouse.PeriRectF(10.0, 20.0)))
	println("Diagonal of rectangle (integer): ", fmt.Sprintf("%f", gouse.DiagRect(10, 20)))
	println("Diagonal of rectangle (float): ", fmt.Sprintf("%f", gouse.DiagRectF(10.0, 20.0)))
	println("Volume of rectangular (integer): ", gouse.VolRect(10, 20, 30))
	println("Volume of rectangular (float): ", fmt.Sprintf("%f", gouse.VolRectF(10.0, 20.0, 30.0)))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='17. math rhombus' />



```go
func MathRhombus() {
	println("Area of rhombus (integer): ", gouse.AreaRhombus(10, 20))
	println("Area of rhombus (float): ", fmt.Sprintf("%f", gouse.AreaRhombusF(10.0, 20.0)))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='18. math sphere' />



```go
func MathSphere() {
	println("Area of sphere (integer): ", fmt.Sprintf("%f", gouse.AreaSphere(10)))
	println("Area of sphere (float): ", fmt.Sprintf("%f", gouse.AreaSphereF(10.0)))
	println("Volume of sphere (integer): ", fmt.Sprintf("%f", gouse.VolSphere(10)))
	println("Volume of sphere (float): ", fmt.Sprintf("%f", gouse.VolSphereF(10.0)))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='19. math square' />



```go
func MathSquare() {
	println("Area of square (integer): ", gouse.AreaSquare(10))
	println("Area of square (float): ", fmt.Sprintf("%f", gouse.AreaSquareF(10.0)))
	println("Perimeter of square (integer): ", gouse.PeriSquare(10))
	println("Perimeter of square (float): ", fmt.Sprintf("%f", gouse.PeriSquareF(10.0)))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='20. math trapezoid' />



```go
func MathTrapezoid() {
	println("Area of trapezoid (integer): ", fmt.Sprintf("%f", gouse.AreaTrapezoid(10, 20, 30)))
	println("Area of trapezoid (float): ", fmt.Sprintf("%f", gouse.AreaTrapezoidF(10.0, 20.0, 30.0)))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='21. math triangle' />



```go
func MathTriangle() {
	println("Area of triangle (integer): ", gouse.AreaTriangle(10, 20))
	println("Area of triangle (float): ", fmt.Sprintf("%f", gouse.AreaTriangleF(10.0, 20.0)))
	println("Perimeter of triangle (integer): ", gouse.PeriTriangle(10, 20, 30))
	println("Perimeter of triangle (float): ", fmt.Sprintf("%f", gouse.PeriTriangleF(10.0, 20.0, 30.0)))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='22. math abs' />



```go
func MathAbs() {
	var num = -10
	println("Absolute number: ", gouse.Abs(num))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='23. math add' />



```go
func MathAdd() {
	var num1, num2 = 10, -2
	println("Add numbers: ", gouse.Add(num1, num2))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='24. math ceil' />



```go
func MathCeil() {
	var num = 10.49
	println("Ceil number: ", gouse.Ceil(num))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='25. math operators' />



```go
func MathOperators() {
	var num1, num2 = 10, -2
	println("Quotient of numbers: ", gouse.Divide(num1, num2))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='26. math factorial' />



```go
func MathFactorial() {
	var num = 5
	println("Factorial of number: ", gouse.Factorial(num))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='27. math floor' />



```go
func MathFloor() {
	var num = 10.49
	println("Floor number: ", gouse.Floor(num))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='28. math min' />



```go
func MathMin() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Min number: ", gouse.Min(num1, num2, num3, num4))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='29. math max' />



```go
func MathMax() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Max number: ", gouse.Max(num1, num2, num3, num4))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='30. math mean' />



```go
func MathMean() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Average/Mean of numbers: ", gouse.Mean(num1, num2, num3, num4))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='31. math multi' />



```go
func MathMulti() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Multiply of numbers: ", gouse.Multi(num1, num2, num3, num4))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='32. math power' />



```go
func MathPower() {
	println("Power square of number: ", gouse.Pow2(2))
	println("Power of integer numbers: ", gouse.Pow(2, 3))
	println("Power of float numbers: ", gouse.PowF(2.0, 3.0))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='33. math remainder' />



```go
func MathRemainder() {
	var num1, num2 = 10, -2
	println("Remainder of numbers: ", gouse.Remainder(num1, num2))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='34. math root' />



```go
func MathRoot() {
	println("Square-Root of integer number: ", gouse.Sqrt(16))
	println("Square-Root of float number: ", fmt.Sprintf("%f", gouse.SqrtF(20.0)))

	println("Cube-Root of integer number: ", gouse.Cbrt(27))
	println("Cube-Root of float number: ", fmt.Sprintf("%f", gouse.CbrtF(20.0)))

	println("Nth-Root of integer number: ", gouse.Root(4, 2))
	println("Nth-Root of float number: ", fmt.Sprintf("%f", gouse.RootF(20.0, 3.0)))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='35. math round' />



```go
func MathRound() {
	var num1, num2 = 10.49, 10.51
	println("Round number: ", gouse.Round(num1), gouse.Round(num2))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='36. math sub' />



```go
func MathSub() {
	var num1, num2 = 10, -2
	println("Subtract of numbers: ", gouse.Sub(num1, num2))
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='37. math sum' />



```go
func MathSum() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Sum of numbers: ", gouse.Sum(num1, num2, num3, num4))
}
```

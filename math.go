package gouse

import "math"

/* Formulas */

func Log(number, base int) int {
	return int(math.Log(float64(number)) / math.Log(float64(base)))
}

func LogF(number, base float64) float64 {
	return math.Log(number) / math.Log(base)
}

func Log2(number int) int {
	return Log(number, 2)
}

func Log2F(number float64) float64 {
	return LogF(number, 2)
}

func Log10(number int) int {
	return Log(number, 10)
}

func Log10F(number float64) float64 {
	return LogF(number, 10)
}

func Pytago(side1, side2 int) float64 {
	return SqrtF(float64(Pow(side1, 2) + Pow(side2, 2)))
}

func PytagoF(side1, side2 float64) float64 {
	return SqrtF(float64(PowF(side1, 2) + PowF(side2, 2)))
}

func Sin(angle int) float64 {
	return math.Sin(float64(angle))
}

func SinF(angle float64) float64 {
	return math.Sin(angle)
}

func Cos(angle int) float64 {
	return math.Cos(float64(angle))
}

func CosF(angle float64) float64 {
	return math.Cos(angle)
}

func Tan(angle int) float64 {
	return math.Tan(float64(angle))
}

func TanF(angle float64) float64 {
	return math.Tan(angle)
}

func Speed(distance, time float64) float64 {
	return DivideF(distance, time)
}

func Distance(speed, time float64) float64 {
	return MultiF(speed, time)
}

func Time(distance, speed float64) float64 {
	return DivideF(distance, speed)
}

/* Functions */

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func AbsF(num float64) float64 {
	if num < 0 {
		return -num
	}
	return num
}

// down to the nearest integer
func Floor(num float64) int {
	return int(num)
}

// up to the nearest integer
func Ceil(num float64) int {
	if num == float64(int(num)) {
		return int(num)
	}
	if num > 0 {
		return int(num) + 1
	}
	return int(num) - 1
}

// round to the nearest integer
func Round(num float64) int {
	if num < 0 {
		return int(num)
	}

	if num-float64(int(num)) < 0.5 {
		return int(num)
	}

	return int(num) + 1
}

func Max(nums ...int) int {
	var max int
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func MaxF(nums ...float64) float64 {
	var max float64
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func Min(nums ...int) int {
	min := nums[0]

	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}

func MinF(nums ...float64) float64 {
	min := nums[0]

	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}

func Sum(nums ...int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}

func SumF(nums ...float64) float64 {
	var sum float64
	for _, v := range nums {
		sum += v
	}
	return sum
}

func Add(num1, num2 int) int {
	return num1 + num2
}

func AddF(num1, num2 float64) float64 {
	return num1 + num2
}

func Sub(num1, num2 int) int {
	return num1 - num2
}

func SubF(num1, num2 float64) float64 {
	return num1 - num2
}

func Multi(nums ...int) int {
	product := 1
	for _, v := range nums {
		product *= v
	}
	return product
}

func MultiF(nums ...float64) float64 {
	product := 1.0
	for _, v := range nums {
		product *= v
	}
	return product
}

func Divide(num1, num2 int) int {
	return num1 / num2
}

func DivideF(num1, num2 float64) float64 {
	return num1 / num2
}

func Remainder(num1, num2 int) int {
	return num1 % num2
}

func Mean(nums ...int) int {
	return Sum(nums...) / len(nums)
}

func MeanF(nums ...float64) float64 {
	return SumF(nums...) / float64(len(nums))
}

func Pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func PowF(base, exp float64) float64 {
	result := 1.0
	for i := 0; i < int(exp); i++ {
		result *= base
	}
	return result
}

func Pow2(base int) int {
	return Multi(base, base)
}

func Pow2F(base float64) float64 {
	return MultiF(base, base)
}

func Pow3(base int) int {
	return Multi(base, base, base)
}

func Pow3F(base float64) float64 {
	return MultiF(base, base, base)
}

func Factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * Factorial(num-1)
}

func Root(number, n int) int {
	return int(math.Pow(float64(number), 1.0/float64(n)))
}

func RootF(number, n float64) float64 {
	return math.Pow(number, 1.0/n)
}

func Sqrt(number int) int {
	return int(math.Sqrt(float64(number)))
}

func SqrtF(number float64) float64 {
	return math.Sqrt(number)
}

func Cbrt(number int) int {
	return Root(number, 3)
}

func CbrtF(number float64) float64 {
	return RootF(number, 3)
}

/* Geometry */

func AreaRect(length, width int) int {
	return Multi(length, width)
}

func AreaRectF(length, width float64) float64 {
	return MultiF(length, width)
}

func PeriRect(length, width int) int {
	return Multi(2, Add(length, width))
}

func PeriRectF(length, width float64) float64 {
	return MultiF(2, AddF(length, width))
}

func DiagRect(length, width int) float64 {
	return Pytago(length, width)
}

func DiagRectF(length, width float64) float64 {
	return PytagoF(length, width)
}

func VolRect(length, width, height int) int {
	return Multi(length, width, height)
}

func VolRectF(length, width, height float64) float64 {
	return MultiF(length, width, height)
}

func AreaCircle(radius int) float64 {
	return math.Pi * float64(Pow2(radius))
}

func AreaCircleF(radius float64) float64 {
	return math.Pi * PowF(radius, 2)
}

func PeriCircle(radius int) float64 {
	return 2 * math.Pi * float64(radius)
}

func PeriCircleF(radius float64) float64 {
	return 2 * math.Pi * radius
}

func AreaTriangle(base, height int) int {
	return Divide(Multi(base, height), 2)
}

func AreaTriangleF(base, height float64) float64 {
	return DivideF(MultiF(base, height), 2)
}

func PeriTriangle(side1, side2, side3 int) int {
	return Sum(side1, side2, side3)
}

func PeriTriangleF(side1, side2, side3 float64) float64 {
	return SumF(side1, side2, side3)
}

func AreaSquare(side int) int {
	return Pow2(side)
}

func AreaSquareF(side float64) float64 {
	return PowF(side, 2)
}

func PeriSquare(side int) int {
	return Multi(side, 4)
}

func PeriSquareF(side float64) float64 {
	return MultiF(side, 4)
}

func AreaCube(side int) int {
	return Multi(Pow2(side), 6)
}

func AreaCubeF(side float64) float64 {
	return MultiF(PowF(side, 2), 6)
}

func PeriCube(side int) int {
	return Multi(side, 12)
}

func PeriCubeF(side float64) float64 {
	return MultiF(side, 12)
}

func VolCube(side int) int {
	return Pow3(side)
}

func VolCubeF(side float64) float64 {
	return PowF(side, 3)
}

func AreaSphere(radius int) float64 {
	return 4 * math.Pi * float64(Pow2(radius))
}

func AreaSphereF(radius float64) float64 {
	return 4 * math.Pi * PowF(radius, 2)
}

func VolSphere(radius int) float64 {
	return (4 / 3) * math.Pi * float64(Pow2(radius))
}

func VolSphereF(radius float64) float64 {
	return (4 / 3) * math.Pi * PowF(radius, 2)
}

func AreaCylinder(radius, height int) float64 {
	return (2 * math.Pi * float64(radius)) * float64(height)
}

func AreaCylinderF(radius, height float64) float64 {
	return (2 * math.Pi * radius) * height
}

func VolCylinder(radius, height int) float64 {
	return math.Pi * float64(Pow2(radius)) * float64(height)
}

func VolCylinderF(radius, height float64) float64 {
	return math.Pi * PowF(radius, 2) * height
}

func AreaCone(radius, height int) float64 {
	return math.Pi * float64(radius) * (float64(radius) + math.Sqrt(float64(Pow2(height)+Pow2(radius))))
}

func AreaConeF(radius, height float64) float64 {
	return math.Pi * radius * (radius + math.Sqrt(PowF(height, 2)+PowF(radius, 2)))
}

func VolCone(radius, height int) float64 {
	return DivideF(1, 3) * math.Pi * float64(Pow2(radius)) * float64(height)
}

func VolConeF(radius, height float64) float64 {
	return DivideF(1, 3) * math.Pi * PowF(radius, 2) * height
}

func AreaTrapezoid(base1, base2, height int) float64 {
	return 0.5 * float64(Add(base1, base2)) * float64(height)
}

func AreaTrapezoidF(base1, base2, height float64) float64 {
	return 0.5 * (base1 + base2) * height
}

func AreaParallelogram(base, height int) int {
	return Multi(base, height)
}

func AreaParallelogramF(base, height float64) float64 {
	return MultiF(base, height)
}

func AreaRhombus(diag1, diag2 int) int {
	return Divide(Multi(diag1, diag2), 2)
}

func AreaRhombusF(diag1, diag2 float64) float64 {
	return DivideF(MultiF(diag1, diag2), 2)
}

func AreaEllipse(major, minor int) float64 {
	return math.Pi * float64(major) * float64(minor)
}

func AreaEllipseF(major, minor float64) float64 {
	return math.Pi * major * minor
}

func AreaPolygon(lenSide float64, numSide int) float64 {
	return DivideF(MultiF(0.25, float64(numSide), Pow2F(lenSide)), AbsF(TanF(DivideF(180, float64(numSide)))))
}

/* Checks */

func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func IsEven(num int) bool {
	return num%2 == 0
}

func IsOdd(num int) bool {
	return num%2 != 0
}

func IsPerfectSquare(num int) bool {
	for i := 1; i <= num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}

/* Utilities */

// func SumBy(nums []int, fn func(int) int) int {
// 	var sum int
// 	for _, v := range nums {
// 		sum += fn(v)
// 	}
// 	return sum
// }

// func AverageBy(nums []int, fn func(int) int) int {
// 	return SumBy(nums, fn) / len(nums)
// }

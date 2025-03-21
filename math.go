package gouse

import (
	"math"

	"golang.org/x/exp/constraints"
)

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

func Pytago[T int | float64](side1, side2 T) T {
	return T(math.Sqrt(float64(side1*side1 + side2*side2)))
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
	return Div(distance, time)
}

func Distance(speed, time float64) float64 {
	return Multi(speed, time)
}

func Time(distance, speed float64) float64 {
	return Div(distance, speed)
}

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

func Max[T int | float64](nums ...T) T {
	var max T
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func Min[T int | float64](nums ...T) T {
	min := nums[0]

	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}

func Sum[T int | float64](nums ...T) T {
	var sum T
	for _, v := range nums {
		sum += v
	}
	return sum
}

func Add[T int | float64](num1, num2 T) T {
	return num1 + num2
}

func Sub[T int | float64](num1, num2 T) T {
	return num1 - num2
}

func Multi[T int | float64](nums ...T) T {
	var product T = 1
	for _, v := range nums {
		product *= v
	}
	return product
}

func Div[T int | float64](num1, num2 T) T {
	return num1 / num2
}

func Mod(num1, num2 int) int {
	return num1 % num2
}

func ModF(num1, num2 float64) float64 {
	return math.Mod(num1, num2)
}

func Pow[T int | float64](base, exp T) T {
	var result T = 1
	switch exp := any(exp).(type) {
	case int:
		for i := 0; i < exp; i++ {
			result *= base
		}
	case float64:
		for i := 0.0; i < exp; i++ {
			result *= base
		}
	}
	return result
}

func Pow2[T int | float64](base T) T {
	return Multi(base, base)
}

func Pow2F(base float64) float64 {
	return Multi(base, base)
}

func Pow3(base int) int {
	return Multi(base, base, base)
}

func Pow3F(base float64) float64 {
	return Multi(base, base, base)
}

func Factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * Factorial(num-1)
}

func Root[T int | float64](number, n T) T {
	var result T
	switch num := any(number).(type) {
	case int:
		result = T(int(math.Pow(float64(num), 1.0/float64(n))))
	case float64:
		result = T(math.Pow(num, 1.0/float64(n)))
	}
	return result
}

func Sqrt[T int | float64](number T) T {
	var result T
	switch num := any(number).(type) {
	case int:
		result = T(int(math.Sqrt(float64(num))))
	case float64:
		result = T(math.Sqrt(num))
	}
	return result
}

func Cbrt[T int | float64](number T) T {
	return Root(number, 3)
}

func Mean[T int | float64](nums ...T) T {
	return Sum(nums...) / T(len(nums))
}

func Median[T int | float64](nums ...T) T {
	SortNum(nums)
	n := len(nums)

	if n%2 == 0 {
		return (nums[n/2-1] + nums[n/2]) / 2
	}

	return nums[n/2]
}

func Mode[T int | float64](nums ...T) T {
	if len(nums) == 0 {
		var zero T
		return zero
	}

	counts := make(map[T]int)
	for _, v := range nums {
		counts[v]++
	}

	var mode T
	maxCount := 0

	for k, v := range counts {
		if v > maxCount {
			mode = k
			maxCount = v
		}
	}

	return mode
}

func AreaRect[T int | float64](length, width T) T {
	return Multi(length, width)
}

func PeriRect[T int | float64](length, width T) T {
	return Multi(2, Add(length, width))
}

func VolRect[T int | float64](length, width, height T) T {
	return Multi(length, width, height)
}

func DiagRect[T int | float64](length, width T) T {
	return Pytago(length, width)
}

func AreaCir[T int | float64](radius T) float64 {
	return math.Pi * float64(Pow2(radius))
}

func PeriCir[T int | float64](radius T) float64 {
	return 2 * math.Pi * float64(radius)
}

func AreaTri[T int | float64](base, height T) T {
	return Div(Multi(base, height), 2)
}

func PeriTri[T int | float64](side1, side2, side3 T) T {
	return Sum(side1, side2, side3)
}

func AreaSquare[T int | float64](side T) T {
	return Pow2(side)
}

func PeriSquare[T int | float64](side T) T {
	return Multi(side, 4)
}

func DiagSquare[T constraints.Integer | constraints.Float](side T) float64 {
	return math.Sqrt(float64(side)*float64(side) + float64(side)*float64(side))
}

func AreaCube[T int | float64](side T) T {
	return Multi(Pow2(side), 6)
}

func PeriCube[T int | float64](side T) T {
	return Multi(side, 12)
}

func DiagCube(side int) float64 {
	return Sqrt(float64(Pow3(side) + Pow3(side)))
}

func DiagCubeF(side float64) float64 {
	return Sqrt(Pow(side, 3) + Pow(side, 3))
}

func VolCube(side int) int {
	return Pow3(side)
}

func VolCubeF(side float64) float64 {
	return Pow(side, 3)
}

func AreaSphere(radius int) float64 {
	return 4 * math.Pi * float64(Pow2(radius))
}

func AreaSphereF(radius float64) float64 {
	return 4 * math.Pi * Pow(radius, 2)
}

func VolSphere(radius int) float64 {
	return (4 / 3) * math.Pi * float64(Pow2(radius))
}

func VolSphereF(radius float64) float64 {
	return (4 / 3) * math.Pi * Pow2(radius)
}

func AreaCyl[T constraints.Integer | constraints.Float](radius, height T) float64 {
	return 2 * math.Pi * float64(radius) * float64(height)
}

func VolCyl[T constraints.Integer | constraints.Float](radius, height T) float64 {
	return math.Pi * Pow2(float64(radius)) * float64(height)
}

func AreaCone[T constraints.Integer | constraints.Float](radius, height T) float64 {
	return math.Pi * float64(radius) * (float64(radius) + math.Sqrt(float64((radius*radius)+(height*height))))
}

func VolCone[T constraints.Integer | constraints.Float](radius, height T) float64 {
	return (1.0 / 3.0) * math.Pi * Pow2(float64(radius)) * float64(height)
}

func AreaTrapezoid[T constraints.Integer | constraints.Float](base1, base2, height T) float64 {
	return 0.5 * float64(base1+base2) * float64(height)
}

func AreaParallelogram[T int | float64](base, height T) T {
	return Multi(base, height)
}

func AreaRhombus[T int | float64](diag1, diag2 T) T {
	return Div(Multi(diag1, diag2), 2)
}

func AreaEllipse[T constraints.Integer | constraints.Float](major, minor T) float64 {
	return math.Pi * float64(major) * float64(minor)
}

func AreaPolygon(lenSide float64, numSide int) float64 {
	return Div(Multi(0.25, float64(numSide), Pow2F(lenSide)), AbsF(TanF(Div(180, float64(numSide)))))
}

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
		if Pow2(i) == num {
			return true
		}
	}
	return false
}

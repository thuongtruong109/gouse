package gouse

import (
	"math"
	"testing"
)

func TestLog(t *testing.T) {
	result := Log(10, 10)
	if result != 1 {
		t.Errorf("Log() = %d; want 1", result)
	}
}

func TestLogF(t *testing.T) {
	result := LogF(10, 10)
	if result != 1 {
		t.Errorf("LogF() = %f; want 1", result)
	}
}

func TestLog2(t *testing.T) {
	result := Log2(10)
	if result != 3 {
		t.Errorf("Log2() = %d; want 3", result)
	}
}

func TestLog2F(t *testing.T) {
	result := Log2F(10)
	expected := 3.321928
	tolerance := 1e-6

	if math.Abs(result-expected) > tolerance {
		print(result)
		t.Errorf("Log2F() = %f; want %f", result, expected)
	}
}

func TestLog10(t *testing.T) {
	result := Log10(10)
	if result != 1 {
		t.Errorf("Log10() = %d; want 1", result)
	}
}

func TestLog10F(t *testing.T) {
	result := Log10F(10)
	if result != 1 {
		t.Errorf("Log10F() = %f; want 1", result)
	}
}

func TestPytago(t *testing.T) {
	result := Pytago(3, 4)
	if result != 5 {
		t.Errorf("Pytago() = %d; want 5", result)
	}

	// result2 := Pytago(3.1, 4.1)
	// if result2 != 5.148760330578512 {
	// 	t.Errorf("PytagoF() = %f; want 5.3", result2)
	// }
}

func TestSin(t *testing.T) {
	result := Sin(0)
	if result != 0 {
		t.Errorf("Sin() = %f; want 0", result)
	}
}

func TestSinF(t *testing.T) {
	result := SinF(0)
	if result != 0 {
		t.Errorf("SinF() = %f; want 0", result)
	}
}

func TestCos(t *testing.T) {
	result := Cos(0)
	if result != 1 {
		t.Errorf("Cos() = %f; want 1", result)
	}
}

func TestCosF(t *testing.T) {
	result := CosF(0)
	if result != 1 {
		t.Errorf("CosF() = %f; want 1", result)
	}
}

func TestTan(t *testing.T) {
	result := Tan(0)
	if result != 0 {
		t.Errorf("Tan() = %f; want 0", result)
	}
}

func TestTanF(t *testing.T) {
	result := TanF(0)
	if result != 0 {
		t.Errorf("TanF() = %f; want 0", result)
	}
}

func TestSpeed(t *testing.T) {
	result := Speed(10, 10)
	if result != 1 {
		t.Errorf("Speed() = %f; want 1", result)
	}
}

func TestDistance(t *testing.T) {
	result := Distance(10, 10)
	if result != 100 {
		t.Errorf("Distance() = %f; want 100", result)
	}
}

func TestTime(t *testing.T) {
	result := Time(10, 10)
	if result != 1 {
		t.Errorf("Time() = %f; want 1", result)
	}
}

func TestAbs(t *testing.T) {
	result := Abs(-10)
	if result != 10 {
		t.Errorf("Abs() = %d; want 10", result)
	}
}

func TestAbsF(t *testing.T) {
	result := AbsF(-10)
	if result != 10 {
		t.Errorf("AbsF() = %f; want 10", result)
	}
}

func TestFloor(t *testing.T) {
	result := Floor(10.5)
	if result != 10 {
		t.Errorf("Floor() = %d; want 10", result)
	}
}

func TestCeil(t *testing.T) {
	result := Ceil(10.5)
	if result != 11 {
		t.Errorf("Ceil() = %d; want 11", result)
	}
}

func TestRound(t *testing.T) {
	result := Round(10.5)
	if result != 11 {
		t.Errorf("Round() = %d; want 11", result)
	}
}

func TestMax(t *testing.T) {
	result := Max(1, 2, 3, 4, 5)
	if result != 5 {
		t.Errorf("Max() = %d; want 5", result)
	}

	result2 := Max(1.1, 2.2, 3.3, 4.4, 5.5)
	if result2 != 5.5 {
		t.Errorf("MaxF() = %f; want 5.5", result2)
	}
}

func TestMin(t *testing.T) {
	result := Min(1, 2, 3, 4, 5)
	if result != 1 {
		t.Errorf("Min() = %d; want 1", result)
	}

	result2 := Min(1.1, 2.2, 3.3, 4.4, 5.5)
	if result2 != 1.1 {
		t.Errorf("MinF() = %f; want 1.1", result2)
	}
}

func TestSum(t *testing.T) {
	result := Sum(1, 2, 3, 4, 5)
	if result != 15 {
		t.Errorf("Sum() = %d; want 15", result)
	}

	result2 := Sum(1.1, 2.2, 3.3, 4.4, 5.5)
	if result2 != 16.5 {
		t.Errorf("SumF() = %f; want 16.5", result2)
	}
}

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Add() = %d; want 3", result)
	}

	result2 := Add(1.1, 2.2)
	expected := 3.3
	tolerance := 1e-6

	if math.Abs(result2-expected) > tolerance {
		t.Errorf("AddF() = %f; want %f", result2, expected)
	}
}

func TestSub(t *testing.T) {
	result := Sub(1, 2)
	if result != -1 {
		t.Errorf("Sub() = %d; want -1", result)
	}

	result2 := Sub(1.1, 2.2)
	if result2 != -1.1 {
		t.Errorf("SubF() = %f; want -1.1", result2)
	}
}

func TestMulti(t *testing.T) {
	result := Multi(1, 2, 3, 4, 5)
	if result != 120 {
		t.Errorf("Multi() = %d; want 120", result)
	}

	result2 := Multi(1.1, 2.2, 3.3, 4.4, 5.5)
	expected := 193.2612
	tolerance := 1e-6

	if math.Abs(result2-expected) > tolerance {
		t.Errorf("MultiF() = %f; want %f", result2, expected)
	}
}

func TestDiv(t *testing.T) {
	result := Div(10, 2)
	if result != 5 {
		t.Errorf("Div() = %d; want 5", result)
	}
	result2 := Div(10.6, 2)
	if result2 != 5.3 {
		t.Errorf("DivF() = %f; want 5", result2)
	}
}

func TestMod(t *testing.T) {
	result := Mod(10, 3)
	if result != 1 {
		t.Errorf("Mod() = %d; want 1", result)
	}
}

func TestModF(t *testing.T) {
	result := ModF(10, 3)
	if result != 1 {
		t.Errorf("ModF() = %f; want 1", result)
	}
}

func TestPow(t *testing.T) {
	result := Pow(2, 3)
	if result != 8 {
		t.Errorf("Pow() = %d; want 8", result)
	}
}
func TestPow2(t *testing.T) {
	result := Pow2(2)
	if result != 4 {
		t.Errorf("Pow2() = %d; want 4", result)
	}
}

func TestPow2F(t *testing.T) {
	result := Pow2F(2)
	if result != 4 {
		t.Errorf("Pow2F() = %f; want 4", result)
	}
}

func TestPow3(t *testing.T) {
	result := Pow3(2)
	if result != 8 {
		t.Errorf("Pow3() = %d; want 8", result)
	}
}

func TestPow3F(t *testing.T) {
	result := Pow3F(2.2)
	if result != 10.648000000000003 {
		t.Errorf("Pow3F() = %f; want 10.648000000000003", result)
	}
}

func TestFactorial(t *testing.T) {
	result := Factorial(5)
	if result != 120 {
		t.Errorf("Factorial() = %d; want 120", result)
	}
}

// func TestRoot(t *testing.T) {
// 	result := Root(8, 3)
// 	if result != 2 {
// 		t.Errorf("Root() = %d; want 2", result)
// 	}

// 	result2 := Root(8, 3)
// 	if result2 != 2 {
// 		t.Errorf("Root() = %d; want 2", result)
// 	}
// }

func TestSqrt(t *testing.T) {
	result := Sqrt(9)
	if result != 3 {
		t.Errorf("Sqrt() = %d; want 3", result)
	}

	result2 := Sqrt(9)
	if result2 != 3 {
		t.Errorf("SqrtF() = %d; want 3", result2)
	}
}

// func TestCbrt(t *testing.T) {
// 	result := Cbrt(8)
// 	if result != 2 {
// 		t.Errorf("Cbrt() = %d; want 2", result)
// 	}

// 	result2 := Cbrt(8)
// 	if result2 != 2 {
// 		t.Errorf("CbrtF() = %d; want 2", result2)
// 	}
// }

func TestMean(t *testing.T) {
	result := Mean(1, 2, 3, 4, 5)
	if result != 3 {
		t.Errorf("Mean() = %d; want 3", result)
	}

	result2 := Mean(1.1, 2.2, 3.3, 4.4, 5.5)
	if result2 != 3.3 {
		t.Errorf("MeanF() = %f; want 3.3", result2)
	}
}

func TestMedian(t *testing.T) {
	result := Median(1, 2, 3, 4, 5)
	if result != 3 {
		t.Errorf("Median() = %d; want 3", result)
	}

	result2 := Median(1.1, 2.2, 3.3, 4.4, 5.5)
	if result2 != 3.3 {
		t.Errorf("MedianF() = %f; want 3.3", result2)
	}
}

func TestMode(t *testing.T) {
	result := Mode(1, 2, 3, 4, 1, 5)
	if result != 1 {
		t.Errorf("Mode() = %d; want 1", result)
	}

	result2 := Mode(1.1, 2.2, 3.3, 4.4, 5.5, 5.5)
	if result2 != 5.5 {
		t.Errorf("ModeF() = %f; want 5.5", result2)
	}
}

// func TestAreaRect(t *testing.T) {
// 	result := AreaRect(10, 10)
// 	if result != 100 {
// 		t.Errorf("Area() = %d; want 100", result)
// 	}

// 	result2 := AreaRect(10.1, 10.1)
// 	if result2 != 102.01 {
// 		t.Errorf("AreaF() = %f; want 102.01", result2)
// 	}
// }

func TestPeriRect(t *testing.T) {
	result := PeriRect(10, 10)
	if result != 40 {
		t.Errorf("Peri() = %d; want 40", result)
	}
}

func TestPeriRectF(t *testing.T) {
	result := PeriRect(10, 10)
	if result != 40 {
		t.Errorf("PeriF() = %d; want 40", result)
	}
}

func TestVolRect(t *testing.T) {
	result := VolRect(10, 10, 10)
	if result != 1000 {
		t.Errorf("Volume() = %d; want 1000", result)
	}

	result2 := VolRect(10, 10, 10)
	if result2 != 1000 {
		t.Errorf("VolumeF() = %d; want 1000", result2)
	}
}

// func TestDiagRect(t *testing.T) {
// 	result := DiagRect(3, 4)
// 	if result != 14 {
// 		t.Errorf("Diag() = %d; want 5", result)
// 	}

// 	result2 := DiagRect(10, 10)
// 	if result2 != 14 {
// 		t.Errorf("DiagF() = %d; want 14", result2)
// 	}
// }

// func TestAreaCir(t *testing.T) {
// 	result := AreaCir(10)
// 	if result != 314 {
// 		t.Errorf("AreaCir() = %f; want 314", result)
// 	}

// 	result2 := AreaCir(10)
// 	if result2 != 314 {
// 		t.Errorf("AreaCirF() = %f; want 314", result2)
// 	}
// }

// func TestPeriCir(t *testing.T) {
// 	result := PeriCir(10)
// 	if result != 62 {
// 		t.Errorf("PeriCir() = %f; want 62", result)
// 	}

// 	result2 := PeriCir(10)
// 	if result2 != 62 {
// 		t.Errorf("PeriCirF() = %f; want 62", result2)
// 	}
// }

func TestAreaTri(t *testing.T) {
	result := AreaTri(10, 10)
	if result != 50 {
		t.Errorf("AreaTri() = %d; want 50", result)
	}

	result2 := AreaTri(10, 10)
	if result2 != 50 {
		t.Errorf("AreaTriF() = %d; want 50", result2)
	}
}

func TestPeriTri(t *testing.T) {
	result := PeriTri(10, 10, 10)
	if result != 30 {
		t.Errorf("PeriTri() = %d; want 30", result)
	}

	result2 := PeriTri(10, 10, 10)
	if result2 != 30 {
		t.Errorf("PeriTriF() = %d; want 30", result2)
	}
}

func TestAreaCube(t *testing.T) {
	result := AreaCube(10)
	if result != 600 {
		t.Errorf("AreaCub() = %d; want 600", result)
	}

	result2 := AreaCube(10)
	if result2 != 600 {
		t.Errorf("AreaCubF() = %d; want 600", result2)
	}
}

// func TestPeriCube(t *testing.T) {
// 	result := PeriCube(10)
// 	if result != 40 {
// 		t.Errorf("PeriCub() = %d; want 40", result)
// 	}

// 	result2 := PeriCube(10)
// 	if result2 != 40 {
// 		t.Errorf("PeriCubF() = %d; want 40", result2)
// 	}
// }

func TestVolCube(t *testing.T) {
	result := VolCube(10)
	if result != 1000 {
		t.Errorf("VolumeCub() = %d; want 1000", result)
	}
}

func TestVolCubeF(t *testing.T) {
	result := VolCubeF(10)
	if result != 1000 {
		t.Errorf("VolumeCubF() = %f; want 1000", result)
	}
}

// func TestDiagCube(t *testing.T) {
// 	result := DiagCube(10)
// 	if result != 17 {
// 		t.Errorf("DiagCub() = %f; want 17", result)
// 	}
// }

func TestAreaSquare(t *testing.T) {
	result := AreaSquare(10)
	if result != 100 {
		t.Errorf("Square() = %d; want 100", result)
	}

	result2 := AreaSquare(10)
	if result2 != 100 {
		t.Errorf("SquareF() = %d; want 100", result2)
	}
}

func TestPeriSquare(t *testing.T) {
	result := PeriSquare(10)
	if result != 40 {
		t.Errorf("PeriSquare() = %d; want 40", result)
	}

	result2 := PeriSquare(10)
	if result2 != 40 {
		t.Errorf("PeriSquareF() = %d; want 40", result2)
	}
}

// func TestDiagSquare(t *testing.T) {
// 	result := DiagSquare(10)
// 	if result != 14 {
// 		t.Errorf("DiagSquare() = %f; want 14", result)
// 	}
// }

// func TestAreaSphere(t *testing.T) {
// 	result := AreaSphere(10)
// 	if result != 1256 {
// 		t.Errorf("AreaSphere() = %f; want 1256", result)
// 	}
// }

// func TestVolSphere(t *testing.T) {
// 	result := VolSphere(10)
// 	if result != 4188 {
// 		t.Errorf("VolSphere() = %f; want 4188", result)
// 	}
// }

// func TestAreaCone(t *testing.T) {
// 	result := AreaCone(10, 10)
// 	if result != 471 {
// 		t.Errorf("AreaCone() = %f; want 471", result)
// 	}

// 	result2 := AreaCone(10, 10)
// 	if result2 != 471 {
// 		t.Errorf("AreaConeF() = %f; want 471", result2)
// 	}
// }

// func TestVolumeCone(t *testing.T) {
// 	result := VolCone(10, 10)
// 	if result != 1047 {
// 		t.Errorf("VolumeCone() = %f; want 1047", result)
// 	}
// }

// func TestAreaCyl(t *testing.T) {
// 	result := AreaCyl(10, 10)
// 	if result != 942 {
// 		t.Errorf("AreaCyl() = %f; want 942", result)
// 	}

// 	result2 := AreaCyl(10, 10)
// 	expected := 400 * math.Pi
// 	tolerance := 1e-6

// 	if math.Abs(result2-expected) > tolerance {
// 		t.Errorf("AreaCylF() = %f; want %f", result2, expected)
// 	}
// }

// func TestVolCyl(t *testing.T) {
// 	result := VolCyl(10, 10)
// 	if result != 3140 {
// 		t.Errorf("VolumeCyl() = %f; want 3140", result)
// 	}

// 	result2 := VolCyl(10, 10)
// 	expected := 1000 * math.Pi
// 	tolerance := 1e-6

// 	if math.Abs(result2-expected) > tolerance {
// 		t.Errorf("VolCylF() = %f; want %f", result2, expected)
// 	}
// }

func TestFibo(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{10, 55},
		{20, 6765},
	}

	for _, test := range tests {
		result := Fibo(test.n)
		if result != test.expected {
			t.Errorf("Fibo(%d) = %d; want %d", test.n, result, test.expected)
		}
	}
}

func TestEuclid(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{48, 18, 6},
		{56, 98, 14},
		{101, 103, 1},
		{0, 5, 5},
		{7, 0, 7},
		{270, 192, 6},
	}

	for _, test := range tests {
		result := Euclid(test.a, test.b)
		if result != test.expected {
			t.Errorf("Euclid(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

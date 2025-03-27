package tests

import (
	. "github.com/michaelzhao820/raytracer/raytracer"
	"math"
	"testing"
)

func TestMatrix4x4Construction(t *testing.T) {
	m := NewMatrix(4, 4)

	// Populate the matrix with the specified values
	testValues := []struct {
		row, col int
		value    float64
	}{
		{0, 0, 1},
		{0, 1, 2},
		{0, 2, 3},
		{0, 3, 4},
		{1, 0, 5.5},
		{1, 1, 6.5},
		{1, 2, 7.5},
		{1, 3, 8.5},
		{2, 0, 9},
		{2, 1, 10},
		{2, 2, 11},
		{2, 3, 12},
		{3, 0, 13.5},
		{3, 1, 14.5},
		{3, 2, 15.5},
		{3, 3, 16.5},
	}

	// Set each value
	for _, tv := range testValues {
		var err error
		m, err = m.Set(tv.row, tv.col, tv.value)
		if err != nil {
			t.Errorf("Failed to set value at [%d,%d]: %v", tv.row, tv.col, err)
		}
	}

	// Verify each value
	for _, tv := range testValues {
		val, err := m.Get(tv.row, tv.col)
		if err != nil {
			t.Errorf("Failed to get value at [%d,%d]: %v", tv.row, tv.col, err)
		}
		if val != tv.value {
			t.Errorf("Incorrect value at [%d,%d]. Expected %f, got %f",
				tv.row, tv.col, tv.value, val)
		}
	}
}

func TestMatrix2x2Construction(t *testing.T) {
	m := NewMatrix(2, 2)

	testValues := []struct {
		row, col int
		value    float64
	}{
		{0, 0, -3},
		{0, 1, 5},
		{1, 0, 1},
		{1, 1, -2},
	}

	// Set each value
	for _, tv := range testValues {
		var err error
		m, err = m.Set(tv.row, tv.col, tv.value)
		if err != nil {
			t.Errorf("Failed to set value at [%d,%d]: %v", tv.row, tv.col, err)
		}
	}

	// Verify each value
	for _, tv := range testValues {
		val, err := m.Get(tv.row, tv.col)
		if err != nil {
			t.Errorf("Failed to get value at [%d,%d]: %v", tv.row, tv.col, err)
		}
		if val != tv.value {
			t.Errorf("Incorrect value at [%d,%d]. Expected %f, got %f",
				tv.row, tv.col, tv.value, val)
		}
	}
}

func TestMatrix3x3Construction(t *testing.T) {
	m := NewMatrix(3, 3)

	testValues := []struct {
		row, col int
		value    float64
	}{
		{0, 0, -3},
		{0, 1, 5},
		{0, 2, 0},
		{1, 0, 1},
		{1, 1, -2},
		{1, 2, -7},
		{2, 0, 0},
		{2, 1, 1},
		{2, 2, 1},
	}

	// Set each value
	for _, tv := range testValues {
		var err error
		m, err = m.Set(tv.row, tv.col, tv.value)
		if err != nil {
			t.Errorf("Failed to set value at [%d,%d]: %v", tv.row, tv.col, err)
		}
	}

	// Verify each value
	for _, tv := range testValues {
		val, err := m.Get(tv.row, tv.col)
		if err != nil {
			t.Errorf("Failed to get value at [%d,%d]: %v", tv.row, tv.col, err)
		}
		if val != tv.value {
			t.Errorf("Incorrect value at [%d,%d]. Expected %f, got %f",
				tv.row, tv.col, tv.value, val)
		}
	}
}

func TestMatrixEquality(t *testing.T) {
	// Test Case 1: Identical Matrices
	testCase1 := []struct {
		name     string
		matrixA  Matrix
		matrixB  Matrix
		expected bool
	}{
		{
			name: "Identical 4x4 Matrices",
			matrixA: func() Matrix {
				m := NewMatrix(4, 4)
				testData := []float64{
					1, 2, 3, 4,
					5, 6, 7, 8,
					9, 8, 7, 6,
					5, 4, 3, 2,
				}
				for i, val := range testData {
					m.Set(i/4, i%4, val)
				}
				return m
			}(),
			matrixB: func() Matrix {
				m := NewMatrix(4, 4)
				testData := []float64{
					1, 2, 3, 4,
					5, 6, 7, 8,
					9, 8, 7, 6,
					5, 4, 3, 2,
				}
				for i, val := range testData {
					m.Set(i/4, i%4, val)
				}
				return m
			}(),
			expected: true,
		},
	}

	for _, tc := range testCase1 {
		t.Run(tc.name, func(t *testing.T) {
			if result := tc.matrixA.Equals(tc.matrixB); result != tc.expected {
				t.Errorf("Expected matrices to be equal, but got %v", result)
			}
		})
	}

	// Test Case 2: Different Matrices
	testCase2 := []struct {
		name     string
		matrixA  Matrix
		matrixB  Matrix
		expected bool
	}{
		{
			name: "Different 4x4 Matrices",
			matrixA: func() Matrix {
				m := NewMatrix(4, 4)
				testData := []float64{
					1, 2, 3, 4,
					5, 6, 7, 8,
					9, 8, 7, 6,
					5, 4, 3, 2,
				}
				for i, val := range testData {
					m, _ = m.Set(i/4, i%4, val)
				}
				return m
			}(),
			matrixB: func() Matrix {
				m := NewMatrix(4, 4)
				testData := []float64{
					2, 3, 4, 5,
					6, 7, 8, 9,
					8, 7, 6, 5,
					4, 3, 2, 1,
				}
				for i, val := range testData {
					m.Set(i/4, i%4, val)
				}
				return m
			}(),
			expected: false,
		},
	}

	for _, tc := range testCase2 {
		t.Run(tc.name, func(t *testing.T) {
			if result := tc.matrixA.Equals(tc.matrixB); result != tc.expected {
				t.Errorf("Expected matrices to be different, but got %v", result)
			}
		})
	}
}

func TestMatrixMultiplication(t *testing.T) {

	t.Run("Incompatible Matrix Dimensions", func(t *testing.T) {
		a := NewMatrix(2, 3)
		b := NewMatrix(4, 3)

		a, err := a.MultiplyMatrices(b)
		if err == nil {
			t.Errorf("Expected error for incompatible matrix dimensions")
		}
	})

	t.Run("4x4 Matrix Multiplication", func(t *testing.T) {
		// Create 4x4 matrices with test data
		a := NewMatrix(4, 4)
		b := NewMatrix(4, 4)

		// Populate matrices with test values
		testDataA := []float64{
			1, 2, 3, 4,
			5, 6, 7, 8,
			9, 10, 11, 12,
			13, 14, 15, 16,
		}
		testDataB := []float64{
			17, 18, 19, 20,
			21, 22, 23, 24,
			25, 26, 27, 28,
			29, 30, 31, 32,
		}

		for i, val := range testDataA {
			a, _ = a.Set(i/4, i%4, val)
		}
		for i, val := range testDataB {
			b, _ = b.Set(i/4, i%4, val)
		}

		result, err := a.MultiplyMatrices(b)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// You would calculate the expected result manually or with a separate method
		expectedData := []float64{
			250, 260, 270, 280,
			618, 644, 670, 696,
			986, 1028, 1070, 1112,
			1354, 1412, 1470, 1528,
		}

		expected := NewMatrix(4, 4)
		for i, val := range expectedData {
			expected, _ = expected.Set(i/4, i%4, val)
		}

		if !result.Equals(expected) {
			t.Errorf("4x4 Matrix multiplication result incorrect")
		}
	})
}

func TestMatrixMultiplyWithTuple(t *testing.T) {
	// Test Case: Matrix multiplication with tuple
	t.Run("Matrix Multiplication with Tuple", func(t *testing.T) {
		// Create matrix A
		a := NewMatrix(4, 4)
		matrixData := []float64{
			1, 2, 3, 4,
			2, 4, 4, 2,
			8, 6, 4, 1,
			0, 0, 0, 1,
		}

		// Populate matrix A
		for i, val := range matrixData {
			var err error
			if a, err = a.Set(i/4, i%4, val); err != nil {
				t.Fatalf("Error setting matrix value: %v", err)
			}
		}

		// Create tuple b
		b := NewTuple(0, 0, 0, 0)
		b[0] = 1
		b[1] = 2
		b[2] = 3
		b[3] = 1

		// Expected result
		expectedResult := NewTuple(0, 0, 0, 0)
		expectedResult[0] = 18
		expectedResult[1] = 24
		expectedResult[2] = 33
		expectedResult[3] = 1

		// Perform multiplication
		result, err := a.MultiplyWithTuple(b)
		if err != nil {
			t.Fatalf("Unexpected error during multiplication: %v", err)
		}

		// Compare results
		for i := 0; i < 4; i++ {
			if math.Abs(result[i]-expectedResult[i]) > 1e-6 {
				t.Errorf("Mismatch at index %d: expected %f, got %f",
					i, expectedResult[i], result[i])
			}
		}
	})
}

func TestMatrixInverse(t *testing.T) {
	// Helper function to compare float64 values with a small margin of error
	closeEnough := func(a, b float64) bool {
		epsilon := 1e-5
		return math.Abs(a-b) < epsilon
	}

	t.Run("First Inverse Matrix Scenario", func(t *testing.T) {
		// Create the input matrix
		A := NewMatrix(4, 4)
		A, _ = A.Set(0, 0, 8)
		A, _ = A.Set(0, 1, -5)
		A, _ = A.Set(0, 2, 9)
		A, _ = A.Set(0, 3, 2)
		A, _ = A.Set(1, 0, 7)
		A, _ = A.Set(1, 1, 5)
		A, _ = A.Set(1, 2, 6)
		A, _ = A.Set(1, 3, 1)
		A, _ = A.Set(2, 0, -6)
		A, _ = A.Set(2, 1, 0)
		A, _ = A.Set(2, 2, 9)
		A, _ = A.Set(2, 3, 6)
		A, _ = A.Set(3, 0, -3)
		A, _ = A.Set(3, 1, 0)
		A, _ = A.Set(3, 2, -9)
		A, _ = A.Set(3, 3, -4)

		// Expected inverse matrix
		expectedInverse := NewMatrix(4, 4)
		expectedInverse, _ = expectedInverse.Set(0, 0, -0.15385)
		expectedInverse, _ = expectedInverse.Set(0, 1, -0.15385)
		expectedInverse, _ = expectedInverse.Set(0, 2, -0.28205)
		expectedInverse, _ = expectedInverse.Set(0, 3, -0.53846)
		expectedInverse, _ = expectedInverse.Set(1, 0, -0.07692)
		expectedInverse, _ = expectedInverse.Set(1, 1, 0.12308)
		expectedInverse, _ = expectedInverse.Set(1, 2, 0.02564)
		expectedInverse, _ = expectedInverse.Set(1, 3, 0.03077)
		expectedInverse, _ = expectedInverse.Set(2, 0, 0.35897)
		expectedInverse, _ = expectedInverse.Set(2, 1, 0.35897)
		expectedInverse, _ = expectedInverse.Set(2, 2, 0.43590)
		expectedInverse, _ = expectedInverse.Set(2, 3, 0.92308)
		expectedInverse, _ = expectedInverse.Set(3, 0, -0.69231)
		expectedInverse, _ = expectedInverse.Set(3, 1, -0.69231)
		expectedInverse, _ = expectedInverse.Set(3, 2, -0.76923)
		expectedInverse, _ = expectedInverse.Set(3, 3, -1.92308)

		// Calculate the actual inverse
		actualInverse, err := A.Inverse()
		if err != nil {
			t.Fatalf("Error calculating inverse: %v", err)
		}

		// Compare each element of the inverse matrix
		for row := 0; row < 4; row++ {
			for col := 0; col < 4; col++ {
				actualVal, _ := actualInverse.Get(row, col)
				expectedVal, _ := expectedInverse.Get(row, col)
				if !closeEnough(actualVal, expectedVal) {
					t.Errorf("Mismatch at (%d, %d): got %f, want %f",
						row, col, actualVal, expectedVal)
				}
			}
		}
	})

	t.Run("Second Inverse Matrix Scenario", func(t *testing.T) {
		// Create the input matrix
		A := NewMatrix(4, 4)
		A, _ = A.Set(0, 0, 9)
		A, _ = A.Set(0, 1, 3)
		A, _ = A.Set(0, 2, 0)
		A, _ = A.Set(0, 3, 9)
		A, _ = A.Set(1, 0, -5)
		A, _ = A.Set(1, 1, -2)
		A, _ = A.Set(1, 2, -6)
		A, _ = A.Set(1, 3, -3)
		A, _ = A.Set(2, 0, -4)
		A, _ = A.Set(2, 1, 9)
		A, _ = A.Set(2, 2, 6)
		A, _ = A.Set(2, 3, 4)
		A, _ = A.Set(3, 0, -7)
		A, _ = A.Set(3, 1, 6)
		A, _ = A.Set(3, 2, 6)
		A, _ = A.Set(3, 3, 2)

		// Expected inverse matrix
		expectedInverse := NewMatrix(4, 4)
		expectedInverse, _ = expectedInverse.Set(0, 0, -0.04074)
		expectedInverse, _ = expectedInverse.Set(0, 1, -0.07778)
		expectedInverse, _ = expectedInverse.Set(0, 2, 0.14444)
		expectedInverse, _ = expectedInverse.Set(0, 3, -0.22222)
		expectedInverse, _ = expectedInverse.Set(1, 0, -0.07778)
		expectedInverse, _ = expectedInverse.Set(1, 1, 0.03333)
		expectedInverse, _ = expectedInverse.Set(1, 2, 0.36667)
		expectedInverse, _ = expectedInverse.Set(1, 3, -0.33333)
		expectedInverse, _ = expectedInverse.Set(2, 0, -0.02901)
		expectedInverse, _ = expectedInverse.Set(2, 1, -0.14630)
		expectedInverse, _ = expectedInverse.Set(2, 2, -0.10926)
		expectedInverse, _ = expectedInverse.Set(2, 3, 0.12963)
		expectedInverse, _ = expectedInverse.Set(3, 0, 0.17778)
		expectedInverse, _ = expectedInverse.Set(3, 1, 0.06667)
		expectedInverse, _ = expectedInverse.Set(3, 2, -0.26667)
		expectedInverse, _ = expectedInverse.Set(3, 3, 0.33333)

		// Calculate the actual inverse
		actualInverse, err := A.Inverse()
		if err != nil {
			t.Fatalf("Error calculating inverse: %v", err)
		}

		// Compare each element of the inverse matrix
		for row := 0; row < 4; row++ {
			for col := 0; col < 4; col++ {
				actualVal, _ := actualInverse.Get(row, col)
				expectedVal, _ := expectedInverse.Get(row, col)
				if !closeEnough(actualVal, expectedVal) {
					t.Errorf("Mismatch at (%d, %d): got %f, want %f",
						row, col, actualVal, expectedVal)
				}
			}
		}
	})

	t.Run("Multiply Matrix by Its Inverse", func(t *testing.T) {
		// Create matrix A
		A := NewMatrix(4, 4)
		A, _ = A.Set(0, 0, 3)
		A, _ = A.Set(0, 1, -9)
		A, _ = A.Set(0, 2, 7)
		A, _ = A.Set(0, 3, 3)
		A, _ = A.Set(1, 0, 3)
		A, _ = A.Set(1, 1, -8)
		A, _ = A.Set(1, 2, 2)
		A, _ = A.Set(1, 3, -9)
		A, _ = A.Set(2, 0, -4)
		A, _ = A.Set(2, 1, 4)
		A, _ = A.Set(2, 2, 4)
		A, _ = A.Set(2, 3, 1)
		A, _ = A.Set(3, 0, -6)
		A, _ = A.Set(3, 1, 5)
		A, _ = A.Set(3, 2, -1)
		A, _ = A.Set(3, 3, 1)

		// Create matrix B
		B := NewMatrix(4, 4)
		B, _ = B.Set(0, 0, 8)
		B, _ = B.Set(0, 1, 2)
		B, _ = B.Set(0, 2, 2)
		B, _ = B.Set(0, 3, 2)
		B, _ = B.Set(1, 0, 3)
		B, _ = B.Set(1, 1, -1)
		B, _ = B.Set(1, 2, 7)
		B, _ = B.Set(1, 3, 0)
		B, _ = B.Set(2, 0, 7)
		B, _ = B.Set(2, 1, 0)
		B, _ = B.Set(2, 2, 5)
		B, _ = B.Set(2, 3, 4)
		B, _ = B.Set(3, 0, 6)
		B, _ = B.Set(3, 1, -2)
		B, _ = B.Set(3, 2, 0)
		B, _ = B.Set(3, 3, 5)

		// Multiply A and B
		C, err := A.MultiplyMatrices(B)
		if err != nil {
			t.Fatalf("Error multiplying matrices: %v", err)
		}

		// Get the inverse of B
		BInverse, err := B.Inverse()
		if err != nil {
			t.Fatalf("Error calculating inverse of B: %v", err)
		}

		// Multiply C by B's inverse
		result, err := C.MultiplyMatrices(BInverse)
		if err != nil {
			t.Fatalf("Error multiplying C by B's inverse: %v", err)
		}

		// Check if the result is close to A
		if !result.Equals(A) {
			t.Errorf("Result of multiplying C by B's inverse is not equal to A")
		}
	})

	t.Run("Non-Invertible Matrix", func(t *testing.T) {
		// Create a singular matrix (determinant is zero)
		A := NewMatrix(4, 4)
		for i := 0; i < 4; i++ {
			A, _ = A.Set(i, i, 0)
		}

		// Attempt to calculate inverse
		_, err := A.Inverse()
		if err == nil {
			t.Errorf("Expected an error for non-invertible matrix, got none")
		}
	})

	t.Run("Identity Matrix Inverse", func(t *testing.T) {
		// Create a 4x4 identity matrix
		A := NewMatrix(4, 4)
		for i := 0; i < 4; i++ {
			A, _ = A.Set(i, i, 1.0)
		}

		// Calculate inverse
		inverseA, err := A.Inverse()
		if err != nil {
			t.Fatalf("Error calculating inverse of identity matrix: %v", err)
		}

		// Check if the inverse is the same as the original matrix
		if !inverseA.Equals(A) {
			t.Errorf("Inverse of identity matrix is not the identity matrix")
		}
	})
}

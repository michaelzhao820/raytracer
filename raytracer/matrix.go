package raytracer

import "fmt"

type Matrix struct {
	width  int
	height int
	data   []float64
}

func NewMatrix(width, height int) Matrix {
	return Matrix{
		width:  width,
		height: height,
		data:   make([]float64, width*height),
	}
}

// Set
// Currently deciding to do a functional approach to Set, added memory
// overhead is limited as the matrices do not seem like they will get
// too large, may be incorrect.
// /*
func (m Matrix) Set(row, col int, val float64) (Matrix, error) {
	newMatrix := NewMatrix(m.width, m.height)
	copy(newMatrix.data, m.data)

	index := row*m.width + col
	if index < 0 || index >= len(newMatrix.data) {
		return Matrix{}, fmt.Errorf("index is not correct")
	}

	newMatrix.data[index] = val
	return newMatrix, nil
}

func (m Matrix) Get(row, col int) (float64, error) {

	index := row*m.width + col
	if index < 0 || index >= len(m.data) {
		return 0, fmt.Errorf("index is not correct")
	}
	return m.data[index], nil
}

func (m Matrix) Equals(other Matrix) bool {
	if m.width != other.width || m.height != other.height {
		return false
	}
	for i := range m.data {
		if !equalsWithMargin(m.data[i], other.data[i]) {
			return false
		}
	}
	return true
}

func (m Matrix) MultiplyMatrices(other Matrix) (Matrix, error) {
	// Check if matrices can be multiplied
	if m.width != other.height {
		return Matrix{}, fmt.Errorf("matrix dimensions are incompatible for multiplication: "+
			"first matrix width (%d) must equal second matrix height (%d)", m.width, other.height)
	}

	n := NewMatrix(m.height, other.width)

	for row := 0; row < m.height; row++ {
		for col := 0; col < other.width; col++ {
			var totalSum float64
			for k := 0; k < m.width; k++ {
				totalSum += m.data[row*m.width+k] * other.data[k*other.width+col]
			}
			var err error
			if n, err = n.Set(row, col, totalSum); err != nil {
				return Matrix{}, err
			}
		}
	}
	return n, nil
}

func (m Matrix) MultiplyWithTuple(other Tuple) (Tuple, error) {

	if m.width != len(other) {
		return Tuple{}, fmt.Errorf("matrix width (%d) must equal tuple length (%d)", m.width, len(other))
	}

	n := make(Tuple, m.height)

	for row := 0; row < m.height; row++ {
		var totalSum float64
		for col := 0; col < m.width; col++ {
			matrixIndex := row*m.width + col
			totalSum += m.data[matrixIndex] * other[col]
		}
		n[row] = totalSum
	}
	return n, nil
}

func (m Matrix) MultiplyByIdentity() (Matrix, error) {

	n := NewMatrix(m.height, m.height)
	for i := 0; i < n.width; i++ {
		n.Set(i, i, 1.0)
	}
	k, _ := m.MultiplyMatrices(n)
	return k, nil
}

func (m Matrix) Transpose() (Matrix, error) {

	n := NewMatrix(m.width, m.height)

	for row := 0; row < m.height; row++ {
		for col := 0; col < m.width; col++ {
			val, _ := m.Get(row, col)
			var err error
			if n, err = n.Set(col, row, val); err != nil {
				return Matrix{}, fmt.Errorf("error setting transposed value at (%d, %d): %v", col, row, err)
			}
		}
	}
	return n, nil
}

func (m Matrix) Inverse() (Matrix, error) {
	det, _ := m.determinant()

	if equalsWithMargin(det, 0.0) {
		return Matrix{}, fmt.Errorf("matrix is not invertible (determinant is zero)")
	}

	// finding the cofactor matrix and dividing by determinant
	newMatrix := NewMatrix(m.width, m.height)
	for i := 0; i < newMatrix.height; i++ {
		for j := 0; j < newMatrix.width; j++ {
			cofactor, _ := m.cofactor(i, j)
			newMatrix, _ = newMatrix.Set(i, j, cofactor/det)
		}
	}

	newMatrix, _ = newMatrix.Transpose()
	return newMatrix, nil
}

func (m Matrix) determinantOf2x2() (float64, error) {
	if m.width != 2 || m.height != 2 {
		return 0.0, fmt.Errorf("matrix dimensions must be 2x2")
	}
	a, _ := m.Get(0, 0)
	b, _ := m.Get(0, 1)
	c, _ := m.Get(1, 0)
	d, _ := m.Get(1, 1)
	return a*d - b*c, nil
}

func (m Matrix) submatrix(row, column int) (Matrix, error) {

	if row < 0 || row >= m.height || column < 0 || column >= m.width {
		return Matrix{}, fmt.Errorf("invalid row (%d) or column (%d) index", row, column)
	}

	rowindex, colindex := 0, 0
	n := NewMatrix(m.width-1, m.height-1)
	for i := 0; i < m.height; i++ {
		if i == row {
			continue
		}
		colindex = 0
		for j := 0; j < m.width; j++ {
			if j == column {
				continue
			}
			val, _ := m.Get(i, j)
			n, _ = n.Set(rowindex, colindex, val)
			colindex++
		}
		rowindex++
	}
	return n, nil
}

func (m Matrix) minor(row, column int) (float64, error) {
	subMatrix, err := m.submatrix(row, column)
	if err != nil {
		return 0, fmt.Errorf("error creating submatrix: %v", err)
	}
	det, err := subMatrix.determinant()
	if err != nil {
		return 0, fmt.Errorf("error calculating minor: %v", err)
	}

	return det, nil
}

func (m Matrix) cofactor(row, column int) (float64, error) {
	minor, err := m.minor(row, column)
	if err != nil {
		return 0, err
	}
	if (row+column)%2 == 1 {
		return -minor, nil
	}
	return minor, nil
}

func (m Matrix) determinant() (float64, error) {
	if m.width == 2 && m.height == 2 {
		return m.determinantOf2x2()
	}
	var determinant float64
	for i := 0; i < m.width; i++ {
		cofactor, _ := m.cofactor(0, i)
		value, _ := m.Get(0, i)
		determinant += cofactor * value
	}
	return determinant, nil
}

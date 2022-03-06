package complexMatrix

type mutable [][]complex128

// Creates a mutable matrix from a 2-d array of complex numbers
func NewMutable(table [][]complex128) M {
	if len(table) == 0 {
		return nil
	}
	for _, row := range table[1:] {
		if len(row) != len(table[0]) {
			panic("complexMatrix.NewMutable parameter error: values are not in rectangular shape")
		}
	}
	return mutable(table)
}

// Creates a mutable matrix from 2-d arrays of real and imaginary parts
func CombineIntoMutable(real [][]float64, imag [][]float64) M {
	return NewMutable(combine(real, imag))
}

// Returns an mutable matrix with given dimensions filled with zero values
func EmptyMutable(height int, width int) M {
	n := make(mutable, 0, height)
	for len(n) < height {
		row := make([]complex128, width)
		n = append(n, row)
	}
	return n
}

func (m mutable) copy() mutable {
	if m == nil {
		return nil
	}
	n := make(mutable, 0, len(m))
	for _, row := range m {
		nrow := make([]complex128, len(row))
		copy(nrow, row)
		n = append(n, nrow)
	}
	return n
}

func (m mutable) Dim() (int, int) {
	if m == nil || len(m) == 0 {
		return 0, 0
	}
	return len(m), len(m[0])
}

func (m mutable) Get(i int, j int) complex128 {
	if m == nil {
		return 0
	}
	return m[i][j]
}

func (m mutable) Set(c complex128, i int, j int) M {
	m[i][j] = c
	return m
}

func (m mutable) Scale(v complex128) M {
	for i, row := range m {
		for j := range row {
			m[i][j] = m[i][j] * v
		}
	}
	return m
}

func (m mutable) Add(o M) M {
	{
		mR, mC := m.Dim()
		oR, oC := o.Dim()
		if mR != oR || mC != oC {
			panic("dimesion mismatch, can only add matricies of the same dimentions")
		}
	}
	for i, row := range m {
		for j := range row {
			m[i][j] = m[i][j] + o.Get(i, j)
		}
	}
	return m
}

func (m mutable) Transpose() M {
	return transpose{
		wrap: m,
	}
}

func (m mutable) Dot(B M) M {
	return dot(m, B, m)
}

func (m mutable) Build(d [][]complex128) M {
	return NewMutable(d)
}

func (m mutable) String() string {
	return SPrintCustom(m, "(", "), ", ", ")
}

func (m mutable) Map(f func(v complex128, r int, c int) complex128) M {
	for i, row := range m {
		for j := range row {
			m[i][j] = f(m[i][j], i, j)
		}
	}
	return m
}

func (m mutable) Resize(R int, C int) M {
	n := make(mutable, R)
	mR, mC := m.Dim()
	for i := range n {
		n[i] = make([]complex128, C)
		for j := range n[i] {
			if i >= mR || j >= mC {
				n[i][j] = 0
			} else {
				n[i][j] = m[i][j]
			}
		}
	}
	return n
}

func (m mutable) Immutable() M {
	return NewImmutable(m)
}

func (m mutable) Mutable() M {
	return m.copy()
}

package complexMatrix

import "testing"

func TestEqual(t *testing.T) {
	vals := make([]*complex128, 0, 10)
	for v := float64(1.0); v <= 10.0; v++ {
		temp := new(complex128)
		*temp = complex(v, v)
		vals = append(vals, temp)
	}
	a := &immutable{
		R: 2,
		C: 2,
		Data: [][]*complex128{
			[]*complex128{vals[0], vals[1]},
			[]*complex128{vals[2], vals[3]},
		},
	}
	if !Equal(a, a) {
		t.Errorf("reflexive equality check failed")
	}
	b := &immutable{
		R: 3,
		C: 1,
		Data: [][]*complex128{
			[]*complex128{vals[4]},
			[]*complex128{vals[5]},
			[]*complex128{vals[6]},
		},
	}
	if Equal(a, b) {
		t.Errorf("Failed to distinguish difference in dimensions")
	}
	c := &immutable{
		R: 2,
		C: 2,
		Data: [][]*complex128{
			[]*complex128{vals[7], vals[8]},
			[]*complex128{vals[9], vals[0]},
		},
	}
	if Equal(a, c) {
		t.Errorf("Failed to check values")
	}
}

func TestImmutable(t *testing.T) {
	a := NewImmutable([][]complex128{
		[]complex128{complex(1, 1), complex(2, 2)},
		[]complex128{complex(3, 3), complex(4, 4)},
	})
	if !Equal(a, a) || a.Get(0, 1) != complex(2, 2) {
		t.Errorf("Incorrect immutable construction")
	}
	t.Run("ConstructOnlyRectangles", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("non rectangle not caught")
			}
		}()
		NewImmutable([][]complex128{
			[]complex128{complex(1, 1), complex(2, 2)},
			[]complex128{complex(3, 3)},
		})
	})
	if !Equal(a, a.(*immutable).copy()) {
		t.Errorf("copy failed to create replica")
	}
	t.Logf("a = %v", a)
	b := a.Set(complex(5, 5), 0, 0)
	if b.Get(0, 0) != complex(5, 5) {
		t.Logf("b = %v", b)
		t.Errorf("Set failed to update")
	}
	if a.Get(0, 0) == complex(5, 5) {
		t.Logf("a = %v", a)
		t.Errorf("Set altered original")
	}
	c := a.Scale(complex(2, 1))
	if a.Get(0, 0) != complex(1, 1) {
		t.Errorf("Scale altered original, a = %v", a)
	}
	if c.Get(0, 0) != complex(1, 3) {
		t.Errorf("Scale incorrect result, c = %v", c)
	}
	d := a.Add(c)
	if a.Get(0, 0) != complex(1, 1) {
		t.Errorf("Add altered original, a = %v", a)
	}
	if d.Get(0, 0) != complex(2, 4) {
		t.Errorf("Add incorrect result, d = %v", d)
	}
	e := a.Transpose()
	if a.Get(0, 0) != complex(1, 1) {
		t.Errorf("Transpose altered original, a = %v", a)
	}
	if e.Get(0, 0) != a.Get(0, 0) || e.Get(0, 1) != a.Get(1, 0) {
		t.Errorf("Transpose incorrect, e = %v", e)
	}
	f := a.(*immutable).Build([][]complex128{
		[]complex128{complex(1, 0)},
		[]complex128{complex(0, 1)},
	})
	if a.Get(0, 0) != complex(1, 1) {
		t.Errorf("Build altered original, a = %v", a)
	}
	if fR, fC := f.Dim(); fR != 2 || fC != 1 || f.Get(0, 0) != complex(1, 0) {
		t.Errorf("Build incorrect result, f = %v", f)
	}
	g := a.Dot(f)
	if a.Get(0, 0) != complex(1, 1) {
		t.Errorf("Dot altered original, a = %v", a)
	}
	if gR, gC := g.Dim(); gR != 2 || gC != 1 || g.Get(0, 0) != complex(-1, 3) {
		t.Errorf("Dot incorrect result, g = %v", g)
	}
}

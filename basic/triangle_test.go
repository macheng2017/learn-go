package test_triangle

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 171},
		{12, 35, 0},
		{30000, 40000, 50000},
	}

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); tt.c != actual {
			t.Errorf("calcTriangle(%d,%d); got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

//=== RUN   TestTriangle
//--- FAIL: TestTriangle (0.00s)
//triangle_test.go:16: calcTriangle(8,15); got 17; expected 171
//triangle_test.go:16: calcTriangle(12,35); got 37; expected 0
//FAIL

package newmath

import "testing"

func TestSqrt(t *testing.T) {
    const in, out = 4, 2
    if x := Sqrt(in); x != out {
        t.Errorf("Sqrt(%v) = %v, want %v", in, x, out )
    }
}

func TestSqrtMore( t *testing.T) {
    var x float64 = 1
    for x <= 12 {
        if x != Sqrt(x*x) {
            t.Errorf("Sqrt(%v) = %v, want %v", x*x, Sqrt(x*x), x )
        }
        x++
    }
}

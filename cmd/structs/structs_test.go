package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{23.20, 77.00}
	got := Perimeter(rectangle)
	want := 200.40

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{23.00, 77.00}
		got := rectangle.Area()
		want := 1771.00

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793
		checkArea(t, got, want)
	})
}

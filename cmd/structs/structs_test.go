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

// func TestArea(t *testing.T) {
// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}
// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{23.00, 77.00}
// 		want := 1771.00
// 		checkArea(t, rectangle, want)
// 	})
// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		want := 314.1592653589793
// 		checkArea(t, circle, want)
// 	})
// }

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{23.5, 79.6}, want: 1870.6},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{10.6, 13.8}, want: 73.14},
	}

	for _, test := range areaTests {
		got := test.shape.Area()
		if got != test.want {
			t.Errorf("got %g want %g", got, test.want)
		}
	}
}

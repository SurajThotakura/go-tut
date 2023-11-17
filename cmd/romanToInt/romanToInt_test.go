package romanToInt

import "testing"

func TestRomanToInt(t *testing.T) {
	t.Run("caseOne", func(t *testing.T) {
		roman := "LVIII"

		got := romanToInt(roman)
		want := 58
		validate(t, got, want)
	})
	t.Run("caseTwo", func(t *testing.T) {
		roman := "MCMXCIV"

		got := romanToInt(roman)
		want := 1994
		validate(t, got, want)
	})
}
func validate(t testing.TB, got int, want int) {
	if got != want {
		t.Errorf("want %d got %d", want, got)
	}
}

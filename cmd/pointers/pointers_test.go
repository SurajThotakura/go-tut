package pointers

import "testing"

func TestWallet(t *testing.T) {

	checkBalance := func(t testing.TB, got Bitcoin, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{Bitcoin(0)}
		wallet.Deposit(10)
		got := wallet.Balance()
		want := Bitcoin(10)
		checkBalance(t, got, want)
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(30)}
		wallet.Withdraw(10)
		got := wallet.Balance()
		want := Bitcoin(20)
		checkBalance(t, got, want)
	})
	t.Run("Withdraw unavailable funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(0)}
		err := wallet.Withdraw(10)

		if err == nil {
			t.Error("error expected but did not get any")
		}
	})
}

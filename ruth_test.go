package ruth

import "testing"

func TestWhoIsRuth(t *testing.T) {
	want := "Ruth is a girl I used to chat with."
	if got := WhoIsRuth(); got != want {
		t.Errorf("WhoIsRuth() = got %q, want %q", got, want)
	}
}

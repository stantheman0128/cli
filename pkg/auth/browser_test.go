package auth

import (
	"os"
	"testing"

	"golang.org/x/term"
)

func TestCanCompleteBrowserLogin_MatchesIsTerminal(t *testing.T) {
	want := term.IsTerminal(int(os.Stdin.Fd()))
	if got := CanCompleteBrowserLogin(); got != want {
		t.Fatalf("CanCompleteBrowserLogin() = %v, term.IsTerminal(stdin) = %v", got, want)
	}
}
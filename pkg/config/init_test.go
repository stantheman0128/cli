package config

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestCreateConfigFile_Mode600(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "zeabur", "cli.yaml")
	if err := os.Chmod(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	createConfigFile(path)
	st, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if runtime.GOOS == "windows" {
		// Windows ACLs ignore Unix mode bits; the OpenFile/Chmod path is still
		// exercised above. Assert modes on POSIX only.
		return
	}
	if got := st.Mode().Perm(); got != 0o600 {
		t.Fatalf("new file mode = %o, want 0600", got)
	}

	// Simulate the old os.Create umask hole, then start again.
	if err := os.Chmod(path, 0o644); err != nil {
		t.Fatal(err)
	}
	createConfigFile(path)
	st, err = os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if got := st.Mode().Perm(); got != 0o600 {
		t.Fatalf("tightened file mode = %o, want 0600", got)
	}
}
package provider

import (
	"os"
	"path/filepath"
	"testing"
)

func TestClaudeAccountDirOnlyWhenItExists(t *testing.T) {
	isolate(t)
	h := t.TempDir()
	t.Setenv("HOME", h)
	t.Setenv("USERPROFILE", h)
	t.Setenv("XDG_CONFIG_HOME", filepath.Join(h, ".config"))
	if got := ClaudeAccountDir("Saved@Example.com"); got != "" {
		t.Fatalf("no directory yet: got %q, want empty", got)
	}
	dir := claudeAccountDir("saved@example.com")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		t.Fatal(err)
	}
	if got := ClaudeAccountDir(" Saved@Example.com "); got != dir {
		t.Fatalf("got %q, want %q", got, dir)
	}
}

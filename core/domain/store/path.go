package store

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/mitchellh/go-homedir"
)

const (
	// DataDirName is the name of the data directory.
	DataDirName = ".:bin"

	// DefaultFileMode is the file mode used for the data directory.
	DefaultFileMode = os.FileMode(0o700)
)

// customHome is used for testing via the WithFakeHome function.
var customHome string

// Store contains the name of the subdirectory in the data directory
// For example Store(certs) would return ~/.:bin/certs.
type Store string

// Open is a convenience method to open a file from the current Store.
func (s Store) Open(name string, flags int, perm os.FileMode) (*os.File, error) {
	dataDir, err := s.Dir()
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(path.Join(dataDir, name), flags, perm)
	if err != nil {
		return nil, fmt.Errorf("store: %w", err)
	}

	return f, nil
}

// Dir returns the store directory for the current user.
func (s Store) Dir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", fmt.Errorf("store: %w", err)
	}

	if customHome != "" {
		home = customHome
	}

	dir := path.Join(home, DataDirName, string(s))

	if err = os.MkdirAll(dir, DefaultFileMode); err != nil {
		return "", fmt.Errorf("store: %w", err)
	}

	return dir, nil
}

func WithFakeHome(t *testing.T, callback func(fakeHome string)) {
	t.Helper()

	home, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatal(err)
	}

	customHome = home

	callback(home)

	_ = os.RemoveAll(customHome)
	customHome = ""
}

package store

import (
	"fmt"
	"github.com/docker/docker/pkg/homedir"
	"os"
	"path"
	"testing"
)

const (
	// DataDirName is the name of the data directory.
	DataDirName = ".:bin"

	// DefaultFileMode is the file mode used for the data directory.
	DefaultFileMode = os.FileMode(0700)
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

// DataDir returns the store directory for the current user.
func (s Store) Dir() (string, error) {
	dataDir, err := DataDir()
	if err != nil {
		return "", err
	}

	dir := path.Join(dataDir, string(s))

	if err = os.MkdirAll(dir, DefaultFileMode); err != nil {
		return "", fmt.Errorf("store: %w", err)
	}

	return dir, nil
}

// DataDir returns the path to the data directory for the current user
// Usually, this is ~/.:bin.
func DataDir() (string, error) {
	dataDir := path.Join(userHome(), DataDirName)

	if err := os.MkdirAll(dataDir, DefaultFileMode); err != nil {
		return "", fmt.Errorf("store: %w", err)
	}

	return dataDir, nil
}

func userHome() string {
	if customHome != "" {
		return customHome
	}

	return homedir.Get()
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

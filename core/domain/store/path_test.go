package store_test

import (
	"github.com/owner/repository/core/domain/store"
	"gotest.tools/v3/assert"
	"os"
	"path"
	"testing"
)

func TestStore_Dir(t *testing.T) {
	store.WithFakeHome(t, func(fakeHome string) {
		s := store.Store("test")
		dir, err := s.Dir()
		assert.NilError(t, err)

		assert.Equal(t, dir, path.Join(fakeHome, store.DataDirName, "test"))
	})
}

func TestStore_Open(t *testing.T) {
	store.WithFakeHome(t, func(fakeHome string) {
		s := store.Store("test")
		f, err := s.Open("test.txt", os.O_CREATE, 0700)
		assert.NilError(t, err)
		defer f.Close()

		assert.Equal(t, f.Name(), path.Join(fakeHome, store.DataDirName, "test", "test.txt"))
	})
}

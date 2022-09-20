package log_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/owner/repository/core/domain/log"
	"github.com/owner/repository/core/domain/store"
	"github.com/vite-cloud/go-zoup"

	panics "github.com/magiconair/properties/assert"
	"gotest.tools/v3/assert"
)

func TestLog(t *testing.T) {
	store.WithFakeHome(t, func(fakeHome string) {
		log.Log(zoup.DebugLevel, "hello world", zoup.Fields{
			"_stack": "@", // simplifies testing
			"_time":  "@", // simplifies testing
			"key":    "value",
		})

		dir, err := log.Store.Dir()
		assert.NilError(t, err)

		contents, err := os.ReadFile(path.Join(dir, log.DefaultLogFile))
		assert.NilError(t, err)

		assert.Equal(t, string(contents), "_stack=@ _time=@ key=value level=debug message=\"hello world\"\n")
	})
}

func TestLog4(t *testing.T) {
	store.WithFakeHome(t, func(fakeHome string) {
		dir, err := log.Store.Dir()
		assert.NilError(t, err)

		err = os.Mkdir(dir+"/"+log.DefaultLogFile, 0o600)
		assert.NilError(t, err)

		panics.Panic(t, func() {
			log.Log(zoup.DebugLevel, "hello world", zoup.Fields{})
		}, fmt.Sprintf("open %s: is a directory", dir+"/"+log.DefaultLogFile))
	})
}

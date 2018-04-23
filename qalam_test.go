package qalam

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

var (
	tmpDir, err = ioutil.TempDir("", "cronowriter")
)

func Test_Write(t *testing.T) {
	tests := []struct {
		pattern        string
		expectedSuffix string
	}{
		{"test.log.%Y%m%d%H%M", "test.log.201804231918"},
	}

	for _, test := range tests {
		c := New(filepath.Join(tmpDir, test.pattern))
		for i := 0; i < 2; i++ {
			if _, err := c.Write([]byte("testeh?")); err != nil {
				t.Fatal(err)
			}
			time.Sleep(1 * time.Second)
		}

		if _, err := os.Stat(c.path); err != nil {
			t.Fatal(err)
		}

		if !strings.HasSuffix(c.path, test.expectedSuffix) {
			t.Fatalf("Expected suffix %s, got %s", test.expectedSuffix, c.path)
		}
	}

}

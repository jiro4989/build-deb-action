package main

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainFunc(t *testing.T) {
	assert := assert.New(t)

	f1 := filepath.Join("template", "control")
	f2 := filepath.Join(os.TempDir(), "test1.tmp")
	defer os.Remove(f2)

	t.Setenv("INPUT_PACKAGE", "nimjson")
	t.Setenv("INPUT_VERSION", "v2.0.0")
	t.Setenv("INPUT_INSTALLED_SIZE", "9999")
	t.Setenv("INPUT_ARCH", "amd64")
	t.Setenv("INPUT_MAINTAINER", "jiro4989")
	t.Setenv("INPUT_DEPENDS", "git")
	t.Setenv("INPUT_HOMEPAGE", "https://github.com/jiro4989/nimjson")
	t.Setenv("INPUT_SECTION", "unknown")
	t.Setenv("INPUT_DESC", "sample1\nsample2\n\nsample3")

	got := Main(f1, f2)

	assert.Equal(exitCodeOK, got)
	f3, err := os.Open(f2)
	assert.NoError(err)
	b, err := io.ReadAll(f3)
	assert.NoError(err)
	gotstr := string(b)
	want := `Package: nimjson
Version: 2.0.0
Installed-Size: 9999
Architecture: amd64
Maintainer: jiro4989
Depends: git
Homepage: https://github.com/jiro4989/nimjson
Section: unknown
Description: sample1
 sample2
 .
 sample3
`
	assert.Equal(want, gotstr)
}

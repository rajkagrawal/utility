package file

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsFileExist(t *testing.T) {
	assert.True(t,IsFileExist("file_test.go") )
}

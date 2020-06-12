package sort

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSortStringArray(t *testing.T) {
	s := []string{"stri", "raj", "apple"}
	sortStringArray(s)
	require.Equal(t, []string{"apple","raj","stri"},s)
}

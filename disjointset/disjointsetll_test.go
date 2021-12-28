package disjointset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func initSetLL() *DisjointSetLL {
	set := NewDisjointSetLL(10)
	set.Union(3, 4)
	set.Union(1, 3)
	set.Union(2, 5)
	set.Union(7, 9)
	set.Union(0, 5)
	set.Union(2, 9) // 2-5-9-7-0 3-4-1 6 8
	return set
}

func TestInitializationAllDisconnectedLL(t *testing.T) {
	set := NewDisjointSetLL(5)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			assert.Equal(t, set.Connected(i, j), i == j
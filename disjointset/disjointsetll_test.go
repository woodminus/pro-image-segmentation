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
	
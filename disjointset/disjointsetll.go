
package disjointset

import (
	"fmt"
)

/**
 * Node for the linked list
 */
type DisjointSetLLNode struct {
	id   int
	next *DisjointSetLLNode
}

/**
 * Each component of the disjoint set.
 * Contains the id of the region, its parent (if the parent has the same
 * address, it means that it's the root), a pointer to the last element of the
 * LL, a pointer to the first element of the LL and the size of the region
 */
type DisjointSetLLRegion struct {
	id     int
	parent *DisjointSetLLRegion
	head   *DisjointSetLLNode
	last   *DisjointSetLLNode
	size   int
	rank   int
}

/**
 * DisjointSetLL datatype, it contains a list of all the regions and
 * the number of components that it's storing
 */
type DisjointSetLL struct {
	regions         []*DisjointSetLLRegion
	totalComponents int
}

/**
 * Instantiates a new DisjointSetLL with 'size' elements
 */
func NewDisjointSetLL(size int) *DisjointSetLL {
	set := new(DisjointSetLL)
	set.regions = make([]*DisjointSetLLRegion, size, size)
	for i, _ := range set.regions {
		set.regions[i] = new(DisjointSetLLRegion)
		region := set.regions[i]
		region.id = i
		region.parent = region
		region.size = 1
		region.rank = 0
		region.head = new(DisjointSetLLNode)
		region.head.id = i
		region.head.next = region.head
		region.last = region.head
	}
	set.totalComponents = size
	return set
}

/**
 * Return the number of components that are stored
 */
func (set *DisjointSetLL) TotalComponents() int {
	return set.totalComponents
}

/**
 * Return the number of elements that are stored
 */
func (set *DisjointSetLL) TotalElements() int {
	return len(set.regions)
}

/**
 * Returns the id of the component to which the node v belongs
 */
func (set *DisjointSetLL) Find(v int) int {
	var findNode func(int) *DisjointSetLLRegion
	findNode = func(v int) *DisjointSetLLRegion {
		if set.regions[v].parent.id != v {
			set.regions[v].parent = findNode(set.regions[v].parent.id)
		}
		return set.regions[v].parent
	}
	return findNode(v).id
}

/**
 * Return the size of the region that v belongs to
 */
func (set *DisjointSetLL) Size(v int) int {
	return set.regions[set.Find(v)].size
}

/*
 * Merge the regions that a and b belong to
 */
func (set *DisjointSetLL) Union(a, b int) {
	region1 := set.regions[set.Find(a)]
	region2 := set.regions[set.Find(b)]
	if region1.id == region2.id {
		return
	}
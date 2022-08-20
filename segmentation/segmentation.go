/**
 * This is probably the most important package in the repository. It contains
 * all the segmentation algorithms and the utilities to transform from an
 * image to a graph and from a DisjointSet forest to an image.
 */
package segmentation

import (
	"fmt"
	"github.com/miguelfrde/image-segmentation/disjointset"
	"github.com/miguelfrde/image-segmentation/graph"
	"github.com/miguelfrde/imaging"
	"image"
	"time"
)

/**
 *
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
 * Type used to run all the segmentation algorithms.
 * It stores the graph, the resultset, the original image, the graph
 * obtained from the image and if it will generate a result image with
 * random colors.
 */
type Segmenter struct {
	rando
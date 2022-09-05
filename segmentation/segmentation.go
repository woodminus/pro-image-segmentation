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
	randomColors bool
	img          image.Image
	graph        *graph.Graph
	resultset    *disjointset.DisjointSet
	graphType    graph.GraphType
	weightfn     graph.WeightFn
}

/**
 * Returns a new Segmenter, generates a graph of the given graph type from
 * the given image using the gi
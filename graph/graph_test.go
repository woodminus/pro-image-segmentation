package graph

import (
	"github.com/stretchr/testify/assert"
	"image"
	_ "image/png"
	"os"
	"testing"
)

/*
 * Helper functions
 */

func loadGraphFromImage(imgName string, graphType GraphType) *Graph {
	f, _ := os.Open(imgName)
	defer f.Close()
	img, _, _ := image.Decode(f)
	return FromImage(img, func(p, q Pixel) float64 {
		return 1.0
	}, graphType)
}

/*
 * Tests
 */

func TestInitializationGridGraph(t *testing.T) {
	graph := New(5, 6, GRIDGRAP
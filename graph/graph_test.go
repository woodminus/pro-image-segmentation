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

func loadGraphFromImage(imgName string, graphType GraphType) *G
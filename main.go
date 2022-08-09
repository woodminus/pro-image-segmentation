
package main

import (
	"fmt"
	"github.com/miguelfrde/image-segmentation/graph"
	"github.com/miguelfrde/image-segmentation/segmentation"
	"html/template"
	"image"
	"image/png"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

const RANDOM_STR_SIZE = 25

var templates = template.Must(template.ParseGlob("web/templates/*"))
var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789")

func randomString() string {
	chars := make([]byte, RANDOM_STR_SIZE)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func loadImageFromFile(filename string) image.Image {
	f, _ := os.Open(filename)
	defer f.Close()
	img, _, _ := image.Decode(f)
	return img
}

func createFileInFS(file multipart.File, extension string) (string, error) {
	defer file.Close()
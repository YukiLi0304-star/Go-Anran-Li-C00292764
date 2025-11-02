package main

import (
	"image/color"
	"log"
	"math/rand"
	"sync"

	"github.com/hajimehoshi/ebiten"
)

const scale int = 2
const width = 300
const height = 300
const segments = 3

var blue color.Color = color.RGBA{69, 145, 196, 255}
var yellow color.Color = color.RGBA{255, 230, 120, 255}
var grid [width][height]uint8 = [width][height]uint8{}
var buffer [width][height]uint8 = [width][height]uint8{}
var count int = 0

type Segment struct {
	StartY int
	EndY   int
}

func Segments() []Segment {
	segList := make([]Segment, segments)
	segH := (height - 2) / segments
	for i := 0; i < segments; i++ {
		start := i*segH + 1
		end := (i+1)*segH + 1
		if i == segments-1 {
			end = height - 1
		}
		segList[i] = Segment{StartY: start, EndY: end}
	}
	return segList
}

func update() error {
	segList := Segments()
	var wg sync.WaitGroup

	for _, seg := range segList {
		wg.Add(1)
		go func(seg Segment) {
			defer wg.Done()
			for x := 1; x < width-1; x++ {
				for y := seg.StartY; y < seg.EndY; y++ {

					n := grid[x-1][y-1] + grid[x-1][y] +
						grid[x-1][y+1] + grid[x][y-1] +
						grid[x][y+1] + grid[x+1][y-1] +
						grid[x+1][y] + grid[x+1][y+1]

					if grid[x][y] == 0 && n == 3 {
						buffer[x][y] = 1
					} else if n < 2 || n > 3 {
						buffer[x][y] = 0
					} else {
						buffer[x][y] = grid[x][y]
					}
				}
			}
		}(seg)
	}

	wg.Wait()

	temp := buffer
	buffer = grid
	grid = temp
	return nil
}

func display(window *ebiten.Image) {
	window.Fill(blue)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			for i := 0; i < scale; i++ {
				for j := 0; j < scale; j++ {
					if grid[x][y] == 1 {
						window.Set(x*scale+i, y*scale+j, yellow)
					}
				}
			}
		}
	}
}

func frame(window *ebiten.Image) error {
	count++
	var err error = nil
	if count == 20 {
		err = update()
		count = 0
	}
	if !ebiten.IsDrawingSkipped() {
		display(window)
	}

	return err
}

func main() {
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			if rand.Float32() < 0.5 {
				grid[x][y] = 1
			}
		}
	}

	if err := ebiten.Run(frame, width, height, 2, "Game of Life"); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"cart/w4"
	"math"
	"math/rand"
	"strconv"
)

type rigid struct {
	x      float64
	y      float64
	width  float64
	height float64
	speed  float64
	angle  float64
	xspeed float64
	yspeed float64
}

var p1 = bar{rigid{x: 8, y: 32, width: 8, height: 32}}
var p2 = bar{rigid{x: 160 - 16, y: 64, width: 8, height: 32}}
var b = ball{rigid{x: 80 - 4, y: 80 - 4, width: 8, height: 8, angle: rand.Float64() * 2 * math.Pi, speed: 1}}

// var old_b = b
// var old_old_b = old_b
var counter int64 = 0
var p1score, p2score int

//go:export start
func start() {
	w4.PALETTE[0] = 0xfff6d3
	w4.PALETTE[1] = 0xf9a875
	w4.PALETTE[2] = 0xeb6b6f
	w4.PALETTE[3] = 0x7c3f58
}

//go:export update
func update() {
	p1.update(*w4.GAMEPAD1)
	p2.update(*w4.GAMEPAD2)
	b.update()

	p1.draw()
	p2.draw()
	b.draw()

	*w4.DRAW_COLORS = 4
	w4.Text(strconv.Itoa(p1score)+" - "+strconv.Itoa(p2score), 80-16, 0)
}

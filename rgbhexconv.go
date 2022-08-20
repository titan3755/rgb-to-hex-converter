package main

import (
	"fmt"
	"math"
	"strconv"
)

func RGB(r, g, b int) string {
	if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
		if r < 0 || r > 255 {
			if math.Abs(float64(r - 0)) > math.Abs(float64(r - 255)) {
				r = 255
			} else {
				r = 0
			}
		}
		if g < 0 || g > 255 {
			if math.Abs(float64(g - 0)) > math.Abs(float64(g - 255)) {
				g = 255
			} else {
				g = 0
			}
		}
		if b < 0 || b > 255 {
			if math.Abs(float64(b - 0)) > math.Abs(float64(b - 255)) {
				b = 255
			} else {
				b = 0
			}
		}
	}
	result := RGBCompute(r, g, b)
	return result
}

func RGBCompute(r, g, b int) string {
	result := ""
	var hexMap = map[int]string {
		10: "A",
		11: "B",
		12: "C",
		13: "D",
		14: "E",
		15: "F",
	}
	rd := float64(r) / 16
	gr := float64(g) / 16
	bl := float64(b) / 16

	if comp := math.Floor(rd); comp > 9 {
		result += hexMap[int(comp)]
	} else {
		result += strconv.Itoa(int(comp))
	}

	if comp := (rd - math.Floor(rd)) * 16; comp > 9 {
		result += hexMap[int(comp)]
	} else {
		result += strconv.Itoa(int(comp))
	}

	if comp := math.Floor(gr); comp > 9 {
		result += hexMap[int(comp)]
	} else {
		result += strconv.Itoa(int(comp))
	}

	if comp := (gr - math.Floor(gr)) * 16; comp > 9 {
		result += hexMap[int(comp)]
	} else {
		result += strconv.Itoa(int(comp))
	}

	if comp := math.Floor(bl); comp > 9 {
		result += hexMap[int(comp)]
	} else {
		result += strconv.Itoa(int(comp))
	}

	if comp := (bl - math.Floor(bl)) * 16; comp > 9 {
		result += hexMap[int(comp)]
	} else {
		result += strconv.Itoa(int(comp))
	}
	return result
}
package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Point struct {
	x int
	y int
}
type direction int

const (
	LEFT direction = iota
	RIGHT
	TOP
	DOWN
)

type IDirection interface {
	UpdateDirection(direction rune) direction
	UpdatePoint(p *Point, operation rune)
}

type LeftDirection struct{}

func (l *LeftDirection) UpdateDirection(d rune) direction {
	if d == 'L' {
		return DOWN
	}
	if d == 'R' {
		return TOP
	}
	return LEFT
}

func (l *LeftDirection) UpdatePoint(p *Point, operation rune) {
	if operation == 'F' {
		p.x = p.x - 1
	}
	if operation == 'F' {
		p.x = p.x + 1
	}
}

type Robort struct {
	// 记录当前方向
	Direction direction
	director  IDirection
	p         *Point
}

func ParseCmd(str string) string {
	data := []rune(str)
	var srcCmd []rune
	index := 0
	for {
		if index == len(data) {
			break
		}
		fmt.Println(data[index])
		if unicode.IsLetter(data[index]) {
			srcCmd = append(srcCmd, data[index])
			index++
			continue
		}

		if unicode.IsNumber(data[index]) {
			for i := 0; i < int(data[index]-'0'); i++ {
				srcCmd = append(srcCmd, data[index+2], data[index+3])
			}
			index = index + 5
		}
	}
	fmt.Printf("%s", string(srcCmd))
	return string(srcCmd)
}

func (r *Robort) HandleOperation(path string) {
	for _, val := range path {
		// 方向操作
		if strings.ContainsRune("RL", val) {
			r.setDirection(r.director.UpdateDirection(val))
		}
		if strings.ContainsRune("FB", val) {
			r.director.UpdatePoint(r.p, val)
		}
	}
}

func (r *Robort) setDirection(d direction) {
}

func main() {
	ParseCmd("R2(LF)")
}

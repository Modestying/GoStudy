package main

import "fmt"

//class Player:
//Must:
//ID
//Race
//Sex
//Profession
//
//Hair:
//color
//model
type Color int

const (
	Red Color = iota
	Green
	White
	Orange
)

type Hair struct {
	color string
	Model int
}

type Player struct {
	ID         string
	Profession string
	hair       Hair
}

type HairOption interface {
	apply(hair *Hair)
}

type HairOptionFunc struct {
	f func(hari *Hair)
}

func (H *HairOptionFunc) apply(hair *Hair) {
	H.f(hair)
}

func NewHairOption(f func(hair *Hair)) *HairOptionFunc {
	return &HairOptionFunc{
		f: f,
	}
}

func WithColor(color string) HairOption {
	return NewHairOption(func(hair *Hair) {
		hair.color = color
	})
}

func WithModel(model int) HairOption {
	return NewHairOption(func(hair *Hair) {
		hair.Model = model
	})
}

func NewPlayer(opts ...HairOption) *Player {
	player := &Player{
		ID: "ss",
	}
	for _, opt := range opts {
		opt.apply(&player.hair)
	}

	return player
}

func TestPlayer() {
	fmt.Println(NewPlayer([]HairOption{WithColor("red"), WithModel(1)}...))
}

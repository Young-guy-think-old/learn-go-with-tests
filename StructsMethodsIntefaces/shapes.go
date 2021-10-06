package main

import "math"

type Rectangle struct {
    Width float64
    Height float64
}

type Square struct {
    Width float64
}

type Circle struct {
    Radius float64
}

type Shape interface {
    Area() float64
}


func Perimeter(rectangle Rectangle) float64  {
	return (rectangle.Height + rectangle.Width) * 2
}

func Area(rectangle Rectangle) float64  {
	return rectangle.Height * rectangle.Width
}

func (c Circle) Area() float64  {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64  {
	return r.Height * r.Width
}

func (s Square) Area() float64  {
	return s.Width * s.Width
}



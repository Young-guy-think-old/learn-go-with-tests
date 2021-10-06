package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
    want := 40.0
    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}


func TestArea(t *testing.T) {

    areaTests := []struct {
        name string
        shape Shape
        want float64
    }{
        {name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
        {name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
        {name: "Square", shape: Square{Width: 10}, want: 100},

    }

    for _, area := range areaTests {
        t.Run(area.name, func(t *testing.T) {
            got := area.shape.Area()
            if got != area.want {
                t.Errorf("got %g want %g", got, area.want)
            }
        })
    }
}
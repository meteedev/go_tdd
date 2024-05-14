package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {

	r := Rectangle{
		Width:  10.0,
		Length: 10.0,
	}

	got := Perimeter(r)

	want := 40.0
	assert.Equal(t, want, got)

}

func TestArea(t *testing.T) {

	t.Run("Rectangle", func(t *testing.T) {

		r := Rectangle{
			Width:  10.0,
			Length: 10.0,
		}

		got := r.Area()

		want := 100.0
		assert.Equal(t, want, got)

	})

	t.Run("Circle", func(t *testing.T) {

		c := Circle{
			Radius: 10.0,
		}

		got := c.Area()

		want := 314.1592653589793
		assert.Equal(t, want, got)

	})

}

func TestAreaPolymorp(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Length: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
	}

	for _, tt := range areaTests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.hasArea, tt.shape.Area())
			})
	}

}

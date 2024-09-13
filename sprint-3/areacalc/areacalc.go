package areacalc

const pi = 3.14159

type Shape interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	length float64
	width float64
	name string
}

func NewRectangle(length float64, width float64, name string) *Rectangle {

	return &Rectangle{
        length: length,
        width: width,
		name: name,
    }
}

func (r Rectangle) Area() float64{
	var area = r.length*r.width
	return area
}

func (r Rectangle) Type() string{
	return r.name
}

func (c Circle) Area() float64{
	var area = pi * c.radius * c.radius
	return area
}

func (c Circle) Type() string{
	return c.name
}

type Circle struct {
	radius float64
	name string
}

func NewCircle(radius float64, name string) *Circle {
	return &Circle{
        radius: radius,
		name: name,
    }
}

func AreaCalculator(figures []Shape) (string, float64) {
	
	var names =""
	var area = 0.0

	if len(figures)!= 0 {
		for number, figure := range figures{
			names = names + figure.Type()
			if number != len(figures)-1{
				names = names + "-"
			}
			area = area + figure.Area()
		} 
	}

	return names, area
}

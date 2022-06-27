package golang_united_school_homework

import "errors"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity == len(b.shapes) {
		return errors.New("Full capacity!")
	}

	b.shapes = append(b.shapes, shape)

	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errors.New("Index out of range")
	}

	for index, value := range b.shapes {
		if i == index {
			return value, nil
		}
	}

	return nil, errors.New("No shape with this index")

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	shape, err := b.GetByIndex(i)

	if err != nil {
		return nil, err
	}

	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)

	return shape, nil

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	_, err := b.GetByIndex(i)

	if err != nil {
		return nil, err
	}

	result := make([]Shape, len(b.shapes))

	copy(result, b.shapes)

	b.shapes[i] = shape

	return result[i], nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var result float64

	for _, value := range b.shapes {
		result += value.CalcPerimeter()
	}

	return result
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var result float64

	for _, value := range b.shapes {
		result += value.CalcArea()
	}

	return result
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var isExist bool

	for i := 0; i < len(b.shapes); i++ {
		_, ok := b.shapes[i].(*Circle)

		if ok {
			b.ExtractByIndex(i)
			isExist = true
			i--
		}
	}

	if !isExist {
		return errors.New("Found no ircles")
	}

	return nil
}

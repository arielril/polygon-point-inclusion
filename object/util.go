package object

// GenerateRandomPoints create a list of random points
func GenerateRandomPoints(n int) []Point {
	list := make([]Point, 0)

	for i := 0; i < n; i++ {
		x := float32(rnd.Int31()%100) / 10
		y := float32(rnd.Int31()%100) / 10

		list = append(list, NewPoint2D(x, y))
	}

	return list
}

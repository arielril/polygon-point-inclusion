package program

import "github.com/arielril/polygon-point-inclusion/object"

func isLinePointsInsideStripe(l object.Line, s object.Stripe) bool {
	lp1 := l.Raw().P1.Raw()
	lp2 := l.Raw().P2.Raw()

	bot := s.Raw().Bottom
	top := s.Raw().Top

	// the first point is inside the stripe
	if lp1.Y >= bot && lp1.Y < top {
		return true
	}

	// the second point is inside the stripe
	if lp2.Y >= bot && lp2.Y < top {
		return true
	}

	return false
}

func isLineGoingThroughStripe(l object.Line, s object.Stripe) bool {
	lp1 := l.Raw().P1.Raw()
	lp2 := l.Raw().P2.Raw()

	bot := s.Raw().Bottom
	top := s.Raw().Top

	// the first point of the line is outside the stripe it self
	if lp1.Y < bot && lp2.Y > top {
		return true
	}

	// the second point of the line is outside the stripe it self
	if lp2.Y < bot && lp1.Y > top {
		return true
	}

	return false
}

func setEdgeToStripes(polygonLines []object.Line, stripeList []object.Stripe) {
	for _, s := range stripeList {
		for _, l := range polygonLines {
			if isLinePointsInsideStripe(l, s) {
				s.AddLine(l)
				// fmt.Printf("line (%v) added to stripe (%v)\n", l, s)
			} else if isLineGoingThroughStripe(l, s) {
				s.AddLine(l)
				// fmt.Printf("line (%v) added to stripe (%v)\n", l, s)
			}
		}
	}
}

func createStripes() []object.Stripe {
	stripeQty := 10
	stripeList := make([]object.Stripe, 0)

	step := float64(stripeMaxY-stripeMinY) / float64(stripeQty)

	for i := 0.0; i < float64(stripeQty); i += step {
		s := object.NewStripe(float32(i), float32(i+step))
		stripeList = append(stripeList, s)
	}

	return stripeList
}

// SetRunningBenchmark that
func SetRunningBenchmark(b bool) {
	runningBenchmark = b
}

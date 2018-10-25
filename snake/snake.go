package snake

type Snake struct {
	elements  []int // TODO: Consider linked-list implementation
	direction string
}

var opposites = map[string]string{
	"left":  "right",
	"right": "left",
	"up":    "down",
	"down":  "up",
}

func (s *Snake) SetDirection(direction string) bool {
	if direction == s.direction || opposites[direction] == s.direction {
		return false
	}

	if opposites[direction] == "" {
		return false
	}

	s.direction = direction

	return true
}

func (s *Snake) move(grow bool) {
	nextLength := len(s.elements)
	if grow {
		nextLength += 2
	}

	newElements := make([]int, nextLength)

	for idx, coord := range s.elements {
		if grow {
			newElements[idx] = coord
		} else if idx > 1 {
			newElements[idx-2] = coord
		}
	}

	oldFrontX, oldFrontY := s.front()
	xpos := nextLength - 2
	ypos := xpos + 1

	switch s.direction {
	case "right":
		newElements[xpos] = oldFrontX + 1
		newElements[ypos] = oldFrontY
		break
	case "left":
		newElements[xpos] = oldFrontX - 1
		newElements[ypos] = oldFrontY
		break
	case "up":
		newElements[xpos] = oldFrontX
		newElements[ypos] = oldFrontY - 1
		break
	case "down":
		newElements[xpos] = oldFrontX
		newElements[ypos] = oldFrontY + 1
		break
	}

	s.elements = newElements
}

func (s Snake) hasSelfCollision() bool {
	fx, fy := s.front()
	for idx := 0; idx < len(s.elements)-2; idx += 2 {
		if fx == s.elements[idx] && fy == s.elements[idx+1] {
			return true
		}
	}

	return false
}

func (s Snake) front() (x, y int) {
	length := len(s.elements)
	return s.elements[length-2], s.elements[length-1]
}

func (s Snake) back() (x, y int) {
	return s.elements[0], s.elements[1]
}

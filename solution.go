package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const Pi = 3.14

const SidesCircle = 0
const SidesTriangle = 3
const SidesSquare = 4

func CalcSquare(sideLen float64, sidesNum int8) float64 {
	switch sidesNum {
	case SidesCircle:
		return Pi * math.Pow(sideLen, 2)
	case SidesTriangle:
		return (math.Pow(sideLen, sideLen) * math.Sqrt(3)) / 4
	case SidesSquare:
		return math.Pow(sideLen, 2)
	}
	return 0
}

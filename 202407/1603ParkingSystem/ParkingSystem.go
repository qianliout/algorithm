package main

func main() {

}

type ParkingSystem struct {
	Big     int
	Medium  int
	Small   int
	B, M, S int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		Big:    big,
		Medium: medium,
		Small:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {

	case 1:
		if this.B < this.Big {
			this.B++
			return true
		}
		return false

	case 2:
		if this.M < this.Medium {
			this.M++
			return true
		}
		return false
	case 3:
		if this.S < this.Small {
			this.S++
			return true
		}
		return false
	}
	return false
}

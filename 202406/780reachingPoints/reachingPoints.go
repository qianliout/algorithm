package main

func main() {

}

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for sx < tx && sy < ty {
		if ty > tx {
			ty = ty % tx
		} else {
			tx = tx % ty
		}
	}
	if tx < sx || ty < sy {
		return false
	}
	if sx == tx && ty == sy {
		return true
	}
	if sx == tx && ty > sy && (ty-sy)%sx == 0 {
		return true
	}

	if sy == ty && tx > sx && (tx-sx)%sy == 0 {
		return true
	}

	return false
}

package greatestCommonDivisor

func GCDEuclidSub(a int, b int) int {
	for a != b {
		if a == b {
			return a
		}
		if a > b {
			a, b = b, a-b
		} else {
			a, b = b, b-a
		}
	}
	return a
}

func GCDEuclidDiv(a int, b int) int {
	for a != 0 && b != 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}
	if a == 0 {
		return b
	}
	return a
}

func GCDEuclidBinBit(a int, b int) int {
	nod := 1
	if a == 0 {
		return b
	} else if b == 0 {
		return a
	} else if a == b {
		return a
	} else if a == 1 || b == 1 {
		return 1
	}
	for a != 0 && b != 0 {
		if (a&1 | b&1) == 0 {
			nod <<= 1
			a >>= 1
			b >>= 1
			continue
		} else if (a&1 == 0) && (b&1 == 1) {
			a >>= 1
			continue
		} else if (a&1 == 1) && (b&1 == 0) {
			b >>= 1
			continue
		}
		if a > b {
			a = (a - b) >> 1
		} else {
			b = (b - a) >> 1
		}
	}
	if a == 0 {
		return nod * b
	} else {
		return nod * a
	}
}

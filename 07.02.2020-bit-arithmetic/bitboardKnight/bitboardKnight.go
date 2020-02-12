package main

func BitBoardKing(position uint64) []uint64 {

	k := uint64(1 << position)
	nA := uint64(0xfefefefefefefefe)
	nB := uint64(0xfdfdfdfdfdfdfdfd)

	nG := uint64(0xbfbfbfbfbfbfbfbf)
	nH := uint64(0x7f7f7f7f7f7f7f7f)

	p1 := (k & nA) << 15
	p2 := (k & nH) << 17
	p3 := (k & nB) & (k & nA) << 6
	p4 := (k & nG) & (k & nH) << 10
	p5 := (k & nB) & (k & nA) >> 10
	p6 := (k & nG) & (k & nH) >> 6
	p7 := (k & nA) >> 17
	p8 := (k & nH) >> 15

	mask := p1 | p2 | p3 | p4 | p5 | p6 | p7 | p8
	count := CountQuantity(mask)
	result := []uint64{count, mask}
	return result
}

func CountQuantity(mask uint64) (count uint64) {
	for mask > 1 {
		mask &= mask - 1
		count++
	}
	return count
}

package bitboardKing

func BitBoardKing(position uint64) []uint64 {

	k := uint64(1 << position)
	nA := uint64(0xfefefefefefefefe)
	nH := uint64(0x7f7f7f7f7f7f7f7f)
	p7 := (k & nA) << 7
	p8 := k << 8
	p9 := (k & nH) << 9
	p4 := (k & nA) >> 1
	p6 := (k & nH) << 1
	p1 := (k & nA) >> 9
	p2 := k >> 8
	p3 := (k & nH) >> 7
	mask := p7 | p8 | p9 | p4 | p6 | p1 | p2 | p3
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

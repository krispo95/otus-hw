package main

import (
	"fmt"
	"strings"
)

func BitBoardTrucker(fen string) string {
	maskRook := uint64(0)
	maskBishop := uint64(0)
	maskQueen := uint64(0)
	maskBlack := uint64(0)
	maskWhite := uint64(0)

	fenElems := strings.Split(fen, "")
	i := 56
	for _, item := range fenElems {
		switch item {
		case "p":
			maskBlack += CountPosition(i)
			i++
			break
		case "n":
			maskBlack += CountPosition(i)
			i++
			break
		case "b":
			maskBlack += CountPosition(i)
			i++
			break
		case "r":
			maskBlack += CountPosition(i)
			i++
			break
		case "q":
			maskBlack += CountPosition(i)
			i++
			break
		case "k":
			maskBlack += CountPosition(i)
			i++
			break
		case "P":
			maskWhite += CountPosition(i)
			i++
			break
		case "N":
			maskWhite += CountPosition(i)
			i++
			break
		case "B":
			maskWhite += CountPosition(i)
			i++
			break
		case "R":
			maskWhite += CountPosition(i)
			i++
			break
		case "Q":
			maskWhite += CountPosition(i)
			i++
			break
		case "K":
			maskWhite += CountPosition(i)
			i++
			break
		case "/":
			i -= 16
			break
		case "1":
			i++
			break
		case "2":
			i += 2
			break
		case "3":
			i += 3
			break
		case "4":
			i += 4
			break
		case "5":
			i += 5
			break
		case "6":
			i += 6
			break
		case "7":
			i += 7
			break
		case "8":
			i += 8
			break
		}
	}

	i = 56
	for _, item := range fenElems {
		switch item {
		case "p":
			i++
			break
		case "n":
			i++
			break
		case "b":
			i++
			break
		case "r":
			i++
			break
		case "q":
			i++
			break
		case "k":
			i++
			break
		case "P":
			i++
			break
		case "N":
			i++
			break
		case "B":
			maskBishop += CountBishop(uint64(i), maskWhite, maskBlack)
			i++
			break
		case "R":
			maskRook += CountRook(uint64(i), maskWhite, maskBlack)
			i++
			break
		case "Q":
			maskQueen += CountQueen(uint64(i), maskWhite, maskBlack)
			i++
			break
		case "K":
			i++
			break
		case "/":
			i -= 16
			break
		case "1":
			i++
			break
		case "2":
			i += 2
			break
		case "3":
			i += 3
			break
		case "4":
			i += 4
			break
		case "5":
			i += 5
			break
		case "6":
			i += 6
			break
		case "7":
			i += 7
			break
		case "8":
			i += 8
			break
		}
	}
	return MakeString(maskRook, maskBishop, maskQueen)
}

func CountPosition(mask int) (position uint64) {
	position = 1 << uint64(mask)
	return position
}

func CountRook(position uint64, whiteMask uint64, blackMask uint64) uint64 {
	mask := uint64(0)
	nA := uint64(0xfefefefefefefefe)
	nH := uint64(0x7f7f7f7f7f7f7f7f)
	tmpPos := int(position + 1)
	tmpPp := uint64(1 << uint64(tmpPos))
	for { //right
		if (tmpPp&nA) != tmpPp || tmpPos > 63 {
			break
		}
		if tmpPp&whiteMask != 0 {
			break
		}
		if tmpPp&blackMask != 0 {
			mask += uint64(1 << uint64(tmpPos))
			break
		}
		mask += uint64(1 << uint64(tmpPos))
		tmpPos += 1
		tmpPp = uint64(1 << uint64(tmpPos))
	}
	tmpPos = int(position - 1)
	tmpPp = uint64(1 << uint64(tmpPos))
	for { //left
		if (tmpPp&nH) != tmpPp || tmpPos < 0 {
			break
		}
		if tmpPp&whiteMask != 0 {
			break
		}
		if tmpPp&blackMask != 0 {
			mask += uint64(1 << uint64(tmpPos))
			break
		}
		mask += uint64(1 << uint64(tmpPos))
		tmpPos -= 1
		tmpPp = uint64(1 << uint64(tmpPos))
	}
	tmpPos = int(position - 8)
	tmpPp = uint64(1 << uint64(tmpPos))
	for { //down
		if tmpPos < 0 {
			break
		}
		if tmpPp&whiteMask != 0 {
			break
		}
		if tmpPp&blackMask != 0 {
			mask += uint64(1 << uint64(tmpPos))
			break
		}
		mask += uint64(1 << uint64(tmpPos))
		tmpPos -= 8
		tmpPp = uint64(1 << uint64(tmpPos))
	}
	tmpPos = int(position + 8)
	tmpPp = uint64(1 << uint64(tmpPos))
	for { //up
		if tmpPos > 63 {
			break
		}
		if tmpPp&whiteMask != 0 {
			break
		}
		if tmpPp&blackMask != 0 {
			mask += uint64(1 << uint64(tmpPos))
			break
		}
		mask += uint64(1 << uint64(tmpPos))
		tmpPos += 8
		tmpPp = uint64(1 << uint64(tmpPos))

	}
	return mask
}

func CountBishop(position uint64, whiteMask uint64, blackMask uint64) uint64 {
	mask := uint64(0)
	nA := uint64(0xfefefefefefefefe)
	nH := uint64(0x7f7f7f7f7f7f7f7f)
	tmpPos := int(position + 9)
	tmpPp := uint64(1 << uint64(tmpPos))
	for { //right up
		if (tmpPp&nA) != tmpPp || tmpPos > 63 {
			break
		}
		if tmpPp&whiteMask != 0 {
			break
		}
		if tmpPp&blackMask != 0 {
			mask += uint64(1 << uint64(tmpPos))
			break
		}

		mask += uint64(1 << uint64(tmpPos))
		tmpPos += 9
		tmpPp = uint64(1 << uint64(tmpPos))
	}
	tmpPos = int(position + 7)
	tmpPp = uint64(1 << uint64(tmpPos))
	for { //left up
		if (tmpPp&nH) != tmpPp || tmpPos > 63 {
			break
		}
		if tmpPp&whiteMask != 0 {
			break
		}
		if tmpPp&blackMask != 0 {
			mask += uint64(1 << uint64(tmpPos))
			break
		}
		mask += uint64(1 << uint64(tmpPos))
		tmpPos += 7
		tmpPp = uint64(1 << uint64(tmpPos))

	}
	tmpPos = int(position - 7)
	tmpPp = uint64(1 << uint64(tmpPos))
	for { //right down
		if (tmpPp&nA) != tmpPp || tmpPos < 0 {
			break
		}
		if tmpPp&whiteMask != 0 {
			break
		}
		if tmpPp&blackMask != 0 {
			mask += uint64(1 << uint64(tmpPos))
			break
		}
		mask += uint64(1 << uint64(tmpPos))
		tmpPos -= 7
		tmpPp = uint64(1 << uint64(tmpPos))

	}
	tmpPos = int(position - 9)
	tmpPp = uint64(1 << uint64(tmpPos))
	for { //left down
		if (tmpPp&nH) != tmpPp || tmpPos < 0 {
			break
		}
		if tmpPp&whiteMask != 0 {
			break
		}
		if tmpPp&blackMask != 0 {
			mask += uint64(1 << uint64(tmpPos))
			break
		}

		mask += uint64(1 << uint64(tmpPos))
		tmpPos -= 9
		tmpPp = uint64(1 << uint64(tmpPos))

	}
	return mask
}

func CountQueen(position uint64, whiteMask uint64, blackMask uint64) uint64 {
	mask := CountBishop(position, whiteMask, blackMask)
	mask += CountRook(position, whiteMask, blackMask)
	return mask
}

func MakeString(f, s, t uint64) string {
	return fmt.Sprint(f, "\n", s, "\n", t)
}

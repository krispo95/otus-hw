package main

import (
	"fmt"
	"strings"
)

type figures struct {
	whitePawns   uint64
	whiteKnights uint64
	whiteBishops uint64
	whiteRooks   uint64
	whiteQueens  uint64
	whiteKing    uint64

	blackPawns   uint64
	blackKnights uint64
	blackBishops uint64
	blackRooks   uint64
	blackQueens  uint64
	blackKing    uint64
}

func BitBoardFEN(fen string) string {
	mask := figures{}
	fenElems := strings.Split(fen, "")
	i := 56
	for _, item := range fenElems {
		switch item {
		case "p":
			mask.blackPawns += CountPosition(uint64(i))
			i++
			break
		case "n":
			mask.blackKnights += CountPosition(uint64(i))
			i++
			break
		case "b":
			mask.blackBishops += CountPosition(uint64(i))
			i++
			break
		case "r":
			mask.blackRooks += CountPosition(uint64(i))
			i++
			break
		case "q":
			mask.blackQueens += CountPosition(uint64(i))
			i++
			break
		case "k":
			mask.blackKing += CountPosition(uint64(i))
			i++
			break
		case "P":
			mask.whitePawns += CountPosition(uint64(i))
			i++
			break
		case "N":
			mask.whiteKnights += CountPosition(uint64(i))
			i++
			break
		case "B":
			mask.whiteBishops += CountPosition(uint64(i))
			i++
			break
		case "R":
			mask.whiteRooks += CountPosition(uint64(i))
			i++
			break
		case "Q":
			mask.whiteQueens += CountPosition(uint64(i))
			i++
			break
		case "K":
			mask.whiteKing += CountPosition(uint64(i))
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
	return mask.MakeString()
}

func CountPosition(mask uint64) (position uint64) {
	position = 1 << mask
	return position
}

func (f figures) MakeString() string {
	return fmt.Sprint(f.whitePawns, "\n", f.whiteKnights, "\n", f.whiteBishops, "\n",
		f.whiteRooks, "\n", f.whiteQueens, "\n", f.whiteKing, "\n",
		f.blackPawns, "\n", f.blackKnights, "\n", f.blackBishops, "\n",
		f.blackRooks, "\n", f.blackQueens, "\n", f.blackKing)
}

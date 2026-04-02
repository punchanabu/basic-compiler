package main

const (
	BLINE  = 10
	BID    = 11
	BCONST = 12
	BIF    = 13
	BGOTO  = 14
	BPRINT = 15
	BSTOP  = 16
	BOP    = 17
)

var opValue = map[string]int{
	"+": 1,
	"-": 2,
	"<": 3,
	"=": 4,
}

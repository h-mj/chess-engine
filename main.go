package main

import (
	"fmt"

	"github.com/h-mj/chess-engine/pkg/board"
)

func main() {
	var mailbox board.Mailbox

	mailbox.Set(board.A1, board.WhiteRook)
	mailbox.Set(board.B1, board.WhiteKnight)
	mailbox.Set(board.C1, board.WhiteBishop)
	mailbox.Set(board.D1, board.WhiteQueen)
	mailbox.Set(board.E1, board.WhiteKing)
	mailbox.Set(board.F1, board.WhiteBishop)
	mailbox.Set(board.G1, board.WhiteKnight)
	mailbox.Set(board.H1, board.WhiteRook)

	mailbox.Set(board.A2, board.WhitePawn)
	mailbox.Set(board.B2, board.WhitePawn)
	mailbox.Set(board.C2, board.WhitePawn)
	mailbox.Set(board.D2, board.WhitePawn)
	mailbox.Set(board.E2, board.WhitePawn)
	mailbox.Set(board.F2, board.WhitePawn)
	mailbox.Set(board.G2, board.WhitePawn)
	mailbox.Set(board.H2, board.WhitePawn)

	mailbox.Set(board.A7, board.BlackPawn)
	mailbox.Set(board.B7, board.BlackPawn)
	mailbox.Set(board.C7, board.BlackPawn)
	mailbox.Set(board.D7, board.BlackPawn)
	mailbox.Set(board.E7, board.BlackPawn)
	mailbox.Set(board.F7, board.BlackPawn)
	mailbox.Set(board.G7, board.BlackPawn)
	mailbox.Set(board.H7, board.BlackPawn)

	mailbox.Set(board.A8, board.BlackRook)
	mailbox.Set(board.B8, board.BlackKnight)
	mailbox.Set(board.C8, board.BlackBishop)
	mailbox.Set(board.D8, board.BlackQueen)
	mailbox.Set(board.E8, board.BlackKing)
	mailbox.Set(board.F8, board.BlackBishop)
	mailbox.Set(board.G8, board.BlackKnight)
	mailbox.Set(board.H8, board.BlackRook)

	fmt.Println(mailbox)
}

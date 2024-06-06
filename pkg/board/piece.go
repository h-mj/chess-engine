package board

type Piece uint8

const (
	NoPiece     = Piece(0)
	WhitePawn   = Piece(White) | Piece(1<<1) | Piece(Pawn)<<2
	WhiteKnight = Piece(White) | Piece(1<<1) | Piece(Knight)<<2
	WhiteBishop = Piece(White) | Piece(1<<1) | Piece(Bishop)<<2
	WhiteRook   = Piece(White) | Piece(1<<1) | Piece(Rook)<<2
	WhiteQueen  = Piece(White) | Piece(1<<1) | Piece(Queen)<<2
	WhiteKing   = Piece(White) | Piece(1<<1) | Piece(King)<<2
	BlackPawn   = Piece(Black) | Piece(1<<1) | Piece(Pawn)<<2
	BlackKnight = Piece(Black) | Piece(1<<1) | Piece(Knight)<<2
	BlackBishop = Piece(Black) | Piece(1<<1) | Piece(Bishop)<<2
	BlackRook   = Piece(Black) | Piece(1<<1) | Piece(Rook)<<2
	BlackQueen  = Piece(Black) | Piece(1<<1) | Piece(Queen)<<2
	BlackKing   = Piece(Black) | Piece(1<<1) | Piece(King)<<2
)

func (p Piece) Color() Color {
	return Color(p & 0b0000001)
}

func (p Piece) PieceType() PieceType {
	return PieceType(p >> 2)
}

func (p Piece) Identifier() string {
	const pieceIdentifiers = "-?Pp??Nn??Bb??Rr??Qq??Kk"

	return string(pieceIdentifiers[p])
}

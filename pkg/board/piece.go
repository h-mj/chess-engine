package board

type Piece uint8

const (
	NoPiece     = Piece(0)
	WhitePawn   = Piece(White) | Piece(Pawn)<<1
	WhiteKnight = Piece(White) | Piece(Knight)<<1
	WhiteBishop = Piece(White) | Piece(Bishop)<<1
	WhiteRook   = Piece(White) | Piece(Rook)<<1
	WhiteQueen  = Piece(White) | Piece(Queen)<<1
	WhiteKing   = Piece(White) | Piece(King)<<1
	BlackPawn   = Piece(Black) | Piece(Pawn)<<1
	BlackKnight = Piece(Black) | Piece(Knight)<<1
	BlackBishop = Piece(Black) | Piece(Bishop)<<1
	BlackRook   = Piece(Black) | Piece(Rook)<<1
	BlackQueen  = Piece(Black) | Piece(Queen)<<1
	BlackKing   = Piece(Black) | Piece(King)<<1
)

func (p Piece) Color() Color {
	return Color(p & 0b0000001)
}

func (p Piece) PieceType() PieceType {
	return PieceType(p >> 1)
}

func (p Piece) Identifier() string {
	const pieceIdentifiers = "- PpNnBbRrQqKk"

	return string(pieceIdentifiers[p])
}

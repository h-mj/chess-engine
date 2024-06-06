package board

import "strings"

type Mailbox [64]Piece

func (m *Mailbox) Get(square Square) Piece {
	return m[square]
}

func (m *Mailbox) Set(square Square, piece Piece) {
	m[square] = piece
}

func (m Mailbox) String() string {
	var builder strings.Builder

	writePieces(&builder, "  +---+---+---+---+---+---+---+---+\n", nil)
	writePieces(&builder, "8 | # | # | # | # | # | # | # | # |\n", m[A8:H8+1])
	writePieces(&builder, "  +---+---+---+---+---+---+---+---+\n", nil)
	writePieces(&builder, "7 | # | # | # | # | # | # | # | # |\n", m[A7:H7+1])
	writePieces(&builder, "  +---+---+---+---+---+---+---+---+\n", nil)
	writePieces(&builder, "6 | # | # | # | # | # | # | # | # |\n", m[A6:H6+1])
	writePieces(&builder, "  +---+---+---+---+---+---+---+---+\n", nil)
	writePieces(&builder, "5 | # | # | # | # | # | # | # | # |\n", m[A5:H5+1])
	writePieces(&builder, "  +---+---+---+---+---+---+---+---+\n", nil)
	writePieces(&builder, "4 | # | # | # | # | # | # | # | # |\n", m[A4:H4+1])
	writePieces(&builder, "  +---+---+---+---+---+---+---+---+\n", nil)
	writePieces(&builder, "3 | # | # | # | # | # | # | # | # |\n", m[A3:H3+1])
	writePieces(&builder, "  +---+---+---+---+---+---+---+---+\n", nil)
	writePieces(&builder, "2 | # | # | # | # | # | # | # | # |\n", m[A2:H2+1])
	writePieces(&builder, "  +---+---+---+---+---+---+---+---+\n", nil)
	writePieces(&builder, "1 | # | # | # | # | # | # | # | # |\n", m[A1:H1+1])
	writePieces(&builder, "  +---+---+---+---+---+---+---+---+\n", nil)
	writePieces(&builder, "    a   b   c   d   e   f   g   h", nil)

	return builder.String()
}

func writePieces(builder *strings.Builder, template string, pieces []Piece) {
	pieceCount := len(pieces)

	if pieceCount == 0 {
		builder.WriteString(template)
		return
	}

	parts := strings.SplitN(template, "#", pieceCount+1)

	for i, piece := range pieces {
		builder.WriteString(parts[i])

		if piece != NoPiece {
			builder.WriteString(piece.Identifier())
		} else {
			builder.WriteString(" ")
		}
	}

	builder.WriteString(parts[pieceCount])
}

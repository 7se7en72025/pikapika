package ast

type Span struct {
	start  int
	end    int
	line   int
	column int
}

type Node interface {
	nodeKind() string // node ka type boltaa
	Span() Span       // gives node ka strt nd end posns
}

type Stmt interface {
	Node
}

type Expr interface {
	Node
}

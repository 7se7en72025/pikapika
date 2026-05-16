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

type Param struct {
	Name string
	S    Span
}

type VarDecl struct {
	Name  string
	Value Expr
	S     Span
}

type BlockStmt struct {
	Stmts []Stmt
	S     Span
}
type IfStmt struct {
	Cond Expr
	Then *BlockStmt
	Else *BlockStmt
	S    Span
}

type FnDecl struct {
	Name   string
	Params []Param
	Body   *BlockStmt
	S      Span
}

func (f *FnDecl) nodeKind() string { return "FnDecl" }
func (f *FnDecl) Span() Span       { return f.S }

func (p *Param) nodeKind() string { return "Param" }
func (p *Param) Span() Span       { return p.S }

func (v *VarDecl) nodeKind() string { return "VarDecl" }
func (v *VarDecl) Span() Span       { return v.S }

func (i *IfStmt) nodeKind() string { return "IfStmt" }
func (i *IfStmt) Span() Span       { return i.S }

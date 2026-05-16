package ast

type Span struct{ Start, End, Line, Column int }
type Node interface {
	nodeKind() string
	Span() Span
}
type Stmt interface{ Node }
type Expr interface{ Node }

type FnDecl struct {
	Name   string
	Params []Param
	Body   *BlockStmt
	S      Span
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

type IfStmt struct {
	Condition  Expr
	Then, Else *BlockStmt
	S          Span
}

type ForStmt struct {
	Init      Stmt
	Condition Expr
	Step      Stmt
	Body      *BlockStmt
	S         Span
}

type WhileStmt struct {
	Condition Expr
	Body      *BlockStmt
	S         Span
}

type ReturnStmt struct {
	Value Expr
	S     Span
}

type BlockStmt struct {
	Stmts []Stmt
	S     Span
}

type ExprStmt struct {
	Expr Expr
	S    Span
}

type BinaryExpr struct {
	Left  Expr
	Op    string
	Right Expr
	S     Span
}

type UnaryExpr struct {
	Op      string
	Operand Expr
	S       Span
}

type CallExpr struct {
	Callee Expr
	Args   []Expr
	S      Span
}

type IndexExpr struct {
	Object, Index Expr
	S             Span
}

type MemberExpr struct {
	Object Expr
	Field  string
	S      Span
}

type Ident struct {
	Name string
	S    Span
}

type IntLit struct {
	Value int64
	S     Span
}

type FloatLit struct {
	Value float64
	S     Span
}

type StringLit struct {
	Value string
	S     Span
}

type BoolLit struct {
	Value bool
	S     Span
}

type NullLit struct{ S Span }

type ArrayLit struct {
	Elements []Expr
	S        Span
}

type MapLit struct {
	Pairs []MapPair
	S     Span
}

type MapPair struct{ Key, Value Expr }

type MatchExpr struct {
	Subject Expr
	Arms    []MatchArm
	S       Span
}

type MatchArm struct {
	Pattern Pattern
	Body    Expr
	S       Span
}

type Pattern interface{ patternKind() string }

type WildcardPattern struct{ S Span }

type LiteralPattern struct {
	Value Expr
	S     Span
}

type IdentPattern struct {
	Name string
	S    Span
}

type RangePattern struct {
	Low, High Expr
	Inclusive bool
	S         Span
}

type ArrayPattern struct {
	Elements []Pattern
	Rest     string
	S        Span
}

type MapPattern struct {
	Fields map[string]Pattern
	S      Span
}

type PipeExpr struct {
	Left, Right Expr
	S           Span
}

type RangeExpr struct {
	Low, High Expr
	Inclusive bool
	S         Span
}

type DestructureStmt struct {
	Pattern Pattern
	Value   Expr
	S       Span
}

type ErrorLit struct {
	Message Expr
	S       Span
}

type ForInStmt struct {
	Name string
	Iter Expr
	Body *BlockStmt
	S    Span
}

type AssignStmt struct {
	Name  string
	Value Expr
	S     Span
}

func (f *FnDecl) nodeKind() string    { return "FnDecl" }
func (f *FnDecl) Span() Span          { return f.S }
func (p *Param) nodeKind() string     { return "Param" }
func (p *Param) Span() Span           { return p.S }
func (v *VarDecl) nodeKind() string   { return "VarDecl" }
func (v *VarDecl) Span() Span         { return v.S }
func (i *IfStmt) nodeKind() string    { return "IfStmt" }
func (i *IfStmt) Span() Span          { return i.S }
func (f *ForStmt) nodeKind() string   { return "ForStmt" }
func (f *ForStmt) Span() Span         { return f.S }
func (w *WhileStmt) nodeKind() string { return "WhileStmt" }
func (w *WhileStmt) Span() Span       { return w.S }
func (r *ReturnStmt) nodeKind() string {
	return "ReturnStmt"
}
func (r *ReturnStmt) Span() Span       { return r.S }
func (b *BlockStmt) nodeKind() string  { return "BlockStmt" }
func (b *BlockStmt) Span() Span        { return b.S }
func (e *ExprStmt) nodeKind() string   { return "ExprStmt" }
func (e *ExprStmt) Span() Span         { return e.S }
func (b *BinaryExpr) nodeKind() string { return "BinaryExpr" }
func (b *BinaryExpr) Span() Span       { return b.S }
func (u *UnaryExpr) nodeKind() string  { return "UnaryExpr" }
func (u *UnaryExpr) Span() Span        { return u.S }
func (c *CallExpr) nodeKind() string   { return "CallExpr" }
func (c *CallExpr) Span() Span         { return c.S }
func (i *IndexExpr) nodeKind() string  { return "IndexExpr" }
func (i *IndexExpr) Span() Span        { return i.S }
func (m *MemberExpr) nodeKind() string { return "MemberExpr" }
func (m *MemberExpr) Span() Span       { return m.S }
func (id *Ident) nodeKind() string     { return "Ident" }
func (id *Ident) Span() Span           { return id.S }
func (i *IntLit) nodeKind() string     { return "IntLit" }
func (i *IntLit) Span() Span           { return i.S }
func (f *FloatLit) nodeKind() string   { return "FloatLit" }
func (f *FloatLit) Span() Span         { return f.S }
func (s *StringLit) nodeKind() string  { return "StringLit" }
func (s *StringLit) Span() Span        { return s.S }
func (b *BoolLit) nodeKind() string    { return "BoolLit" }
func (b *BoolLit) Span() Span          { return b.S }
func (n *NullLit) nodeKind() string    { return "NullLit" }
func (n *NullLit) Span() Span          { return n.S }
func (a *ArrayLit) nodeKind() string   { return "ArrayLit" }
func (a *ArrayLit) Span() Span         { return a.S }
func (m *MapLit) nodeKind() string     { return "MapLit" }
func (m *MapLit) Span() Span           { return m.S }
func (mp *MapPair) nodeKind() string   { return "MapPair" }
func (mp *MapPair) Span() Span         { return Span{} }
func (m *MatchExpr) nodeKind() string  { return "MatchExpr" }
func (m *MatchExpr) Span() Span        { return m.S }
func (ma *MatchArm) nodeKind() string  { return "MatchArm" }
func (ma *MatchArm) Span() Span        { return ma.S }
func (w *WildcardPattern) patternKind() string {
	return "Wildcard"
}
func (w *WildcardPattern) Span() Span { return w.S }
func (l *LiteralPattern) patternKind() string {
	return "Literal"
}
func (l *LiteralPattern) Span() Span { return l.S }
func (i *IdentPattern) patternKind() string {
	return "Ident"
}
func (i *IdentPattern) Span() Span { return i.S }
func (r *RangePattern) patternKind() string {
	return "Range"
}
func (r *RangePattern) Span() Span { return r.S }
func (a *ArrayPattern) patternKind() string {
	return "ArrayPattern"
}
func (a *ArrayPattern) Span() Span { return a.S }
func (m *MapPattern) patternKind() string {
	return "MapPattern"
}
func (m *MapPattern) Span() Span     { return m.S }
func (p *PipeExpr) nodeKind() string { return "PipeExpr" }
func (p *PipeExpr) Span() Span       { return p.S }
func (r *RangeExpr) nodeKind() string {
	return "RangeExpr"
}
func (r *RangeExpr) Span() Span { return r.S }
func (d *DestructureStmt) nodeKind() string {
	return "DestructureStmt"
}
func (d *DestructureStmt) Span() Span { return d.S }
func (e *ErrorLit) nodeKind() string  { return "ErrorLit" }
func (e *ErrorLit) Span() Span        { return e.S }
func (f *ForInStmt) nodeKind() string { return "ForInStmt" }
func (f *ForInStmt) Span() Span       { return f.S }

func (a *AssignStmt) nodeKind() string { return "AssignStmt" }
func (a *AssignStmt) Span() Span       { return a.S }

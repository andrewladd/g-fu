package gfu

import (
  //"log"
  "strings"
)

type NilType struct {
  BasicType
}

func (t *NilType) Init(id *Sym) *NilType {
  t.BasicType.Init(id)
  return t
}

func (t *NilType) Call(g *G, val Val, args []Val, env *Env) (Val, Error) {
  return g.NIL, g.NewError(g.Pos, "Nil call")
}
  
func (t *NilType) Dump(_ Val, out *strings.Builder) {
  out.WriteRune('_')
}

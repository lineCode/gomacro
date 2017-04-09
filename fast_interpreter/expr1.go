/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * expr1.go
 *
 *  Created on Apr 03, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	"go/constant"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func LitValue(value I) Lit {
	return Lit{Type: r.TypeOf(value), Value: value}
}

func ExprUntypedLit(kind r.Kind, value constant.Value) *Expr {
	return &Expr{Lit: Lit{Type: r.TypeOf(value), Value: UntypedLit{Kind: kind, Obj: value}}}
}

func ExprValue(value I) *Expr {
	return &Expr{Lit: Lit{Type: r.TypeOf(value), Value: value}, IsNil: value == nil}
}

func ExprLit(lit Lit) *Expr {
	return &Expr{Lit: lit, IsNil: lit.Value == nil}
}

func ExprFun(t r.Type, fun I) *Expr {
	return &Expr{Lit: Lit{Type: t}, Fun: fun}
}

func ExprX1(t r.Type, fun func(env *Env) r.Value) *Expr {
	return &Expr{Lit: Lit{Type: t}, Fun: fun}
}

func ExprXV(t r.Type, fun func(env *Env) (r.Value, []r.Value)) *Expr {
	return &Expr{Lit: Lit{Type: t}, Fun: fun}
}

func ExprBool(fun func(env *Env) bool) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfBool}, Fun: fun}
}

func ExprInt(fun func(env *Env) int) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfInt}, Fun: fun}
}

func ExprInt8(fun func(env *Env) int8) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfInt8}, Fun: fun}
}

func ExprInt16(fun func(env *Env) int16) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfInt16}, Fun: fun}
}

func ExprInt32(fun func(env *Env) int32) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfInt32}, Fun: fun}
}

func ExprInt64(fun func(env *Env) int64) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfInt64}, Fun: fun}
}

func ExprUint(fun func(env *Env) uint) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfUint}, Fun: fun}
}

func ExprUint8(fun func(env *Env) uint8) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfUint8}, Fun: fun}
}

func ExprUint16(fun func(env *Env) uint16) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfUint16}, Fun: fun}
}

func ExprUint32(fun func(env *Env) uint32) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfUint32}, Fun: fun}
}

func ExprUint64(fun func(env *Env) uint64) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfUint64}, Fun: fun}
}

func ExprUintptr(fun func(env *Env) uintptr) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfUintptr}, Fun: fun}
}

func ExprFloat32(fun func(env *Env) float32) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfFloat32}, Fun: fun}
}

func ExprFloat64(fun func(env *Env) float64) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfFloat64}, Fun: fun}
}

func ExprComplex64(fun func(env *Env) complex64) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfComplex64}, Fun: fun}
}

func ExprComplex128(fun func(env *Env) complex128) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfComplex128}, Fun: fun}
}

func ExprString(fun func(env *Env) string) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfString}, Fun: fun}
}

func (expr *Expr) EvalConst(opts CompileOptions) I {
	if expr == nil {
		return nil
	}
	if expr.Const() {
		return expr.Value
	}
	ret, rets := AsXV(expr.Fun, opts)(nil)
	if ret == None {
		Errorf("constant should evaluate to a single value, found no values at all")
		return nil
	}
	if len(rets) > 1 {
		Errorf("constant should evaluate to a single value, found %d values: %v", len(rets), rets)
		return nil
	}
	t1 := expr.Type
	t2 := ValueType(ret)
	if t1 != t2 {
		Errorf("constant should evaluate to <%v>, found: %v <%v>", t1, t2, ret)
		return nil
	}
	var value I
	if ret != Nil {
		value = ret.Interface()
	}
	expr.Value = value
	expr.IsNil = value == nil
	expr.Fun = nil // no longer needed.
	return value
}

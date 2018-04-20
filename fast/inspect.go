/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * inspect.go
 *
 *  Created on: Apr 20, 2017
 *      Author: Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (ir *Interp) Inspect(str string) {
	c := ir.Comp
	expr := c.Compile(c.Parse(str))
	v := ir.RunExpr1(expr)
	xt := expr.Type
	var t r.Type
	if v.IsValid() && v.CanInterface() && v != None {
		t = r.TypeOf(v.Interface()) // show concrete type
	} else if xt != nil {
		t = xt.ReflectType()
	}
	ip := NewInspector(str, v, t, xt, ir.Comp.Globals)
	ip.Show()
	ip.Repl()
}

// this file was generated by gomacro command: import "debug/gosym"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"debug/gosym"
)

func init() {
	Binds["debug/gosym"] = map[string]Value{
		"NewLineTable":	ValueOf(gosym.NewLineTable),
		"NewTable":	ValueOf(gosym.NewTable),
	}
	Types["debug/gosym"] = map[string]Type{
		"DecodingError":	TypeOf((*gosym.DecodingError)(nil)).Elem(),
		"Func":	TypeOf((*gosym.Func)(nil)).Elem(),
		"LineTable":	TypeOf((*gosym.LineTable)(nil)).Elem(),
		"Obj":	TypeOf((*gosym.Obj)(nil)).Elem(),
		"Sym":	TypeOf((*gosym.Sym)(nil)).Elem(),
		"Table":	TypeOf((*gosym.Table)(nil)).Elem(),
		"UnknownFileError":	TypeOf((*gosym.UnknownFileError)(nil)).Elem(),
		"UnknownLineError":	TypeOf((*gosym.UnknownLineError)(nil)).Elem(),
	}
	Proxies["debug/gosym"] = map[string]Type{
	}
}
// this file was generated by gomacro command: import _i "github.com/cosmos72/gomacro/fast"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package fast

import (
	r "reflect"

	"github.com/cosmos72/gomacro/imports"
)

// reflection: allow interpreted code to import "github.com/cosmos72/gomacro/fast"
func init() {
	imports.Packages["github.com/cosmos72/gomacro/fast"] = imports.Package{
		Binds: map[string]r.Value{
			"AnyDepth":            r.ValueOf(AnyDepth),
			"CompileDefaults":     r.ValueOf(CompileDefaults),
			"CompileKeepUntyped":  r.ValueOf(CompileKeepUntyped),
			"ConstBind":           r.ValueOf(ConstBind),
			"ConstBindDescriptor": r.ValueOf(ConstBindDescriptor),
			"FileDepth":           r.ValueOf(FileDepth),
			"FuncBind":            r.ValueOf(FuncBind),
			"IntBind":             r.ValueOf(IntBind),
			"Interrupt":           r.ValueOf(&Interrupt).Elem(),
			"MakeBindDescriptor":  r.ValueOf(MakeBindDescriptor),
			"New":                 r.ValueOf(New),
			"NewComp":             r.ValueOf(NewComp),
			"NewCompEnv":          r.ValueOf(NewCompEnv),
			"NewCompEnvTop":       r.ValueOf(NewCompEnvTop),
			"NewEnv":              r.ValueOf(NewEnv),
			"NewEnv4Func":         r.ValueOf(NewEnv4Func),
			"NewThreadGlobals":    r.ValueOf(NewThreadGlobals),
			"NoIndex":             r.ValueOf(NoIndex),
			"PlaceAddress":        r.ValueOf(PlaceAddress),
			"PlaceSettable":       r.ValueOf(PlaceSettable),
			"PoolCapacity":        r.ValueOf(PoolCapacity),
			"SigDefer":            r.ValueOf(SigDefer),
			"SigNone":             r.ValueOf(SigNone),
			"SigReturn":           r.ValueOf(SigReturn),
			"TopDepth":            r.ValueOf(TopDepth),
			"VarBind":             r.ValueOf(VarBind),
		},
		Types: map[string]r.Type{
			"Bind":               r.TypeOf((*Bind)(nil)).Elem(),
			"BindClass":          r.TypeOf((*BindClass)(nil)).Elem(),
			"BindDescriptor":     r.TypeOf((*BindDescriptor)(nil)).Elem(),
			"Builtin":            r.TypeOf((*Builtin)(nil)).Elem(),
			"Call":               r.TypeOf((*Call)(nil)).Elem(),
			"Code":               r.TypeOf((*Code)(nil)).Elem(),
			"Comp":               r.TypeOf((*Comp)(nil)).Elem(),
			"CompEnv":            r.TypeOf((*Interp)(nil)).Elem(),
			"CompThreadGlobals":  r.TypeOf((*CompThreadGlobals)(nil)).Elem(),
			"CompileOptions":     r.TypeOf((*CompileOptions)(nil)).Elem(),
			"Env":                r.TypeOf((*Env)(nil)).Elem(),
			"Expr":               r.TypeOf((*Expr)(nil)).Elem(),
			"FuncInfo":           r.TypeOf((*FuncInfo)(nil)).Elem(),
			"Function":           r.TypeOf((*Function)(nil)).Elem(),
			"I":                  r.TypeOf((*I)(nil)).Elem(),
			"Lit":                r.TypeOf((*Lit)(nil)).Elem(),
			"LoopInfo":           r.TypeOf((*LoopInfo)(nil)).Elem(),
			"NamedType":          r.TypeOf((*NamedType)(nil)).Elem(),
			"Place":              r.TypeOf((*Place)(nil)).Elem(),
			"PlaceOption":        r.TypeOf((*PlaceOption)(nil)).Elem(),
			"Signal":             r.TypeOf((*Signal)(nil)).Elem(),
			"Stmt":               r.TypeOf((*Stmt)(nil)).Elem(),
			"Symbol":             r.TypeOf((*Symbol)(nil)).Elem(),
			"ThreadGlobals":      r.TypeOf((*ThreadGlobals)(nil)).Elem(),
			"TypeAssertionError": r.TypeOf((*TypeAssertionError)(nil)).Elem(),
			"UntypedLit":         r.TypeOf((*UntypedLit)(nil)).Elem(),
			"Var":                r.TypeOf((*Var)(nil)).Elem(),
		},
		Proxies: map[string]r.Type{
			"I": r.TypeOf((*I_github_com_cosmos72_gomacro_fast)(nil)).Elem(),
		}}
}

// --------------- proxy for github.com/cosmos72/gomacro/fast.I ---------------
type I_github_com_cosmos72_gomacro_fast struct {
	Object interface{}
}

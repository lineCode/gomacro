// this file was generated by gomacro command: import "container/ring"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"container/ring"
)

func init() {
	Binds["container/ring"] = map[string]Value{
		"New":	ValueOf(ring.New),
	}
	Types["container/ring"] = map[string]Type{
		"Ring":	TypeOf((*ring.Ring)(nil)).Elem(),
	}
	Proxies["container/ring"] = map[string]Type{
	}
}
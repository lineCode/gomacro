// this file was generated by gomacro command: import "encoding/base32"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"encoding/base32"
)

func init() {
	Binds["encoding/base32"] = map[string]Value{
		"HexEncoding":	ValueOf(&base32.HexEncoding).Elem(),
		"NewDecoder":	ValueOf(base32.NewDecoder),
		"NewEncoder":	ValueOf(base32.NewEncoder),
		"NewEncoding":	ValueOf(base32.NewEncoding),
		"StdEncoding":	ValueOf(&base32.StdEncoding).Elem(),
	}
	Types["encoding/base32"] = map[string]Type{
		"CorruptInputError":	TypeOf((*base32.CorruptInputError)(nil)).Elem(),
		"Encoding":	TypeOf((*base32.Encoding)(nil)).Elem(),
	}
	Proxies["encoding/base32"] = map[string]Type{
	}
}
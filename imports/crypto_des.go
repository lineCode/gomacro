// this file was generated by gomacro command: import "crypto/des"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"crypto/des"
)

func init() {
	Binds["crypto/des"] = map[string]Value{
		"BlockSize":	ValueOf(des.BlockSize),
		"NewCipher":	ValueOf(des.NewCipher),
		"NewTripleDESCipher":	ValueOf(des.NewTripleDESCipher),
	}
	Types["crypto/des"] = map[string]Type{
		"KeySizeError":	TypeOf((*des.KeySizeError)(nil)).Elem(),
	}
	Proxies["crypto/des"] = map[string]Type{
	}
}
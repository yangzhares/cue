// Code generated by cuelang.org/go/pkg/gen. DO NOT EDIT.

package base64

import (
	"cuelang.org/go/internal/core/adt"
	"cuelang.org/go/pkg/internal"

	_ "embed"
)

func init() {
	internal.Register("encoding/base64", pkg)
}

//go:embed pkg.cue
var cueDecls string

var _ = adt.TopKind // in case the adt package isn't used

var pkg = &internal.Package{
	Funcs: map[string]func(c *internal.CallCtxt){
		"EncodedLen": func(c *internal.CallCtxt) {
			encoding, n := c.Value(0), c.Int(1)
			if c.Do() {
				c.Ret, c.Err = EncodedLen(encoding, n)
			}
		},
		"DecodedLen": func(c *internal.CallCtxt) {
			encoding, x := c.Value(0), c.Int(1)
			if c.Do() {
				c.Ret, c.Err = DecodedLen(encoding, x)
			}
		},
		"Encode": func(c *internal.CallCtxt) {
			encoding, src := c.Value(0), c.Bytes(1)
			if c.Do() {
				c.Ret, c.Err = Encode(encoding, src)
			}
		},
		"Decode": func(c *internal.CallCtxt) {
			encoding, s := c.Value(0), c.String(1)
			if c.Do() {
				c.Ret, c.Err = Decode(encoding, s)
			}
		},
	},
	CUE: cueDecls,
}

// Code generated by cuelang.org/go/pkg/gen. DO NOT EDIT.

package regexp

import (
	"cuelang.org/go/internal/core/adt"
	"cuelang.org/go/pkg/internal"

	_ "embed"
)

func init() {
	internal.Register("regexp", pkg)
}

//go:embed pkg.cue
var cueDecls string

var _ = adt.TopKind // in case the adt package isn't used

var pkg = &internal.Package{
	Funcs: map[string]func(c *internal.CallCtxt){
		"Find": func(c *internal.CallCtxt) {
			pattern, s := c.String(0), c.String(1)
			if c.Do() {
				c.Ret, c.Err = Find(pattern, s)
			}
		},
		"FindAll": func(c *internal.CallCtxt) {
			pattern, s, n := c.String(0), c.String(1), c.Int(2)
			if c.Do() {
				c.Ret, c.Err = FindAll(pattern, s, n)
			}
		},
		"FindAllNamedSubmatch": func(c *internal.CallCtxt) {
			pattern, s, n := c.String(0), c.String(1), c.Int(2)
			if c.Do() {
				c.Ret, c.Err = FindAllNamedSubmatch(pattern, s, n)
			}
		},
		"FindAllSubmatch": func(c *internal.CallCtxt) {
			pattern, s, n := c.String(0), c.String(1), c.Int(2)
			if c.Do() {
				c.Ret, c.Err = FindAllSubmatch(pattern, s, n)
			}
		},
		"FindNamedSubmatch": func(c *internal.CallCtxt) {
			pattern, s := c.String(0), c.String(1)
			if c.Do() {
				c.Ret, c.Err = FindNamedSubmatch(pattern, s)
			}
		},
		"FindSubmatch": func(c *internal.CallCtxt) {
			pattern, s := c.String(0), c.String(1)
			if c.Do() {
				c.Ret, c.Err = FindSubmatch(pattern, s)
			}
		},
		"ReplaceAll": func(c *internal.CallCtxt) {
			pattern, src, repl := c.String(0), c.String(1), c.String(2)
			if c.Do() {
				c.Ret, c.Err = ReplaceAll(pattern, src, repl)
			}
		},
		"ReplaceAllLiteral": func(c *internal.CallCtxt) {
			pattern, src, repl := c.String(0), c.String(1), c.String(2)
			if c.Do() {
				c.Ret, c.Err = ReplaceAllLiteral(pattern, src, repl)
			}
		},
		"Valid": func(c *internal.CallCtxt) {
			pattern := c.String(0)
			if c.Do() {
				c.Ret, c.Err = Valid(pattern)
			}
		},
		"Match": func(c *internal.CallCtxt) {
			pattern, s := c.String(0), c.String(1)
			if c.Do() {
				c.Ret, c.Err = Match(pattern, s)
			}
		},
		"QuoteMeta": func(c *internal.CallCtxt) {
			s := c.String(0)
			if c.Do() {
				c.Ret = QuoteMeta(s)
			}
		},
	},
	CUE: cueDecls,
}

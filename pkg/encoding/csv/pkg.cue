// Code generated by cuelang.org/go/pkg/gen. DO NOT EDIT.
package csv

funcs: Encode: {
	in: [{
		name: "x"
		type: _
	}]
	out: string
}
funcs: Decode: {
	in: [{
		name: "r"
		type: bytes | string
	}]
	out: [...]
}

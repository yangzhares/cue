// Code generated by cuelang.org/go/pkg/gen. DO NOT EDIT.
package json

funcs: Valid: {
	in: [{
		name: "data"
		type: bytes | string
	}]
	out: bool
}
funcs: Compact: {
	in: [{
		name: "src"
		type: bytes | string
	}]
	out: string
}
funcs: Indent: {
	in: [{
		name: "src"
		type: bytes | string
	}, {
		name: "prefix"
		type: string
	}, {
		name: "indent"
		type: string
	}]
	out: string
}
funcs: HTMLEscape: {
	in: [{
		name: "src"
		type: bytes | string
	}]
	out: string
}
funcs: Marshal: {
	in: [{
		name: "v"
		type: _
	}]
	out: string
}
funcs: MarshalStream: {
	in: [{
		name: "v"
		type: _
	}]
	out: string
}
funcs: UnmarshalStream: {
	in: [{
		name: "data"
		type: bytes | string
	}]
	out: _
}
funcs: Unmarshal: {
	in: [{
		name: "b"
		type: bytes | string
	}]
	out: _
}
funcs: Validate: {
	in: [{
		name: "b"
		type: bytes | string
	}, {
		name: "v"
		type: _
	}]
	out: bool
}

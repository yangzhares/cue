// Code generated by cuelang.org/go/pkg/gen. DO NOT EDIT.
package ed25519

PublicKeySize: 32
funcs: Valid: {
	in: [{
		name: "publicKey"
		type: bytes | string
	}, {
		name: "message"
		type: bytes | string
	}, {
		name: "signature"
		type: bytes | string
	}]
	out: bool
}

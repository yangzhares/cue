// Code generated by cuelang.org/go/pkg/gen. DO NOT EDIT.
package hmac

MD5:        "MD5"
SHA1:       "SHA1"
SHA224:     "SHA224"
SHA256:     "SHA256"
SHA384:     "SHA384"
SHA512:     "SHA512"
SHA512_224: "SHA512_224"
SHA512_256: "SHA512_256"
funcs: Sign: {
	in: [{
		name: "hashName"
		type: string
	}, {
		name: "key"
		type: bytes | string
	}, {
		name: "data"
		type: bytes | string
	}]
	out: bytes | string
}

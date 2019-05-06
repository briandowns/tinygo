// +build freebsd

package main

// commands used by the compilation process might have different file names on FreeBSD than those used on Linux.
var commands = map[string]string{
	"clang":   "clang80",
	"ld.lld":  "ld.lld",
	"wasm-ld": "wasm-ld",
}

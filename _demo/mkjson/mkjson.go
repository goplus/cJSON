package main

import (
	"github.com/goplus/cjson"
	"github.com/goplus/llgo/c"
)

func main() {
	mod := cjson.Object()
	mod.SetItem(c.Str("name"), cjson.String(c.Str("math")))

	syms := cjson.Array()

	fn := cjson.Object()
	fn.SetItem(c.Str("name"), cjson.String(c.Str("sqrt")))
	fn.SetItem(c.Str("sig"), cjson.String(c.Str("(x, /)")))
	syms.AddItem(fn)

	v := cjson.Object()
	v.SetItem(c.Str("name"), cjson.String(c.Str("pi")))
	syms.AddItem(v)

	mod.SetItem(c.Str("items"), syms)

	c.Printf(c.Str("%s\n"), mod.CStr())
	mod.Delete()
}

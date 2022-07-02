package cue

import (
	"fmt"
	"testing"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

func TestCtx(t *testing.T) {
	ctx := cuecontext.New()

	v := ctx.CompileString(`
		x: a
		a: 2
		b: 3
		"a+b": a + b
	`)

	p := func(format string, args ...interface{}) {
		fmt.Printf(format+"\n", args...)
	}

	p("lookups")
	p("a:     %v", v.LookupPath(cue.ParsePath("a")))
	p("b:     %v", v.LookupPath(cue.ParsePath("b")))
	p(`"a+b": %v`, v.LookupPath(cue.ParsePath(`"a+b"`)))
	p(`"x": %v`, v.LookupPath(cue.ParsePath(`"x"`)))
	p("")
	p("expressions")
	p("a + b: %v", ctx.CompileString("a + b", cue.Scope(v)))
	p("a * b: %v", ctx.CompileString("a * b", cue.Scope(v)))
}

func TestAllow(t *testing.T) {
	ctx := cuecontext.New()

	const file = `
a: [1, 2, ...int]

b: #Point
#Point: {
	x:  int
	y:  int
	z?: int
}

c: [string]: int

d: #C
#C: [>"m"]: int
`

	v := ctx.CompileString(file)

	a := v.LookupPath(cue.ParsePath("a"))
	fmt.Println("a allows:")
	fmt.Println("  index 4:       ", a.Allows(cue.Index(4)))
	fmt.Println("  any index:     ", a.Allows(cue.AnyIndex))
	fmt.Println("  any string:    ", a.Allows(cue.AnyString))

	b := v.LookupPath(cue.ParsePath("b"))
	fmt.Println("b allows:")
	fmt.Println("  field x:       ", b.Allows(cue.Str("x")))
	fmt.Println("  field z:       ", b.Allows(cue.Str("z")))
	fmt.Println("  field foo:     ", b.Allows(cue.Str("foo")))
	fmt.Println("  index 4:       ", b.Allows(cue.Index(4)))
	fmt.Println("  any string:    ", b.Allows(cue.AnyString))

	c := v.LookupPath(cue.ParsePath("c"))
	fmt.Println("c allows:")
	fmt.Println("  field z:       ", c.Allows(cue.Str("z")))
	fmt.Println("  field foo:     ", c.Allows(cue.Str("foo")))
	fmt.Println("  index 4:       ", c.Allows(cue.Index(4)))
	fmt.Println("  any string:    ", c.Allows(cue.AnyString))

	d := v.LookupPath(cue.ParsePath("d"))
	fmt.Println("d allows:")
	fmt.Println("  field z:       ", d.Allows(cue.Str("z")))
	fmt.Println("  field foo:     ", d.Allows(cue.Str("foo")))
	fmt.Println("  index 4:       ", d.Allows(cue.Index(4)))
	fmt.Println("  any string:    ", d.Allows(cue.AnyString))

}

func TestFormat(t *testing.T) {
	ctx := cuecontext.New()

	inputValue := ctx.CompileString(`c: 1+3`)
	fmt.Println(inputValue.LookupPath(cue.ParsePath("c")))

	v := ctx.CompileString(`
		a: 2 + b
		b: *3 | int
		s: "foo\nbar"
		dSpec: {
			kind: int
		}
		d: dSpec // define constraint
		d: {
        	kind:  c 
		}   // use value from outside
	`, cue.Scope(inputValue))

	fmt.Println("### ALL")
	fmt.Println(v)
	fmt.Println("---")
	fmt.Printf("%#v\n", v)
	fmt.Println("---")
	fmt.Printf("%+v\n", v)

	a := v.LookupPath(cue.ParsePath("a"))
	fmt.Println("\n### INT")
	fmt.Printf("%%v:   %v\n", a)
	fmt.Printf("%%05d: %05d\n", a)

	s := v.LookupPath(cue.ParsePath("s"))
	fmt.Println("\n### STRING")
	fmt.Printf("%%v: %v\n", s)
	fmt.Printf("%%s: %s\n", s)
	fmt.Printf("%%q: %q\n", s)

	v = ctx.CompileString(`
		#Def: a: [string]: int
		b: #Def
		b: a: {
			a: 3
			b: 3
		}
	`)
	b := v.LookupPath(cue.ParsePath("b.a"))
	fmt.Println("\n### DEF")
	fmt.Println(b)
	fmt.Println("---")
	// This will indicate that the result is closed by including a hidden
	// definition.
	fmt.Printf("%#v\n", b)
}

package pjsval

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"

	"github.com/achiku/varfmt"
	"github.com/lestrrat/go-jshschema"
	"github.com/lestrrat/go-jspointer"
	"github.com/lestrrat/go-jsval"
	"github.com/lestrrat/go-jsval/builder"
)

type validatorList []*jsval.JSVal

func (vl validatorList) Len() int {
	return len(vl)
}
func (vl validatorList) Swap(i, j int) {
	vl[i], vl[j] = vl[j], vl[i]
}
func (vl validatorList) Less(i, j int) bool {
	return vl[i].Name < vl[j].Name
}

// Generate validator source code
func Generate(in io.Reader, out io.Writer, pkg string) error {
	var m map[string]interface{}
	if err := json.NewDecoder(in).Decode(&m); err != nil {
		return err
	}
	validators := validatorList{}
	for k, v := range m["properties"].(map[string]interface{}) {
		ptr := v.(map[string]interface{})["$ref"].(string)
		resolver, err := jspointer.New(ptr[1:])
		if err != nil {
			return err
		}
		resolved, err := resolver.Get(m)
		if err != nil {
			return err
		}
		hsc := hschema.New()
		hsc.Extract(resolved.(map[string]interface{}))
		for _, link := range hsc.Links {
			var v *jsval.JSVal
			if link.Schema == nil {
				v = jsval.New()
				v.SetRoot(jsval.Any())
			} else {
				b := builder.New()
				v, err = b.BuildWithCtx(link.Schema, m)
				if err != nil {
					return err
				}
			}
			v.Name = varfmt.PublicVarName(k + varfmt.PublicVarName(link.Rel) + "Validator")
			validators = append(validators, v)
		}
	}
	g := jsval.NewGenerator()
	var src bytes.Buffer
	fmt.Fprintln(&src, "package "+pkg)
	fmt.Fprintf(&src, "import \"github.com/lestrrat/go-jsval\"")
	g.Process(&src, validators...)
	b, err := format.Source(src.Bytes())
	if err != nil {
		return err
	}
	out.Write(b)
	return nil
}

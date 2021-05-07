package email

import (
	"fmt"
	"text/template"
)

var tmpls map[string]*template.Template

func tmpl(path string, tmpl string) (*template.Template, error) {
	if tmpls == nil {
		tmpls = make(map[string]*template.Template)
	}

	f := fmt.Sprintf("%s/%s", path, tmpl)

	t := tmpls[f]
	if t == nil {
		var err error
		t, err = template.ParseFiles(f)
		if err != nil {
			return nil, ErrInvalidTemplate
		}

		tmpls[f] = t
	}

	return t, nil
}

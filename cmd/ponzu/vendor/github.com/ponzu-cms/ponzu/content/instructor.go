package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Instructor struct {
	item.Item

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Twitter   string `json:"twitter"`
	Email     string `json:"email"`
	Linkedin  string `json:"linkedin"`
}

func (i *Instructor) String() string {
	return i.FirstName + " " + i.LastName
}

// MarshalEditor writes a buffer of html to edit a Instructor within the CMS
// and implements editor.Editable
func (i *Instructor) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(i,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Instructor field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("FirstName", i, map[string]string{
				"label":       "FirstName",
				"type":        "text",
				"placeholder": "Enter the First Name here",
			}),
		},
		editor.Field{
			View: editor.Input("LastName", i, map[string]string{
				"label":       "LastName",
				"type":        "text",
				"placeholder": "Enter the Last Name here",
			}),
		},
		editor.Field{
			View: editor.Input("Twitter", i, map[string]string{
				"label":       "Twitter",
				"type":        "text",
				"placeholder": "Enter the Twitter here",
			}),
		},
		editor.Field{
			View: editor.Input("Email", i, map[string]string{
				"label":       "Email",
				"type":        "text",
				"placeholder": "Enter the Email here",
			}),
		},
		editor.Field{
			View: editor.Input("Linkedin", i, map[string]string{
				"label":       "Linkedin",
				"type":        "text",
				"placeholder": "Enter the Linkedin here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Instructor editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Instructor"] = func() interface{} { return new(Instructor) }
}

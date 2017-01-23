package content

import (
	"fmt"
	"net/http"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Job struct {
	item.Item

	Title        string `json:"title"`
	Description  string `json:"description"`
	Website      string `json:"website"`
	Requirements string `json:"requirements"`
	ContactPhone string `json:"contact_phone"`
	ContactEmail string `json:"contact_email"`
	ContactName  string `json:"contact_name"`
}

func (j *Job) Accept(r *http.Request) error {
	return nil
}

// MarshalEditor writes a buffer of html to edit a Job within the CMS
// and implements editor.Editable
func (j *Job) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(j,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Job field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", j, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Input("Description", j, map[string]string{
				"label":       "Description",
				"type":        "text",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.Input("Website", j, map[string]string{
				"label":       "Website",
				"type":        "text",
				"placeholder": "Enter the Website here",
			}),
		},
		editor.Field{
			View: editor.Input("Requirements", j, map[string]string{
				"label":       "Requirements",
				"type":        "text",
				"placeholder": "Enter the Requirements here",
			}),
		},
		editor.Field{
			View: editor.Input("ContactPhone", j, map[string]string{
				"label":       "ContactPhone",
				"type":        "text",
				"placeholder": "Enter the ContactPhone here",
			}),
		},
		editor.Field{
			View: editor.Input("ContactEmail", j, map[string]string{
				"label":       "ContactEmail",
				"type":        "text",
				"placeholder": "Enter the ContactEmail here",
			}),
		},
		editor.Field{
			View: editor.Input("ContactName", j, map[string]string{
				"label":       "ContactName",
				"type":        "text",
				"placeholder": "Enter the ContactName here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Job editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Job"] = func() interface{} { return new(Job) }
}

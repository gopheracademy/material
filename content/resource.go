package content

import (
	"fmt"
	"net/http"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Resource struct {
	item.Item

	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Category    string `json:"category"`
}

func (res *Resource) Accept(w http.ResponseWriter, r *http.Request) error {
	http.Redirect(w, r, "https://gopheracademy.com/resources/thanks", 301)
	return nil
}

func (res *Resource) Approve(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (res *Resource) AfterApprove(w http.ResponseWriter, r *http.Request) error {
	return nil

}

// MarshalEditor writes a buffer of html to edit a Resource within the CMS
// and implements editor.Editable
func (r *Resource) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(r,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Resource field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", r, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Input("Description", r, map[string]string{
				"label":       "Description",
				"type":        "text",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.Input("Link", r, map[string]string{
				"label":       "Link",
				"type":        "text",
				"placeholder": "Enter the Link here",
			}),
		},
		editor.Field{
			View: editor.Input("Category", r, map[string]string{
				"label":       "Category",
				"type":        "text",
				"placeholder": "Enter the Category here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Resource editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Resource"] = func() interface{} { return new(Resource) }
}

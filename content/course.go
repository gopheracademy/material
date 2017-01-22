package content

import (
	"fmt"

	"github.com/bosssauce/reference"
	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Course struct {
	item.Item

	Name       string   `json:"name"`
	Title      string   `json:"title"`
	Subtitle   string   `json:"subtitle"`
	VideoCode  string   `json:"video_code"`
	Instructor string   `json:"instructor"`
	Modules    []string `json:"modules"`
}
type CourseListResult struct {
	Data []Course `json:"data"`
}

type CourseResult struct {
	Course `json:"data"`
}

func (c *Course) String() string {
	return c.Name
}

// MarshalEditor writes a buffer of html to edit a Course within the CMS
// and implements editor.Editable
func (c *Course) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(c,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Course field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:

		editor.Field{
			View: editor.Input("Name", c, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},

		editor.Field{
			View: editor.Input("Title", c, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Input("Subtitle", c, map[string]string{
				"label":       "Subtitle",
				"type":        "text",
				"placeholder": "Enter the Subtitle here",
			}),
		},
		editor.Field{
			View: editor.Input("VideoCode", c, map[string]string{
				"label":       "Video Code",
				"type":        "text",
				"placeholder": "Enter the Video Codehere",
			}),
		},
		editor.Field{
			View: reference.SelectRepeater("Modules", c, map[string]string{
				"label":       "Module",
				"placeholder": "Enter the module here",
			},
				"Module",
				"{{ .name }}",
			),
		},
		editor.Field{
			View: reference.Select("Instructor", c, map[string]string{
				"label":       "Instructor",
				"placeholder": "Enter the instructor here",
			},
				"Instructor",
				"{{ .first_name }}",
			),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Course editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Course"] = func() interface{} { return new(Course) }
}

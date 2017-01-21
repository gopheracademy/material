package content

import (
	"fmt"

	"github.com/bosssauce/reference"
	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Module struct {
	item.Item

	Name        string   `json:"name"`
	Title       string   `json:"title"`
	Subtitle    string   `json:"subtitle"`
	Description string   `json:"description"`
	VideoCode   string   `json:"video_code"`
	Topic       string   `json:"topic"`
	Level       string   `json:"level"`
	Lessons     []string `json:"lessons"`
}

type ModuleResult struct {
	Data []Module `json:"data"`
}

func (m *Module) String() string {
	return m.Name
}

// MarshalEditor writes a buffer of html to edit a Modules within the CMS
// and implements editor.Editable
func (m *Module) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(m,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Modules field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:

		editor.Field{
			View: editor.Select("Level", m, map[string]string{
				"label":       "Level",
				"type":        "text",
				"placeholder": "Enter the Level here",
			}, map[string]string{
				"Beginner":     "Beginner",
				"Intermediate": "Intermediate",
				"Advanced":     "Advanced",
				"Expert":       "Expert",
			}),
		},
		editor.Field{
			View: editor.Input("Name", m, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},

		editor.Field{
			View: editor.Input("Title", m, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Input("Subtitle", m, map[string]string{
				"label":       "Subtitle",
				"type":        "text",
				"placeholder": "Enter the Subtitle here",
			}),
		},

		editor.Field{
			View: editor.Select("Topic", m, map[string]string{
				"label":       "Topic",
				"type":        "text",
				"placeholder": "Enter the Topic here",
			}, map[string]string{
				"Go":                    "Go",
				"Kubernetes":            "Kubernetes",
				"Git":                   "Git",
				"Docker":                "Docker",
				"Distributed Computing": "Distributed Computing",
			}),
		},
		editor.Field{
			View: editor.Input("Description", m, map[string]string{
				"label":       "Description",
				"type":        "text",
				"placeholder": "Enter the description here",
			}),
		},
		editor.Field{
			//	View: editor.Input("Lessons", m, map[string]string{
			//		"label":       "Lessons",
			//		"type":        "text",
			//		"placeholder": "Enter the Lessons here",
			//	}),
			View: reference.SelectRepeater("Lessons", m, map[string]string{
				"label":       "Lesson",
				"type":        "text",
				"placeholder": "Enter the Lesson here",
			},
				"Lesson",
				"{{ .name }}",
			),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Module editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Module"] = func() interface{} { return new(Module) }
}

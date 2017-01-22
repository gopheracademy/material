package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Lesson struct {
	item.Item

	Name         string `json:"name"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	Content      string `json:"content"`
	VideoCode    string `json:"video_code"`
	GithubLesson string `json:"github_lesson"`
	Thumb        string `json:"thumb"`
}

type LessonListResult struct {
	Data []Lesson `json:"data"`
}

type LessonResult struct {
	Course `json:"data"`
}

func (l *Lesson) String() string {
	return l.Name
}

// MarshalEditor writes a buffer of html to edit a Lesson within the CMS
// and implements editor.Editable
func (l *Lesson) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(l,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Lesson field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", l, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Input("Title", l, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Input("Subtitle", l, map[string]string{
				"label":       "Subtitle",
				"type":        "text",
				"placeholder": "Enter the Subtitle here",
			}),
		},
		editor.Field{
			View: editor.Input("VideoCode", l, map[string]string{
				"label":       "Video Code",
				"type":        "text",
				"placeholder": "Enter the Video Code here",
			}),
		},
		editor.Field{
			View: editor.Input("GithubLesson", l, map[string]string{
				"label":       "Github Lesson",
				"type":        "text",
				"placeholder": "Enter the URL to the Github Lesson Repository",
			}),
		},
		editor.Field{
			View: editor.Textarea("Content", l, map[string]string{
				"label":       "Content",
				"type":        "text",
				"placeholder": "Enter the Content here",
			}),
		},

		editor.Field{
			View: editor.File("Thumb", l, map[string]string{
				"label": "Thumb",
				"type":  "text",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Lesson editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Lesson"] = func() interface{} { return new(Lesson) }
}

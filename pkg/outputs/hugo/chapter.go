package hugo

import (
	"fmt"

	"github.com/feloy/kubernetes-api-reference/pkg/formats/markdown"
	"github.com/feloy/kubernetes-api-reference/pkg/outputs"
)

// Chapter of a Hugo output
// implements the outputs.Chapter interface
type Chapter struct {
	hugo *Hugo
	part *Part
	name string
}

// SetAPIVersion writes the APIVersion for a chapter
func (o Chapter) SetAPIVersion(s string) error {
	err := o.hugo.addContent(o.part.name, o.name, markdown.Code("apiVersion: "+s))
	if err != nil {
		return fmt.Errorf("Error adding GV for chapter %s/%s: %s", o.part.name, o.name, err)
	}
	return nil
}

// SetGoImport writes the Go import for a chapter
func (o Chapter) SetGoImport(s string) error {
	err := o.hugo.addContent(o.part.name, o.name, markdown.Code("import \""+s+"\""))
	if err != nil {
		return fmt.Errorf("Error adding Go Import for chapter %s/%s: %s", o.part.name, o.name, err)
	}
	return nil
}

// AddSection adds a section to the chapter
func (o Chapter) AddSection(i int, name string) (outputs.Section, error) {
	err := o.hugo.addSection(o.part.name, o.name, name)
	if err != nil {
		return Section{}, err
	}
	return Section{
		hugo:    o.hugo,
		part:    o.part,
		chapter: &o,
	}, nil
}

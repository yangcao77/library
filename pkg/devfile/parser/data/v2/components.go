package v2

import (
	v1 "github.com/devfile/api/pkg/apis/workspaces/v1alpha2"
	"github.com/devfile/parser/pkg/devfile/parser/data/v2/common"

	v210 "github.com/devfile/parser/pkg/devfile/parser/data/v2/2.1.0"
	v220 "github.com/devfile/parser/pkg/devfile/parser/data/v2/2.2.0"
)

// GetComponents returns the slice of Component objects parsed from the Devfile
func (d *DevfileV2) GetComponents() []v1.Component {
	return d.Components
}

// AddComponents adds the slice of Component objects to the devfile's components
// if a component is already defined, error out
func (d *DevfileV2) AddComponents(components []v1.Component) error {

	// different map for volume and container component as a volume and a container with same name
	// can exist in devfile
	containerMap := make(map[string]bool)
	volumeMap := make(map[string]bool)

	for _, component := range d.Components {
		if component.Volume != nil {
			volumeMap[component.Name] = true
		}
		if component.Container != nil {
			containerMap[component.Name] = true
		}
	}

	for _, component := range components {

		if component.Volume != nil {
			if _, ok := volumeMap[component.Name]; !ok {
				d.Components = append(d.Components, component)
			} else {
				return &common.FieldAlreadyExistError{Name: component.Name, Field: "component"}
			}
		}

		if component.Container != nil {
			if _, ok := containerMap[component.Name]; !ok {
				d.Components = append(d.Components, component)
			} else {
				return &common.FieldAlreadyExistError{Name: component.Name, Field: "component"}
			}
		}
	}
	return nil
}

// UpdateComponent updates the component with the given name
func (d *DevfileV2) UpdateComponent(component v1.Component) {
	index := -1
	for i := range d.Components {
		if d.Components[i].Name == component.Name {
			index = i
			break
		}
	}
	if index != -1 {
		d.Components[index] = component
	}
}

func (d *DevfileV2) GetCustomType210() string {

	// This feature was introduced in 210; so any version 210 and up should use the 210 implementation
	switch d.SchemaVersion {
	case "2.0.0":
		return ""
	default:
		return v210.GetCustomType210(d.Commands[1].Id)
	}

}

func (d *DevfileV2) GetCustomType220() string {

	// This feature was introduced in 220; so any version below 2.2.0 will return an empty obj
	switch d.SchemaVersion {
	case "2.2.0":
		return v220.GetCustomType220(d.Commands[0].Id, d.Commands[1].Id)
	default:
		return ""
	}
}

package parser

import (
	"fmt"
	"reflect"

	devfilev1 "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	"github.com/devfile/library/v2/pkg/devfile/parser/data/v2/common"
)

// GetDeployComponents gets the default deploy command associated components
func GetDeployComponents(devfileObj DevfileObj) (map[string]string, error) {
	deployCommandFilter := common.DevfileOptions{
		CommandOptions: common.CommandOptions{
			CommandGroupKind: devfilev1.DeployCommandGroupKind,
		},
	}
	deployCommands, err := devfileObj.Data.GetCommands(deployCommandFilter)
	if err != nil {
		return nil, err
	}

	deployAssociatedComponents := make(map[string]string)
	var deployAssociatedSubCommands []string

	for _, command := range deployCommands {
		if command.Apply != nil {
			if len(deployCommands) > 1 && command.Apply.Group.IsDefault != nil && !*command.Apply.Group.IsDefault {
				continue
			}
			deployAssociatedComponents[command.Apply.Component] = command.Apply.Component
		} else if command.Composite != nil {
			if len(deployCommands) > 1 && command.Composite.Group.IsDefault != nil && !*command.Composite.Group.IsDefault {
				continue
			}
			deployAssociatedSubCommands = append(deployAssociatedSubCommands, command.Composite.Commands...)
		}
	}

	applyCommandFilter := common.DevfileOptions{
		CommandOptions: common.CommandOptions{
			CommandType: devfilev1.ApplyCommandType,
		},
	}
	applyCommands, err := devfileObj.Data.GetCommands(applyCommandFilter)
	if err != nil {
		return nil, err
	}

	for _, command := range applyCommands {
		if command.Apply != nil {
			for _, deployCommand := range deployAssociatedSubCommands {
				if deployCommand == command.Id {
					deployAssociatedComponents[command.Apply.Component] = command.Apply.Component
				}
			}

		}
	}

	return deployAssociatedComponents, nil
}

// GetImageBuildComponent gets the image build component from the deploy associated components
func GetImageBuildComponent(devfileObj DevfileObj, deployAssociatedComponents map[string]string) (devfilev1.Component, error) {
	imageComponentFilter := common.DevfileOptions{
		ComponentOptions: common.ComponentOptions{
			ComponentType: devfilev1.ImageComponentType,
		},
	}

	imageComponents, err := devfileObj.Data.GetComponents(imageComponentFilter)
	if err != nil {
		return devfilev1.Component{}, err
	}

	var imageBuildComponent devfilev1.Component
	for _, component := range imageComponents {
		if _, ok := deployAssociatedComponents[component.Name]; ok && component.Image != nil {
			if reflect.DeepEqual(imageBuildComponent, devfilev1.Component{}) {
				imageBuildComponent = component
			} else {
				errMsg := "expected to find one devfile image component with a deploy command for build. Currently there is more than one image component"
				return devfilev1.Component{}, fmt.Errorf(errMsg)
			}
		}
	}

	// If there is not one image component defined in the deploy command, err out
	if reflect.DeepEqual(imageBuildComponent, devfilev1.Component{}) {
		errMsg := "expected to find one devfile image component with a deploy command for build. Currently there is no image component"
		return devfilev1.Component{}, fmt.Errorf(errMsg)
	}

	return imageBuildComponent, nil
}

package testingutil

import (
	"errors"
	"fmt"
	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/yaml"
	// "sigs.k8s.io/controller-runtime/pkg/client"
	"context"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
)

type FakeK8sClient struct {
	// client.Client         // To satisfy interface; override all used methods
	clientset.Clientset
	DevWorkspaceResources map[string]v1alpha2.DevWorkspaceTemplate
	Errors                map[string]string
}

type FakeK8sClientInterface struct {
	rest.Interface
	DevWorkspaceResources map[string]v1alpha2.DevWorkspaceTemplate
	Errors                map[string]string
}

type FakeRequest struct {
	rest.Request
	NameValue             string
	NamespaceValue        string
	DevWorkspaceResources map[string]v1alpha2.DevWorkspaceTemplate
	Errors                map[string]string
}

func (client *FakeK8sClient) RESTClient() FakeK8sClientInterface {
	return FakeK8sClientInterface{
		DevWorkspaceResources: client.DevWorkspaceResources,
		Errors:                client.Errors,
	}
}

func (inter FakeK8sClientInterface) Get() *FakeRequest {
	return &FakeRequest{
		DevWorkspaceResources: inter.DevWorkspaceResources,
		Errors:                inter.Errors,
	}
}

func (request *FakeRequest) Namespace(namespace string) *FakeRequest {
	return &FakeRequest{
		NameValue:             request.NameValue,
		NamespaceValue:        namespace,
		DevWorkspaceResources: request.DevWorkspaceResources,
		Errors:                request.Errors,
	}
}

func (request *FakeRequest) Name(name string) *FakeRequest {
	return &FakeRequest{
		NameValue:             name,
		NamespaceValue:        request.NamespaceValue,
		DevWorkspaceResources: request.DevWorkspaceResources,
		Errors:                request.Errors,
	}
}

func (request *FakeRequest) DoRaw(_ context.Context) ([]byte, error) {
	if element, ok := request.DevWorkspaceResources[request.NameValue]; ok {
		yamlData, err := yaml.Marshal(element)
		if err != nil {
			return nil, fmt.Errorf("called Get() in fake client with error")
		}
		return yamlData, nil
	}
	if err, ok := request.Errors[request.NameValue]; ok {
		return nil, errors.New(err)
	}

	return nil, fmt.Errorf("test does not define an entry for %s", request.NameValue)

}

//func (client *FakeK8sClient) Get(_ context.Context, namespacedName client.ObjectKey, obj runtime.Object) error {
//	template, ok := obj.(*v1alpha2.DevWorkspaceTemplate)
//	if !ok {
//		return fmt.Errorf("called Get() in fake client with non-DevWorkspaceTemplate")
//	}
//	if element, ok := client.DevWorkspaceResources[namespacedName.Name]; ok {
//		*template = element
//		return nil
//	}
//
//	if err, ok := client.Errors[namespacedName.Name]; ok {
//		return errors.New(err)
//	}
//	return fmt.Errorf("test does not define an entry for %s", namespacedName.Name)
//}

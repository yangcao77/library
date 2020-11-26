package v2

import (
	"reflect"
	"testing"

	v1 "github.com/devfile/api/pkg/apis/workspaces/v1alpha2"
	devfilepkg "github.com/devfile/api/pkg/devfile"
)

func TestDevfile200_GetSchemaVersion(t *testing.T) {

	type args struct {
		name string
	}
	tests := []struct {
		name                  string
		expectedSchemaVersion string
		devfilev2             *DevfileV2
	}{
		{
			name: "case 1: Get the schema version",
			devfilev2: &DevfileV2{
				v1.Devfile{
					DevfileHeader: devfilepkg.DevfileHeader{
						SchemaVersion: "1.0.0",
					},
				},
			},
			expectedSchemaVersion: "1.0.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			version := tt.devfilev2.GetSchemaVersion()
			if version != tt.expectedSchemaVersion {
				t.Errorf("TestDevfile200_GetSchemaVersion error - schema version did not match. Expected %s, got %s", tt.expectedSchemaVersion, version)
			}
		})
	}
}

func TestDevfile200_SetSchemaVersion(t *testing.T) {

	type args struct {
		name string
	}
	tests := []struct {
		name              string
		schemaVersion     string
		devfilev2         *DevfileV2
		expectedDevfilev2 *DevfileV2
	}{
		{
			name: "case 1: empty header",
			devfilev2: &DevfileV2{
				v1.Devfile{
					DevfileHeader: devfilepkg.DevfileHeader{},
				},
			},
			schemaVersion: "2.0.0",
			expectedDevfilev2: &DevfileV2{
				v1.Devfile{
					DevfileHeader: devfilepkg.DevfileHeader{
						SchemaVersion: "2.0.0",
					},
				},
			},
		},
		{
			name: "case 2: override existing header",
			devfilev2: &DevfileV2{
				v1.Devfile{
					DevfileHeader: devfilepkg.DevfileHeader{
						SchemaVersion: "1.0.0",
					},
				},
			},
			schemaVersion: "2.0.0",
			expectedDevfilev2: &DevfileV2{
				v1.Devfile{
					DevfileHeader: devfilepkg.DevfileHeader{
						SchemaVersion: "2.0.0",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.devfilev2.SetSchemaVersion(tt.schemaVersion)
			if !reflect.DeepEqual(tt.devfilev2, tt.expectedDevfilev2) {
				t.Errorf("TestDevfile200_SetSchemaVersion() expected %v, got %v", tt.expectedDevfilev2, tt.devfilev2)
			}
		})
	}
}

func TestDevfile200_SetSetMetadata(t *testing.T) {

	type args struct {
		name string
	}
	tests := []struct {
		name              string
		metadataName      string
		metadataVersion   string
		devfilev2         *DevfileV2
		expectedDevfilev2 *DevfileV2
	}{
		{
			name: "case 1: empty header",
			devfilev2: &DevfileV2{
				v1.Devfile{
					DevfileHeader: devfilepkg.DevfileHeader{},
				},
			},
			metadataName:    "nodejs",
			metadataVersion: "2.0.0",
			expectedDevfilev2: &DevfileV2{
				v1.Devfile{
					DevfileHeader: devfilepkg.DevfileHeader{
						Metadata: devfilepkg.DevfileMetadata{
							Name:    "nodejs",
							Version: "2.0.0",
						},
					},
				},
			},
		},
		{
			name: "case 2: override existing header",
			devfilev2: &DevfileV2{
				v1.Devfile{
					DevfileHeader: devfilepkg.DevfileHeader{
						SchemaVersion: "2.0.0",
					},
				},
			},
			metadataName:    "nodejs",
			metadataVersion: "2.0.0",
			expectedDevfilev2: &DevfileV2{
				v1.Devfile{
					DevfileHeader: devfilepkg.DevfileHeader{
						SchemaVersion: "2.0.0",
						Metadata: devfilepkg.DevfileMetadata{
							Name:    "nodejs",
							Version: "2.0.0",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.devfilev2.SetMetadata(tt.metadataName, tt.metadataVersion)
			if !reflect.DeepEqual(tt.devfilev2, tt.expectedDevfilev2) {
				t.Errorf("TestDevfile200_SetSchemaVersion() expected %v, got %v", tt.expectedDevfilev2, tt.devfilev2)
			}
		})
	}
}

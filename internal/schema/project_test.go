package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"terraform-provider-capella/internal/errors"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TestProjectSchemaValidate(t *testing.T) {
	type test struct {
		name                   string
		input                  Project
		expectedProjectId      string
		expectedOrganizationId string
		expectedErr            error
	}

	tests := []test{
		{
			name: "[POSITIVE] organization ID and project ID are passed via terraform apply",
			input: Project{
				Id:             basetypes.NewStringValue("100"),
				OrganizationId: basetypes.NewStringValue("200"),
			},
			expectedProjectId:      "100",
			expectedOrganizationId: "200",
		},
		{
			name: "[POSITIVE] IDs are passed via terraform import",
			input: Project{
				Id: basetypes.NewStringValue("id=100,organization_id=200"),
			},
			expectedProjectId:      "100",
			expectedOrganizationId: "200",
		},
		{
			name: "[NEGATIVE] only project ID is passed via terraform apply",
			input: Project{
				Id: basetypes.NewStringValue("100"),
			},
			expectedErr: errors.ErrInvalidImport,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			IDs, err := test.input.Validate()

			if test.expectedErr != nil {
				assert.ErrorContains(t, err, test.expectedErr.Error())
				return
			}

			assert.Equal(t, test.expectedProjectId, IDs[Id])
			assert.Equal(t, test.expectedOrganizationId, IDs[OrganizationId])
		})
	}
}
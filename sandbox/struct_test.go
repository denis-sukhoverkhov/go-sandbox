package sandbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangeField(t *testing.T) {
	tests := []struct {
		name      string
		data      interface{}
		fieldName string
		newValue  interface{}
		expected  interface{}
		shouldErr bool
	}{
		{
			name: "Change string field",
			data: &struct {
				Name string
				Type string
			}{Name: "name1", Type: "type1"},
			fieldName: "Name",
			newValue:  "new name",
			expected: &struct {
				Name string
				Type string
			}{Name: "new name", Type: "type1"},
			shouldErr: false,
		},
		{
			name: "Change map field",
			data: &struct {
				ID  string
				Arr []int
				Map map[string]interface{}
				Any interface{}
			}{ID: "id1", Arr: []int{1, 2, 3}, Map: map[string]interface{}{"key": "value"}, Any: "any"},
			fieldName: "Map",
			newValue:  map[string]interface{}{"test1": 5, "test2": 9},
			expected: &struct {
				ID  string
				Arr []int
				Map map[string]interface{}
				Any interface{}
			}{ID: "id1", Arr: []int{1, 2, 3}, Map: map[string]interface{}{"test1": 5, "test2": 9}, Any: "any"},
			shouldErr: false,
		},
		{
			name: "Invalid field name",
			data: &struct {
				ID  string
				Arr []int
				Map map[string]interface{}
				Any interface{}
			}{ID: "id1", Arr: []int{1, 2, 3}, Map: map[string]interface{}{"key": "value"}, Any: "any"},
			fieldName: "NonExistentField",
			newValue:  "some value",
			shouldErr: true,
		},
		{
			name: "Type mismatch",
			data: &struct {
				Name string
			}{Name: "name1"},
			fieldName: "Name",
			newValue:  123, // This should cause an error
			shouldErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ChangeStructFieldValue(tt.data, tt.fieldName, tt.newValue)
			if tt.shouldErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expected, tt.data)
			}
		})
	}
}

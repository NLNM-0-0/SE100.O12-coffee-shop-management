package ingredientmodel

import (
	"backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIngredientUpdatePrice_TableName(t *testing.T) {
	type fields struct {
		Price float32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of IngredientUpdateAmount successfully",
			fields: fields{
				Price: float32(0),
			},
			want: common.TableIngredient,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ingredient := &IngredientUpdatePrice{
				Price: &tt.fields.Price,
			}
			got := ingredient.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

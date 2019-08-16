package expr

import (
	"encoding/json"
	"testing"

	_expr "github.com/antonmedv/expr"

	"github.com/stretchr/testify/assert"
)

func TestExprJSON(t *testing.T) {
	var stuff map[string]interface{}
	err := json.Unmarshal([]byte(`
	{
		"name": "Robert",
		"age": 42
	}
	`), &stuff)
	assert.NoError(t, err)
	r, err := _expr.Eval(`"Hello " + name`, stuff)
	assert.NoError(t, err)
	assert.Equal(t, "Hello Robert", r)
	r, err = _expr.Eval(`age -2`, stuff)
	assert.NoError(t, err)
	assert.Equal(t, float64(40), r)
}

func TestNil(t *testing.T) {
	r, err := _expr.Eval(`a`, map[string]interface{}{})
	assert.NoError(t, err)
	assert.Equal(t, nil, r)
}

func TestFunc(t *testing.T) {
	r, err := _expr.Eval(`add(41, 1)`, map[string]interface{}{
		"add": func(i, j int) interface{} {
			return i + j
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, 42, r)
}

func TestDefault(t *testing.T) {
	// only ternary operator can handle nil
	r, err := _expr.Eval(`a != nil ? a : 42`, map[string]interface{}{})
	assert.NoError(t, err)
	assert.Equal(t, 42, r)
	r, err = _expr.Eval(`default(a, 42)`, map[string]interface{}{
		"default": func(v, def interface{}) interface{} {
			if v == nil {
				return def
			}
			return v
		},
	})
	assert.Error(t, err)
}

package rule

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

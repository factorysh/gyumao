package states

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStates(t *testing.T) {
	states := &States{}
	state := State{
		id:     "1",
		values: nil,
	}
	states.Set(&state)
	test := states.Get("1")
	assert.Equal(t, &state, test)
	assert.Equal(t, []string{"1"}, states.All())
}

func TestState(t *testing.T) {
	state := State{
		id:     "1",
		values: nil,
	}
	assert.Equal(t, "1", state.Id())
	state.Set("1", []string{"test1"})
	state.Set("2", []string{"test2"})
	state.Set("3", []string{"test3"})
	assert.Equal(t, []string{"test2"}, state.Get("2"))
	assert.Equal(t, []string{"test1"}, state.Get("1"))
	assert.Equal(t, []string{"test3"}, state.Get("3"))
	assert.Equal(t, []string{"1", "2", "3"}, state.Keys())
}

package states

/*
Manage context, the REST way



GET /environment/{collection}/{id}/{key}
	Value is raw JSON, aka interface{}

GET /environment/{collection}/{id}
	map[string]interface{}
*/

// type States interface {
// 	Get(id string) *State
// 	Set(state *State)
// 	All() []string // id
// }
//
// type State interface {
// 	Id() string // primary key
// 	Get(key string) interface{}
// 	Set(key string, value interface{})
// 	Keys() []string
// }

type States struct {
	states	map[string]*State
}

type State struct {
	id    	string
	values	map[string]interface{}
}

func (s *States) Get(id string) *State {
	return s.states[id]
}

func (s *States) Set(state *State) {
	s.states[state.Id()] = state
}

func (s *States) All() []string {
	var ids []string
	for id := range s.states {
		ids = append(ids, id)
	}
	return ids
}

func (s *State) Id() string {
	return s.id
}

func (s *State) Keys() []string {
	var keys []string
	for key := range s.values {
		keys = append(keys, key)
	}
	return keys
}

func (s *State) Get(key string) (interface{}) {
	return s.values[key]
}

func (s *State) Set(key string, value interface{}) {
	s.values[key] = value
}

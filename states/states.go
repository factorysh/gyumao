package states

/*
Manage context, the REST way



GET /environment/{collection}/{id}/{key}
	Value is raw JSON, aka interface{}

GET /environment/{collection}/{id}
	map[string]interface{}
*/

/*
type States interface {
	Get(id string) *State
	Set(state *State)
	All() []string // id
}

type State interface {
	Id() string // primary key
	Get(key string) interface{}
	Set(key string, value interface{})
	Keys() []string
}
*/

type States struct {
	States	map[string]*State
}

type State struct {
	Id    	string
	Values	map[string]interface{}
}

func (s *States) Get(id string) *State {
	return s.States[id]
}

func (s *States) Set(state *State) {
	if s.States == nil {
		s.States = make(map[string]*State)
	}
	s.States[state.GetId()] = state
}

func (s *States) All() []string {
	var ids []string
	for id := range s.States {
		ids = append(ids, id)
	}
	return ids
}

func (s *State) GetId() string {
	return s.Id
}

func (s *State) Keys() []string {
	var keys []string
	for key := range s.Values {
		keys = append(keys, key)
	}
	return keys
}

func (s *State) Get(key string) (interface{}) {
	return s.Values[key]
}

func (s *State) Set(key string, value interface{}) {
	if s.Values == nil {
		s.Values = make(map[string]interface{})
	}
	s.Values[key] = value
}

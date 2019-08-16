package states

/*
Manage context, the REST way



GET /environment/{collection}/{id}/{key}
	Value is raw JSON, aka interface{}

GET /environment/{collection}/{id}
	map[string]interface{}
*/

type States interface {
	Get(name string) *State
	Set(state *State)
}

type State interface {
	Id() string // primary key
	Get(key string) interface{}
	Set(key string, value interface{})
	Keys() []string
}

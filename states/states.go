package states

/*

/environment/{collection}/{key}

Value is raw JSON, aka interface{}

*/

type States struct {
	states map[string]map[string]interface{}
}

func (s *States) Get(key string) map[string]interface{} {
	return s.states[key]
}

package statesrest

/*
Manage context, the REST way



GET /environment/{collection}/{id}/{key}
	Value is raw JSON, aka interface{}

GET /environment/{collection}/{id}
	map[string]interface{}
*/

import (
	"encoding/json"
	"net/http"

	"github.com/factorysh/gyumao/states"
	"github.com/factorysh/gyumao/statesbolt"
	"github.com/gorilla/mux"
)

// STATE_TEST
// func init() {
// 	db, err := statesbolt.New("states_db")
// 	if err != nil {
// 		fmt.Println("fail to open db")
// 	}
// 	defer db.Close()
// 	value := make(map[string]interface{})
// 	value["test1"] = []string{"test", "test", "test"}
// 	s := &states.States{}
// 	state := states.State{
// 		Id:     "1",
// 		Values: value,
// 	}
// 	state.Set("key1", []string{"test", "test", "test"})
// 	state.Set("key2", []string{"test", "test", "test"})
// 	state.Set("key3", []string{"test", "test", "test"})
// 	s.Set(&state)
// 	err = db.Set("collection", s)
// 	if err != nil {
// 		fmt.Println("db set failed :", err)
// 	}
// }

// ServeHTTP is a HTTP handler implementation
func StatesRESTHandlerId(w http.ResponseWriter, r *http.Request) {
	db, err := statesbolt.New("states_db")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer db.Close()
	params := mux.Vars(r)
	collection, err := db.Get(params["collection"])
	if err != nil || collection == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	s := collection.(states.States)
	result := s.Get(params["id"])
	if result == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(result.Values)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// ServeHTTP is a HTTP handler implementation
func StatesRESTHandlerKey(w http.ResponseWriter, r *http.Request) {
	db, err := statesbolt.New("states_db")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer db.Close()
	params := mux.Vars(r)
	collection, err := db.Get(params["collection"])
	if err != nil || collection == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	s := collection.(states.States)
	result := s.Get(params["id"])
	if result == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key := result.Get(params["key"])
	if key == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

package main

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
)

// ProtoToMap roundtrips a proto object to JSON
// and back to map[string]interface{} for Firestore.
func ProtoToMap(v interface{})  map[string]interface{} {
	b, _ := json.Marshal(v)  // intentionally unhandled
	var out map[string]interface{}
	_ = json.Unmarshal(b, &out) // intentionally unhandled
	return out
}

// add: PROTO --> JSON --> map --> firestore
// get: firestore --> map --> JSON --> PROTO

// ProtoToMap roundtrips a proto object to JSON
// and back to map[string]interface{} for Firestore.
func MapToProto(v map[string]interface{}, out proto.Message)  {
	b, _ := json.Marshal(v)
	_ = json.Unmarshal(b, out) // ignored intentionally
}

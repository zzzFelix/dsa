package collection_test

import (
	"encoding/json"
	"testing"

	"github.com/zzzFelix/dsa/collection"
)

type ctype struct {
	Set collection.CSet[string] `json:"set"`
}

func TestUnmarshal(t *testing.T) {
	byt := []byte(`{"set":["num","strs"]}`)
	var dat ctype
	if err := json.Unmarshal(byt, &dat); err != nil {
		t.Fatalf("unmarshalling failed: %v", err)
	}
	if dat.Set.String() != "[num strs]" {
		t.Fatal("unexpected result.")
	}
}

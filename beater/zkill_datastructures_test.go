package beater

import (
	"testing"
	"encoding/json"
	"os"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestDatastructureLoad (t *testing.T) {
	f, err := os.Open("../tests/data/example_1.json")
	assert.Nil(t, err)
	assert.NotNil(t, f)
	var ret ZkillPackage
	err = json.NewDecoder(f).Decode(&ret)
	b, err := json.MarshalIndent(ret, "","\t")
	assert.Nil(t, err)
	assert.NotNil(t, ret)
	fmt.Println(string(b))
	assert.NotNil(t, nil)
}
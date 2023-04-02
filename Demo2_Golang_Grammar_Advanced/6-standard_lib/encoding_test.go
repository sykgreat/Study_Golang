package __standard_lib

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"testing"
)

func Test_Encoding(t *testing.T) {
	type user struct {
		ID   int
		Name string
		Age  int
	}

	u1 := user{ID: 1, Name: "nick", Age: 18}
	marshal, err := json.Marshal(u1)
	t.Log(marshal, err)
	var u2 user
	err = json.Unmarshal(marshal, &u2)
	t.Log(u2, err)

	// base64编解码
	toString := base64.StdEncoding.EncodeToString(marshal)
	t.Log(toString)
	decodeString, err := base64.StdEncoding.DecodeString(toString)
	t.Log(decodeString, err)

	// 16进制编解码
	encodeToString := hex.EncodeToString(marshal)
	t.Log(encodeToString)
	decodeString, err = hex.DecodeString(encodeToString)
	t.Log(decodeString, err)
}

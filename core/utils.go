package core

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// EncodingJSON export
func EncodingJSON(value interface{}) []byte {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

// DecodingJSON export
func DecodingJSON(value []byte, arg interface{}) {
	if err := json.Unmarshal(value, &arg); err != nil {
		panic(err)
	}
}

// BindMap export
func BindMap(c *gin.Context) (m map[string]interface{}) {
	decoder := json.NewDecoder(c.Request.Body)
	defer c.Request.Body.Close()
	_ = decoder.Decode(&m)
	return m
}

// GetJSON
func GetJSON(c *gin.Context, param interface{}) {
	if err := c.ShouldBindJSON(&param); err != nil {
		panic(err)
	}
}

package redisCache

import (
	"encoding/json"
	"github.com/vmihailenco/bufpool"
)

const (
	compressionThreshold = 64
	timeLen              = 4
)

var bp = bufpool.Pool{}

func Marshal(value interface{}) ([]byte, error) {
	return _marshal(value)
}

func _marshal(value interface{}) ([]byte, error) {
	switch value := value.(type) {
	case nil:
		return nil, nil
	case []byte:
		return value, nil
	case string:
		return []byte(value), nil
	}

	//buf := bp.Get()
	//defer bp.Put(buf)
	//
	//enc := msgpack.GetEncoder()
	//enc.Reset(buf)
	//enc.UseCompactInts(true)
	//
	//err := enc.Encode(value)
	//
	//msgpack.PutEncoder(enc)
	//if err != nil {
	//	return nil, err
	//}
	//return buf.Bytes(), nil

	/********默认使用json**********/
	return json.Marshal(value)
}

func Unmarshal(b []byte, value interface{}) error {
	return _unmarshal(b, value)
}

func _unmarshal(b []byte, value interface{}) error {
	if len(b) == 0 {
		return nil
	}

	switch value := value.(type) {
	case nil:
		return nil
	case *[]byte:
		clone := make([]byte, len(b))
		copy(clone, b)
		*value = clone
		return nil
	case *string:
		*value = string(b)
		return nil
	}
	//return msgpack.Unmarshal(b, value)

	/********默认使用json**********/
	return json.Unmarshal(b, value)
}

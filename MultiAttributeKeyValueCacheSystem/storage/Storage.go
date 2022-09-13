package storage

import (
	"errors"
	"fmt"
	"reflect"
)

type Storage struct {
	KeyStore map[string][]AttributeValue
	//{"delhi":[{"pollution-level": "very high","population":      "10 million"}]}
	AttributeType map[string]reflect.Type
	//{"pollution-level":"string","population":"string"}
	AttributeStore map[string]map[string]interface{}
	//{"pollution-level": {"delhi":"high"}}
}

type AttributeValue struct {
	AttributeName string
	Value         interface{}
}

func NewStorage() Storage {
	return Storage{
		KeyStore:       make(map[string][]AttributeValue),
		AttributeStore: make(map[string]map[string]interface{}),
		AttributeType:  make(map[string]reflect.Type),
	}
}

func (s *Storage) Add(key string, value map[string]interface{}) error {
	var attrValues []AttributeValue
	for attrName, attrVal := range value {
		attrType := reflect.TypeOf(attrVal)
		if attp, ok := s.AttributeType[attrName]; ok && attrType != attp {
			return fmt.Errorf("invalid attribute value type got : %s , want : %s", attrType, attp)
		}
	}
	for attrName, attrVal := range value {
		attrType := reflect.TypeOf(attrVal)
		attrValues = append(attrValues, AttributeValue{AttributeName: attrName, Value: attrVal})
		s.AttributeType[attrName] = attrType
		mp, ok := s.AttributeStore[attrName]
		if !ok {
			mp = make(map[string]interface{})
		}
		mp[key] = attrVal
		s.AttributeStore[attrName] = mp
	}
	s.KeyStore[key] = attrValues
	return nil
}

func (s *Storage) GetValueByKey(key string) ([]AttributeValue, error) {
	val, ok := s.KeyStore[key]
	if ok {
		return val, nil
	}
	return nil, errors.New("key not present")
}

func (s *Storage) GetKeyByAttributeValue(attribute string, value interface{}) ([]string, error) {
	val, ok := s.AttributeStore[attribute]
	if !ok {
		return nil, errors.New("attribute not present")
	}
	if val, ok := s.AttributeType[attribute]; ok && val != reflect.TypeOf(value) {
		return nil, errors.New("invalid type of attribute value")
	}
	var result []string
	for k, v := range val {
		if v == value {
			result = append(result, k)
		}
	}
	return result, nil
}

func (s *Storage) DeleteKey(key string) {
	delete(s.KeyStore, key)
	for _, v := range s.AttributeStore {
		delete(v, key)
	}
}

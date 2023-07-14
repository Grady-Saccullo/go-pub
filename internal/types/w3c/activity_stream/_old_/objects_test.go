package _old_

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"
)

//func TestObject_LoadJson(t *testing.T) {
//
//	jsonData := []byte(`
//	{
//		"@context": ["https://hackerman-labratories.com"],
//		"type": "Note",
//		"actor": {
//			"type": "Application",
//			"meh": 123
//		}
//	}
//`)
//
//	o, err := NewObjectFromJson(jsonData)
//	if err != nil {
//		t.Fatalf(err.Error())
//	}
//
//	if err := o.ValidateType(); err != nil {
//		t.Fatalf(err.Error())
//	}
//
//	a := Activity{}
//
//	if err := o.GetExtension(&a); err != nil {
//		t.Fatalf(err.Error())
//	}
//	data, _ := json.MarshalIndent(a, "", "  ")
//
//	t.Log(string(data))
//}

func TestNewArticleFromJson(t *testing.T) {
	S()
}

/**
* TODO: possibly always convert to slice on non functional types so that all we need to do is loop over
*  data and do introspection onto each item aka (GetObject, GetString, GetActor, GetLink)? These
*  functions could return *value. Something like the below
 */

type SomeInterface interface {
	GetObject() *Object
	GetActor() *Actor
	GetLink() *Link
	GetString() *string
}

type SomeStruct struct {
	SomeValue []SomeInterface
}

func SomeHandler(si SomeStruct) {
	for _, v := range si.SomeValue {
		if subV := v.GetLink(); subV != nil {
			// do something
		} else if subV := v.GetObject(); subV != nil {
			// do something
		} else if subV := v.GetActor(); subV != nil {
			// do something
		} else if subV := v.GetString(); subV != nil {
			// do something
		} else {
			panic("unknown object")
		}
	}
}

func S() {
	a := PropertyAttributedTo2[Actor]{
		AttributedTo: &Actor{
			ObjectCore: ObjectCore{},
			Type:       "Application",
		},
	}
	b := Actor{}
	Get(a, &b)
	data, _ := json.MarshalIndent(b, "", "  ")
	log.Println(string(data))
}

type PropertyAttributedTo2[T AttributedTo] struct {
	AttributedTo *T `json:"attributedTo"`
}

type AttributedTo interface {
	Actor | Object | Link | string | []interface{}
}

func Get[A interface{}, T AttributedTo](v A, newType *T) {
	it := reflect.TypeOf(newType)
	s := reflect.ValueOf(v)
	for i := 0; i < s.NumField(); i++ {
		sf := s.Field(i)
		if sf.Type() == it {
			fmt.Println(sf.Type())
			tmp := sf.Interface().(*T)
			*newType = *tmp
		}
	}
}

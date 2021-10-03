package di

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fatih/structtag"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

type ITest interface {
	Get() string
}
type testStruct struct {
	T string
}

func (s *testStruct) Ctor() {
	s.T = "Ctor"
}

func (s *testStruct) Get() string {
	return s.T
}

func (s *testStruct) Set(name string) {
	s.T = name
}

func TestReflectionCallMethod(t *testing.T) {
	rtPtr := reflect.TypeOf(&testStruct{})
	rtElem := rtPtr.Elem()
	fmt.Printf("rtPtr=%v, rtElem=%v", rtPtr, rtElem)

	method, ok := rtPtr.MethodByName("Set")
	assert.True(t, ok)
	fmt.Printf("method=%v", method)

	method, ok = rtPtr.MethodByName("Ctor")
	assert.True(t, ok)
	fmt.Printf("method=%v", method)

	obj := &testStruct{}

	rtSetValue := reflect.ValueOf(obj).MethodByName("Set")

	inputs := make([]reflect.Value, 1)
	inputs[0] = reflect.ValueOf("dog")
	rtSetValue.Call(inputs)
	assert.Equal(t, "dog", obj.Get())

	inputs[0] = reflect.ValueOf("cat")
	_, err := invoke(obj, "Set", "cat")
	assert.NoError(t, err)
	assert.Equal(t, "cat", obj.Get())

	_, err = invoke(obj, "Ctor")
	assert.NoError(t, err)
	assert.Equal(t, "Ctor", obj.Get())

}

func TestReflectionTypeConvert(t *testing.T) {
	type t2 struct {
		Test       ITest       `inject:"ITest2"`
		Tests      []ITest     `inject:""`
		TestStruct *testStruct `inject:"testStruct"`
		T          string      `json:"foo,omitempty,string" xml:"foo"`
	}
	rtT := reflect.TypeOf((*ITest)(nil))
	rtTElem := rtT.Elem()

	dst := &t2{}
	v := reflect.ValueOf(dst).Elem()
	f := v.FieldByName("Test")
	if f.IsValid() {
		// A Value can be changed only if it is
		// addressable and was not obtained by
		// the use of unexported struct fields.
		if f.CanSet() {
			obj := &testStruct{
				T: "hi",
			}
			objValue := reflect.ValueOf(obj)
			f.Set(objValue)
		}
	}
	f = v.FieldByName("Tests")
	if f.IsValid() {
		// A Value can be changed only if it is
		// addressable and was not obtained by
		// the use of unexported struct fields.
		if f.CanSet() {

			d := []ITest{
				&testStruct{
					T: "hi",
				},
			}
			objValue := reflect.ValueOf(d)
			f.Set(objValue)
		}
	}
	assert.Equal(t, "hi", dst.Test.Get())
	assert.Equal(t, "hi", dst.Tests[0].Get())

	var dest = &t2{}
	v2 := reflect.ValueOf(dest).Elem()
	field := v2.FieldByName("Test")
	var ts interface{} = &testStruct{
		T: "hi",
	}
	tsV := reflect.ValueOf(ts)
	field.Set(tsV.Convert(field.Type()))

	field = v2.FieldByName("Tests")
	sliceType := reflect.SliceOf(rtTElem)
	sliceV := reflect.New(sliceType).Elem()
	sliceV = reflect.Append(sliceV, tsV)
	field.Set(sliceV)

	assert.Equal(t, "hi", dest.Test.Get())
	assert.Equal(t, "hi", dest.Tests[0].Get())

	f = v.FieldByName("Tests")
	if f.IsValid() {
		// A Value can be changed only if it is
		// addressable and was not obtained by
		// the use of unexported struct fields.
		if f.CanSet() {
			sliceType := reflect.SliceOf(rtTElem)
			sliceV := reflect.New(sliceType).Elem()

			d := []interface{}{
				&testStruct{
					T: "hi",
				},
			}
			for _, v := range d {
				tsV := reflect.ValueOf(v)
				sliceV = reflect.Append(sliceV, tsV)
			}

			f.Set(sliceV)
		}
	}
	assert.Equal(t, "hi", dst.Tests[0].Get())
}
func TestTypeTags(t *testing.T) {

	type t2 struct {
		Test2      ITest       `inject:"ITest2"`
		Tests      []ITest     `inject:""`
		TestStruct *testStruct `inject:"testStruct"`
		T          string      `json:"foo,omitempty,string" xml:"foo"`
	}

	rtT := reflect.TypeOf((*ITest)(nil))
	rtTElem := rtT.Elem()
	rtTestStruct := reflect.TypeOf(&testStruct{})
	rtTestStructElem := rtTestStruct.Elem()
	ta := reflect.SliceOf(rtTElem)
	var tt []ITest
	ta2 := reflect.TypeOf(tt)
	fmt.Printf("rtT:%v,rtTElem:%v\n", rtT, rtTElem)
	fmt.Printf("rtTestStruct:%v,rtTestStructElem:%v\n", rtTestStruct, rtTestStructElem)
	fmt.Printf("ta:%v,ta2:%v\n", ta, ta2)

	var ttElem = ta2.Elem()
	ttElemKind := ttElem.Kind()
	fmt.Printf("ttElem:%v,ttElemKind:%v\n", ttElem, ttElemKind)
	assert.Equal(t, rtTElem, ttElem)
	assert.Equal(t, ta2, ta)

	fmt.Printf("rtT:%v,elem:%v\n", rtT, rtTElem)
	fmt.Printf("rtTestStruct:%v,elem:%v\n", rtTestStruct, rtTestStructElem)

	dst := &t2{}
	rtDstPtr := reflect.TypeOf(dst)
	rtDstPtrElem := reflect.TypeOf(dst).Elem()

	log.Info().
		Str("kind.rtDstPtr", fmt.Sprintf("%v", rtDstPtr)).
		Interface("kind.rtDstPtrElem", fmt.Sprintf("%v", rtDstPtrElem)).
		Send()
	// get field tag
	rt := reflect.TypeOf(dst).Elem()
	fmt.Printf("rt:%v,elem:%v\n", rt, reflect.TypeOf(t2{}))

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		tag := field.Tag
		rtField := field.Type
		switch kind := rtField.Kind(); kind {
		case reflect.Interface:
			fmt.Println("reflect.Interface")

		case reflect.Struct:
			fmt.Println("reflect.Struct")
		case reflect.Array, reflect.Slice:
			fmt.Printf("reflect.Struct:%v\n", rtField)

		case reflect.Ptr:
			fmt.Println("reflect.Ptr")

		}

		fmt.Printf("rtField:%v\n", rtField)

		//	fmt.Printf("rtField:%v,elem: %v\n", rtField, rtField.Elem())
		// ... and start using structtag by parsing the tag
		tags, err := structtag.Parse(string(tag))
		if err != nil {
			panic(err)
		}

		fmt.Println(tags)
		for _, tTags := range tags.Tags() {
			fmt.Printf("tag: %v:%v\n", tTags.Key, tTags.Value())
			fmt.Println(tTags.Key)     // Output: json
			fmt.Println(tTags.Name)    // Output: foo
			fmt.Println(tTags.Options) // Output: [omitempty string]
			if tTags.Key == "inject" {
				if tTags.Name == "ITest2" {
					//			assert.Equal(t, rtT, rtField)
				}
				if tTags.Name == "testStruct" {
					//		assert.Equal(t, rtT, rtTestStruct)
				}

			}
		}

	}

}

package __standard_lib

import (
	"errors"
	"reflect"
	"testing"
)

func Test_Reflect(t *testing.T) {
	type A struct {
		ID    int
		Name  string
		Hobby []string
	}

	type B struct {
		ID    int
		Name  string
		Hobby []string
	}

	a := A{ID: 1, Name: "nick", Hobby: []string{"game", "music"}}
	b := &B{}
	err := copyObject(b, a)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(b, a)

	listA := []A{
		{ID: 1, Name: "nick", Hobby: []string{"game", "music"}},
		{ID: 2, Name: "nick2", Hobby: []string{"game2", "music2"}},
		{ID: 3, Name: "nick3", Hobby: []string{"game3", "music3"}},
	}

	colum := sliceColum(listA, "Hobby")
	t.Log(colum)
}

func copyObject(dest any, source any) error {
	st := reflect.TypeOf(source)
	sv := reflect.ValueOf(source)

	if st.Kind() == reflect.Ptr {
		st = st.Elem()
		sv = sv.Elem()
	}

	dt := reflect.TypeOf(dest)
	dv := reflect.ValueOf(dest)
	if dt.Kind() != reflect.Ptr {
		return errors.New("dest must be a pointer")
	}
	dt = dt.Elem()
	dv = dv.Elem()

	if st.Kind() != reflect.Struct {
		return errors.New("source must be a struct")
	}

	if dt.Kind() != reflect.Struct {
		return errors.New("dest must be a struct")
	}

	destObj := reflect.New(dt)
	for i := 0; i < dt.NumField(); i++ {
		df := dt.Field(i)
		if _, ok := st.FieldByName(df.Name); ok {
			value := sv.FieldByName(df.Name)
			destObj.Elem().FieldByName(df.Name).Set(value)
		}
	}
	dv.Set(destObj.Elem())
	return nil
}

func sliceColum(list any, colum string) any {
	listType := reflect.TypeOf(list)
	listValue := reflect.ValueOf(list)
	if listType.Kind() == reflect.Ptr {
		listType = listType.Elem()
		listValue = listValue.Elem()
	}

	if listType.Kind() != reflect.Slice {
		return nil
	}
	listType = listType.Elem()
	if listType.Kind() == reflect.Ptr {
		listType = listType.Elem()
	}

	f, _ := listType.FieldByName(colum)
	sliceType := reflect.SliceOf(f.Type)
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)
	for i := 0; i < listValue.Len(); i++ {
		item := listValue.Index(i)
		if item.Kind() == reflect.Struct {
			val := item.FieldByName(colum)
			sliceValue = reflect.Append(sliceValue, val)
		}
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
			val := item.FieldByName(colum)
			sliceValue = reflect.Append(sliceValue, val)
		}
	}
	return sliceValue.Interface()
}

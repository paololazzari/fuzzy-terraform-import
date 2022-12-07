package util

import (
	"reflect"
	"sort"
)

// build a slice of structs from a slice of maps
func BuildSliceOfStructs(sliceOfMaps []map[string]interface{}) []any {
	var result []any
	for _, m := range sliceOfMaps {
		fields := make([]reflect.StructField, 0, len(m))

		for k, v := range m {
			f := reflect.StructField{
				Name: k,
				Type: reflect.TypeOf(v),
			}
			fields = append(fields, f)
		}

		st := reflect.StructOf(fields)
		sv := reflect.New(st)

		for k, v := range m {
			sv.Elem().FieldByName(k).Set(reflect.ValueOf(v))
		}

		result = append(result, sv.Elem().Interface())
	}
	return result
}

// create the strings required for fuzzyfinder menu
func FormatFuzzyInput(r interface{}) (string, []interface{}) {

	l := reflect.TypeOf(r).NumField()

	properties := make(map[string]string)
	keys := []string{}

	for idx := 0; idx < l; idx++ {
		k := reflect.TypeOf(r).Field(idx).Name
		v := reflect.ValueOf(r).FieldByName(k).String()
		properties[k] = v
		keys = append(keys, reflect.TypeOf(r).Field(idx).Name)
	}

	// Name or id properties go first and tags go last
	format := ""
	values := []interface{}{}
	name := properties["Name"]
	id := properties["Id"]

	if name != "" {
		format += "Name: %s\n"
		values = append(values, name)
		delete(properties, "Name")
	} else {
		format += "Id: %s\n"
		values = append(values, id)
		delete(properties, "Id")
	}

	tags := properties["Tags"]
	delete(properties, "Tags")

	// All other properties are sorted alphabetically
	sort.Strings(keys)
	for _, k := range keys {
		if k != "Name" && k != "Id" && k != "Tags" {
			format += k
			format += ":"
			format += " %s \n"
			values = append(values, properties[k])
		}
	}

	// Not all resources support tags
	if tags != "" {
		format += "Tags: %s"
		values = append(values, tags)
	}

	return format, values
}

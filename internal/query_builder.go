package minimax

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

type QueryBuilder interface {
	// StructToParams .
	StructToParams(v any) string
	// Build .
	Build(fullURL string, data any) string
}

type URLQueryBuilder struct{}

func NewURLQueryBuilder() QueryBuilder {
	return &URLQueryBuilder{}
}

func (u *URLQueryBuilder) Build(fullURL string, data any) string {
	if u.isZeroOfUnderlyingType(data) {
		return fullURL
	}

	return fmt.Sprintf("%s%s", fullURL, u.StructToParams(data))
}

func (u *URLQueryBuilder) StructToParams(data any) string {
	v := make(url.Values)
	st := reflect.TypeOf(data).Elem()
	sv := reflect.ValueOf(data).Elem()

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		value := sv.Field(i)

		if u.isZeroOfUnderlyingType(value.Interface()) {
			continue
		}

		var strValue string
		switch value.Kind() {
		case reflect.String:
			strValue = value.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			strValue = strconv.FormatInt(value.Int(), 10)
		case reflect.Bool:
			strValue = strconv.FormatBool(value.Bool())
		default:
			continue
		}

		v.Set(field.Tag.Get("json"), strValue)
	}

	if len(v) == 0 {
		return ""
	}

	return "&" + v.Encode()
}

func (u *URLQueryBuilder) isZeroOfUnderlyingType(x any) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

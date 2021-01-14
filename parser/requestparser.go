package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"time"

	valid "github.com/martinsd3v/gobrvalid"
)

//Parser reponsable for sing a struct
func Parser(request *http.Request, data interface{}) {
	if request != nil && request.Body != nil {
		request.ParseMultipartForm(32)

		//If request form is empty try parse json
		if len(request.Form) < 1 {
			decodeJSON(request.Body, data)
		} else {
			vl := reflect.ValueOf(data)
			bindValues(vl, FormSlice(request.Form))
		}
	}
}

func decodeJSON(r io.Reader, obj interface{}) (err error) {
	decoder := json.NewDecoder(r)
	decoder.UseNumber()
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&obj)
	return
}

func bindValues(vl reflect.Value, form FormSlice) {
	//must be a pointer
	if vl.Kind() == reflect.Ptr || vl.Kind() == reflect.Struct {

		var elm reflect.Value

		if vl.Kind() == reflect.Ptr {
			elm = vl.Elem()
		} else {
			elm = vl
		}

		for i := 0; i < elm.NumField(); i++ {
			fieldValue := elm.Field(i)
			field := elm.Type().Field(i).Name
			tagField := field

			//try to get custom tags
			if tag := elm.Type().Field(i).Tag.Get("json"); tag != "" {
				tagField = tag
			} else if tag := elm.Type().Field(i).Tag.Get("form"); tag != "" {
				tagField = tag
			}

			//verify field is public
			if valid.Matches(strings.Split(field, "")[0], "[A-Z]") {
				//If the field type is a struct then validate individual
				if fieldValue.Kind() == reflect.Struct {
					switch fieldValue.Interface().(type) {
					case time.Time:
						valueInsert := getValueInForm(tagField, form)
						if valueInsert[field] != nil {
							dateInsert := valueInsert[field][0]
							valid, dt := valid.IsDate(dateInsert)
							if valid {
								elm.Field(i).Set(reflect.ValueOf(dt))
							}
						}
					default:
						normalizedForm := normalizeFormStructs(tagField, form)
						bindValues(fieldValue, normalizedForm)
					}
				} else
				//If the type is a slice then loop and validate one by one
				if fieldValue.Kind() == reflect.Slice {
					sliceType := elm.Type().Field(i).Type

					if sliceType.Elem().Kind() == reflect.Struct {
						sl := prepareFormStructSlices(tagField, sliceType, form)
						elm.Field(i).Set(sl)
					} else {
						// Here is any slice
						valueInsert := normalizeFormSimpleSlices(tagField, form)
						item := elm.Field(i)
						trySetValue(item, valueInsert, field)
					}
				} else {
					if tagField != "" {
						field = tagField
					}
					valueInsert := getValueInForm(tagField, form)
					item := elm.Field(i)
					trySetValue(item, valueInsert, field)
				}
			}
		}
	}
}

//FormSlice typing of data received via form-data
type FormSlice map[string][]string

//responsible for normalizing data sent in the form
func normalizeFormStructs(key string, formIn FormSlice) FormSlice {
	keyLowerCase := strings.ToLower(key)
	regexFind := fmt.Sprintf("(%s\\[\\w+\\])", keyLowerCase)

	formOut := FormSlice{}

	for k := range formIn {
		keyFormLowerCase := strings.ToLower(k)
		if valid.Matches(keyFormLowerCase, regexFind) {
			//This is a slice
			if valid.Matches(k, "\\[\\]$") {
				cleanKey := valid.ReplacePattern(keyFormLowerCase, "\\[\\]$", "")
				cleanKey = valid.ReplacePattern(cleanKey, "^"+keyLowerCase, "")
				cleanKey = valid.ReplacePattern(cleanKey, "\\]", "")
				cleanKey = valid.ReplacePattern(cleanKey, "\\[", "&")
				cleanKey = valid.ReplacePattern(cleanKey, "^&", "")

				nkeys := strings.Split(cleanKey, "&")
				indexKey := ""

				if len(nkeys) > 1 {
					indexKey = nkeys[0] + "[" + strings.Join(nkeys[1:], "][") + "][]"
				} else {
					indexKey = nkeys[0] + "[]"
				}

				formOut[indexKey] = formIn[k]

			} else {
				//This is a struct
				cleanKey := valid.ReplacePattern(keyFormLowerCase, "\\[\\]$", "&")
				cleanKey = valid.ReplacePattern(cleanKey, "^"+keyLowerCase, "")
				cleanKey = valid.ReplacePattern(cleanKey, "\\]", "")
				cleanKey = valid.ReplacePattern(cleanKey, "\\[", "&")
				cleanKey = valid.ReplacePattern(cleanKey, "^&", "")

				nkeys := strings.Split(cleanKey, "&")

				indexKey := ""
				if len(nkeys) > 1 {
					indexKey = nkeys[0] + "[" + strings.Join(nkeys[1:], "][") + "]"
				} else {
					indexKey = nkeys[0]
					indexKey = valid.ReplacePattern(indexKey, "\\[\\]", "")
				}
				formOut[indexKey] = formIn[k]
			}
		}
	}

	return formOut
}

//responsible for normalizing data sent in the form
func normalizeFormSimpleSlices(key string, formIn FormSlice) FormSlice {
	keyLowerCase := strings.ToLower(key)
	regexFind1 := fmt.Sprintf("^%s\\[\\]", keyLowerCase)
	regexFind2 := fmt.Sprintf("^%s\\[\\w+\\]", keyLowerCase)

	formOut := FormSlice{}

	for k := range formIn {
		keyFormLowerCase := strings.ToLower(k)
		if valid.Matches(keyFormLowerCase, regexFind1) || valid.Matches(keyFormLowerCase, regexFind2) {
			formOut[key] = formIn[k]
		}
	}

	return formOut
}

func normalizeFormStructSlices(key string, formIn FormSlice) map[int]FormSlice {
	keyLowerCase := strings.ToLower(key)

	mpForm := map[int]FormSlice{}

	stIndex := 0
	mpIndex := map[string]int{}

	regexFind1 := fmt.Sprintf("^%s\\[\\]", keyLowerCase)
	regexFind2 := fmt.Sprintf("^%s\\[\\w+\\]", keyLowerCase)

	for k := range formIn {
		keyFormLowerCase := strings.ToLower(k)
		if valid.Matches(keyFormLowerCase, regexFind1) || valid.Matches(keyFormLowerCase, regexFind2) {

			cleanKey := valid.ReplacePattern(keyFormLowerCase, "\\]", "")
			cleanKey = valid.ReplacePattern(cleanKey, "\\[", "&")

			nkeys := strings.Split(cleanKey, "&")
			indexKey := ""
			if len(nkeys) > 3 {
				indexKey = nkeys[2] + "[" + strings.Join(nkeys[3:], "][") + "]"
			} else {
				indexKey = nkeys[2]
				indexKey = valid.ReplacePattern(indexKey, "\\[\\]", "")
			}

			index, exits := mpIndex[nkeys[1]]
			if exits {
				fm, _ := mpForm[index]
				fm[indexKey] = formIn[k]
			} else {

				fm := FormSlice{}
				fm[indexKey] = formIn[k]
				mpForm[stIndex] = fm

				mpIndex[nkeys[1]] = stIndex
				stIndex++
			}
		}
	}

	return mpForm
}

//responsible to prepare struct of slices
func prepareFormStructSlices(key string, slice reflect.Type, formIn FormSlice) reflect.Value {
	mpForm := normalizeFormStructSlices(key, formIn)
	mpElem := []reflect.Value{}

	typ := slice.Elem()

	for z := 0; z < len(mpForm); z++ {
		nw := reflect.New(typ)
		bindValues(nw, mpForm[z])
		mpElem = append(mpElem, nw)
	}

	sl := reflect.MakeSlice(slice, len(mpForm), len(mpForm))

	for q := 0; q < len(mpForm); q++ {
		sl.Index(q).Set(mpElem[q].Elem())
	}

	return sl
}

//responsible for taking the values passed in the form
func getValueInForm(key string, formIn FormSlice) FormSlice {
	formOut := FormSlice{}
	keyLowerCase := strings.ToLower(key)

	for k := range formIn {
		keyFormLowerCase := strings.ToLower(k)
		if keyFormLowerCase == keyLowerCase {
			if formIn[k][0] != "" {
				formOut[key] = formIn[k]
			}
		}
	}

	return formOut
}

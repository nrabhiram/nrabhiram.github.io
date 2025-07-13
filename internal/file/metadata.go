package file

import (
	"reflect"
	"strings"
	"time"
)

type Metadata struct {
	Date       time.Time         `json:"date,omitempty"`
	Slug       string            `json:"slug"`
	Categories []string          `json:"categories"`
	Thumbnail  string            `json:"thumbnail,omitempty"`
	Title      string            `json:"title,omitempty"`
	Summary    string            `json:"summary,omitempty"`
	Index      uint              `json:"index,omitempty"`
	Others     map[string]string `json:"others,omitempty"`
}

type MetadataField struct {
	timeValue   *time.Time
	stringValue *string
	sliceValue  *[]string
	uintValue   *uint
}

func (mf MetadataField) Time() (time.Time, bool) {
	if mf.timeValue != nil {
		return *mf.timeValue, true
	}
	return time.Time{}, false
}
func (mf MetadataField) String() (string, bool) {
	if mf.stringValue != nil {
		return *mf.stringValue, true
	}
	return "", false
}
func (mf MetadataField) StringSlice() ([]string, bool) {
	if mf.sliceValue != nil {
		return *mf.sliceValue, true
	}
	return nil, false
}
func (mf MetadataField) Uint() (uint, bool) {
	if mf.uintValue != nil {
		return *mf.uintValue, true
	}
	return 0, false
}

func MakeMetadata(
	date time.Time,
	slug string,
	categories []string,
	thumbnail string,
	title string,
	summary string,
	index uint,
	others map[string]string,
) *Metadata {
	return &Metadata{
		Date:       date,
		Slug:       slug,
		Categories: categories,
		Thumbnail:  thumbnail,
		Title:      title,
		Summary:    summary,
		Index:      index,
		Others:     others,
	}
}

func (m *Metadata) GetMetadataField(field string) MetadataField {
	if len(field) == 0 {
		return MetadataField{}
	}
	capitalizedField := strings.ToUpper(string(field[0])) + field[1:]
	v := reflect.ValueOf(m).Elem()
	fieldVal := v.FieldByName(capitalizedField)
	if fieldVal.IsValid() {
		switch fieldVal.Kind() {
		case reflect.Struct:
			if fieldVal.Type() == reflect.TypeOf(time.Time{}) {
				t := fieldVal.Interface().(time.Time)
				return MetadataField{timeValue: &t}
			}
		case reflect.String:
			s := fieldVal.String()
			return MetadataField{stringValue: &s}
		case reflect.Slice:
			if fieldVal.Type() == reflect.TypeOf([]string{}) {
				s := fieldVal.Interface().([]string)
				return MetadataField{sliceValue: &s}
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			u := uint(fieldVal.Uint())
			return MetadataField{uintValue: &u}
		}
	}
	if value, ok := m.Others[field]; ok {
		return MetadataField{stringValue: &value}
	}
	return MetadataField{}
}

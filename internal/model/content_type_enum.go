// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

const (
	ContentTypeMovie    ContentType = "movie"
	ContentTypeTvShow   ContentType = "tv_show"
	ContentTypeMusic    ContentType = "music"
	ContentTypeGame     ContentType = "game"
	ContentTypeSoftware ContentType = "software"
	ContentTypeBook     ContentType = "book"
	ContentTypeXxx      ContentType = "xxx"
)

var ErrInvalidContentType = fmt.Errorf("not a valid ContentType, try [%s]", strings.Join(_ContentTypeNames, ", "))

var _ContentTypeNames = []string{
	string(ContentTypeMovie),
	string(ContentTypeTvShow),
	string(ContentTypeMusic),
	string(ContentTypeGame),
	string(ContentTypeSoftware),
	string(ContentTypeBook),
	string(ContentTypeXxx),
}

// ContentTypeNames returns a list of possible string values of ContentType.
func ContentTypeNames() []string {
	tmp := make([]string, len(_ContentTypeNames))
	copy(tmp, _ContentTypeNames)
	return tmp
}

// ContentTypeValues returns a list of the values for ContentType
func ContentTypeValues() []ContentType {
	return []ContentType{
		ContentTypeMovie,
		ContentTypeTvShow,
		ContentTypeMusic,
		ContentTypeGame,
		ContentTypeSoftware,
		ContentTypeBook,
		ContentTypeXxx,
	}
}

// String implements the Stringer interface.
func (x ContentType) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x ContentType) IsValid() bool {
	_, err := ParseContentType(string(x))
	return err == nil
}

var _ContentTypeValue = map[string]ContentType{
	"movie":    ContentTypeMovie,
	"tv_show":  ContentTypeTvShow,
	"music":    ContentTypeMusic,
	"game":     ContentTypeGame,
	"software": ContentTypeSoftware,
	"book":     ContentTypeBook,
	"xxx":      ContentTypeXxx,
}

// ParseContentType attempts to convert a string to a ContentType.
func ParseContentType(name string) (ContentType, error) {
	if x, ok := _ContentTypeValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _ContentTypeValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return ContentType(""), fmt.Errorf("%s is %w", name, ErrInvalidContentType)
}

// MarshalText implements the text marshaller method.
func (x ContentType) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ContentType) UnmarshalText(text []byte) error {
	tmp, err := ParseContentType(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var errContentTypeNilPtr = errors.New("value pointer is nil") // one per type for package clashes

// Scan implements the Scanner interface.
func (x *ContentType) Scan(value interface{}) (err error) {
	if value == nil {
		*x = ContentType("")
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case string:
		*x, err = ParseContentType(v)
	case []byte:
		*x, err = ParseContentType(string(v))
	case ContentType:
		*x = v
	case *ContentType:
		if v == nil {
			return errContentTypeNilPtr
		}
		*x = *v
	case *string:
		if v == nil {
			return errContentTypeNilPtr
		}
		*x, err = ParseContentType(*v)
	default:
		return errors.New("invalid type for ContentType")
	}

	return
}

// Value implements the driver Valuer interface.
func (x ContentType) Value() (driver.Value, error) {
	return x.String(), nil
}

type NullContentType struct {
	ContentType ContentType
	Valid       bool
	Set         bool
}

func NewNullContentType(val interface{}) (x NullContentType) {
	err := x.Scan(val) // yes, we ignore this error, it will just be an invalid value.
	_ = err            // make any errcheck linters happy
	return
}

// Scan implements the Scanner interface.
func (x *NullContentType) Scan(value interface{}) (err error) {
	if value == nil {
		x.ContentType, x.Valid = ContentType(""), false
		return
	}

	err = x.ContentType.Scan(value)
	x.Valid = (err == nil)
	return
}

// Value implements the driver Valuer interface.
func (x NullContentType) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	return x.ContentType.String(), nil
}

// MarshalJSON correctly serializes a NullContentType to JSON.
func (n NullContentType) MarshalJSON() ([]byte, error) {
	const nullStr = "null"
	if n.Valid {
		return json.Marshal(n.ContentType)
	}
	return []byte(nullStr), nil
}

// UnmarshalJSON correctly deserializes a NullContentType from JSON.
func (n *NullContentType) UnmarshalJSON(b []byte) error {
	n.Set = true
	var x interface{}
	err := json.Unmarshal(b, &x)
	if err != nil {
		return err
	}
	err = n.Scan(x)
	return err
}

// MarshalGQL correctly serializes a NullContentType to GraphQL.
func (n NullContentType) MarshalGQL(w io.Writer) {
	bytes, err := json.Marshal(n)
	if err == nil {
		_, _ = w.Write(bytes)
	}
}

// UnmarshalGQL correctly deserializes a NullContentType from GraphQL.
func (n *NullContentType) UnmarshalGQL(v any) error {
	if v == nil {
		return nil
	}
	str, ok := v.(string)
	if !ok {
		return errors.New("value is not a string")
	}
	return n.UnmarshalJSON([]byte(str))
}

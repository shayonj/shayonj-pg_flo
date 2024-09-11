package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/jackc/pgtype"
)

// ParseTimestamp attempts to parse a timestamp string using multiple layouts
func ParseTimestamp(value string) (time.Time, error) {
	layouts := []string{
		time.RFC3339Nano,
		"2006-01-02 15:04:05.999999-07",
		"2006-01-02 15:04:05.999999Z07:00",
		"2006-01-02 15:04:05.999999",
		"2006-01-02T15:04:05.999999Z",
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05Z",
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, value); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse timestamp: %s", value)
}

// OidToTypeName maps PostgreSQL OIDs to their corresponding type names
var OidToTypeName = map[uint32]string{
	pgtype.BoolOID:             "bool",
	pgtype.ByteaOID:            "bytea",
	pgtype.Int8OID:             "int8",
	pgtype.Int2OID:             "int2",
	pgtype.Int4OID:             "int4",
	pgtype.TextOID:             "text",
	pgtype.JSONOID:             "json",
	pgtype.Float4OID:           "float4",
	pgtype.Float8OID:           "float8",
	pgtype.BoolArrayOID:        "bool[]",
	pgtype.Int2ArrayOID:        "int2[]",
	pgtype.Int4ArrayOID:        "int4[]",
	pgtype.TextArrayOID:        "text[]",
	pgtype.ByteaArrayOID:       "bytea[]",
	pgtype.Int8ArrayOID:        "int8[]",
	pgtype.Float4ArrayOID:      "float4[]",
	pgtype.Float8ArrayOID:      "float8[]",
	pgtype.BPCharOID:           "bpchar",
	pgtype.VarcharOID:          "varchar",
	pgtype.DateOID:             "date",
	pgtype.TimeOID:             "time",
	pgtype.TimestampOID:        "timestamp",
	pgtype.TimestampArrayOID:   "timestamp[]",
	pgtype.DateArrayOID:        "date[]",
	pgtype.TimestamptzOID:      "timestamptz",
	pgtype.TimestamptzArrayOID: "timestamptz[]",
	pgtype.IntervalOID:         "interval",
	pgtype.NumericArrayOID:     "numeric[]",
	pgtype.BitOID:              "bit",
	pgtype.VarbitOID:           "varbit",
	pgtype.NumericOID:          "numeric",
	pgtype.UUIDOID:             "uuid",
	pgtype.UUIDArrayOID:        "uuid[]",
	pgtype.JSONBOID:            "jsonb",
	pgtype.JSONBArrayOID:       "jsonb[]",
}

// OIDToString converts a PostgreSQL OID to its string representation
func OIDToString(oid uint32) string {
	if typeName, ok := OidToTypeName[oid]; ok {
		return typeName
	}
	return fmt.Sprintf("unknown_%d", oid)
}

// StringToOID converts a type name to its PostgreSQL OID
func StringToOID(typeName string) uint32 {
	for oid, name := range OidToTypeName {
		if name == typeName {
			return oid
		}
	}
	return 0
}

// ToInt64 converts an interface{} to int64
func ToInt64(v interface{}) (int64, bool) {
	switch v := v.(type) {
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(v).Int(), true
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(v).Uint()), true
	case string:
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return i, true
		}
	}
	return 0, false
}

// ToFloat64 converts an interface{} to float64
func ToFloat64(v interface{}) (float64, bool) {
	switch v := v.(type) {
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(v).Int()), true
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(v).Uint()), true
	case float32, float64:
		return reflect.ValueOf(v).Float(), true
	case string:
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f, true
		}
	}
	return 0, false
}

// ToBool converts various types to bool
func ToBool(v interface{}) (bool, bool) {
	switch v := v.(type) {
	case bool:
		return v, true
	case string:
		if v == "true" || v == "1" {
			return true, true
		}
		if v == "false" || v == "0" {
			return false, true
		}
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(v).Int() != 0, true
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(v).Uint() != 0, true
	case float32, float64:
		return reflect.ValueOf(v).Float() != 0, true
	}
	return false, false
}

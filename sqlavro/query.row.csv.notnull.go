package sqlavro

import (
	"strconv"

	"github.com/khezen/avro"
)

func sql2CSVFieldNotNull(schema avro.Schema, sqlField interface{}) (string, error) {
	switch schema.TypeName() {
	case avro.TypeInt64:
		return strconv.FormatInt(*sqlField.(*int64), 64), nil
	case avro.TypeInt32:
		return strconv.FormatInt(int64(*sqlField.(*int32)), 32), nil
	case avro.Type(avro.LogicalTypeDate):
		return sql2CSVDate(sqlField)
	case avro.Type(avro.LogicalTypeTime):
		return sql2CSVTime(sqlField)
	case avro.Type(avro.LogicalTypeTimestamp):
		return sql2CSVTimestamp(schema, sqlField)
	case avro.TypeFloat64:
		return strconv.FormatFloat(*sqlField.(*float64), 'f', -1, 64), nil
	case avro.TypeFloat32:
		return strconv.FormatFloat(float64(*sqlField.(*float32)), 'f', -1, 64), nil
	case avro.TypeString:
		return *sqlField.(*string), nil
	case avro.TypeBytes, avro.TypeFixed:
		return string(*sqlField.(*[]byte)), nil
	case avro.Type(avro.LogicalTypeDecimal):
		return sql2CSVDecimal(sqlField)
	}
	return "", ErrUnsupportedTypeForSQL
}

func sql2CSVTimestamp(schema avro.Schema, sqlField interface{}) (string, error) {
	switch schema.(*avro.DerivedPrimitiveSchema).Documentation {
	case string(DateTime):
		return *sqlField.(*string), nil
	case "", string(Timestamp):
		return strconv.FormatInt(int64(*sqlField.(*int32)), 32), nil
	default:
		return "", ErrUnsupportedTypeForSQL
	}
}

func sql2CSVTime(sqlField interface{}) (string, error) {
	return *sqlField.(*string), nil
}

func sql2CSVDate(sqlField interface{}) (string, error) {
	return *sqlField.(*string), nil
}

func sql2CSVDecimal(sqlField interface{}) (string, error) {
	field := *sqlField.(*[]byte)
	return string(field), nil
}

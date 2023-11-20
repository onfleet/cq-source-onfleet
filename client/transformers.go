package client

import (
	"github.com/apache/arrow/go/v14/arrow"
	"reflect"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"golang.org/x/exp/slices"
)

func TypeTransformer(field reflect.StructField) (arrow.DataType, error) {
	if isTimestampField(field) {
		return arrow.FixedWidthTypes.Timestamp_us, nil
	}

	return transformers.DefaultTypeTransformer(field)
}

func isTimestampField(field reflect.StructField) bool {
	timestampFieldNames := []string{
		"TimeCreated",
		"TimeLastModified",
		"TimeLastSeen",
		"CompletedAfter",
		"CompleteAfter",
		"CompleteBefore",
		"EstimatedCompletionTime",
		"EstimatedArrivalTime",
	}

	fieldType := field.Type
	if fieldType.Kind() == reflect.Ptr {
		fieldType = fieldType.Elem()
	}

	return fieldType.Kind() == reflect.Int64 &&
		slices.Contains(timestampFieldNames, field.Name)
}

func ResolverTransformer(field reflect.StructField, path string) schema.ColumnResolver {
	if isTimestampField(field) {
		return ResolveTimestampField(path)
	}

	return nil
}

func SharedTransformers() []transformers.StructTransformerOption {
	return []transformers.StructTransformerOption{
		transformers.WithPrimaryKeys("ID"),
		transformers.WithResolverTransformer(ResolverTransformer),
		transformers.WithTypeTransformer(TypeTransformer),
	}
}

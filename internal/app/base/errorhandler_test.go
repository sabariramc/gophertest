package base_test

import (
	"context"
	"fmt"
	"gopertest/internal/app/base"
	"gopertest/internal/errors"
	"testing"

	"gotest.tools/assert"
)

var tc = []struct {
	name     string
	err      *errors.HTTPError
	expected string
}{
	{
		name:     "BadRequest",
		err:      errors.ErrBadRequest,
		expected: `{"Code":"BAD_REQUEST","Message":"Invalid input"}`,
	}, {
		name:     "NotFound",
		err:      errors.ErrNotFound,
		expected: `{"Code":"NOT_FOUND","Message":"URL Not Found"}`,
	}, {
		name:     "MethodNotAllowed",
		err:      errors.ErrMethodNotAllowed,
		expected: `{"Code":"METHOD_NOT_ALLOWED","Message":"Method Not Allowed"}`,
	}, {
		name:     "InternalServerError",
		err:      errors.ErrInternalServerError,
		expected: `{"Code":"INTERNAL_SERVER_ERROR","Message":"Internal Server Error","Description":"Retry after some time, if persist contact technical team"}`,
	}, {
		name:     "CustomError",
		err:      &errors.HTTPError{StatusCode: 400, CustomError: &errors.CustomError{Code: "CUSTOM_ERROR", Message: "Custom Error", Description: "xyz"}},
		expected: `{"Code":"CUSTOM_ERROR","Message":"Custom Error","Description":"xyz"}`,
	}, {
		name:     "CustomErrorWithMap",
		err:      &errors.HTTPError{StatusCode: 400, CustomError: &errors.CustomError{Code: "CUSTOM_ERROR", Message: "Custom Error", Description: map[string]string{"key": "value"}}},
		expected: `{"Code":"CUSTOM_ERROR","Message":"Custom Error","Description":{"key":"value"}}`,
	}, {
		name:     "CustomErrorWithArray",
		err:      &errors.HTTPError{StatusCode: 400, CustomError: &errors.CustomError{Code: "CUSTOM_ERROR", Message: "Custom Error", Description: []any{"a", 1}}},
		expected: `{"Code":"CUSTOM_ERROR","Message":"Custom Error","Description":["a",1]}`,
	}, {
		name:     "CustomErrorWithStruct",
		err:      &errors.HTTPError{StatusCode: 400, CustomError: &errors.CustomError{Code: "CUSTOM_ERROR", Message: "Custom Error", Description: struct{ A, B string }{"a", "b"}}},
		expected: `{"Code":"CUSTOM_ERROR","Message":"Custom Error","Description":{"A":"a","B":"b"}}`,
	}, {
		name:     "CustomErrorWithNestedStruct",
		err:      &errors.HTTPError{StatusCode: 400, CustomError: &errors.CustomError{Code: "CUSTOM_ERROR", Message: "Custom Error", Description: struct{ A, B struct{ C, D string } }{struct{ C, D string }{"c", "d"}, struct{ C, D string }{"e", "f"}}}},
		expected: `{"Code":"CUSTOM_ERROR","Message":"Custom Error","Description":{"A":{"C":"c","D":"d"},"B":{"C":"e","D":"f"}}}`,
	}, {
		name:     "CustomErrorWithNestedMap",
		err:      &errors.HTTPError{StatusCode: 400, CustomError: &errors.CustomError{Code: "CUSTOM_ERROR", Message: "Custom Error", Description: map[string]map[string]string{"a": {"b": "c"}}}},
		expected: `{"Code":"CUSTOM_ERROR","Message":"Custom Error","Description":{"a":{"b":"c"}}}`,
	}, {
		name:     "CustomErrorWithNestedArray",
		err:      &errors.HTTPError{StatusCode: 400, CustomError: &errors.CustomError{Code: "CUSTOM_ERROR", Message: "Custom Error", Description: []map[string]string{{"a": "b"}}}},
		expected: `{"Code":"CUSTOM_ERROR","Message":"Custom Error","Description":[{"a":"b"}]}`,
	}, {
		name:     "CustomWithIntegers",
		err:      &errors.HTTPError{StatusCode: 400, CustomError: &errors.CustomError{Code: "CUSTOM_ERROR", Message: "Custom Error", Description: 123}},
		expected: `{"Code":"CUSTOM_ERROR","Message":"Custom Error","Description":123}`,
	}, {
		name:     "CustomWithFloats",
		err:      &errors.HTTPError{StatusCode: 400, CustomError: &errors.CustomError{Code: "CUSTOM_ERROR", Message: "Custom Error", Description: 123.45}},
		expected: `{"Code":"CUSTOM_ERROR","Message":"Custom Error","Description":123.45}`,
	}, {
		name:     "CustomWithBooleans",
		err:      &errors.HTTPError{StatusCode: 400, CustomError: &errors.CustomError{Code: "CUSTOM_ERROR", Message: "Custom Error", Description: true}},
		expected: `{"Code":"CUSTOM_ERROR","Message":"Custom Error","Description":true}`,
	}, {
		name:     "CustomWithNull",
		err:      &errors.HTTPError{StatusCode: 400, CustomError: &errors.CustomError{Code: "CUSTOM_ERROR", Message: "Custom Error", Description: nil}},
		expected: `{"Code":"CUSTOM_ERROR","Message":"Custom Error"}`,
	},
}

func TestErrorEncoding(t *testing.T) {
	ctx := context.Background()
	var err error
	for _, tt := range tc {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			statusCode, blob, err = base.ProcessError(ctx, tt.err)
			assert.NilError(t, err)
			assert.Equal(t, tt.expected, string(blob))
		})
	}
}

var getIter = func() func() error {
	idx := 0
	return func() error {
		if idx >= len(tc) {
			idx = 0
		}
		data := tc[idx]
		idx++
		return data.err
	}
}

func BenchmarkErrorEncoding(b *testing.B) {
	cnt = 0
	ctx := context.Background()
	b.ResetTimer()
	nextErr := getIter()
	for i := 0; i < b.N; i++ {
		statusCode, blob, _ = base.ProcessError(ctx, nextErr())
	}
}

func BenchmarkErrorEncodingParallel(b *testing.B) {
	cnt = 0
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < 8; i += 1 {
		b.Run(fmt.Sprintf("Iter-%d", i), func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				nextErr := getIter()
				for pb.Next() {
					statusCode, blob, _ = base.ProcessError(ctx, nextErr())
					cnt++
				}
			})
		})
	}
	fmt.Println(cnt)
}

var cnt int
var blob []byte
var statusCode int

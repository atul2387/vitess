/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sqltypes

import (
	"reflect"
	"testing"

	"vitess.io/vitess/go/test/utils"

	querypb "vitess.io/vitess/go/vt/proto/query"
)

func TestMakeRowTrusted(t *testing.T) {
	fields := MakeTestFields(
		"some_int|some_text|another_int",
		"int8|varchar|int8",
	)

	values := []byte{}
	hw := []byte("hello, world")
	values = append(values, hw...)
	values = append(values, byte(42))

	row := &querypb.Row{
		Lengths: []int64{-1, int64(len(hw)), 1},
		Values:  values,
	}

	want := []Value{
		MakeTrusted(querypb.Type_NULL_TYPE, nil),
		MakeTrusted(querypb.Type_VARCHAR, []byte("hello, world")),
		MakeTrusted(querypb.Type_INT8, []byte{byte(42)}),
	}

	result := MakeRowTrusted(fields, row)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("MakeRowTrusted:\ngot: %#v\nwant: %#v", result, want)
	}
}

func TestMakeRowTrustedDoesNotPanicOnNewColumns(t *testing.T) {
	fields := MakeTestFields(
		"some_int|some_text",
		"int8|varchar",
	)

	values := []byte{byte(123)}
	hw := []byte("hello, world")
	values = append(values, hw...)
	values = append(values, byte(42))

	row := &querypb.Row{
		Lengths: []int64{1, int64(len(hw)), 1},
		Values:  values,
	}

	want := []Value{
		MakeTrusted(querypb.Type_INT8, []byte{byte(123)}),
		MakeTrusted(querypb.Type_VARCHAR, []byte("hello, world")),
	}

	result := MakeRowTrusted(fields, row)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("MakeRowTrusted:\ngot: %#v\nwant: %#v", result, want)
	}
}

func TestRepair(t *testing.T) {
	fields := []*querypb.Field{{
		Type: Int64,
	}, {
		Type: VarChar,
	}}
	in := Result{
		Rows: [][]Value{
			{TestValue(VarBinary, "1"), TestValue(VarBinary, "aa")},
			{TestValue(VarBinary, "2"), TestValue(VarBinary, "bb")},
		},
	}
	want := Result{
		Rows: [][]Value{
			{TestValue(Int64, "1"), TestValue(VarChar, "aa")},
			{TestValue(Int64, "2"), TestValue(VarChar, "bb")},
		},
	}
	in.Repair(fields)
	if !reflect.DeepEqual(in, want) {
		t.Errorf("Repair:\n%#v, want\n%#v", in, want)
	}
}

func TestCopy(t *testing.T) {
	in := &Result{
		Fields: []*querypb.Field{{
			Type: Int64,
		}, {
			Type: VarChar,
		}},
		InsertID:     1,
		RowsAffected: 2,
		Rows: [][]Value{
			{TestValue(Int64, "1"), MakeTrusted(Null, nil)},
			{TestValue(Int64, "2"), MakeTrusted(VarChar, nil)},
			{TestValue(Int64, "3"), TestValue(VarChar, "")},
		},
	}
	out := in.Copy()
	utils.MustMatch(t, in, out)
}

func TestTruncate(t *testing.T) {
	in := &Result{
		Fields: []*querypb.Field{{
			Type: Int64,
		}, {
			Type: VarChar,
		}},
		InsertID:     1,
		RowsAffected: 2,
		Rows: [][]Value{
			{TestValue(Int64, "1"), MakeTrusted(Null, nil)},
			{TestValue(Int64, "2"), MakeTrusted(VarChar, nil)},
			{TestValue(Int64, "3"), TestValue(VarChar, "")},
		},
	}

	out := in.Truncate(0)
	if !reflect.DeepEqual(out, in) {
		t.Errorf("Truncate(0):\n%v, want\n%v", out, in)
	}

	out = in.Truncate(1)
	want := &Result{
		Fields: []*querypb.Field{{
			Type: Int64,
		}},
		InsertID:     1,
		RowsAffected: 2,
		Rows: [][]Value{
			{TestValue(Int64, "1")},
			{TestValue(Int64, "2")},
			{TestValue(Int64, "3")},
		},
	}
	if !reflect.DeepEqual(out, want) {
		t.Errorf("Truncate(1):\n%v, want\n%v", out, want)
	}
}

func TestStripMetaData(t *testing.T) {
	testcases := []struct {
		name           string
		in             *Result
		expected       *Result
		includedFields querypb.ExecuteOptions_IncludedFields
	}{{
		name:     "no fields",
		in:       &Result{},
		expected: &Result{},
	}, {
		name: "empty fields",
		in: &Result{
			Fields: []*querypb.Field{},
		},
		expected: &Result{
			Fields: []*querypb.Field{},
		},
	}, {
		name:           "no name",
		includedFields: querypb.ExecuteOptions_TYPE_ONLY,
		in: &Result{
			Fields: []*querypb.Field{{
				Type: Int64,
			}, {
				Type: VarChar,
			}},
		},
		expected: &Result{
			Fields: []*querypb.Field{{
				Type: Int64,
			}, {
				Type: VarChar,
			}},
		},
	}, {
		name:           "names",
		includedFields: querypb.ExecuteOptions_TYPE_ONLY,
		in: &Result{
			Fields: []*querypb.Field{{
				Name: "field1",
				Type: Int64,
			}, {
				Name: "field2",
				Type: VarChar,
			}},
		},
		expected: &Result{
			Fields: []*querypb.Field{{
				Type: Int64,
			}, {
				Type: VarChar,
			}},
		},
	}, {
		name:           "all fields - strip to type",
		includedFields: querypb.ExecuteOptions_TYPE_ONLY,
		in: &Result{
			Fields: []*querypb.Field{{
				Name:         "field1",
				Table:        "table1",
				OrgTable:     "orgtable1",
				OrgName:      "orgname1",
				ColumnLength: 5,
				Charset:      63,
				Decimals:     0,
				Flags:        2,
				Type:         Int64,
			}, {
				Name:         "field2",
				Table:        "table2",
				OrgTable:     "orgtable2",
				OrgName:      "orgname2",
				ColumnLength: 5,
				Charset:      63,
				Decimals:     0,
				Flags:        2,
				Type:         VarChar,
			}},
		},
		expected: &Result{
			Fields: []*querypb.Field{{
				Type: Int64,
			}, {
				Type: VarChar,
			}},
		},
	}, {
		name:           "all fields - not stripped",
		includedFields: querypb.ExecuteOptions_ALL,
		in: &Result{
			Fields: []*querypb.Field{{
				Name:         "field1",
				Table:        "table1",
				OrgTable:     "orgtable1",
				OrgName:      "orgname1",
				ColumnLength: 5,
				Charset:      63,
				Decimals:     0,
				Flags:        2,
				Type:         Int64,
			}, {
				Name:         "field2",
				Table:        "table2",
				OrgTable:     "orgtable2",
				OrgName:      "orgname2",
				ColumnLength: 5,
				Charset:      63,
				Decimals:     0,
				Flags:        2,
				Type:         VarChar,
			}},
		},
		expected: &Result{
			Fields: []*querypb.Field{{
				Name:         "field1",
				Table:        "table1",
				OrgTable:     "orgtable1",
				OrgName:      "orgname1",
				ColumnLength: 5,
				Charset:      63,
				Decimals:     0,
				Flags:        2,
				Type:         Int64,
			}, {
				Name:         "field2",
				Table:        "table2",
				OrgTable:     "orgtable2",
				OrgName:      "orgname2",
				ColumnLength: 5,
				Charset:      63,
				Decimals:     0,
				Flags:        2,
				Type:         VarChar,
			}},
		},
	}, {
		name: "all fields - strip to type and name",
		in: &Result{
			Fields: []*querypb.Field{{
				Name:         "field1",
				Table:        "table1",
				OrgTable:     "orgtable1",
				OrgName:      "orgname1",
				ColumnLength: 5,
				Charset:      63,
				Decimals:     0,
				Flags:        2,
				Type:         Int64,
			}, {
				Name:         "field2",
				Table:        "table2",
				OrgTable:     "orgtable2",
				OrgName:      "orgname2",
				ColumnLength: 5,
				Charset:      63,
				Decimals:     0,
				Flags:        2,
				Type:         VarChar,
			}},
		},
		expected: &Result{
			Fields: []*querypb.Field{{
				Name: "field1",
				Type: Int64,
			}, {
				Name: "field2",
				Type: VarChar,
			}},
		},
	}}
	for _, tcase := range testcases {
		inCopy := tcase.in.Copy()
		out := inCopy.StripMetadata(tcase.includedFields)
		if !reflect.DeepEqual(out, tcase.expected) {
			t.Errorf("StripMetaData unexpected result for %v: %v", tcase.name, out)
		}
		if len(tcase.in.Fields) > 0 {
			// check the out array is different than the in array.
			if out.Fields[0] == inCopy.Fields[0] && tcase.includedFields != querypb.ExecuteOptions_ALL {
				t.Errorf("StripMetaData modified original Field for %v", tcase.name)
			}
		}
		// check we didn't change the original result.
		utils.MustMatch(t, tcase.in, inCopy)
	}
}

func TestAppendResult(t *testing.T) {
	src := &Result{
		Fields: []*querypb.Field{{
			Type: Int64,
		}, {
			Type: VarChar,
		}},
		InsertID:     1,
		RowsAffected: 2,
		Rows: [][]Value{
			{TestValue(Int64, "2"), MakeTrusted(VarChar, nil)},
			{TestValue(Int64, "3"), TestValue(VarChar, "")},
		},
	}

	result := &Result{
		Fields: []*querypb.Field{{
			Type: Int64,
		}, {
			Type: VarChar,
		}},
		InsertID:     3,
		RowsAffected: 4,
		Rows: [][]Value{
			{TestValue(Int64, "1"), MakeTrusted(Null, nil)},
		},
	}

	want := &Result{
		Fields: []*querypb.Field{{
			Type: Int64,
		}, {
			Type: VarChar,
		}},
		InsertID:     1,
		RowsAffected: 6,
		Rows: [][]Value{
			{TestValue(Int64, "1"), MakeTrusted(Null, nil)},
			{TestValue(Int64, "2"), MakeTrusted(VarChar, nil)},
			{TestValue(Int64, "3"), TestValue(VarChar, "")},
		},
	}

	result.AppendResult(src)

	if !reflect.DeepEqual(result, want) {
		t.Errorf("Got:\n%#v, want:\n%#v", result, want)
	}
}

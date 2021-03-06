// Copyright 2015 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.
//
// Author: Peter Mattis (peter@cockroachlabs.com)

package sql_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/cockroachdb/cockroach/sql"
	"github.com/cockroachdb/cockroach/util/leaktest"
)

func TestAllocateIDs(t *testing.T) {
	defer leaktest.AfterTest(t)

	desc := sql.TableDescriptor{
		ID:   sql.MaxReservedDescID + 1,
		Name: "foo",
		Columns: []sql.ColumnDescriptor{
			{Name: "a"},
			{Name: "b"},
			{Name: "c"},
		},
		PrimaryIndex: sql.IndexDescriptor{Name: "c", ColumnNames: []string{"a", "b"}},
		Indexes: []sql.IndexDescriptor{
			{Name: "d", ColumnNames: []string{"b", "a"}},
			{Name: "e", ColumnNames: []string{"b"}},
		},
		Privileges: sql.NewDefaultPrivilegeDescriptor(),
	}
	if err := desc.AllocateIDs(); err != nil {
		t.Fatal(err)
	}

	expected := sql.TableDescriptor{
		ID:   sql.MaxReservedDescID + 1,
		Name: "foo",
		Columns: []sql.ColumnDescriptor{
			{ID: 1, Name: "a"},
			{ID: 2, Name: "b"},
			{ID: 3, Name: "c"},
		},
		PrimaryIndex: sql.IndexDescriptor{
			ID: 1, Name: "c", ColumnIDs: []sql.ColumnID{1, 2}, ColumnNames: []string{"a", "b"}},
		Indexes: []sql.IndexDescriptor{
			{ID: 2, Name: "d", ColumnIDs: []sql.ColumnID{2, 1}, ColumnNames: []string{"b", "a"}},
			{ID: 3, Name: "e", ColumnIDs: []sql.ColumnID{2}, ColumnNames: []string{"b"},
				ImplicitColumnIDs: []sql.ColumnID{1}},
		},
		Privileges:   sql.NewDefaultPrivilegeDescriptor(),
		NextColumnID: 4,
		NextIndexID:  4,
	}
	if !reflect.DeepEqual(expected, desc) {
		a, _ := json.MarshalIndent(expected, "", "  ")
		b, _ := json.MarshalIndent(desc, "", "  ")
		t.Fatalf("expected %s, but found %s", a, b)
	}
}

func TestValidateTableDesc(t *testing.T) {
	defer leaktest.AfterTest(t)

	testData := []struct {
		err  string
		desc sql.TableDescriptor
	}{
		{`empty table name`,
			sql.TableDescriptor{}},
		{`invalid table ID 0`,
			sql.TableDescriptor{ID: 0, Name: "foo"}},
		{`table must contain at least 1 column`,
			sql.TableDescriptor{ID: 1, Name: "foo"}},
		{`empty column name`,
			sql.TableDescriptor{
				ID:   1,
				Name: "foo",
				Columns: []sql.ColumnDescriptor{
					{ID: 0},
				},
				NextColumnID: 2,
			}},
		{`invalid column ID 0`,
			sql.TableDescriptor{
				ID:   1,
				Name: "foo",
				Columns: []sql.ColumnDescriptor{
					{ID: 0, Name: "bar"},
				},
				NextColumnID: 2,
			}},
		{`table must contain a primary key`,
			sql.TableDescriptor{
				ID:   1,
				Name: "foo",
				Columns: []sql.ColumnDescriptor{
					{ID: 1, Name: "bar"},
				},
				NextColumnID: 2,
			}},
		{`duplicate column name: "bar"`,
			sql.TableDescriptor{
				ID:   1,
				Name: "foo",
				Columns: []sql.ColumnDescriptor{
					{ID: 1, Name: "bar"},
					{ID: 1, Name: "bar"},
				},
				NextColumnID: 2,
			}},
		{`column "blah" duplicate ID of column "bar": 1`,
			sql.TableDescriptor{
				ID:   1,
				Name: "foo",
				Columns: []sql.ColumnDescriptor{
					{ID: 1, Name: "bar"},
					{ID: 1, Name: "blah"},
				},
				NextColumnID: 2,
			}},
		{`table must contain a primary key`,
			sql.TableDescriptor{
				ID:   1,
				Name: "foo",
				Columns: []sql.ColumnDescriptor{
					{ID: 1, Name: "bar"},
				},
				PrimaryIndex: sql.IndexDescriptor{ID: 0},
				NextColumnID: 2,
			}},
		{`invalid index ID 0`,
			sql.TableDescriptor{
				ID:   1,
				Name: "foo",
				Columns: []sql.ColumnDescriptor{
					{ID: 1, Name: "bar"},
				},
				PrimaryIndex: sql.IndexDescriptor{ID: 0, Name: "bar", ColumnIDs: []sql.ColumnID{0}},
				NextColumnID: 2,
			}},
		{`index "bar" must contain at least 1 column`,
			sql.TableDescriptor{
				ID:   1,
				Name: "foo",
				Columns: []sql.ColumnDescriptor{
					{ID: 1, Name: "bar"},
				},
				PrimaryIndex: sql.IndexDescriptor{ID: 1, Name: "primary", ColumnIDs: []sql.ColumnID{1}, ColumnNames: []string{"bar"}},
				Indexes: []sql.IndexDescriptor{
					{ID: 2, Name: "bar"},
				},

				NextColumnID: 2,
				NextIndexID:  3,
			}},
		{`mismatched column IDs (1) and names (0)`,
			sql.TableDescriptor{
				ID:   1,
				Name: "foo",
				Columns: []sql.ColumnDescriptor{
					{ID: 1, Name: "bar"},
				},
				PrimaryIndex: sql.IndexDescriptor{ID: 1, Name: "bar", ColumnIDs: []sql.ColumnID{1}},
				NextColumnID: 2,
				NextIndexID:  2,
			}},
		{`mismatched column IDs (1) and names (2)`,
			sql.TableDescriptor{
				ID:   1,
				Name: "foo",
				Columns: []sql.ColumnDescriptor{
					{ID: 1, Name: "bar"},
					{ID: 2, Name: "blah"},
				},
				PrimaryIndex: sql.IndexDescriptor{ID: 1, Name: "bar", ColumnIDs: []sql.ColumnID{1}, ColumnNames: []string{"bar", "blah"}},
				NextColumnID: 3,
				NextIndexID:  2,
			}},
		{`duplicate index name: "bar"`,
			sql.TableDescriptor{
				ID:   1,
				Name: "foo",
				Columns: []sql.ColumnDescriptor{
					{ID: 1, Name: "bar"},
				},
				PrimaryIndex: sql.IndexDescriptor{ID: 1, Name: "bar", ColumnIDs: []sql.ColumnID{1}, ColumnNames: []string{"bar"}},
				Indexes: []sql.IndexDescriptor{
					{ID: 1, Name: "bar", ColumnIDs: []sql.ColumnID{1}, ColumnNames: []string{"bar"}},
				},
				NextColumnID: 2,
				NextIndexID:  2,
			}},
		{`index "blah" duplicate ID of index "bar": 1`,
			sql.TableDescriptor{
				ID:   1,
				Name: "foo",
				Columns: []sql.ColumnDescriptor{
					{ID: 1, Name: "bar"},
				},
				PrimaryIndex: sql.IndexDescriptor{ID: 1, Name: "bar", ColumnIDs: []sql.ColumnID{1}, ColumnNames: []string{"bar"}},
				Indexes: []sql.IndexDescriptor{
					{ID: 1, Name: "blah", ColumnIDs: []sql.ColumnID{1}, ColumnNames: []string{"bar"}},
				},
				NextColumnID: 2,
				NextIndexID:  2,
			}},
		{`index "bar" column "bar" should have ID 1, but found ID 2`,
			sql.TableDescriptor{
				ID:   1,
				Name: "foo",
				Columns: []sql.ColumnDescriptor{
					{ID: 1, Name: "bar"},
				},
				PrimaryIndex: sql.IndexDescriptor{ID: 1, Name: "bar", ColumnIDs: []sql.ColumnID{2}, ColumnNames: []string{"bar"}},
				NextColumnID: 2,
				NextIndexID:  2,
			}},
		{`index "bar" contains unknown column "blah"`,
			sql.TableDescriptor{
				ID:   1,
				Name: "foo",
				Columns: []sql.ColumnDescriptor{
					{ID: 1, Name: "bar"},
				},
				PrimaryIndex: sql.IndexDescriptor{ID: 1, Name: "bar", ColumnIDs: []sql.ColumnID{1}, ColumnNames: []string{"blah"}},
				NextColumnID: 2,
				NextIndexID:  2,
			}},
	}
	for i, d := range testData {
		if err := d.desc.Validate(); err == nil {
			t.Errorf("%d: expected \"%s\", but found success: %+v", i, d.err, d.desc)
		} else if d.err != err.Error() {
			t.Errorf("%d: expected \"%s\", but found \"%s\"", i, d.err, err.Error())
		}
	}
}

func TestColumnTypeSQLString(t *testing.T) {
	defer leaktest.AfterTest(t)

	testData := []struct {
		colType     sql.ColumnType
		expectedSQL string
	}{
		{sql.ColumnType{Kind: sql.ColumnType_INT}, "INT"},
		{sql.ColumnType{Kind: sql.ColumnType_INT, Width: 2}, "INT(2)"},
		{sql.ColumnType{Kind: sql.ColumnType_FLOAT}, "FLOAT"},
		{sql.ColumnType{Kind: sql.ColumnType_FLOAT, Precision: 3}, "FLOAT(3)"},
		{sql.ColumnType{Kind: sql.ColumnType_DECIMAL}, "DECIMAL"},
		{sql.ColumnType{Kind: sql.ColumnType_DECIMAL, Precision: 6}, "DECIMAL(6)"},
		{sql.ColumnType{Kind: sql.ColumnType_DECIMAL, Precision: 7, Width: 8}, "DECIMAL(7,8)"},
		{sql.ColumnType{Kind: sql.ColumnType_DATE}, "DATE"},
		{sql.ColumnType{Kind: sql.ColumnType_TIMESTAMP}, "TIMESTAMP"},
		{sql.ColumnType{Kind: sql.ColumnType_INTERVAL}, "INTERVAL"},
		{sql.ColumnType{Kind: sql.ColumnType_STRING}, "STRING"},
		{sql.ColumnType{Kind: sql.ColumnType_STRING, Width: 10}, "STRING(10)"},
		{sql.ColumnType{Kind: sql.ColumnType_BYTES}, "BYTES"},
	}
	for i, d := range testData {
		sql := d.colType.SQLString()
		if d.expectedSQL != sql {
			t.Errorf("%d: expected %s, but got %s", i, d.expectedSQL, sql)
		}
	}
}

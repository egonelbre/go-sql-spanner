// Copyright 2021 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mock_server_tests

import (
	"time"

	"cloud.google.com/go/spanner"
)

type Singer struct {
	SingerId  int64 `gorm:"primaryKey:true;autoIncrement:false"`
	FirstName *string
	LastName  string
	BirthDate spanner.NullDate `gorm:"type:DATE"`
}

type Album struct {
	AlbumId  int64 `gorm:"primaryKey:true;autoIncrement:false"`
	SingerId *int64
	Singer   *Singer
	Title    string
}

type AllBasicTypes struct {
	Id int64 `gorm:"primaryKey:true;autoIncrement:false"`
	Int64 int64
	Float64 float64
	Numeric spanner.NullNumeric
	String string
	Bytes []byte
	Date spanner.NullDate
	Timestamp time.Time
	Json spanner.NullJSON
}

type AllSpannerNullableTypes struct {
	Id int64 `gorm:"primaryKey:true;autoIncrement:false"`
	Int64 spanner.NullInt64
	Float64 spanner.NullFloat64
	Numeric spanner.NullNumeric
	String spanner.NullString
	Bytes []byte
	Date spanner.NullDate
	Timestamp spanner.NullTime
	Json spanner.NullJSON
}
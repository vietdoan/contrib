// Copyright 2019-present Facebook
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

package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// MultiWordSchema holds the schema definition for the MultiWordSchema entity.
type MultiWordSchema struct {
	ent.Schema
}

// Fields of the MultiWordSchema.
func (MultiWordSchema) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("unit").
			Values("m", "ft").
			Default("m").
			Annotations(
				entproto.Field(2),
				entproto.Enum(map[string]int32{
					"m":  0,
					"ft": 1,
				}),
			),
	}
}

// Edges of the MultiWordSchema.
func (MultiWordSchema) Edges() []ent.Edge {
	return nil
}

func (MultiWordSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}

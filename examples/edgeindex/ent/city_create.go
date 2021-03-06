// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/city"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/street"
	"github.com/facebookincubator/ent/schema/field"
)

// CityCreate is the builder for creating a City entity.
type CityCreate struct {
	config
	name    *string
	streets map[int]struct{}
}

// SetName sets the name field.
func (cc *CityCreate) SetName(s string) *CityCreate {
	cc.name = &s
	return cc
}

// AddStreetIDs adds the streets edge to Street by ids.
func (cc *CityCreate) AddStreetIDs(ids ...int) *CityCreate {
	if cc.streets == nil {
		cc.streets = make(map[int]struct{})
	}
	for i := range ids {
		cc.streets[ids[i]] = struct{}{}
	}
	return cc
}

// AddStreets adds the streets edges to Street.
func (cc *CityCreate) AddStreets(s ...*Street) *CityCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cc.AddStreetIDs(ids...)
}

// Save creates the City in the database.
func (cc *CityCreate) Save(ctx context.Context) (*City, error) {
	if cc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	return cc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CityCreate) SaveX(ctx context.Context) *City {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CityCreate) sqlSave(ctx context.Context) (*City, error) {
	var (
		c     = &City{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: city.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: city.FieldID,
			},
		}
	)
	if value := cc.name; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: city.FieldName,
		})
		c.Name = *value
	}
	if nodes := cc.streets; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: street.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	c.ID = int(id)
	return c, nil
}

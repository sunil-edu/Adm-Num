// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"adm-num/ent/admnumber"
	"adm-num/ent/mststudent"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vmihailenco/msgpack/v5"
)

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

// MarshalGQL implements graphql.Marshaler interface.
func (o OrderDirection) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(o.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (o *OrderDirection) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("order direction %T must be a string", val)
	}
	*o = OrderDirection(str)
	return o.Validate()
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

func cursorsToPredicates(direction OrderDirection, after, before *Cursor, field, idField string) []func(s *sql.Selector) {
	var predicates []func(s *sql.Selector)
	if after != nil {
		if after.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeGT
			} else {
				predicate = sql.CompositeLT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					after.Value, after.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.GT
			} else {
				predicate = sql.LT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					after.ID,
				))
			})
		}
	}
	if before != nil {
		if before.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeLT
			} else {
				predicate = sql.CompositeGT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					before.Value, before.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.LT
			} else {
				predicate = sql.GT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					before.ID,
				))
			})
		}
	}
	return predicates
}

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID    int   `msgpack:"i"`
	Value Value `msgpack:"v,omitempty"`
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	quote := []byte{'"'}
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	_ = msgpack.NewEncoder(wc).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// AdmNumberEdge is the edge representation of AdmNumber.
type AdmNumberEdge struct {
	Node   *AdmNumber `json:"node"`
	Cursor Cursor     `json:"cursor"`
}

// AdmNumberConnection is the connection containing edges to AdmNumber.
type AdmNumberConnection struct {
	Edges      []*AdmNumberEdge `json:"edges"`
	PageInfo   PageInfo         `json:"pageInfo"`
	TotalCount int              `json:"totalCount"`
}

func (c *AdmNumberConnection) build(nodes []*AdmNumber, pager *admnumberPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *AdmNumber
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *AdmNumber {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *AdmNumber {
			return nodes[i]
		}
	}
	c.Edges = make([]*AdmNumberEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &AdmNumberEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// AdmNumberPaginateOption enables pagination customization.
type AdmNumberPaginateOption func(*admnumberPager) error

// WithAdmNumberOrder configures pagination ordering.
func WithAdmNumberOrder(order *AdmNumberOrder) AdmNumberPaginateOption {
	if order == nil {
		order = DefaultAdmNumberOrder
	}
	o := *order
	return func(pager *admnumberPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultAdmNumberOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithAdmNumberFilter configures pagination filter.
func WithAdmNumberFilter(filter func(*AdmNumberQuery) (*AdmNumberQuery, error)) AdmNumberPaginateOption {
	return func(pager *admnumberPager) error {
		if filter == nil {
			return errors.New("AdmNumberQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type admnumberPager struct {
	order  *AdmNumberOrder
	filter func(*AdmNumberQuery) (*AdmNumberQuery, error)
}

func newAdmNumberPager(opts []AdmNumberPaginateOption) (*admnumberPager, error) {
	pager := &admnumberPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultAdmNumberOrder
	}
	return pager, nil
}

func (p *admnumberPager) applyFilter(query *AdmNumberQuery) (*AdmNumberQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *admnumberPager) toCursor(an *AdmNumber) Cursor {
	return p.order.Field.toCursor(an)
}

func (p *admnumberPager) applyCursors(query *AdmNumberQuery, after, before *Cursor) *AdmNumberQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultAdmNumberOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *admnumberPager) applyOrder(query *AdmNumberQuery, reverse bool) *AdmNumberQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultAdmNumberOrder.Field {
		query = query.Order(direction.orderFunc(DefaultAdmNumberOrder.Field.field))
	}
	return query
}

func (p *admnumberPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultAdmNumberOrder.Field {
			b.Comma().Ident(DefaultAdmNumberOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to AdmNumber.
func (an *AdmNumberQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...AdmNumberPaginateOption,
) (*AdmNumberConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newAdmNumberPager(opts)
	if err != nil {
		return nil, err
	}
	if an, err = pager.applyFilter(an); err != nil {
		return nil, err
	}
	conn := &AdmNumberConnection{Edges: []*AdmNumberEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = an.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	an = pager.applyCursors(an, after, before)
	an = pager.applyOrder(an, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		an.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := an.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := an.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// AdmNumberOrderField defines the ordering field of AdmNumber.
type AdmNumberOrderField struct {
	field    string
	toCursor func(*AdmNumber) Cursor
}

// AdmNumberOrder defines the ordering of AdmNumber.
type AdmNumberOrder struct {
	Direction OrderDirection       `json:"direction"`
	Field     *AdmNumberOrderField `json:"field"`
}

// DefaultAdmNumberOrder is the default ordering of AdmNumber.
var DefaultAdmNumberOrder = &AdmNumberOrder{
	Direction: OrderDirectionAsc,
	Field: &AdmNumberOrderField{
		field: admnumber.FieldID,
		toCursor: func(an *AdmNumber) Cursor {
			return Cursor{ID: an.ID}
		},
	},
}

// ToEdge converts AdmNumber into AdmNumberEdge.
func (an *AdmNumber) ToEdge(order *AdmNumberOrder) *AdmNumberEdge {
	if order == nil {
		order = DefaultAdmNumberOrder
	}
	return &AdmNumberEdge{
		Node:   an,
		Cursor: order.Field.toCursor(an),
	}
}

// MstStudentEdge is the edge representation of MstStudent.
type MstStudentEdge struct {
	Node   *MstStudent `json:"node"`
	Cursor Cursor      `json:"cursor"`
}

// MstStudentConnection is the connection containing edges to MstStudent.
type MstStudentConnection struct {
	Edges      []*MstStudentEdge `json:"edges"`
	PageInfo   PageInfo          `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}

func (c *MstStudentConnection) build(nodes []*MstStudent, pager *mststudentPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *MstStudent
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *MstStudent {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *MstStudent {
			return nodes[i]
		}
	}
	c.Edges = make([]*MstStudentEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &MstStudentEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// MstStudentPaginateOption enables pagination customization.
type MstStudentPaginateOption func(*mststudentPager) error

// WithMstStudentOrder configures pagination ordering.
func WithMstStudentOrder(order *MstStudentOrder) MstStudentPaginateOption {
	if order == nil {
		order = DefaultMstStudentOrder
	}
	o := *order
	return func(pager *mststudentPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultMstStudentOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithMstStudentFilter configures pagination filter.
func WithMstStudentFilter(filter func(*MstStudentQuery) (*MstStudentQuery, error)) MstStudentPaginateOption {
	return func(pager *mststudentPager) error {
		if filter == nil {
			return errors.New("MstStudentQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type mststudentPager struct {
	order  *MstStudentOrder
	filter func(*MstStudentQuery) (*MstStudentQuery, error)
}

func newMstStudentPager(opts []MstStudentPaginateOption) (*mststudentPager, error) {
	pager := &mststudentPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultMstStudentOrder
	}
	return pager, nil
}

func (p *mststudentPager) applyFilter(query *MstStudentQuery) (*MstStudentQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *mststudentPager) toCursor(ms *MstStudent) Cursor {
	return p.order.Field.toCursor(ms)
}

func (p *mststudentPager) applyCursors(query *MstStudentQuery, after, before *Cursor) *MstStudentQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultMstStudentOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *mststudentPager) applyOrder(query *MstStudentQuery, reverse bool) *MstStudentQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultMstStudentOrder.Field {
		query = query.Order(direction.orderFunc(DefaultMstStudentOrder.Field.field))
	}
	return query
}

func (p *mststudentPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultMstStudentOrder.Field {
			b.Comma().Ident(DefaultMstStudentOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to MstStudent.
func (ms *MstStudentQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...MstStudentPaginateOption,
) (*MstStudentConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newMstStudentPager(opts)
	if err != nil {
		return nil, err
	}
	if ms, err = pager.applyFilter(ms); err != nil {
		return nil, err
	}
	conn := &MstStudentConnection{Edges: []*MstStudentEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = ms.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	ms = pager.applyCursors(ms, after, before)
	ms = pager.applyOrder(ms, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		ms.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := ms.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := ms.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// MstStudentOrderField defines the ordering field of MstStudent.
type MstStudentOrderField struct {
	field    string
	toCursor func(*MstStudent) Cursor
}

// MstStudentOrder defines the ordering of MstStudent.
type MstStudentOrder struct {
	Direction OrderDirection        `json:"direction"`
	Field     *MstStudentOrderField `json:"field"`
}

// DefaultMstStudentOrder is the default ordering of MstStudent.
var DefaultMstStudentOrder = &MstStudentOrder{
	Direction: OrderDirectionAsc,
	Field: &MstStudentOrderField{
		field: mststudent.FieldID,
		toCursor: func(ms *MstStudent) Cursor {
			return Cursor{ID: ms.ID}
		},
	},
}

// ToEdge converts MstStudent into MstStudentEdge.
func (ms *MstStudent) ToEdge(order *MstStudentOrder) *MstStudentEdge {
	if order == nil {
		order = DefaultMstStudentOrder
	}
	return &MstStudentEdge{
		Node:   ms,
		Cursor: order.Field.toCursor(ms),
	}
}

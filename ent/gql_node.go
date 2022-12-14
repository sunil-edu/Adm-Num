// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"adm-num/ent/admnumber"
	"adm-num/ent/mststudent"
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/sync/semaphore"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     int      `json:"id,omitempty"`     // node id.
	Type   string   `json:"type,omitempty"`   // node type.
	Fields []*Field `json:"fields,omitempty"` // node fields.
	Edges  []*Edge  `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string `json:"type,omitempty"` // edge type.
	Name string `json:"name,omitempty"` // edge name.
	IDs  []int  `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (an *AdmNumber) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     an.ID,
		Type:   "AdmNumber",
		Fields: make([]*Field, 8),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(an.AdmStartNo); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "adm_start_no",
		Value: string(buf),
	}
	if buf, err = json.Marshal(an.AdmCurrentNo); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "adm_current_no",
		Value: string(buf),
	}
	if buf, err = json.Marshal(an.IsPrefixed); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "bool",
		Name:  "is_prefixed",
		Value: string(buf),
	}
	if buf, err = json.Marshal(an.PrefixStr); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "prefix_str",
		Value: string(buf),
	}
	if buf, err = json.Marshal(an.SuffixStr); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "suffix_str",
		Value: string(buf),
	}
	if buf, err = json.Marshal(an.Separator); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "separator",
		Value: string(buf),
	}
	if buf, err = json.Marshal(an.PrefillWithZero); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "bool",
		Name:  "prefill_with_zero",
		Value: string(buf),
	}
	if buf, err = json.Marshal(an.PrefillWidth); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "int",
		Name:  "prefill_width",
		Value: string(buf),
	}
	return node, nil
}

func (ms *MstStudent) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ms.ID,
		Type:   "MstStudent",
		Fields: make([]*Field, 15),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(ms.FirstName); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "first_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.MiddleName); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "middle_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.LastName); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "last_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdStudying); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "bool",
		Name:  "std_studying",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdStatus); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "std_status",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdSex); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "std_sex",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdRegNo); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "std_reg_no",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdAdmNo); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "string",
		Name:  "std_adm_no",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdDoa); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "time.Time",
		Name:  "std_doa",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdFresher); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "bool",
		Name:  "std_fresher",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdDob); err != nil {
		return nil, err
	}
	node.Fields[10] = &Field{
		Type:  "time.Time",
		Name:  "std_dob",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdEmail); err != nil {
		return nil, err
	}
	node.Fields[11] = &Field{
		Type:  "string",
		Name:  "std_email",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdMobile); err != nil {
		return nil, err
	}
	node.Fields[12] = &Field{
		Type:  "string",
		Name:  "std_mobile",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdFatherName); err != nil {
		return nil, err
	}
	node.Fields[13] = &Field{
		Type:  "string",
		Name:  "std_father_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdMotherName); err != nil {
		return nil, err
	}
	node.Fields[14] = &Field{
		Type:  "string",
		Name:  "std_mother_name",
		Value: string(buf),
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id int) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, int) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, int) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, int) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id int) (string, error) {
			return c.tables.nodeType(ctx, c.driver, id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id int, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id int) (Noder, error) {
	switch table {
	case admnumber.Table:
		query := c.AdmNumber.Query().
			Where(admnumber.ID(id))
		query, err := query.CollectFields(ctx, "AdmNumber")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case mststudent.Table:
		query := c.MstStudent.Query().
			Where(mststudent.ID(id))
		query, err := query.CollectFields(ctx, "MstStudent")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []int, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]int)
	id2idx := make(map[int][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []int) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[int][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case admnumber.Table:
		query := c.AdmNumber.Query().
			Where(admnumber.IDIn(ids...))
		query, err := query.CollectFields(ctx, "AdmNumber")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case mststudent.Table:
		query := c.MstStudent.Query().
			Where(mststudent.IDIn(ids...))
		query, err := query.CollectFields(ctx, "MstStudent")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}

type tables struct {
	once  sync.Once
	sem   *semaphore.Weighted
	value atomic.Value
}

func (t *tables) nodeType(ctx context.Context, drv dialect.Driver, id int) (string, error) {
	tables, err := t.Load(ctx, drv)
	if err != nil {
		return "", err
	}
	idx := int(id / (1<<32 - 1))
	if idx < 0 || idx >= len(tables) {
		return "", fmt.Errorf("cannot resolve table from id %v: %w", id, errNodeInvalidID)
	}
	return tables[idx], nil
}

func (t *tables) Load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	t.once.Do(func() { t.sem = semaphore.NewWeighted(1) })
	if err := t.sem.Acquire(ctx, 1); err != nil {
		return nil, err
	}
	defer t.sem.Release(1)
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	tables, err := t.load(ctx, drv)
	if err == nil {
		t.value.Store(tables)
	}
	return tables, err
}

func (*tables) load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(drv.Dialect()).
		Select("type").
		From(sql.Table(schema.TypeTable)).
		OrderBy(sql.Asc("id")).
		Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []string
	return tables, sql.ScanSlice(rows, &tables)
}

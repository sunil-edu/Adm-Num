// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"adm-num/ent/migrate"

	"adm-num/ent/admnumber"
	"adm-num/ent/mststudent"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AdmNumber is the client for interacting with the AdmNumber builders.
	AdmNumber *AdmNumberClient
	// MstStudent is the client for interacting with the MstStudent builders.
	MstStudent *MstStudentClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.AdmNumber = NewAdmNumberClient(c.config)
	c.MstStudent = NewMstStudentClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		AdmNumber:  NewAdmNumberClient(cfg),
		MstStudent: NewMstStudentClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		AdmNumber:  NewAdmNumberClient(cfg),
		MstStudent: NewMstStudentClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AdmNumber.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.AdmNumber.Use(hooks...)
	c.MstStudent.Use(hooks...)
}

// AdmNumberClient is a client for the AdmNumber schema.
type AdmNumberClient struct {
	config
}

// NewAdmNumberClient returns a client for the AdmNumber from the given config.
func NewAdmNumberClient(c config) *AdmNumberClient {
	return &AdmNumberClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `admnumber.Hooks(f(g(h())))`.
func (c *AdmNumberClient) Use(hooks ...Hook) {
	c.hooks.AdmNumber = append(c.hooks.AdmNumber, hooks...)
}

// Create returns a builder for creating a AdmNumber entity.
func (c *AdmNumberClient) Create() *AdmNumberCreate {
	mutation := newAdmNumberMutation(c.config, OpCreate)
	return &AdmNumberCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AdmNumber entities.
func (c *AdmNumberClient) CreateBulk(builders ...*AdmNumberCreate) *AdmNumberCreateBulk {
	return &AdmNumberCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AdmNumber.
func (c *AdmNumberClient) Update() *AdmNumberUpdate {
	mutation := newAdmNumberMutation(c.config, OpUpdate)
	return &AdmNumberUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AdmNumberClient) UpdateOne(an *AdmNumber) *AdmNumberUpdateOne {
	mutation := newAdmNumberMutation(c.config, OpUpdateOne, withAdmNumber(an))
	return &AdmNumberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AdmNumberClient) UpdateOneID(id int) *AdmNumberUpdateOne {
	mutation := newAdmNumberMutation(c.config, OpUpdateOne, withAdmNumberID(id))
	return &AdmNumberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AdmNumber.
func (c *AdmNumberClient) Delete() *AdmNumberDelete {
	mutation := newAdmNumberMutation(c.config, OpDelete)
	return &AdmNumberDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AdmNumberClient) DeleteOne(an *AdmNumber) *AdmNumberDeleteOne {
	return c.DeleteOneID(an.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *AdmNumberClient) DeleteOneID(id int) *AdmNumberDeleteOne {
	builder := c.Delete().Where(admnumber.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AdmNumberDeleteOne{builder}
}

// Query returns a query builder for AdmNumber.
func (c *AdmNumberClient) Query() *AdmNumberQuery {
	return &AdmNumberQuery{
		config: c.config,
	}
}

// Get returns a AdmNumber entity by its id.
func (c *AdmNumberClient) Get(ctx context.Context, id int) (*AdmNumber, error) {
	return c.Query().Where(admnumber.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AdmNumberClient) GetX(ctx context.Context, id int) *AdmNumber {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AdmNumberClient) Hooks() []Hook {
	return c.hooks.AdmNumber
}

// MstStudentClient is a client for the MstStudent schema.
type MstStudentClient struct {
	config
}

// NewMstStudentClient returns a client for the MstStudent from the given config.
func NewMstStudentClient(c config) *MstStudentClient {
	return &MstStudentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `mststudent.Hooks(f(g(h())))`.
func (c *MstStudentClient) Use(hooks ...Hook) {
	c.hooks.MstStudent = append(c.hooks.MstStudent, hooks...)
}

// Create returns a builder for creating a MstStudent entity.
func (c *MstStudentClient) Create() *MstStudentCreate {
	mutation := newMstStudentMutation(c.config, OpCreate)
	return &MstStudentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MstStudent entities.
func (c *MstStudentClient) CreateBulk(builders ...*MstStudentCreate) *MstStudentCreateBulk {
	return &MstStudentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MstStudent.
func (c *MstStudentClient) Update() *MstStudentUpdate {
	mutation := newMstStudentMutation(c.config, OpUpdate)
	return &MstStudentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MstStudentClient) UpdateOne(ms *MstStudent) *MstStudentUpdateOne {
	mutation := newMstStudentMutation(c.config, OpUpdateOne, withMstStudent(ms))
	return &MstStudentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MstStudentClient) UpdateOneID(id int) *MstStudentUpdateOne {
	mutation := newMstStudentMutation(c.config, OpUpdateOne, withMstStudentID(id))
	return &MstStudentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MstStudent.
func (c *MstStudentClient) Delete() *MstStudentDelete {
	mutation := newMstStudentMutation(c.config, OpDelete)
	return &MstStudentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MstStudentClient) DeleteOne(ms *MstStudent) *MstStudentDeleteOne {
	return c.DeleteOneID(ms.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *MstStudentClient) DeleteOneID(id int) *MstStudentDeleteOne {
	builder := c.Delete().Where(mststudent.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MstStudentDeleteOne{builder}
}

// Query returns a query builder for MstStudent.
func (c *MstStudentClient) Query() *MstStudentQuery {
	return &MstStudentQuery{
		config: c.config,
	}
}

// Get returns a MstStudent entity by its id.
func (c *MstStudentClient) Get(ctx context.Context, id int) (*MstStudent, error) {
	return c.Query().Where(mststudent.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MstStudentClient) GetX(ctx context.Context, id int) *MstStudent {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MstStudentClient) Hooks() []Hook {
	hooks := c.hooks.MstStudent
	return append(hooks[:len(hooks):len(hooks)], mststudent.Hooks[:]...)
}

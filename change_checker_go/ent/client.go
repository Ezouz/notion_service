// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"gitlab.42paris.fr/notion_service/ent/migrate"

	"gitlab.42paris.fr/notion_service/ent/database"
	"gitlab.42paris.fr/notion_service/ent/status"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Database is the client for interacting with the Database builders.
	Database *DatabaseClient
	// Status is the client for interacting with the Status builders.
	Status *StatusClient
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
	c.Database = NewDatabaseClient(c.config)
	c.Status = NewStatusClient(c.config)
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
		ctx:      ctx,
		config:   cfg,
		Database: NewDatabaseClient(cfg),
		Status:   NewStatusClient(cfg),
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
		ctx:      ctx,
		config:   cfg,
		Database: NewDatabaseClient(cfg),
		Status:   NewStatusClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Database.
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
	c.Database.Use(hooks...)
	c.Status.Use(hooks...)
}

// DatabaseClient is a client for the Database schema.
type DatabaseClient struct {
	config
}

// NewDatabaseClient returns a client for the Database from the given config.
func NewDatabaseClient(c config) *DatabaseClient {
	return &DatabaseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `database.Hooks(f(g(h())))`.
func (c *DatabaseClient) Use(hooks ...Hook) {
	c.hooks.Database = append(c.hooks.Database, hooks...)
}

// Create returns a builder for creating a Database entity.
func (c *DatabaseClient) Create() *DatabaseCreate {
	mutation := newDatabaseMutation(c.config, OpCreate)
	return &DatabaseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Database entities.
func (c *DatabaseClient) CreateBulk(builders ...*DatabaseCreate) *DatabaseCreateBulk {
	return &DatabaseCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Database.
func (c *DatabaseClient) Update() *DatabaseUpdate {
	mutation := newDatabaseMutation(c.config, OpUpdate)
	return &DatabaseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DatabaseClient) UpdateOne(d *Database) *DatabaseUpdateOne {
	mutation := newDatabaseMutation(c.config, OpUpdateOne, withDatabase(d))
	return &DatabaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DatabaseClient) UpdateOneID(id int) *DatabaseUpdateOne {
	mutation := newDatabaseMutation(c.config, OpUpdateOne, withDatabaseID(id))
	return &DatabaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Database.
func (c *DatabaseClient) Delete() *DatabaseDelete {
	mutation := newDatabaseMutation(c.config, OpDelete)
	return &DatabaseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DatabaseClient) DeleteOne(d *Database) *DatabaseDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DatabaseClient) DeleteOneID(id int) *DatabaseDeleteOne {
	builder := c.Delete().Where(database.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DatabaseDeleteOne{builder}
}

// Query returns a query builder for Database.
func (c *DatabaseClient) Query() *DatabaseQuery {
	return &DatabaseQuery{
		config: c.config,
	}
}

// Get returns a Database entity by its id.
func (c *DatabaseClient) Get(ctx context.Context, id int) (*Database, error) {
	return c.Query().Where(database.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DatabaseClient) GetX(ctx context.Context, id int) *Database {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStatus queries the status edge of a Database.
func (c *DatabaseClient) QueryStatus(d *Database) *StatusQuery {
	query := &StatusQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(database.Table, database.FieldID, id),
			sqlgraph.To(status.Table, status.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, database.StatusTable, database.StatusColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DatabaseClient) Hooks() []Hook {
	return c.hooks.Database
}

// StatusClient is a client for the Status schema.
type StatusClient struct {
	config
}

// NewStatusClient returns a client for the Status from the given config.
func NewStatusClient(c config) *StatusClient {
	return &StatusClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `status.Hooks(f(g(h())))`.
func (c *StatusClient) Use(hooks ...Hook) {
	c.hooks.Status = append(c.hooks.Status, hooks...)
}

// Create returns a builder for creating a Status entity.
func (c *StatusClient) Create() *StatusCreate {
	mutation := newStatusMutation(c.config, OpCreate)
	return &StatusCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Status entities.
func (c *StatusClient) CreateBulk(builders ...*StatusCreate) *StatusCreateBulk {
	return &StatusCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Status.
func (c *StatusClient) Update() *StatusUpdate {
	mutation := newStatusMutation(c.config, OpUpdate)
	return &StatusUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StatusClient) UpdateOne(s *Status) *StatusUpdateOne {
	mutation := newStatusMutation(c.config, OpUpdateOne, withStatus(s))
	return &StatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StatusClient) UpdateOneID(id int) *StatusUpdateOne {
	mutation := newStatusMutation(c.config, OpUpdateOne, withStatusID(id))
	return &StatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Status.
func (c *StatusClient) Delete() *StatusDelete {
	mutation := newStatusMutation(c.config, OpDelete)
	return &StatusDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StatusClient) DeleteOne(s *Status) *StatusDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *StatusClient) DeleteOneID(id int) *StatusDeleteOne {
	builder := c.Delete().Where(status.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StatusDeleteOne{builder}
}

// Query returns a query builder for Status.
func (c *StatusClient) Query() *StatusQuery {
	return &StatusQuery{
		config: c.config,
	}
}

// Get returns a Status entity by its id.
func (c *StatusClient) Get(ctx context.Context, id int) (*Status, error) {
	return c.Query().Where(status.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StatusClient) GetX(ctx context.Context, id int) *Status {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *StatusClient) Hooks() []Hook {
	return c.hooks.Status
}

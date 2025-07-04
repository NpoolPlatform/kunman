// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated/capacity"
	"github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated/quota"
	"github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated/subscription"

	stdsql "database/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Capacity is the client for interacting with the Capacity builders.
	Capacity *CapacityClient
	// Quota is the client for interacting with the Quota builders.
	Quota *QuotaClient
	// Subscription is the client for interacting with the Subscription builders.
	Subscription *SubscriptionClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Capacity = NewCapacityClient(c.config)
	c.Quota = NewQuotaClient(c.config)
	c.Subscription = NewSubscriptionClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("generated: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("generated: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Capacity:     NewCapacityClient(cfg),
		Quota:        NewQuotaClient(cfg),
		Subscription: NewSubscriptionClient(cfg),
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
		ctx:          ctx,
		config:       cfg,
		Capacity:     NewCapacityClient(cfg),
		Quota:        NewQuotaClient(cfg),
		Subscription: NewSubscriptionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Capacity.
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
	c.Capacity.Use(hooks...)
	c.Quota.Use(hooks...)
	c.Subscription.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Capacity.Intercept(interceptors...)
	c.Quota.Intercept(interceptors...)
	c.Subscription.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CapacityMutation:
		return c.Capacity.mutate(ctx, m)
	case *QuotaMutation:
		return c.Quota.mutate(ctx, m)
	case *SubscriptionMutation:
		return c.Subscription.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("generated: unknown mutation type %T", m)
	}
}

// CapacityClient is a client for the Capacity schema.
type CapacityClient struct {
	config
}

// NewCapacityClient returns a client for the Capacity from the given config.
func NewCapacityClient(c config) *CapacityClient {
	return &CapacityClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `capacity.Hooks(f(g(h())))`.
func (c *CapacityClient) Use(hooks ...Hook) {
	c.hooks.Capacity = append(c.hooks.Capacity, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `capacity.Intercept(f(g(h())))`.
func (c *CapacityClient) Intercept(interceptors ...Interceptor) {
	c.inters.Capacity = append(c.inters.Capacity, interceptors...)
}

// Create returns a builder for creating a Capacity entity.
func (c *CapacityClient) Create() *CapacityCreate {
	mutation := newCapacityMutation(c.config, OpCreate)
	return &CapacityCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Capacity entities.
func (c *CapacityClient) CreateBulk(builders ...*CapacityCreate) *CapacityCreateBulk {
	return &CapacityCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CapacityClient) MapCreateBulk(slice any, setFunc func(*CapacityCreate, int)) *CapacityCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CapacityCreateBulk{err: fmt.Errorf("calling to CapacityClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CapacityCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CapacityCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Capacity.
func (c *CapacityClient) Update() *CapacityUpdate {
	mutation := newCapacityMutation(c.config, OpUpdate)
	return &CapacityUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CapacityClient) UpdateOne(ca *Capacity) *CapacityUpdateOne {
	mutation := newCapacityMutation(c.config, OpUpdateOne, withCapacity(ca))
	return &CapacityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CapacityClient) UpdateOneID(id uint32) *CapacityUpdateOne {
	mutation := newCapacityMutation(c.config, OpUpdateOne, withCapacityID(id))
	return &CapacityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Capacity.
func (c *CapacityClient) Delete() *CapacityDelete {
	mutation := newCapacityMutation(c.config, OpDelete)
	return &CapacityDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CapacityClient) DeleteOne(ca *Capacity) *CapacityDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CapacityClient) DeleteOneID(id uint32) *CapacityDeleteOne {
	builder := c.Delete().Where(capacity.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CapacityDeleteOne{builder}
}

// Query returns a query builder for Capacity.
func (c *CapacityClient) Query() *CapacityQuery {
	return &CapacityQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCapacity},
		inters: c.Interceptors(),
	}
}

// Get returns a Capacity entity by its id.
func (c *CapacityClient) Get(ctx context.Context, id uint32) (*Capacity, error) {
	return c.Query().Where(capacity.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CapacityClient) GetX(ctx context.Context, id uint32) *Capacity {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CapacityClient) Hooks() []Hook {
	return c.hooks.Capacity
}

// Interceptors returns the client interceptors.
func (c *CapacityClient) Interceptors() []Interceptor {
	return c.inters.Capacity
}

func (c *CapacityClient) mutate(ctx context.Context, m *CapacityMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CapacityCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CapacityUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CapacityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CapacityDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Capacity mutation op: %q", m.Op())
	}
}

// QuotaClient is a client for the Quota schema.
type QuotaClient struct {
	config
}

// NewQuotaClient returns a client for the Quota from the given config.
func NewQuotaClient(c config) *QuotaClient {
	return &QuotaClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `quota.Hooks(f(g(h())))`.
func (c *QuotaClient) Use(hooks ...Hook) {
	c.hooks.Quota = append(c.hooks.Quota, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `quota.Intercept(f(g(h())))`.
func (c *QuotaClient) Intercept(interceptors ...Interceptor) {
	c.inters.Quota = append(c.inters.Quota, interceptors...)
}

// Create returns a builder for creating a Quota entity.
func (c *QuotaClient) Create() *QuotaCreate {
	mutation := newQuotaMutation(c.config, OpCreate)
	return &QuotaCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Quota entities.
func (c *QuotaClient) CreateBulk(builders ...*QuotaCreate) *QuotaCreateBulk {
	return &QuotaCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *QuotaClient) MapCreateBulk(slice any, setFunc func(*QuotaCreate, int)) *QuotaCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &QuotaCreateBulk{err: fmt.Errorf("calling to QuotaClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*QuotaCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &QuotaCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Quota.
func (c *QuotaClient) Update() *QuotaUpdate {
	mutation := newQuotaMutation(c.config, OpUpdate)
	return &QuotaUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *QuotaClient) UpdateOne(q *Quota) *QuotaUpdateOne {
	mutation := newQuotaMutation(c.config, OpUpdateOne, withQuota(q))
	return &QuotaUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *QuotaClient) UpdateOneID(id uint32) *QuotaUpdateOne {
	mutation := newQuotaMutation(c.config, OpUpdateOne, withQuotaID(id))
	return &QuotaUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Quota.
func (c *QuotaClient) Delete() *QuotaDelete {
	mutation := newQuotaMutation(c.config, OpDelete)
	return &QuotaDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *QuotaClient) DeleteOne(q *Quota) *QuotaDeleteOne {
	return c.DeleteOneID(q.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *QuotaClient) DeleteOneID(id uint32) *QuotaDeleteOne {
	builder := c.Delete().Where(quota.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &QuotaDeleteOne{builder}
}

// Query returns a query builder for Quota.
func (c *QuotaClient) Query() *QuotaQuery {
	return &QuotaQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeQuota},
		inters: c.Interceptors(),
	}
}

// Get returns a Quota entity by its id.
func (c *QuotaClient) Get(ctx context.Context, id uint32) (*Quota, error) {
	return c.Query().Where(quota.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *QuotaClient) GetX(ctx context.Context, id uint32) *Quota {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *QuotaClient) Hooks() []Hook {
	return c.hooks.Quota
}

// Interceptors returns the client interceptors.
func (c *QuotaClient) Interceptors() []Interceptor {
	return c.inters.Quota
}

func (c *QuotaClient) mutate(ctx context.Context, m *QuotaMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&QuotaCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&QuotaUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&QuotaUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&QuotaDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Quota mutation op: %q", m.Op())
	}
}

// SubscriptionClient is a client for the Subscription schema.
type SubscriptionClient struct {
	config
}

// NewSubscriptionClient returns a client for the Subscription from the given config.
func NewSubscriptionClient(c config) *SubscriptionClient {
	return &SubscriptionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `subscription.Hooks(f(g(h())))`.
func (c *SubscriptionClient) Use(hooks ...Hook) {
	c.hooks.Subscription = append(c.hooks.Subscription, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `subscription.Intercept(f(g(h())))`.
func (c *SubscriptionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Subscription = append(c.inters.Subscription, interceptors...)
}

// Create returns a builder for creating a Subscription entity.
func (c *SubscriptionClient) Create() *SubscriptionCreate {
	mutation := newSubscriptionMutation(c.config, OpCreate)
	return &SubscriptionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Subscription entities.
func (c *SubscriptionClient) CreateBulk(builders ...*SubscriptionCreate) *SubscriptionCreateBulk {
	return &SubscriptionCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SubscriptionClient) MapCreateBulk(slice any, setFunc func(*SubscriptionCreate, int)) *SubscriptionCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SubscriptionCreateBulk{err: fmt.Errorf("calling to SubscriptionClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SubscriptionCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SubscriptionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Subscription.
func (c *SubscriptionClient) Update() *SubscriptionUpdate {
	mutation := newSubscriptionMutation(c.config, OpUpdate)
	return &SubscriptionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SubscriptionClient) UpdateOne(s *Subscription) *SubscriptionUpdateOne {
	mutation := newSubscriptionMutation(c.config, OpUpdateOne, withSubscription(s))
	return &SubscriptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SubscriptionClient) UpdateOneID(id uint32) *SubscriptionUpdateOne {
	mutation := newSubscriptionMutation(c.config, OpUpdateOne, withSubscriptionID(id))
	return &SubscriptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Subscription.
func (c *SubscriptionClient) Delete() *SubscriptionDelete {
	mutation := newSubscriptionMutation(c.config, OpDelete)
	return &SubscriptionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SubscriptionClient) DeleteOne(s *Subscription) *SubscriptionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SubscriptionClient) DeleteOneID(id uint32) *SubscriptionDeleteOne {
	builder := c.Delete().Where(subscription.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SubscriptionDeleteOne{builder}
}

// Query returns a query builder for Subscription.
func (c *SubscriptionClient) Query() *SubscriptionQuery {
	return &SubscriptionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSubscription},
		inters: c.Interceptors(),
	}
}

// Get returns a Subscription entity by its id.
func (c *SubscriptionClient) Get(ctx context.Context, id uint32) (*Subscription, error) {
	return c.Query().Where(subscription.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SubscriptionClient) GetX(ctx context.Context, id uint32) *Subscription {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SubscriptionClient) Hooks() []Hook {
	return c.hooks.Subscription
}

// Interceptors returns the client interceptors.
func (c *SubscriptionClient) Interceptors() []Interceptor {
	return c.inters.Subscription
}

func (c *SubscriptionClient) mutate(ctx context.Context, m *SubscriptionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SubscriptionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SubscriptionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SubscriptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SubscriptionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Subscription mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Capacity, Quota, Subscription []ent.Hook
	}
	inters struct {
		Capacity, Quota, Subscription []ent.Interceptor
	}
)

// ExecContext allows calling the underlying ExecContext method of the driver if it is supported by it.
// See, database/sql#DB.ExecContext for more information.
func (c *config) ExecContext(ctx context.Context, query string, args ...any) (stdsql.Result, error) {
	ex, ok := c.driver.(interface {
		ExecContext(context.Context, string, ...any) (stdsql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	return ex.ExecContext(ctx, query, args...)
}

// QueryContext allows calling the underlying QueryContext method of the driver if it is supported by it.
// See, database/sql#DB.QueryContext for more information.
func (c *config) QueryContext(ctx context.Context, query string, args ...any) (*stdsql.Rows, error) {
	q, ok := c.driver.(interface {
		QueryContext(context.Context, string, ...any) (*stdsql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	return q.QueryContext(ctx, query, args...)
}

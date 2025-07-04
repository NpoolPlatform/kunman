// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/appcountry"
	"github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/applang"
	"github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/country"
	"github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/lang"
	"github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/message"

	stdsql "database/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AppCountry is the client for interacting with the AppCountry builders.
	AppCountry *AppCountryClient
	// AppLang is the client for interacting with the AppLang builders.
	AppLang *AppLangClient
	// Country is the client for interacting with the Country builders.
	Country *CountryClient
	// Lang is the client for interacting with the Lang builders.
	Lang *LangClient
	// Message is the client for interacting with the Message builders.
	Message *MessageClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.AppCountry = NewAppCountryClient(c.config)
	c.AppLang = NewAppLangClient(c.config)
	c.Country = NewCountryClient(c.config)
	c.Lang = NewLangClient(c.config)
	c.Message = NewMessageClient(c.config)
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
		ctx:        ctx,
		config:     cfg,
		AppCountry: NewAppCountryClient(cfg),
		AppLang:    NewAppLangClient(cfg),
		Country:    NewCountryClient(cfg),
		Lang:       NewLangClient(cfg),
		Message:    NewMessageClient(cfg),
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
		AppCountry: NewAppCountryClient(cfg),
		AppLang:    NewAppLangClient(cfg),
		Country:    NewCountryClient(cfg),
		Lang:       NewLangClient(cfg),
		Message:    NewMessageClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AppCountry.
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
	c.AppCountry.Use(hooks...)
	c.AppLang.Use(hooks...)
	c.Country.Use(hooks...)
	c.Lang.Use(hooks...)
	c.Message.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.AppCountry.Intercept(interceptors...)
	c.AppLang.Intercept(interceptors...)
	c.Country.Intercept(interceptors...)
	c.Lang.Intercept(interceptors...)
	c.Message.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AppCountryMutation:
		return c.AppCountry.mutate(ctx, m)
	case *AppLangMutation:
		return c.AppLang.mutate(ctx, m)
	case *CountryMutation:
		return c.Country.mutate(ctx, m)
	case *LangMutation:
		return c.Lang.mutate(ctx, m)
	case *MessageMutation:
		return c.Message.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("generated: unknown mutation type %T", m)
	}
}

// AppCountryClient is a client for the AppCountry schema.
type AppCountryClient struct {
	config
}

// NewAppCountryClient returns a client for the AppCountry from the given config.
func NewAppCountryClient(c config) *AppCountryClient {
	return &AppCountryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `appcountry.Hooks(f(g(h())))`.
func (c *AppCountryClient) Use(hooks ...Hook) {
	c.hooks.AppCountry = append(c.hooks.AppCountry, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `appcountry.Intercept(f(g(h())))`.
func (c *AppCountryClient) Intercept(interceptors ...Interceptor) {
	c.inters.AppCountry = append(c.inters.AppCountry, interceptors...)
}

// Create returns a builder for creating a AppCountry entity.
func (c *AppCountryClient) Create() *AppCountryCreate {
	mutation := newAppCountryMutation(c.config, OpCreate)
	return &AppCountryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AppCountry entities.
func (c *AppCountryClient) CreateBulk(builders ...*AppCountryCreate) *AppCountryCreateBulk {
	return &AppCountryCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AppCountryClient) MapCreateBulk(slice any, setFunc func(*AppCountryCreate, int)) *AppCountryCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AppCountryCreateBulk{err: fmt.Errorf("calling to AppCountryClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AppCountryCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AppCountryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AppCountry.
func (c *AppCountryClient) Update() *AppCountryUpdate {
	mutation := newAppCountryMutation(c.config, OpUpdate)
	return &AppCountryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AppCountryClient) UpdateOne(ac *AppCountry) *AppCountryUpdateOne {
	mutation := newAppCountryMutation(c.config, OpUpdateOne, withAppCountry(ac))
	return &AppCountryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AppCountryClient) UpdateOneID(id uint32) *AppCountryUpdateOne {
	mutation := newAppCountryMutation(c.config, OpUpdateOne, withAppCountryID(id))
	return &AppCountryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AppCountry.
func (c *AppCountryClient) Delete() *AppCountryDelete {
	mutation := newAppCountryMutation(c.config, OpDelete)
	return &AppCountryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AppCountryClient) DeleteOne(ac *AppCountry) *AppCountryDeleteOne {
	return c.DeleteOneID(ac.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AppCountryClient) DeleteOneID(id uint32) *AppCountryDeleteOne {
	builder := c.Delete().Where(appcountry.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AppCountryDeleteOne{builder}
}

// Query returns a query builder for AppCountry.
func (c *AppCountryClient) Query() *AppCountryQuery {
	return &AppCountryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAppCountry},
		inters: c.Interceptors(),
	}
}

// Get returns a AppCountry entity by its id.
func (c *AppCountryClient) Get(ctx context.Context, id uint32) (*AppCountry, error) {
	return c.Query().Where(appcountry.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AppCountryClient) GetX(ctx context.Context, id uint32) *AppCountry {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AppCountryClient) Hooks() []Hook {
	return c.hooks.AppCountry
}

// Interceptors returns the client interceptors.
func (c *AppCountryClient) Interceptors() []Interceptor {
	return c.inters.AppCountry
}

func (c *AppCountryClient) mutate(ctx context.Context, m *AppCountryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AppCountryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AppCountryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AppCountryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AppCountryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown AppCountry mutation op: %q", m.Op())
	}
}

// AppLangClient is a client for the AppLang schema.
type AppLangClient struct {
	config
}

// NewAppLangClient returns a client for the AppLang from the given config.
func NewAppLangClient(c config) *AppLangClient {
	return &AppLangClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `applang.Hooks(f(g(h())))`.
func (c *AppLangClient) Use(hooks ...Hook) {
	c.hooks.AppLang = append(c.hooks.AppLang, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `applang.Intercept(f(g(h())))`.
func (c *AppLangClient) Intercept(interceptors ...Interceptor) {
	c.inters.AppLang = append(c.inters.AppLang, interceptors...)
}

// Create returns a builder for creating a AppLang entity.
func (c *AppLangClient) Create() *AppLangCreate {
	mutation := newAppLangMutation(c.config, OpCreate)
	return &AppLangCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AppLang entities.
func (c *AppLangClient) CreateBulk(builders ...*AppLangCreate) *AppLangCreateBulk {
	return &AppLangCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AppLangClient) MapCreateBulk(slice any, setFunc func(*AppLangCreate, int)) *AppLangCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AppLangCreateBulk{err: fmt.Errorf("calling to AppLangClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AppLangCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AppLangCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AppLang.
func (c *AppLangClient) Update() *AppLangUpdate {
	mutation := newAppLangMutation(c.config, OpUpdate)
	return &AppLangUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AppLangClient) UpdateOne(al *AppLang) *AppLangUpdateOne {
	mutation := newAppLangMutation(c.config, OpUpdateOne, withAppLang(al))
	return &AppLangUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AppLangClient) UpdateOneID(id uint32) *AppLangUpdateOne {
	mutation := newAppLangMutation(c.config, OpUpdateOne, withAppLangID(id))
	return &AppLangUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AppLang.
func (c *AppLangClient) Delete() *AppLangDelete {
	mutation := newAppLangMutation(c.config, OpDelete)
	return &AppLangDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AppLangClient) DeleteOne(al *AppLang) *AppLangDeleteOne {
	return c.DeleteOneID(al.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AppLangClient) DeleteOneID(id uint32) *AppLangDeleteOne {
	builder := c.Delete().Where(applang.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AppLangDeleteOne{builder}
}

// Query returns a query builder for AppLang.
func (c *AppLangClient) Query() *AppLangQuery {
	return &AppLangQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAppLang},
		inters: c.Interceptors(),
	}
}

// Get returns a AppLang entity by its id.
func (c *AppLangClient) Get(ctx context.Context, id uint32) (*AppLang, error) {
	return c.Query().Where(applang.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AppLangClient) GetX(ctx context.Context, id uint32) *AppLang {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AppLangClient) Hooks() []Hook {
	return c.hooks.AppLang
}

// Interceptors returns the client interceptors.
func (c *AppLangClient) Interceptors() []Interceptor {
	return c.inters.AppLang
}

func (c *AppLangClient) mutate(ctx context.Context, m *AppLangMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AppLangCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AppLangUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AppLangUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AppLangDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown AppLang mutation op: %q", m.Op())
	}
}

// CountryClient is a client for the Country schema.
type CountryClient struct {
	config
}

// NewCountryClient returns a client for the Country from the given config.
func NewCountryClient(c config) *CountryClient {
	return &CountryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `country.Hooks(f(g(h())))`.
func (c *CountryClient) Use(hooks ...Hook) {
	c.hooks.Country = append(c.hooks.Country, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `country.Intercept(f(g(h())))`.
func (c *CountryClient) Intercept(interceptors ...Interceptor) {
	c.inters.Country = append(c.inters.Country, interceptors...)
}

// Create returns a builder for creating a Country entity.
func (c *CountryClient) Create() *CountryCreate {
	mutation := newCountryMutation(c.config, OpCreate)
	return &CountryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Country entities.
func (c *CountryClient) CreateBulk(builders ...*CountryCreate) *CountryCreateBulk {
	return &CountryCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CountryClient) MapCreateBulk(slice any, setFunc func(*CountryCreate, int)) *CountryCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CountryCreateBulk{err: fmt.Errorf("calling to CountryClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CountryCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CountryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Country.
func (c *CountryClient) Update() *CountryUpdate {
	mutation := newCountryMutation(c.config, OpUpdate)
	return &CountryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CountryClient) UpdateOne(co *Country) *CountryUpdateOne {
	mutation := newCountryMutation(c.config, OpUpdateOne, withCountry(co))
	return &CountryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CountryClient) UpdateOneID(id uint32) *CountryUpdateOne {
	mutation := newCountryMutation(c.config, OpUpdateOne, withCountryID(id))
	return &CountryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Country.
func (c *CountryClient) Delete() *CountryDelete {
	mutation := newCountryMutation(c.config, OpDelete)
	return &CountryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CountryClient) DeleteOne(co *Country) *CountryDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CountryClient) DeleteOneID(id uint32) *CountryDeleteOne {
	builder := c.Delete().Where(country.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CountryDeleteOne{builder}
}

// Query returns a query builder for Country.
func (c *CountryClient) Query() *CountryQuery {
	return &CountryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCountry},
		inters: c.Interceptors(),
	}
}

// Get returns a Country entity by its id.
func (c *CountryClient) Get(ctx context.Context, id uint32) (*Country, error) {
	return c.Query().Where(country.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CountryClient) GetX(ctx context.Context, id uint32) *Country {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CountryClient) Hooks() []Hook {
	return c.hooks.Country
}

// Interceptors returns the client interceptors.
func (c *CountryClient) Interceptors() []Interceptor {
	return c.inters.Country
}

func (c *CountryClient) mutate(ctx context.Context, m *CountryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CountryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CountryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CountryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CountryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Country mutation op: %q", m.Op())
	}
}

// LangClient is a client for the Lang schema.
type LangClient struct {
	config
}

// NewLangClient returns a client for the Lang from the given config.
func NewLangClient(c config) *LangClient {
	return &LangClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `lang.Hooks(f(g(h())))`.
func (c *LangClient) Use(hooks ...Hook) {
	c.hooks.Lang = append(c.hooks.Lang, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `lang.Intercept(f(g(h())))`.
func (c *LangClient) Intercept(interceptors ...Interceptor) {
	c.inters.Lang = append(c.inters.Lang, interceptors...)
}

// Create returns a builder for creating a Lang entity.
func (c *LangClient) Create() *LangCreate {
	mutation := newLangMutation(c.config, OpCreate)
	return &LangCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Lang entities.
func (c *LangClient) CreateBulk(builders ...*LangCreate) *LangCreateBulk {
	return &LangCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *LangClient) MapCreateBulk(slice any, setFunc func(*LangCreate, int)) *LangCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &LangCreateBulk{err: fmt.Errorf("calling to LangClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*LangCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &LangCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Lang.
func (c *LangClient) Update() *LangUpdate {
	mutation := newLangMutation(c.config, OpUpdate)
	return &LangUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LangClient) UpdateOne(l *Lang) *LangUpdateOne {
	mutation := newLangMutation(c.config, OpUpdateOne, withLang(l))
	return &LangUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LangClient) UpdateOneID(id uint32) *LangUpdateOne {
	mutation := newLangMutation(c.config, OpUpdateOne, withLangID(id))
	return &LangUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Lang.
func (c *LangClient) Delete() *LangDelete {
	mutation := newLangMutation(c.config, OpDelete)
	return &LangDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LangClient) DeleteOne(l *Lang) *LangDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LangClient) DeleteOneID(id uint32) *LangDeleteOne {
	builder := c.Delete().Where(lang.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LangDeleteOne{builder}
}

// Query returns a query builder for Lang.
func (c *LangClient) Query() *LangQuery {
	return &LangQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLang},
		inters: c.Interceptors(),
	}
}

// Get returns a Lang entity by its id.
func (c *LangClient) Get(ctx context.Context, id uint32) (*Lang, error) {
	return c.Query().Where(lang.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LangClient) GetX(ctx context.Context, id uint32) *Lang {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *LangClient) Hooks() []Hook {
	return c.hooks.Lang
}

// Interceptors returns the client interceptors.
func (c *LangClient) Interceptors() []Interceptor {
	return c.inters.Lang
}

func (c *LangClient) mutate(ctx context.Context, m *LangMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LangCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LangUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LangUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LangDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Lang mutation op: %q", m.Op())
	}
}

// MessageClient is a client for the Message schema.
type MessageClient struct {
	config
}

// NewMessageClient returns a client for the Message from the given config.
func NewMessageClient(c config) *MessageClient {
	return &MessageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `message.Hooks(f(g(h())))`.
func (c *MessageClient) Use(hooks ...Hook) {
	c.hooks.Message = append(c.hooks.Message, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `message.Intercept(f(g(h())))`.
func (c *MessageClient) Intercept(interceptors ...Interceptor) {
	c.inters.Message = append(c.inters.Message, interceptors...)
}

// Create returns a builder for creating a Message entity.
func (c *MessageClient) Create() *MessageCreate {
	mutation := newMessageMutation(c.config, OpCreate)
	return &MessageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Message entities.
func (c *MessageClient) CreateBulk(builders ...*MessageCreate) *MessageCreateBulk {
	return &MessageCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MessageClient) MapCreateBulk(slice any, setFunc func(*MessageCreate, int)) *MessageCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MessageCreateBulk{err: fmt.Errorf("calling to MessageClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MessageCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MessageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Message.
func (c *MessageClient) Update() *MessageUpdate {
	mutation := newMessageMutation(c.config, OpUpdate)
	return &MessageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MessageClient) UpdateOne(m *Message) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessage(m))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MessageClient) UpdateOneID(id uint32) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessageID(id))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Message.
func (c *MessageClient) Delete() *MessageDelete {
	mutation := newMessageMutation(c.config, OpDelete)
	return &MessageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MessageClient) DeleteOne(m *Message) *MessageDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MessageClient) DeleteOneID(id uint32) *MessageDeleteOne {
	builder := c.Delete().Where(message.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MessageDeleteOne{builder}
}

// Query returns a query builder for Message.
func (c *MessageClient) Query() *MessageQuery {
	return &MessageQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMessage},
		inters: c.Interceptors(),
	}
}

// Get returns a Message entity by its id.
func (c *MessageClient) Get(ctx context.Context, id uint32) (*Message, error) {
	return c.Query().Where(message.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MessageClient) GetX(ctx context.Context, id uint32) *Message {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MessageClient) Hooks() []Hook {
	return c.hooks.Message
}

// Interceptors returns the client interceptors.
func (c *MessageClient) Interceptors() []Interceptor {
	return c.inters.Message
}

func (c *MessageClient) mutate(ctx context.Context, m *MessageMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MessageCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MessageUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MessageDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Message mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		AppCountry, AppLang, Country, Lang, Message []ent.Hook
	}
	inters struct {
		AppCountry, AppLang, Country, Lang, Message []ent.Interceptor
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

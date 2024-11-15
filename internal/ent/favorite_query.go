// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/duc-cnzj/mars/v5/internal/ent/favorite"
	"github.com/duc-cnzj/mars/v5/internal/ent/namespace"
	"github.com/duc-cnzj/mars/v5/internal/ent/predicate"
)

// FavoriteQuery is the builder for querying Favorite entities.
type FavoriteQuery struct {
	config
	ctx           *QueryContext
	order         []favorite.OrderOption
	inters        []Interceptor
	predicates    []predicate.Favorite
	withNamespace *NamespaceQuery
	modifiers     []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FavoriteQuery builder.
func (fq *FavoriteQuery) Where(ps ...predicate.Favorite) *FavoriteQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit the number of records to be returned by this query.
func (fq *FavoriteQuery) Limit(limit int) *FavoriteQuery {
	fq.ctx.Limit = &limit
	return fq
}

// Offset to start from.
func (fq *FavoriteQuery) Offset(offset int) *FavoriteQuery {
	fq.ctx.Offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FavoriteQuery) Unique(unique bool) *FavoriteQuery {
	fq.ctx.Unique = &unique
	return fq
}

// Order specifies how the records should be ordered.
func (fq *FavoriteQuery) Order(o ...favorite.OrderOption) *FavoriteQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryNamespace chains the current query on the "namespace" edge.
func (fq *FavoriteQuery) QueryNamespace() *NamespaceQuery {
	query := (&NamespaceClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(favorite.Table, favorite.FieldID, selector),
			sqlgraph.To(namespace.Table, namespace.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, favorite.NamespaceTable, favorite.NamespaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Favorite entity from the query.
// Returns a *NotFoundError when no Favorite was found.
func (fq *FavoriteQuery) First(ctx context.Context) (*Favorite, error) {
	nodes, err := fq.Limit(1).All(setContextOp(ctx, fq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{favorite.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FavoriteQuery) FirstX(ctx context.Context) *Favorite {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Favorite ID from the query.
// Returns a *NotFoundError when no Favorite ID was found.
func (fq *FavoriteQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(1).IDs(setContextOp(ctx, fq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{favorite.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FavoriteQuery) FirstIDX(ctx context.Context) int {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Favorite entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Favorite entity is found.
// Returns a *NotFoundError when no Favorite entities are found.
func (fq *FavoriteQuery) Only(ctx context.Context) (*Favorite, error) {
	nodes, err := fq.Limit(2).All(setContextOp(ctx, fq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{favorite.Label}
	default:
		return nil, &NotSingularError{favorite.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FavoriteQuery) OnlyX(ctx context.Context) *Favorite {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Favorite ID in the query.
// Returns a *NotSingularError when more than one Favorite ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FavoriteQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(2).IDs(setContextOp(ctx, fq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{favorite.Label}
	default:
		err = &NotSingularError{favorite.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FavoriteQuery) OnlyIDX(ctx context.Context) int {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Favorites.
func (fq *FavoriteQuery) All(ctx context.Context) ([]*Favorite, error) {
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryAll)
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Favorite, *FavoriteQuery]()
	return withInterceptors[[]*Favorite](ctx, fq, qr, fq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fq *FavoriteQuery) AllX(ctx context.Context) []*Favorite {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Favorite IDs.
func (fq *FavoriteQuery) IDs(ctx context.Context) (ids []int, err error) {
	if fq.ctx.Unique == nil && fq.path != nil {
		fq.Unique(true)
	}
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryIDs)
	if err = fq.Select(favorite.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FavoriteQuery) IDsX(ctx context.Context) []int {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FavoriteQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryCount)
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fq, querierCount[*FavoriteQuery](), fq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FavoriteQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FavoriteQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryExist)
	switch _, err := fq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FavoriteQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FavoriteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FavoriteQuery) Clone() *FavoriteQuery {
	if fq == nil {
		return nil
	}
	return &FavoriteQuery{
		config:        fq.config,
		ctx:           fq.ctx.Clone(),
		order:         append([]favorite.OrderOption{}, fq.order...),
		inters:        append([]Interceptor{}, fq.inters...),
		predicates:    append([]predicate.Favorite{}, fq.predicates...),
		withNamespace: fq.withNamespace.Clone(),
		// clone intermediate query.
		sql:  fq.sql.Clone(),
		path: fq.path,
	}
}

// WithNamespace tells the query-builder to eager-load the nodes that are connected to
// the "namespace" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FavoriteQuery) WithNamespace(opts ...func(*NamespaceQuery)) *FavoriteQuery {
	query := (&NamespaceClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withNamespace = query
	return fq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Email string `json:"email,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Favorite.Query().
//		GroupBy(favorite.FieldEmail).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *FavoriteQuery) GroupBy(field string, fields ...string) *FavoriteGroupBy {
	fq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FavoriteGroupBy{build: fq}
	grbuild.flds = &fq.ctx.Fields
	grbuild.label = favorite.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Email string `json:"email,omitempty"`
//	}
//
//	client.Favorite.Query().
//		Select(favorite.FieldEmail).
//		Scan(ctx, &v)
func (fq *FavoriteQuery) Select(fields ...string) *FavoriteSelect {
	fq.ctx.Fields = append(fq.ctx.Fields, fields...)
	sbuild := &FavoriteSelect{FavoriteQuery: fq}
	sbuild.label = favorite.Label
	sbuild.flds, sbuild.scan = &fq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FavoriteSelect configured with the given aggregations.
func (fq *FavoriteQuery) Aggregate(fns ...AggregateFunc) *FavoriteSelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *FavoriteQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fq); err != nil {
				return err
			}
		}
	}
	for _, f := range fq.ctx.Fields {
		if !favorite.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FavoriteQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Favorite, error) {
	var (
		nodes       = []*Favorite{}
		_spec       = fq.querySpec()
		loadedTypes = [1]bool{
			fq.withNamespace != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Favorite).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Favorite{config: fq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(fq.modifiers) > 0 {
		_spec.Modifiers = fq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fq.withNamespace; query != nil {
		if err := fq.loadNamespace(ctx, query, nodes, nil,
			func(n *Favorite, e *Namespace) { n.Edges.Namespace = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *FavoriteQuery) loadNamespace(ctx context.Context, query *NamespaceQuery, nodes []*Favorite, init func(*Favorite), assign func(*Favorite, *Namespace)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Favorite)
	for i := range nodes {
		fk := nodes[i].NamespaceID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(namespace.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "namespace_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (fq *FavoriteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	if len(fq.modifiers) > 0 {
		_spec.Modifiers = fq.modifiers
	}
	_spec.Node.Columns = fq.ctx.Fields
	if len(fq.ctx.Fields) > 0 {
		_spec.Unique = fq.ctx.Unique != nil && *fq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FavoriteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(favorite.Table, favorite.Columns, sqlgraph.NewFieldSpec(favorite.FieldID, field.TypeInt))
	_spec.From = fq.sql
	if unique := fq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fq.path != nil {
		_spec.Unique = true
	}
	if fields := fq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, favorite.FieldID)
		for i := range fields {
			if fields[i] != favorite.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if fq.withNamespace != nil {
			_spec.Node.AddColumnOnce(favorite.FieldNamespaceID)
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FavoriteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(favorite.Table)
	columns := fq.ctx.Fields
	if len(columns) == 0 {
		columns = favorite.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.ctx.Unique != nil && *fq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range fq.modifiers {
		m(selector)
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (fq *FavoriteQuery) ForUpdate(opts ...sql.LockOption) *FavoriteQuery {
	if fq.driver.Dialect() == dialect.Postgres {
		fq.Unique(false)
	}
	fq.modifiers = append(fq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return fq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (fq *FavoriteQuery) ForShare(opts ...sql.LockOption) *FavoriteQuery {
	if fq.driver.Dialect() == dialect.Postgres {
		fq.Unique(false)
	}
	fq.modifiers = append(fq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return fq
}

// FavoriteGroupBy is the group-by builder for Favorite entities.
type FavoriteGroupBy struct {
	selector
	build *FavoriteQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FavoriteGroupBy) Aggregate(fns ...AggregateFunc) *FavoriteGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the selector query and scans the result into the given value.
func (fgb *FavoriteGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fgb.build.ctx, ent.OpQueryGroupBy)
	if err := fgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FavoriteQuery, *FavoriteGroupBy](ctx, fgb.build, fgb, fgb.build.inters, v)
}

func (fgb *FavoriteGroupBy) sqlScan(ctx context.Context, root *FavoriteQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fgb.flds)+len(fgb.fns))
		for _, f := range *fgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FavoriteSelect is the builder for selecting fields of Favorite entities.
type FavoriteSelect struct {
	*FavoriteQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FavoriteSelect) Aggregate(fns ...AggregateFunc) *FavoriteSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FavoriteSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fs.ctx, ent.OpQuerySelect)
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FavoriteQuery, *FavoriteSelect](ctx, fs.FavoriteQuery, fs, fs.inters, v)
}

func (fs *FavoriteSelect) sqlScan(ctx context.Context, root *FavoriteQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fs.fns))
	for _, fn := range fs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

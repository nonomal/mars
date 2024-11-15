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
	"github.com/duc-cnzj/mars/v5/internal/ent/changelog"
	"github.com/duc-cnzj/mars/v5/internal/ent/predicate"
	"github.com/duc-cnzj/mars/v5/internal/ent/project"
)

// ChangelogQuery is the builder for querying Changelog entities.
type ChangelogQuery struct {
	config
	ctx         *QueryContext
	order       []changelog.OrderOption
	inters      []Interceptor
	predicates  []predicate.Changelog
	withProject *ProjectQuery
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ChangelogQuery builder.
func (cq *ChangelogQuery) Where(ps ...predicate.Changelog) *ChangelogQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ChangelogQuery) Limit(limit int) *ChangelogQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *ChangelogQuery) Offset(offset int) *ChangelogQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ChangelogQuery) Unique(unique bool) *ChangelogQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ChangelogQuery) Order(o ...changelog.OrderOption) *ChangelogQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryProject chains the current query on the "project" edge.
func (cq *ChangelogQuery) QueryProject() *ProjectQuery {
	query := (&ProjectClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(changelog.Table, changelog.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, changelog.ProjectTable, changelog.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Changelog entity from the query.
// Returns a *NotFoundError when no Changelog was found.
func (cq *ChangelogQuery) First(ctx context.Context) (*Changelog, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{changelog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ChangelogQuery) FirstX(ctx context.Context) *Changelog {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Changelog ID from the query.
// Returns a *NotFoundError when no Changelog ID was found.
func (cq *ChangelogQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{changelog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ChangelogQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Changelog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Changelog entity is found.
// Returns a *NotFoundError when no Changelog entities are found.
func (cq *ChangelogQuery) Only(ctx context.Context) (*Changelog, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{changelog.Label}
	default:
		return nil, &NotSingularError{changelog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ChangelogQuery) OnlyX(ctx context.Context) *Changelog {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Changelog ID in the query.
// Returns a *NotSingularError when more than one Changelog ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ChangelogQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{changelog.Label}
	default:
		err = &NotSingularError{changelog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ChangelogQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Changelogs.
func (cq *ChangelogQuery) All(ctx context.Context) ([]*Changelog, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryAll)
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Changelog, *ChangelogQuery]()
	return withInterceptors[[]*Changelog](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ChangelogQuery) AllX(ctx context.Context) []*Changelog {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Changelog IDs.
func (cq *ChangelogQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryIDs)
	if err = cq.Select(changelog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ChangelogQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ChangelogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryCount)
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ChangelogQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ChangelogQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ChangelogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryExist)
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ChangelogQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ChangelogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ChangelogQuery) Clone() *ChangelogQuery {
	if cq == nil {
		return nil
	}
	return &ChangelogQuery{
		config:      cq.config,
		ctx:         cq.ctx.Clone(),
		order:       append([]changelog.OrderOption{}, cq.order...),
		inters:      append([]Interceptor{}, cq.inters...),
		predicates:  append([]predicate.Changelog{}, cq.predicates...),
		withProject: cq.withProject.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ChangelogQuery) WithProject(opts ...func(*ProjectQuery)) *ChangelogQuery {
	query := (&ProjectClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withProject = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Changelog.Query().
//		GroupBy(changelog.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ChangelogQuery) GroupBy(field string, fields ...string) *ChangelogGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ChangelogGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = changelog.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Changelog.Query().
//		Select(changelog.FieldCreatedAt).
//		Scan(ctx, &v)
func (cq *ChangelogQuery) Select(fields ...string) *ChangelogSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &ChangelogSelect{ChangelogQuery: cq}
	sbuild.label = changelog.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ChangelogSelect configured with the given aggregations.
func (cq *ChangelogQuery) Aggregate(fns ...AggregateFunc) *ChangelogSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ChangelogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !changelog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ChangelogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Changelog, error) {
	var (
		nodes       = []*Changelog{}
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withProject != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Changelog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Changelog{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withProject; query != nil {
		if err := cq.loadProject(ctx, query, nodes, nil,
			func(n *Changelog, e *Project) { n.Edges.Project = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ChangelogQuery) loadProject(ctx context.Context, query *ProjectQuery, nodes []*Changelog, init func(*Changelog), assign func(*Changelog, *Project)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Changelog)
	for i := range nodes {
		fk := nodes[i].ProjectID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(project.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "project_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cq *ChangelogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ChangelogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(changelog.Table, changelog.Columns, sqlgraph.NewFieldSpec(changelog.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, changelog.FieldID)
		for i := range fields {
			if fields[i] != changelog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if cq.withProject != nil {
			_spec.Node.AddColumnOnce(changelog.FieldProjectID)
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ChangelogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(changelog.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = changelog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cq.modifiers {
		m(selector)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (cq *ChangelogQuery) ForUpdate(opts ...sql.LockOption) *ChangelogQuery {
	if cq.driver.Dialect() == dialect.Postgres {
		cq.Unique(false)
	}
	cq.modifiers = append(cq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return cq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (cq *ChangelogQuery) ForShare(opts ...sql.LockOption) *ChangelogQuery {
	if cq.driver.Dialect() == dialect.Postgres {
		cq.Unique(false)
	}
	cq.modifiers = append(cq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return cq
}

// ChangelogGroupBy is the group-by builder for Changelog entities.
type ChangelogGroupBy struct {
	selector
	build *ChangelogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ChangelogGroupBy) Aggregate(fns ...AggregateFunc) *ChangelogGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ChangelogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, ent.OpQueryGroupBy)
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChangelogQuery, *ChangelogGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ChangelogGroupBy) sqlScan(ctx context.Context, root *ChangelogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ChangelogSelect is the builder for selecting fields of Changelog entities.
type ChangelogSelect struct {
	*ChangelogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ChangelogSelect) Aggregate(fns ...AggregateFunc) *ChangelogSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ChangelogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, ent.OpQuerySelect)
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChangelogQuery, *ChangelogSelect](ctx, cs.ChangelogQuery, cs, cs.inters, v)
}

func (cs *ChangelogSelect) sqlScan(ctx context.Context, root *ChangelogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

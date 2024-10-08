// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/vector-ops/lifefolio/ent/institution"
	"github.com/vector-ops/lifefolio/ent/medicalrecord"
	"github.com/vector-ops/lifefolio/ent/predicate"
	"github.com/vector-ops/lifefolio/ent/recordaccess"
)

// RecordAccessQuery is the builder for querying RecordAccess entities.
type RecordAccessQuery struct {
	config
	ctx               *QueryContext
	order             []recordaccess.OrderOption
	inters            []Interceptor
	predicates        []predicate.RecordAccess
	withMedicalrecord *MedicalRecordQuery
	withInstitution   *InstitutionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RecordAccessQuery builder.
func (raq *RecordAccessQuery) Where(ps ...predicate.RecordAccess) *RecordAccessQuery {
	raq.predicates = append(raq.predicates, ps...)
	return raq
}

// Limit the number of records to be returned by this query.
func (raq *RecordAccessQuery) Limit(limit int) *RecordAccessQuery {
	raq.ctx.Limit = &limit
	return raq
}

// Offset to start from.
func (raq *RecordAccessQuery) Offset(offset int) *RecordAccessQuery {
	raq.ctx.Offset = &offset
	return raq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (raq *RecordAccessQuery) Unique(unique bool) *RecordAccessQuery {
	raq.ctx.Unique = &unique
	return raq
}

// Order specifies how the records should be ordered.
func (raq *RecordAccessQuery) Order(o ...recordaccess.OrderOption) *RecordAccessQuery {
	raq.order = append(raq.order, o...)
	return raq
}

// QueryMedicalrecord chains the current query on the "medicalrecord" edge.
func (raq *RecordAccessQuery) QueryMedicalrecord() *MedicalRecordQuery {
	query := (&MedicalRecordClient{config: raq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := raq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := raq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(recordaccess.Table, recordaccess.FieldID, selector),
			sqlgraph.To(medicalrecord.Table, medicalrecord.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, recordaccess.MedicalrecordTable, recordaccess.MedicalrecordColumn),
		)
		fromU = sqlgraph.SetNeighbors(raq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInstitution chains the current query on the "institution" edge.
func (raq *RecordAccessQuery) QueryInstitution() *InstitutionQuery {
	query := (&InstitutionClient{config: raq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := raq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := raq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(recordaccess.Table, recordaccess.FieldID, selector),
			sqlgraph.To(institution.Table, institution.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, recordaccess.InstitutionTable, recordaccess.InstitutionColumn),
		)
		fromU = sqlgraph.SetNeighbors(raq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RecordAccess entity from the query.
// Returns a *NotFoundError when no RecordAccess was found.
func (raq *RecordAccessQuery) First(ctx context.Context) (*RecordAccess, error) {
	nodes, err := raq.Limit(1).All(setContextOp(ctx, raq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{recordaccess.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (raq *RecordAccessQuery) FirstX(ctx context.Context) *RecordAccess {
	node, err := raq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RecordAccess ID from the query.
// Returns a *NotFoundError when no RecordAccess ID was found.
func (raq *RecordAccessQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = raq.Limit(1).IDs(setContextOp(ctx, raq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{recordaccess.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (raq *RecordAccessQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := raq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RecordAccess entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RecordAccess entity is found.
// Returns a *NotFoundError when no RecordAccess entities are found.
func (raq *RecordAccessQuery) Only(ctx context.Context) (*RecordAccess, error) {
	nodes, err := raq.Limit(2).All(setContextOp(ctx, raq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{recordaccess.Label}
	default:
		return nil, &NotSingularError{recordaccess.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (raq *RecordAccessQuery) OnlyX(ctx context.Context) *RecordAccess {
	node, err := raq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RecordAccess ID in the query.
// Returns a *NotSingularError when more than one RecordAccess ID is found.
// Returns a *NotFoundError when no entities are found.
func (raq *RecordAccessQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = raq.Limit(2).IDs(setContextOp(ctx, raq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{recordaccess.Label}
	default:
		err = &NotSingularError{recordaccess.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (raq *RecordAccessQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := raq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RecordAccesses.
func (raq *RecordAccessQuery) All(ctx context.Context) ([]*RecordAccess, error) {
	ctx = setContextOp(ctx, raq.ctx, ent.OpQueryAll)
	if err := raq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RecordAccess, *RecordAccessQuery]()
	return withInterceptors[[]*RecordAccess](ctx, raq, qr, raq.inters)
}

// AllX is like All, but panics if an error occurs.
func (raq *RecordAccessQuery) AllX(ctx context.Context) []*RecordAccess {
	nodes, err := raq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RecordAccess IDs.
func (raq *RecordAccessQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if raq.ctx.Unique == nil && raq.path != nil {
		raq.Unique(true)
	}
	ctx = setContextOp(ctx, raq.ctx, ent.OpQueryIDs)
	if err = raq.Select(recordaccess.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (raq *RecordAccessQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := raq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (raq *RecordAccessQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, raq.ctx, ent.OpQueryCount)
	if err := raq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, raq, querierCount[*RecordAccessQuery](), raq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (raq *RecordAccessQuery) CountX(ctx context.Context) int {
	count, err := raq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (raq *RecordAccessQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, raq.ctx, ent.OpQueryExist)
	switch _, err := raq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (raq *RecordAccessQuery) ExistX(ctx context.Context) bool {
	exist, err := raq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RecordAccessQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (raq *RecordAccessQuery) Clone() *RecordAccessQuery {
	if raq == nil {
		return nil
	}
	return &RecordAccessQuery{
		config:            raq.config,
		ctx:               raq.ctx.Clone(),
		order:             append([]recordaccess.OrderOption{}, raq.order...),
		inters:            append([]Interceptor{}, raq.inters...),
		predicates:        append([]predicate.RecordAccess{}, raq.predicates...),
		withMedicalrecord: raq.withMedicalrecord.Clone(),
		withInstitution:   raq.withInstitution.Clone(),
		// clone intermediate query.
		sql:  raq.sql.Clone(),
		path: raq.path,
	}
}

// WithMedicalrecord tells the query-builder to eager-load the nodes that are connected to
// the "medicalrecord" edge. The optional arguments are used to configure the query builder of the edge.
func (raq *RecordAccessQuery) WithMedicalrecord(opts ...func(*MedicalRecordQuery)) *RecordAccessQuery {
	query := (&MedicalRecordClient{config: raq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	raq.withMedicalrecord = query
	return raq
}

// WithInstitution tells the query-builder to eager-load the nodes that are connected to
// the "institution" edge. The optional arguments are used to configure the query builder of the edge.
func (raq *RecordAccessQuery) WithInstitution(opts ...func(*InstitutionQuery)) *RecordAccessQuery {
	query := (&InstitutionClient{config: raq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	raq.withInstitution = query
	return raq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		RecordID uuid.UUID `json:"record_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RecordAccess.Query().
//		GroupBy(recordaccess.FieldRecordID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (raq *RecordAccessQuery) GroupBy(field string, fields ...string) *RecordAccessGroupBy {
	raq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RecordAccessGroupBy{build: raq}
	grbuild.flds = &raq.ctx.Fields
	grbuild.label = recordaccess.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		RecordID uuid.UUID `json:"record_id,omitempty"`
//	}
//
//	client.RecordAccess.Query().
//		Select(recordaccess.FieldRecordID).
//		Scan(ctx, &v)
func (raq *RecordAccessQuery) Select(fields ...string) *RecordAccessSelect {
	raq.ctx.Fields = append(raq.ctx.Fields, fields...)
	sbuild := &RecordAccessSelect{RecordAccessQuery: raq}
	sbuild.label = recordaccess.Label
	sbuild.flds, sbuild.scan = &raq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RecordAccessSelect configured with the given aggregations.
func (raq *RecordAccessQuery) Aggregate(fns ...AggregateFunc) *RecordAccessSelect {
	return raq.Select().Aggregate(fns...)
}

func (raq *RecordAccessQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range raq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, raq); err != nil {
				return err
			}
		}
	}
	for _, f := range raq.ctx.Fields {
		if !recordaccess.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if raq.path != nil {
		prev, err := raq.path(ctx)
		if err != nil {
			return err
		}
		raq.sql = prev
	}
	return nil
}

func (raq *RecordAccessQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RecordAccess, error) {
	var (
		nodes       = []*RecordAccess{}
		_spec       = raq.querySpec()
		loadedTypes = [2]bool{
			raq.withMedicalrecord != nil,
			raq.withInstitution != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RecordAccess).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RecordAccess{config: raq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, raq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := raq.withMedicalrecord; query != nil {
		if err := raq.loadMedicalrecord(ctx, query, nodes, nil,
			func(n *RecordAccess, e *MedicalRecord) { n.Edges.Medicalrecord = e }); err != nil {
			return nil, err
		}
	}
	if query := raq.withInstitution; query != nil {
		if err := raq.loadInstitution(ctx, query, nodes, nil,
			func(n *RecordAccess, e *Institution) { n.Edges.Institution = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (raq *RecordAccessQuery) loadMedicalrecord(ctx context.Context, query *MedicalRecordQuery, nodes []*RecordAccess, init func(*RecordAccess), assign func(*RecordAccess, *MedicalRecord)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*RecordAccess)
	for i := range nodes {
		fk := nodes[i].RecordID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(medicalrecord.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "record_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (raq *RecordAccessQuery) loadInstitution(ctx context.Context, query *InstitutionQuery, nodes []*RecordAccess, init func(*RecordAccess), assign func(*RecordAccess, *Institution)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*RecordAccess)
	for i := range nodes {
		fk := nodes[i].InstitutionID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(institution.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "institution_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (raq *RecordAccessQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := raq.querySpec()
	_spec.Node.Columns = raq.ctx.Fields
	if len(raq.ctx.Fields) > 0 {
		_spec.Unique = raq.ctx.Unique != nil && *raq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, raq.driver, _spec)
}

func (raq *RecordAccessQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(recordaccess.Table, recordaccess.Columns, sqlgraph.NewFieldSpec(recordaccess.FieldID, field.TypeUUID))
	_spec.From = raq.sql
	if unique := raq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if raq.path != nil {
		_spec.Unique = true
	}
	if fields := raq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, recordaccess.FieldID)
		for i := range fields {
			if fields[i] != recordaccess.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if raq.withMedicalrecord != nil {
			_spec.Node.AddColumnOnce(recordaccess.FieldRecordID)
		}
		if raq.withInstitution != nil {
			_spec.Node.AddColumnOnce(recordaccess.FieldInstitutionID)
		}
	}
	if ps := raq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := raq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := raq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := raq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (raq *RecordAccessQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(raq.driver.Dialect())
	t1 := builder.Table(recordaccess.Table)
	columns := raq.ctx.Fields
	if len(columns) == 0 {
		columns = recordaccess.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if raq.sql != nil {
		selector = raq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if raq.ctx.Unique != nil && *raq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range raq.predicates {
		p(selector)
	}
	for _, p := range raq.order {
		p(selector)
	}
	if offset := raq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := raq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RecordAccessGroupBy is the group-by builder for RecordAccess entities.
type RecordAccessGroupBy struct {
	selector
	build *RecordAccessQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ragb *RecordAccessGroupBy) Aggregate(fns ...AggregateFunc) *RecordAccessGroupBy {
	ragb.fns = append(ragb.fns, fns...)
	return ragb
}

// Scan applies the selector query and scans the result into the given value.
func (ragb *RecordAccessGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ragb.build.ctx, ent.OpQueryGroupBy)
	if err := ragb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RecordAccessQuery, *RecordAccessGroupBy](ctx, ragb.build, ragb, ragb.build.inters, v)
}

func (ragb *RecordAccessGroupBy) sqlScan(ctx context.Context, root *RecordAccessQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ragb.fns))
	for _, fn := range ragb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ragb.flds)+len(ragb.fns))
		for _, f := range *ragb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ragb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ragb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RecordAccessSelect is the builder for selecting fields of RecordAccess entities.
type RecordAccessSelect struct {
	*RecordAccessQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ras *RecordAccessSelect) Aggregate(fns ...AggregateFunc) *RecordAccessSelect {
	ras.fns = append(ras.fns, fns...)
	return ras
}

// Scan applies the selector query and scans the result into the given value.
func (ras *RecordAccessSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ras.ctx, ent.OpQuerySelect)
	if err := ras.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RecordAccessQuery, *RecordAccessSelect](ctx, ras.RecordAccessQuery, ras, ras.inters, v)
}

func (ras *RecordAccessSelect) sqlScan(ctx context.Context, root *RecordAccessQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ras.fns))
	for _, fn := range ras.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ras.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ras.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

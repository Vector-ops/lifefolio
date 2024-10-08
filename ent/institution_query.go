// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
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
	"github.com/vector-ops/lifefolio/ent/user"
)

// InstitutionQuery is the builder for querying Institution entities.
type InstitutionQuery struct {
	config
	ctx               *QueryContext
	order             []institution.OrderOption
	inters            []Interceptor
	predicates        []predicate.Institution
	withRecordaccess  *RecordAccessQuery
	withDoctor        *UserQuery
	withMedicalrecord *MedicalRecordQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InstitutionQuery builder.
func (iq *InstitutionQuery) Where(ps ...predicate.Institution) *InstitutionQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *InstitutionQuery) Limit(limit int) *InstitutionQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *InstitutionQuery) Offset(offset int) *InstitutionQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *InstitutionQuery) Unique(unique bool) *InstitutionQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *InstitutionQuery) Order(o ...institution.OrderOption) *InstitutionQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryRecordaccess chains the current query on the "recordaccess" edge.
func (iq *InstitutionQuery) QueryRecordaccess() *RecordAccessQuery {
	query := (&RecordAccessClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(institution.Table, institution.FieldID, selector),
			sqlgraph.To(recordaccess.Table, recordaccess.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, institution.RecordaccessTable, institution.RecordaccessColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDoctor chains the current query on the "doctor" edge.
func (iq *InstitutionQuery) QueryDoctor() *UserQuery {
	query := (&UserClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(institution.Table, institution.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, institution.DoctorTable, institution.DoctorColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMedicalrecord chains the current query on the "medicalrecord" edge.
func (iq *InstitutionQuery) QueryMedicalrecord() *MedicalRecordQuery {
	query := (&MedicalRecordClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(institution.Table, institution.FieldID, selector),
			sqlgraph.To(medicalrecord.Table, medicalrecord.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, institution.MedicalrecordTable, institution.MedicalrecordColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Institution entity from the query.
// Returns a *NotFoundError when no Institution was found.
func (iq *InstitutionQuery) First(ctx context.Context) (*Institution, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{institution.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *InstitutionQuery) FirstX(ctx context.Context) *Institution {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Institution ID from the query.
// Returns a *NotFoundError when no Institution ID was found.
func (iq *InstitutionQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{institution.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *InstitutionQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Institution entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Institution entity is found.
// Returns a *NotFoundError when no Institution entities are found.
func (iq *InstitutionQuery) Only(ctx context.Context) (*Institution, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{institution.Label}
	default:
		return nil, &NotSingularError{institution.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *InstitutionQuery) OnlyX(ctx context.Context) *Institution {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Institution ID in the query.
// Returns a *NotSingularError when more than one Institution ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *InstitutionQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{institution.Label}
	default:
		err = &NotSingularError{institution.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *InstitutionQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Institutions.
func (iq *InstitutionQuery) All(ctx context.Context) ([]*Institution, error) {
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryAll)
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Institution, *InstitutionQuery]()
	return withInterceptors[[]*Institution](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *InstitutionQuery) AllX(ctx context.Context) []*Institution {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Institution IDs.
func (iq *InstitutionQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryIDs)
	if err = iq.Select(institution.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *InstitutionQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *InstitutionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryCount)
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*InstitutionQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *InstitutionQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *InstitutionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryExist)
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *InstitutionQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InstitutionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *InstitutionQuery) Clone() *InstitutionQuery {
	if iq == nil {
		return nil
	}
	return &InstitutionQuery{
		config:            iq.config,
		ctx:               iq.ctx.Clone(),
		order:             append([]institution.OrderOption{}, iq.order...),
		inters:            append([]Interceptor{}, iq.inters...),
		predicates:        append([]predicate.Institution{}, iq.predicates...),
		withRecordaccess:  iq.withRecordaccess.Clone(),
		withDoctor:        iq.withDoctor.Clone(),
		withMedicalrecord: iq.withMedicalrecord.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithRecordaccess tells the query-builder to eager-load the nodes that are connected to
// the "recordaccess" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InstitutionQuery) WithRecordaccess(opts ...func(*RecordAccessQuery)) *InstitutionQuery {
	query := (&RecordAccessClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withRecordaccess = query
	return iq
}

// WithDoctor tells the query-builder to eager-load the nodes that are connected to
// the "doctor" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InstitutionQuery) WithDoctor(opts ...func(*UserQuery)) *InstitutionQuery {
	query := (&UserClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withDoctor = query
	return iq
}

// WithMedicalrecord tells the query-builder to eager-load the nodes that are connected to
// the "medicalrecord" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InstitutionQuery) WithMedicalrecord(opts ...func(*MedicalRecordQuery)) *InstitutionQuery {
	query := (&MedicalRecordClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withMedicalrecord = query
	return iq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Institution.Query().
//		GroupBy(institution.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *InstitutionQuery) GroupBy(field string, fields ...string) *InstitutionGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InstitutionGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = institution.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Institution.Query().
//		Select(institution.FieldName).
//		Scan(ctx, &v)
func (iq *InstitutionQuery) Select(fields ...string) *InstitutionSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &InstitutionSelect{InstitutionQuery: iq}
	sbuild.label = institution.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InstitutionSelect configured with the given aggregations.
func (iq *InstitutionQuery) Aggregate(fns ...AggregateFunc) *InstitutionSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *InstitutionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !institution.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *InstitutionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Institution, error) {
	var (
		nodes       = []*Institution{}
		_spec       = iq.querySpec()
		loadedTypes = [3]bool{
			iq.withRecordaccess != nil,
			iq.withDoctor != nil,
			iq.withMedicalrecord != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Institution).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Institution{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withRecordaccess; query != nil {
		if err := iq.loadRecordaccess(ctx, query, nodes,
			func(n *Institution) { n.Edges.Recordaccess = []*RecordAccess{} },
			func(n *Institution, e *RecordAccess) { n.Edges.Recordaccess = append(n.Edges.Recordaccess, e) }); err != nil {
			return nil, err
		}
	}
	if query := iq.withDoctor; query != nil {
		if err := iq.loadDoctor(ctx, query, nodes,
			func(n *Institution) { n.Edges.Doctor = []*User{} },
			func(n *Institution, e *User) { n.Edges.Doctor = append(n.Edges.Doctor, e) }); err != nil {
			return nil, err
		}
	}
	if query := iq.withMedicalrecord; query != nil {
		if err := iq.loadMedicalrecord(ctx, query, nodes,
			func(n *Institution) { n.Edges.Medicalrecord = []*MedicalRecord{} },
			func(n *Institution, e *MedicalRecord) { n.Edges.Medicalrecord = append(n.Edges.Medicalrecord, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *InstitutionQuery) loadRecordaccess(ctx context.Context, query *RecordAccessQuery, nodes []*Institution, init func(*Institution), assign func(*Institution, *RecordAccess)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Institution)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(recordaccess.FieldInstitutionID)
	}
	query.Where(predicate.RecordAccess(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(institution.RecordaccessColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.InstitutionID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "institution_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (iq *InstitutionQuery) loadDoctor(ctx context.Context, query *UserQuery, nodes []*Institution, init func(*Institution), assign func(*Institution, *User)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Institution)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.User(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(institution.DoctorColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.institution_doctor
		if fk == nil {
			return fmt.Errorf(`foreign-key "institution_doctor" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "institution_doctor" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (iq *InstitutionQuery) loadMedicalrecord(ctx context.Context, query *MedicalRecordQuery, nodes []*Institution, init func(*Institution), assign func(*Institution, *MedicalRecord)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Institution)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(medicalrecord.FieldInstitutionID)
	}
	query.Where(predicate.MedicalRecord(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(institution.MedicalrecordColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.InstitutionID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "institution_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (iq *InstitutionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *InstitutionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(institution.Table, institution.Columns, sqlgraph.NewFieldSpec(institution.FieldID, field.TypeUUID))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, institution.FieldID)
		for i := range fields {
			if fields[i] != institution.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *InstitutionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(institution.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = institution.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InstitutionGroupBy is the group-by builder for Institution entities.
type InstitutionGroupBy struct {
	selector
	build *InstitutionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *InstitutionGroupBy) Aggregate(fns ...AggregateFunc) *InstitutionGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *InstitutionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, ent.OpQueryGroupBy)
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InstitutionQuery, *InstitutionGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *InstitutionGroupBy) sqlScan(ctx context.Context, root *InstitutionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InstitutionSelect is the builder for selecting fields of Institution entities.
type InstitutionSelect struct {
	*InstitutionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *InstitutionSelect) Aggregate(fns ...AggregateFunc) *InstitutionSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *InstitutionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, ent.OpQuerySelect)
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InstitutionQuery, *InstitutionSelect](ctx, is.InstitutionQuery, is, is.inters, v)
}

func (is *InstitutionSelect) sqlScan(ctx context.Context, root *InstitutionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

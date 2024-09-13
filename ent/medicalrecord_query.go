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

// MedicalRecordQuery is the builder for querying MedicalRecord entities.
type MedicalRecordQuery struct {
	config
	ctx              *QueryContext
	order            []medicalrecord.OrderOption
	inters           []Interceptor
	predicates       []predicate.MedicalRecord
	withUser         *UserQuery
	withInstitution  *InstitutionQuery
	withRecordaccess *RecordAccessQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MedicalRecordQuery builder.
func (mrq *MedicalRecordQuery) Where(ps ...predicate.MedicalRecord) *MedicalRecordQuery {
	mrq.predicates = append(mrq.predicates, ps...)
	return mrq
}

// Limit the number of records to be returned by this query.
func (mrq *MedicalRecordQuery) Limit(limit int) *MedicalRecordQuery {
	mrq.ctx.Limit = &limit
	return mrq
}

// Offset to start from.
func (mrq *MedicalRecordQuery) Offset(offset int) *MedicalRecordQuery {
	mrq.ctx.Offset = &offset
	return mrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mrq *MedicalRecordQuery) Unique(unique bool) *MedicalRecordQuery {
	mrq.ctx.Unique = &unique
	return mrq
}

// Order specifies how the records should be ordered.
func (mrq *MedicalRecordQuery) Order(o ...medicalrecord.OrderOption) *MedicalRecordQuery {
	mrq.order = append(mrq.order, o...)
	return mrq
}

// QueryUser chains the current query on the "user" edge.
func (mrq *MedicalRecordQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: mrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(medicalrecord.Table, medicalrecord.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, medicalrecord.UserTable, medicalrecord.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(mrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInstitution chains the current query on the "institution" edge.
func (mrq *MedicalRecordQuery) QueryInstitution() *InstitutionQuery {
	query := (&InstitutionClient{config: mrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(medicalrecord.Table, medicalrecord.FieldID, selector),
			sqlgraph.To(institution.Table, institution.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, medicalrecord.InstitutionTable, medicalrecord.InstitutionColumn),
		)
		fromU = sqlgraph.SetNeighbors(mrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRecordaccess chains the current query on the "recordaccess" edge.
func (mrq *MedicalRecordQuery) QueryRecordaccess() *RecordAccessQuery {
	query := (&RecordAccessClient{config: mrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(medicalrecord.Table, medicalrecord.FieldID, selector),
			sqlgraph.To(recordaccess.Table, recordaccess.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, medicalrecord.RecordaccessTable, medicalrecord.RecordaccessColumn),
		)
		fromU = sqlgraph.SetNeighbors(mrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MedicalRecord entity from the query.
// Returns a *NotFoundError when no MedicalRecord was found.
func (mrq *MedicalRecordQuery) First(ctx context.Context) (*MedicalRecord, error) {
	nodes, err := mrq.Limit(1).All(setContextOp(ctx, mrq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{medicalrecord.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mrq *MedicalRecordQuery) FirstX(ctx context.Context) *MedicalRecord {
	node, err := mrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MedicalRecord ID from the query.
// Returns a *NotFoundError when no MedicalRecord ID was found.
func (mrq *MedicalRecordQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mrq.Limit(1).IDs(setContextOp(ctx, mrq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{medicalrecord.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mrq *MedicalRecordQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := mrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MedicalRecord entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MedicalRecord entity is found.
// Returns a *NotFoundError when no MedicalRecord entities are found.
func (mrq *MedicalRecordQuery) Only(ctx context.Context) (*MedicalRecord, error) {
	nodes, err := mrq.Limit(2).All(setContextOp(ctx, mrq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{medicalrecord.Label}
	default:
		return nil, &NotSingularError{medicalrecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mrq *MedicalRecordQuery) OnlyX(ctx context.Context) *MedicalRecord {
	node, err := mrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MedicalRecord ID in the query.
// Returns a *NotSingularError when more than one MedicalRecord ID is found.
// Returns a *NotFoundError when no entities are found.
func (mrq *MedicalRecordQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mrq.Limit(2).IDs(setContextOp(ctx, mrq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{medicalrecord.Label}
	default:
		err = &NotSingularError{medicalrecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mrq *MedicalRecordQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := mrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MedicalRecords.
func (mrq *MedicalRecordQuery) All(ctx context.Context) ([]*MedicalRecord, error) {
	ctx = setContextOp(ctx, mrq.ctx, ent.OpQueryAll)
	if err := mrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MedicalRecord, *MedicalRecordQuery]()
	return withInterceptors[[]*MedicalRecord](ctx, mrq, qr, mrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mrq *MedicalRecordQuery) AllX(ctx context.Context) []*MedicalRecord {
	nodes, err := mrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MedicalRecord IDs.
func (mrq *MedicalRecordQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if mrq.ctx.Unique == nil && mrq.path != nil {
		mrq.Unique(true)
	}
	ctx = setContextOp(ctx, mrq.ctx, ent.OpQueryIDs)
	if err = mrq.Select(medicalrecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mrq *MedicalRecordQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := mrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mrq *MedicalRecordQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mrq.ctx, ent.OpQueryCount)
	if err := mrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mrq, querierCount[*MedicalRecordQuery](), mrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mrq *MedicalRecordQuery) CountX(ctx context.Context) int {
	count, err := mrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mrq *MedicalRecordQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mrq.ctx, ent.OpQueryExist)
	switch _, err := mrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mrq *MedicalRecordQuery) ExistX(ctx context.Context) bool {
	exist, err := mrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MedicalRecordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mrq *MedicalRecordQuery) Clone() *MedicalRecordQuery {
	if mrq == nil {
		return nil
	}
	return &MedicalRecordQuery{
		config:           mrq.config,
		ctx:              mrq.ctx.Clone(),
		order:            append([]medicalrecord.OrderOption{}, mrq.order...),
		inters:           append([]Interceptor{}, mrq.inters...),
		predicates:       append([]predicate.MedicalRecord{}, mrq.predicates...),
		withUser:         mrq.withUser.Clone(),
		withInstitution:  mrq.withInstitution.Clone(),
		withRecordaccess: mrq.withRecordaccess.Clone(),
		// clone intermediate query.
		sql:  mrq.sql.Clone(),
		path: mrq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (mrq *MedicalRecordQuery) WithUser(opts ...func(*UserQuery)) *MedicalRecordQuery {
	query := (&UserClient{config: mrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mrq.withUser = query
	return mrq
}

// WithInstitution tells the query-builder to eager-load the nodes that are connected to
// the "institution" edge. The optional arguments are used to configure the query builder of the edge.
func (mrq *MedicalRecordQuery) WithInstitution(opts ...func(*InstitutionQuery)) *MedicalRecordQuery {
	query := (&InstitutionClient{config: mrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mrq.withInstitution = query
	return mrq
}

// WithRecordaccess tells the query-builder to eager-load the nodes that are connected to
// the "recordaccess" edge. The optional arguments are used to configure the query builder of the edge.
func (mrq *MedicalRecordQuery) WithRecordaccess(opts ...func(*RecordAccessQuery)) *MedicalRecordQuery {
	query := (&RecordAccessClient{config: mrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mrq.withRecordaccess = query
	return mrq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		File string `json:"file,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MedicalRecord.Query().
//		GroupBy(medicalrecord.FieldFile).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mrq *MedicalRecordQuery) GroupBy(field string, fields ...string) *MedicalRecordGroupBy {
	mrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MedicalRecordGroupBy{build: mrq}
	grbuild.flds = &mrq.ctx.Fields
	grbuild.label = medicalrecord.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		File string `json:"file,omitempty"`
//	}
//
//	client.MedicalRecord.Query().
//		Select(medicalrecord.FieldFile).
//		Scan(ctx, &v)
func (mrq *MedicalRecordQuery) Select(fields ...string) *MedicalRecordSelect {
	mrq.ctx.Fields = append(mrq.ctx.Fields, fields...)
	sbuild := &MedicalRecordSelect{MedicalRecordQuery: mrq}
	sbuild.label = medicalrecord.Label
	sbuild.flds, sbuild.scan = &mrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MedicalRecordSelect configured with the given aggregations.
func (mrq *MedicalRecordQuery) Aggregate(fns ...AggregateFunc) *MedicalRecordSelect {
	return mrq.Select().Aggregate(fns...)
}

func (mrq *MedicalRecordQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mrq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mrq); err != nil {
				return err
			}
		}
	}
	for _, f := range mrq.ctx.Fields {
		if !medicalrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mrq.path != nil {
		prev, err := mrq.path(ctx)
		if err != nil {
			return err
		}
		mrq.sql = prev
	}
	return nil
}

func (mrq *MedicalRecordQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MedicalRecord, error) {
	var (
		nodes       = []*MedicalRecord{}
		withFKs     = mrq.withFKs
		_spec       = mrq.querySpec()
		loadedTypes = [3]bool{
			mrq.withUser != nil,
			mrq.withInstitution != nil,
			mrq.withRecordaccess != nil,
		}
	)
	if mrq.withUser != nil || mrq.withInstitution != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, medicalrecord.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MedicalRecord).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MedicalRecord{config: mrq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mrq.withUser; query != nil {
		if err := mrq.loadUser(ctx, query, nodes, nil,
			func(n *MedicalRecord, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := mrq.withInstitution; query != nil {
		if err := mrq.loadInstitution(ctx, query, nodes, nil,
			func(n *MedicalRecord, e *Institution) { n.Edges.Institution = e }); err != nil {
			return nil, err
		}
	}
	if query := mrq.withRecordaccess; query != nil {
		if err := mrq.loadRecordaccess(ctx, query, nodes,
			func(n *MedicalRecord) { n.Edges.Recordaccess = []*RecordAccess{} },
			func(n *MedicalRecord, e *RecordAccess) { n.Edges.Recordaccess = append(n.Edges.Recordaccess, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mrq *MedicalRecordQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*MedicalRecord, init func(*MedicalRecord), assign func(*MedicalRecord, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*MedicalRecord)
	for i := range nodes {
		if nodes[i].user_medicalrecord == nil {
			continue
		}
		fk := *nodes[i].user_medicalrecord
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_medicalrecord" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mrq *MedicalRecordQuery) loadInstitution(ctx context.Context, query *InstitutionQuery, nodes []*MedicalRecord, init func(*MedicalRecord), assign func(*MedicalRecord, *Institution)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*MedicalRecord)
	for i := range nodes {
		if nodes[i].institution_medicalrecord == nil {
			continue
		}
		fk := *nodes[i].institution_medicalrecord
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
			return fmt.Errorf(`unexpected foreign-key "institution_medicalrecord" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mrq *MedicalRecordQuery) loadRecordaccess(ctx context.Context, query *RecordAccessQuery, nodes []*MedicalRecord, init func(*MedicalRecord), assign func(*MedicalRecord, *RecordAccess)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*MedicalRecord)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.RecordAccess(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(medicalrecord.RecordaccessColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.medical_record_recordaccess
		if fk == nil {
			return fmt.Errorf(`foreign-key "medical_record_recordaccess" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "medical_record_recordaccess" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mrq *MedicalRecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mrq.querySpec()
	_spec.Node.Columns = mrq.ctx.Fields
	if len(mrq.ctx.Fields) > 0 {
		_spec.Unique = mrq.ctx.Unique != nil && *mrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mrq.driver, _spec)
}

func (mrq *MedicalRecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(medicalrecord.Table, medicalrecord.Columns, sqlgraph.NewFieldSpec(medicalrecord.FieldID, field.TypeUUID))
	_spec.From = mrq.sql
	if unique := mrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mrq.path != nil {
		_spec.Unique = true
	}
	if fields := mrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, medicalrecord.FieldID)
		for i := range fields {
			if fields[i] != medicalrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mrq *MedicalRecordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mrq.driver.Dialect())
	t1 := builder.Table(medicalrecord.Table)
	columns := mrq.ctx.Fields
	if len(columns) == 0 {
		columns = medicalrecord.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mrq.sql != nil {
		selector = mrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mrq.ctx.Unique != nil && *mrq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mrq.predicates {
		p(selector)
	}
	for _, p := range mrq.order {
		p(selector)
	}
	if offset := mrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MedicalRecordGroupBy is the group-by builder for MedicalRecord entities.
type MedicalRecordGroupBy struct {
	selector
	build *MedicalRecordQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mrgb *MedicalRecordGroupBy) Aggregate(fns ...AggregateFunc) *MedicalRecordGroupBy {
	mrgb.fns = append(mrgb.fns, fns...)
	return mrgb
}

// Scan applies the selector query and scans the result into the given value.
func (mrgb *MedicalRecordGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mrgb.build.ctx, ent.OpQueryGroupBy)
	if err := mrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MedicalRecordQuery, *MedicalRecordGroupBy](ctx, mrgb.build, mrgb, mrgb.build.inters, v)
}

func (mrgb *MedicalRecordGroupBy) sqlScan(ctx context.Context, root *MedicalRecordQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mrgb.fns))
	for _, fn := range mrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mrgb.flds)+len(mrgb.fns))
		for _, f := range *mrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MedicalRecordSelect is the builder for selecting fields of MedicalRecord entities.
type MedicalRecordSelect struct {
	*MedicalRecordQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mrs *MedicalRecordSelect) Aggregate(fns ...AggregateFunc) *MedicalRecordSelect {
	mrs.fns = append(mrs.fns, fns...)
	return mrs
}

// Scan applies the selector query and scans the result into the given value.
func (mrs *MedicalRecordSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mrs.ctx, ent.OpQuerySelect)
	if err := mrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MedicalRecordQuery, *MedicalRecordSelect](ctx, mrs.MedicalRecordQuery, mrs, mrs.inters, v)
}

func (mrs *MedicalRecordSelect) sqlScan(ctx context.Context, root *MedicalRecordQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mrs.fns))
	for _, fn := range mrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/poll"
	"backend/ent/polloption"
	"backend/ent/predicate"
	"backend/ent/vote"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PollOptionUpdate is the builder for updating PollOption entities.
type PollOptionUpdate struct {
	config
	hooks    []Hook
	mutation *PollOptionMutation
}

// Where appends a list predicates to the PollOptionUpdate builder.
func (pou *PollOptionUpdate) Where(ps ...predicate.PollOption) *PollOptionUpdate {
	pou.mutation.Where(ps...)
	return pou
}

// SetTitle sets the "title" field.
func (pou *PollOptionUpdate) SetTitle(s string) *PollOptionUpdate {
	pou.mutation.SetTitle(s)
	return pou
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (pou *PollOptionUpdate) SetNillableTitle(s *string) *PollOptionUpdate {
	if s != nil {
		pou.SetTitle(*s)
	}
	return pou
}

// SetCreatedAt sets the "created_at" field.
func (pou *PollOptionUpdate) SetCreatedAt(t time.Time) *PollOptionUpdate {
	pou.mutation.SetCreatedAt(t)
	return pou
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pou *PollOptionUpdate) SetNillableCreatedAt(t *time.Time) *PollOptionUpdate {
	if t != nil {
		pou.SetCreatedAt(*t)
	}
	return pou
}

// SetUpdatedAt sets the "updated_at" field.
func (pou *PollOptionUpdate) SetUpdatedAt(t time.Time) *PollOptionUpdate {
	pou.mutation.SetUpdatedAt(t)
	return pou
}

// SetPollID sets the "poll" edge to the Poll entity by ID.
func (pou *PollOptionUpdate) SetPollID(id int) *PollOptionUpdate {
	pou.mutation.SetPollID(id)
	return pou
}

// SetNillablePollID sets the "poll" edge to the Poll entity by ID if the given value is not nil.
func (pou *PollOptionUpdate) SetNillablePollID(id *int) *PollOptionUpdate {
	if id != nil {
		pou = pou.SetPollID(*id)
	}
	return pou
}

// SetPoll sets the "poll" edge to the Poll entity.
func (pou *PollOptionUpdate) SetPoll(p *Poll) *PollOptionUpdate {
	return pou.SetPollID(p.ID)
}

// AddVoteIDs adds the "votes" edge to the Vote entity by IDs.
func (pou *PollOptionUpdate) AddVoteIDs(ids ...int) *PollOptionUpdate {
	pou.mutation.AddVoteIDs(ids...)
	return pou
}

// AddVotes adds the "votes" edges to the Vote entity.
func (pou *PollOptionUpdate) AddVotes(v ...*Vote) *PollOptionUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return pou.AddVoteIDs(ids...)
}

// Mutation returns the PollOptionMutation object of the builder.
func (pou *PollOptionUpdate) Mutation() *PollOptionMutation {
	return pou.mutation
}

// ClearPoll clears the "poll" edge to the Poll entity.
func (pou *PollOptionUpdate) ClearPoll() *PollOptionUpdate {
	pou.mutation.ClearPoll()
	return pou
}

// ClearVotes clears all "votes" edges to the Vote entity.
func (pou *PollOptionUpdate) ClearVotes() *PollOptionUpdate {
	pou.mutation.ClearVotes()
	return pou
}

// RemoveVoteIDs removes the "votes" edge to Vote entities by IDs.
func (pou *PollOptionUpdate) RemoveVoteIDs(ids ...int) *PollOptionUpdate {
	pou.mutation.RemoveVoteIDs(ids...)
	return pou
}

// RemoveVotes removes "votes" edges to Vote entities.
func (pou *PollOptionUpdate) RemoveVotes(v ...*Vote) *PollOptionUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return pou.RemoveVoteIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pou *PollOptionUpdate) Save(ctx context.Context) (int, error) {
	pou.defaults()
	return withHooks(ctx, pou.sqlSave, pou.mutation, pou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pou *PollOptionUpdate) SaveX(ctx context.Context) int {
	affected, err := pou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pou *PollOptionUpdate) Exec(ctx context.Context) error {
	_, err := pou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pou *PollOptionUpdate) ExecX(ctx context.Context) {
	if err := pou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pou *PollOptionUpdate) defaults() {
	if _, ok := pou.mutation.UpdatedAt(); !ok {
		v := polloption.UpdateDefaultUpdatedAt()
		pou.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pou *PollOptionUpdate) check() error {
	if v, ok := pou.mutation.Title(); ok {
		if err := polloption.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "PollOption.title": %w`, err)}
		}
	}
	return nil
}

func (pou *PollOptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(polloption.Table, polloption.Columns, sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt))
	if ps := pou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pou.mutation.Title(); ok {
		_spec.SetField(polloption.FieldTitle, field.TypeString, value)
	}
	if value, ok := pou.mutation.CreatedAt(); ok {
		_spec.SetField(polloption.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pou.mutation.UpdatedAt(); ok {
		_spec.SetField(polloption.FieldUpdatedAt, field.TypeTime, value)
	}
	if pou.mutation.PollCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   polloption.PollTable,
			Columns: []string{polloption.PollColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(poll.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pou.mutation.PollIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   polloption.PollTable,
			Columns: []string{polloption.PollColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(poll.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pou.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   polloption.VotesTable,
			Columns: []string{polloption.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pou.mutation.RemovedVotesIDs(); len(nodes) > 0 && !pou.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   polloption.VotesTable,
			Columns: []string{polloption.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pou.mutation.VotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   polloption.VotesTable,
			Columns: []string{polloption.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{polloption.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pou.mutation.done = true
	return n, nil
}

// PollOptionUpdateOne is the builder for updating a single PollOption entity.
type PollOptionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PollOptionMutation
}

// SetTitle sets the "title" field.
func (pouo *PollOptionUpdateOne) SetTitle(s string) *PollOptionUpdateOne {
	pouo.mutation.SetTitle(s)
	return pouo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (pouo *PollOptionUpdateOne) SetNillableTitle(s *string) *PollOptionUpdateOne {
	if s != nil {
		pouo.SetTitle(*s)
	}
	return pouo
}

// SetCreatedAt sets the "created_at" field.
func (pouo *PollOptionUpdateOne) SetCreatedAt(t time.Time) *PollOptionUpdateOne {
	pouo.mutation.SetCreatedAt(t)
	return pouo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pouo *PollOptionUpdateOne) SetNillableCreatedAt(t *time.Time) *PollOptionUpdateOne {
	if t != nil {
		pouo.SetCreatedAt(*t)
	}
	return pouo
}

// SetUpdatedAt sets the "updated_at" field.
func (pouo *PollOptionUpdateOne) SetUpdatedAt(t time.Time) *PollOptionUpdateOne {
	pouo.mutation.SetUpdatedAt(t)
	return pouo
}

// SetPollID sets the "poll" edge to the Poll entity by ID.
func (pouo *PollOptionUpdateOne) SetPollID(id int) *PollOptionUpdateOne {
	pouo.mutation.SetPollID(id)
	return pouo
}

// SetNillablePollID sets the "poll" edge to the Poll entity by ID if the given value is not nil.
func (pouo *PollOptionUpdateOne) SetNillablePollID(id *int) *PollOptionUpdateOne {
	if id != nil {
		pouo = pouo.SetPollID(*id)
	}
	return pouo
}

// SetPoll sets the "poll" edge to the Poll entity.
func (pouo *PollOptionUpdateOne) SetPoll(p *Poll) *PollOptionUpdateOne {
	return pouo.SetPollID(p.ID)
}

// AddVoteIDs adds the "votes" edge to the Vote entity by IDs.
func (pouo *PollOptionUpdateOne) AddVoteIDs(ids ...int) *PollOptionUpdateOne {
	pouo.mutation.AddVoteIDs(ids...)
	return pouo
}

// AddVotes adds the "votes" edges to the Vote entity.
func (pouo *PollOptionUpdateOne) AddVotes(v ...*Vote) *PollOptionUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return pouo.AddVoteIDs(ids...)
}

// Mutation returns the PollOptionMutation object of the builder.
func (pouo *PollOptionUpdateOne) Mutation() *PollOptionMutation {
	return pouo.mutation
}

// ClearPoll clears the "poll" edge to the Poll entity.
func (pouo *PollOptionUpdateOne) ClearPoll() *PollOptionUpdateOne {
	pouo.mutation.ClearPoll()
	return pouo
}

// ClearVotes clears all "votes" edges to the Vote entity.
func (pouo *PollOptionUpdateOne) ClearVotes() *PollOptionUpdateOne {
	pouo.mutation.ClearVotes()
	return pouo
}

// RemoveVoteIDs removes the "votes" edge to Vote entities by IDs.
func (pouo *PollOptionUpdateOne) RemoveVoteIDs(ids ...int) *PollOptionUpdateOne {
	pouo.mutation.RemoveVoteIDs(ids...)
	return pouo
}

// RemoveVotes removes "votes" edges to Vote entities.
func (pouo *PollOptionUpdateOne) RemoveVotes(v ...*Vote) *PollOptionUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return pouo.RemoveVoteIDs(ids...)
}

// Where appends a list predicates to the PollOptionUpdate builder.
func (pouo *PollOptionUpdateOne) Where(ps ...predicate.PollOption) *PollOptionUpdateOne {
	pouo.mutation.Where(ps...)
	return pouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pouo *PollOptionUpdateOne) Select(field string, fields ...string) *PollOptionUpdateOne {
	pouo.fields = append([]string{field}, fields...)
	return pouo
}

// Save executes the query and returns the updated PollOption entity.
func (pouo *PollOptionUpdateOne) Save(ctx context.Context) (*PollOption, error) {
	pouo.defaults()
	return withHooks(ctx, pouo.sqlSave, pouo.mutation, pouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pouo *PollOptionUpdateOne) SaveX(ctx context.Context) *PollOption {
	node, err := pouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pouo *PollOptionUpdateOne) Exec(ctx context.Context) error {
	_, err := pouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pouo *PollOptionUpdateOne) ExecX(ctx context.Context) {
	if err := pouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pouo *PollOptionUpdateOne) defaults() {
	if _, ok := pouo.mutation.UpdatedAt(); !ok {
		v := polloption.UpdateDefaultUpdatedAt()
		pouo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pouo *PollOptionUpdateOne) check() error {
	if v, ok := pouo.mutation.Title(); ok {
		if err := polloption.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "PollOption.title": %w`, err)}
		}
	}
	return nil
}

func (pouo *PollOptionUpdateOne) sqlSave(ctx context.Context) (_node *PollOption, err error) {
	if err := pouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(polloption.Table, polloption.Columns, sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt))
	id, ok := pouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PollOption.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, polloption.FieldID)
		for _, f := range fields {
			if !polloption.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != polloption.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pouo.mutation.Title(); ok {
		_spec.SetField(polloption.FieldTitle, field.TypeString, value)
	}
	if value, ok := pouo.mutation.CreatedAt(); ok {
		_spec.SetField(polloption.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pouo.mutation.UpdatedAt(); ok {
		_spec.SetField(polloption.FieldUpdatedAt, field.TypeTime, value)
	}
	if pouo.mutation.PollCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   polloption.PollTable,
			Columns: []string{polloption.PollColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(poll.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pouo.mutation.PollIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   polloption.PollTable,
			Columns: []string{polloption.PollColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(poll.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pouo.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   polloption.VotesTable,
			Columns: []string{polloption.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pouo.mutation.RemovedVotesIDs(); len(nodes) > 0 && !pouo.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   polloption.VotesTable,
			Columns: []string{polloption.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pouo.mutation.VotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   polloption.VotesTable,
			Columns: []string{polloption.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PollOption{config: pouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{polloption.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pouo.mutation.done = true
	return _node, nil
}
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/poll"
	"backend/ent/polloption"
	"backend/ent/predicate"
	"backend/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PollUpdate is the builder for updating Poll entities.
type PollUpdate struct {
	config
	hooks    []Hook
	mutation *PollMutation
}

// Where appends a list predicates to the PollUpdate builder.
func (pu *PollUpdate) Where(ps ...predicate.Poll) *PollUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetQuestion sets the "question" field.
func (pu *PollUpdate) SetQuestion(s string) *PollUpdate {
	pu.mutation.SetQuestion(s)
	return pu
}

// SetNillableQuestion sets the "question" field if the given value is not nil.
func (pu *PollUpdate) SetNillableQuestion(s *string) *PollUpdate {
	if s != nil {
		pu.SetQuestion(*s)
	}
	return pu
}

// SetStatus sets the "status" field.
func (pu *PollUpdate) SetStatus(s string) *PollUpdate {
	pu.mutation.SetStatus(s)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PollUpdate) SetNillableStatus(s *string) *PollUpdate {
	if s != nil {
		pu.SetStatus(*s)
	}
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PollUpdate) SetCreatedAt(t time.Time) *PollUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PollUpdate) SetNillableCreatedAt(t *time.Time) *PollUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PollUpdate) SetUpdatedAt(t time.Time) *PollUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pu *PollUpdate) SetUserID(id int) *PollUpdate {
	pu.mutation.SetUserID(id)
	return pu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (pu *PollUpdate) SetNillableUserID(id *int) *PollUpdate {
	if id != nil {
		pu = pu.SetUserID(*id)
	}
	return pu
}

// SetUser sets the "user" edge to the User entity.
func (pu *PollUpdate) SetUser(u *User) *PollUpdate {
	return pu.SetUserID(u.ID)
}

// AddPollOptionIDs adds the "poll_options" edge to the PollOption entity by IDs.
func (pu *PollUpdate) AddPollOptionIDs(ids ...int) *PollUpdate {
	pu.mutation.AddPollOptionIDs(ids...)
	return pu
}

// AddPollOptions adds the "poll_options" edges to the PollOption entity.
func (pu *PollUpdate) AddPollOptions(p ...*PollOption) *PollUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddPollOptionIDs(ids...)
}

// Mutation returns the PollMutation object of the builder.
func (pu *PollUpdate) Mutation() *PollMutation {
	return pu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pu *PollUpdate) ClearUser() *PollUpdate {
	pu.mutation.ClearUser()
	return pu
}

// ClearPollOptions clears all "poll_options" edges to the PollOption entity.
func (pu *PollUpdate) ClearPollOptions() *PollUpdate {
	pu.mutation.ClearPollOptions()
	return pu
}

// RemovePollOptionIDs removes the "poll_options" edge to PollOption entities by IDs.
func (pu *PollUpdate) RemovePollOptionIDs(ids ...int) *PollUpdate {
	pu.mutation.RemovePollOptionIDs(ids...)
	return pu
}

// RemovePollOptions removes "poll_options" edges to PollOption entities.
func (pu *PollUpdate) RemovePollOptions(p ...*PollOption) *PollUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemovePollOptionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PollUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PollUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PollUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PollUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PollUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := poll.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PollUpdate) check() error {
	if v, ok := pu.mutation.Question(); ok {
		if err := poll.QuestionValidator(v); err != nil {
			return &ValidationError{Name: "question", err: fmt.Errorf(`ent: validator failed for field "Poll.question": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Status(); ok {
		if err := poll.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Poll.status": %w`, err)}
		}
	}
	return nil
}

func (pu *PollUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(poll.Table, poll.Columns, sqlgraph.NewFieldSpec(poll.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Question(); ok {
		_spec.SetField(poll.FieldQuestion, field.TypeString, value)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(poll.FieldStatus, field.TypeString, value)
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(poll.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(poll.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   poll.UserTable,
			Columns: []string{poll.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   poll.UserTable,
			Columns: []string{poll.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.PollOptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poll.PollOptionsTable,
			Columns: []string{poll.PollOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedPollOptionsIDs(); len(nodes) > 0 && !pu.mutation.PollOptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poll.PollOptionsTable,
			Columns: []string{poll.PollOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PollOptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poll.PollOptionsTable,
			Columns: []string{poll.PollOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{poll.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PollUpdateOne is the builder for updating a single Poll entity.
type PollUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PollMutation
}

// SetQuestion sets the "question" field.
func (puo *PollUpdateOne) SetQuestion(s string) *PollUpdateOne {
	puo.mutation.SetQuestion(s)
	return puo
}

// SetNillableQuestion sets the "question" field if the given value is not nil.
func (puo *PollUpdateOne) SetNillableQuestion(s *string) *PollUpdateOne {
	if s != nil {
		puo.SetQuestion(*s)
	}
	return puo
}

// SetStatus sets the "status" field.
func (puo *PollUpdateOne) SetStatus(s string) *PollUpdateOne {
	puo.mutation.SetStatus(s)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PollUpdateOne) SetNillableStatus(s *string) *PollUpdateOne {
	if s != nil {
		puo.SetStatus(*s)
	}
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *PollUpdateOne) SetCreatedAt(t time.Time) *PollUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PollUpdateOne) SetNillableCreatedAt(t *time.Time) *PollUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PollUpdateOne) SetUpdatedAt(t time.Time) *PollUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (puo *PollUpdateOne) SetUserID(id int) *PollUpdateOne {
	puo.mutation.SetUserID(id)
	return puo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (puo *PollUpdateOne) SetNillableUserID(id *int) *PollUpdateOne {
	if id != nil {
		puo = puo.SetUserID(*id)
	}
	return puo
}

// SetUser sets the "user" edge to the User entity.
func (puo *PollUpdateOne) SetUser(u *User) *PollUpdateOne {
	return puo.SetUserID(u.ID)
}

// AddPollOptionIDs adds the "poll_options" edge to the PollOption entity by IDs.
func (puo *PollUpdateOne) AddPollOptionIDs(ids ...int) *PollUpdateOne {
	puo.mutation.AddPollOptionIDs(ids...)
	return puo
}

// AddPollOptions adds the "poll_options" edges to the PollOption entity.
func (puo *PollUpdateOne) AddPollOptions(p ...*PollOption) *PollUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddPollOptionIDs(ids...)
}

// Mutation returns the PollMutation object of the builder.
func (puo *PollUpdateOne) Mutation() *PollMutation {
	return puo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (puo *PollUpdateOne) ClearUser() *PollUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// ClearPollOptions clears all "poll_options" edges to the PollOption entity.
func (puo *PollUpdateOne) ClearPollOptions() *PollUpdateOne {
	puo.mutation.ClearPollOptions()
	return puo
}

// RemovePollOptionIDs removes the "poll_options" edge to PollOption entities by IDs.
func (puo *PollUpdateOne) RemovePollOptionIDs(ids ...int) *PollUpdateOne {
	puo.mutation.RemovePollOptionIDs(ids...)
	return puo
}

// RemovePollOptions removes "poll_options" edges to PollOption entities.
func (puo *PollUpdateOne) RemovePollOptions(p ...*PollOption) *PollUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemovePollOptionIDs(ids...)
}

// Where appends a list predicates to the PollUpdate builder.
func (puo *PollUpdateOne) Where(ps ...predicate.Poll) *PollUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PollUpdateOne) Select(field string, fields ...string) *PollUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Poll entity.
func (puo *PollUpdateOne) Save(ctx context.Context) (*Poll, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PollUpdateOne) SaveX(ctx context.Context) *Poll {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PollUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PollUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PollUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := poll.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PollUpdateOne) check() error {
	if v, ok := puo.mutation.Question(); ok {
		if err := poll.QuestionValidator(v); err != nil {
			return &ValidationError{Name: "question", err: fmt.Errorf(`ent: validator failed for field "Poll.question": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Status(); ok {
		if err := poll.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Poll.status": %w`, err)}
		}
	}
	return nil
}

func (puo *PollUpdateOne) sqlSave(ctx context.Context) (_node *Poll, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(poll.Table, poll.Columns, sqlgraph.NewFieldSpec(poll.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Poll.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, poll.FieldID)
		for _, f := range fields {
			if !poll.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != poll.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Question(); ok {
		_spec.SetField(poll.FieldQuestion, field.TypeString, value)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(poll.FieldStatus, field.TypeString, value)
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(poll.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(poll.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   poll.UserTable,
			Columns: []string{poll.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   poll.UserTable,
			Columns: []string{poll.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.PollOptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poll.PollOptionsTable,
			Columns: []string{poll.PollOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedPollOptionsIDs(); len(nodes) > 0 && !puo.mutation.PollOptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poll.PollOptionsTable,
			Columns: []string{poll.PollOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PollOptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poll.PollOptionsTable,
			Columns: []string{poll.PollOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Poll{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{poll.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
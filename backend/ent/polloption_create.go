// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/poll"
	"backend/ent/polloption"
	"backend/ent/vote"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PollOptionCreate is the builder for creating a PollOption entity.
type PollOptionCreate struct {
	config
	mutation *PollOptionMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (poc *PollOptionCreate) SetTitle(s string) *PollOptionCreate {
	poc.mutation.SetTitle(s)
	return poc
}

// SetCreatedAt sets the "created_at" field.
func (poc *PollOptionCreate) SetCreatedAt(t time.Time) *PollOptionCreate {
	poc.mutation.SetCreatedAt(t)
	return poc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (poc *PollOptionCreate) SetNillableCreatedAt(t *time.Time) *PollOptionCreate {
	if t != nil {
		poc.SetCreatedAt(*t)
	}
	return poc
}

// SetUpdatedAt sets the "updated_at" field.
func (poc *PollOptionCreate) SetUpdatedAt(t time.Time) *PollOptionCreate {
	poc.mutation.SetUpdatedAt(t)
	return poc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (poc *PollOptionCreate) SetNillableUpdatedAt(t *time.Time) *PollOptionCreate {
	if t != nil {
		poc.SetUpdatedAt(*t)
	}
	return poc
}

// SetPollID sets the "poll" edge to the Poll entity by ID.
func (poc *PollOptionCreate) SetPollID(id int) *PollOptionCreate {
	poc.mutation.SetPollID(id)
	return poc
}

// SetNillablePollID sets the "poll" edge to the Poll entity by ID if the given value is not nil.
func (poc *PollOptionCreate) SetNillablePollID(id *int) *PollOptionCreate {
	if id != nil {
		poc = poc.SetPollID(*id)
	}
	return poc
}

// SetPoll sets the "poll" edge to the Poll entity.
func (poc *PollOptionCreate) SetPoll(p *Poll) *PollOptionCreate {
	return poc.SetPollID(p.ID)
}

// AddVoteIDs adds the "votes" edge to the Vote entity by IDs.
func (poc *PollOptionCreate) AddVoteIDs(ids ...int) *PollOptionCreate {
	poc.mutation.AddVoteIDs(ids...)
	return poc
}

// AddVotes adds the "votes" edges to the Vote entity.
func (poc *PollOptionCreate) AddVotes(v ...*Vote) *PollOptionCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return poc.AddVoteIDs(ids...)
}

// Mutation returns the PollOptionMutation object of the builder.
func (poc *PollOptionCreate) Mutation() *PollOptionMutation {
	return poc.mutation
}

// Save creates the PollOption in the database.
func (poc *PollOptionCreate) Save(ctx context.Context) (*PollOption, error) {
	poc.defaults()
	return withHooks(ctx, poc.sqlSave, poc.mutation, poc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (poc *PollOptionCreate) SaveX(ctx context.Context) *PollOption {
	v, err := poc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (poc *PollOptionCreate) Exec(ctx context.Context) error {
	_, err := poc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (poc *PollOptionCreate) ExecX(ctx context.Context) {
	if err := poc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (poc *PollOptionCreate) defaults() {
	if _, ok := poc.mutation.CreatedAt(); !ok {
		v := polloption.DefaultCreatedAt()
		poc.mutation.SetCreatedAt(v)
	}
	if _, ok := poc.mutation.UpdatedAt(); !ok {
		v := polloption.DefaultUpdatedAt()
		poc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (poc *PollOptionCreate) check() error {
	if _, ok := poc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "PollOption.title"`)}
	}
	if v, ok := poc.mutation.Title(); ok {
		if err := polloption.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "PollOption.title": %w`, err)}
		}
	}
	if _, ok := poc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "PollOption.created_at"`)}
	}
	if _, ok := poc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "PollOption.updated_at"`)}
	}
	return nil
}

func (poc *PollOptionCreate) sqlSave(ctx context.Context) (*PollOption, error) {
	if err := poc.check(); err != nil {
		return nil, err
	}
	_node, _spec := poc.createSpec()
	if err := sqlgraph.CreateNode(ctx, poc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	poc.mutation.id = &_node.ID
	poc.mutation.done = true
	return _node, nil
}

func (poc *PollOptionCreate) createSpec() (*PollOption, *sqlgraph.CreateSpec) {
	var (
		_node = &PollOption{config: poc.config}
		_spec = sqlgraph.NewCreateSpec(polloption.Table, sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt))
	)
	if value, ok := poc.mutation.Title(); ok {
		_spec.SetField(polloption.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := poc.mutation.CreatedAt(); ok {
		_spec.SetField(polloption.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := poc.mutation.UpdatedAt(); ok {
		_spec.SetField(polloption.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := poc.mutation.PollIDs(); len(nodes) > 0 {
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
		_node.poll_poll_options = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := poc.mutation.VotesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PollOptionCreateBulk is the builder for creating many PollOption entities in bulk.
type PollOptionCreateBulk struct {
	config
	err      error
	builders []*PollOptionCreate
}

// Save creates the PollOption entities in the database.
func (pocb *PollOptionCreateBulk) Save(ctx context.Context) ([]*PollOption, error) {
	if pocb.err != nil {
		return nil, pocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pocb.builders))
	nodes := make([]*PollOption, len(pocb.builders))
	mutators := make([]Mutator, len(pocb.builders))
	for i := range pocb.builders {
		func(i int, root context.Context) {
			builder := pocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PollOptionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pocb *PollOptionCreateBulk) SaveX(ctx context.Context) []*PollOption {
	v, err := pocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pocb *PollOptionCreateBulk) Exec(ctx context.Context) error {
	_, err := pocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pocb *PollOptionCreateBulk) ExecX(ctx context.Context) {
	if err := pocb.Exec(ctx); err != nil {
		panic(err)
	}
}

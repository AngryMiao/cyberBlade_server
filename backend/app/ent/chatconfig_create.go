// Code generated by ent, DO NOT EDIT.

package ent

import (
	"angrymiao-ai/app/ent/chatconfig"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChatConfigCreate is the builder for creating a ChatConfig entity.
type ChatConfigCreate struct {
	config
	mutation *ChatConfigMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (ccc *ChatConfigCreate) SetCreateTime(t time.Time) *ChatConfigCreate {
	ccc.mutation.SetCreateTime(t)
	return ccc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ccc *ChatConfigCreate) SetNillableCreateTime(t *time.Time) *ChatConfigCreate {
	if t != nil {
		ccc.SetCreateTime(*t)
	}
	return ccc
}

// SetUpdateTime sets the "update_time" field.
func (ccc *ChatConfigCreate) SetUpdateTime(t time.Time) *ChatConfigCreate {
	ccc.mutation.SetUpdateTime(t)
	return ccc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ccc *ChatConfigCreate) SetNillableUpdateTime(t *time.Time) *ChatConfigCreate {
	if t != nil {
		ccc.SetUpdateTime(*t)
	}
	return ccc
}

// SetDeviceCode sets the "device_code" field.
func (ccc *ChatConfigCreate) SetDeviceCode(s string) *ChatConfigCreate {
	ccc.mutation.SetDeviceCode(s)
	return ccc
}

// SetDiscordUserID sets the "discord_user_id" field.
func (ccc *ChatConfigCreate) SetDiscordUserID(s string) *ChatConfigCreate {
	ccc.mutation.SetDiscordUserID(s)
	return ccc
}

// SetNillableDiscordUserID sets the "discord_user_id" field if the given value is not nil.
func (ccc *ChatConfigCreate) SetNillableDiscordUserID(s *string) *ChatConfigCreate {
	if s != nil {
		ccc.SetDiscordUserID(*s)
	}
	return ccc
}

// SetForwardMode sets the "forward_mode" field.
func (ccc *ChatConfigCreate) SetForwardMode(cm chatconfig.ForwardMode) *ChatConfigCreate {
	ccc.mutation.SetForwardMode(cm)
	return ccc
}

// SetNillableForwardMode sets the "forward_mode" field if the given value is not nil.
func (ccc *ChatConfigCreate) SetNillableForwardMode(cm *chatconfig.ForwardMode) *ChatConfigCreate {
	if cm != nil {
		ccc.SetForwardMode(*cm)
	}
	return ccc
}

// Mutation returns the ChatConfigMutation object of the builder.
func (ccc *ChatConfigCreate) Mutation() *ChatConfigMutation {
	return ccc.mutation
}

// Save creates the ChatConfig in the database.
func (ccc *ChatConfigCreate) Save(ctx context.Context) (*ChatConfig, error) {
	var (
		err  error
		node *ChatConfig
	)
	ccc.defaults()
	if len(ccc.hooks) == 0 {
		if err = ccc.check(); err != nil {
			return nil, err
		}
		node, err = ccc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChatConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ccc.check(); err != nil {
				return nil, err
			}
			ccc.mutation = mutation
			if node, err = ccc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ccc.hooks) - 1; i >= 0; i-- {
			if ccc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ccc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ccc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ChatConfig)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ChatConfigMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ccc *ChatConfigCreate) SaveX(ctx context.Context) *ChatConfig {
	v, err := ccc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccc *ChatConfigCreate) Exec(ctx context.Context) error {
	_, err := ccc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccc *ChatConfigCreate) ExecX(ctx context.Context) {
	if err := ccc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccc *ChatConfigCreate) defaults() {
	if _, ok := ccc.mutation.CreateTime(); !ok {
		v := chatconfig.DefaultCreateTime()
		ccc.mutation.SetCreateTime(v)
	}
	if _, ok := ccc.mutation.UpdateTime(); !ok {
		v := chatconfig.DefaultUpdateTime()
		ccc.mutation.SetUpdateTime(v)
	}
	if _, ok := ccc.mutation.ForwardMode(); !ok {
		v := chatconfig.DefaultForwardMode
		ccc.mutation.SetForwardMode(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccc *ChatConfigCreate) check() error {
	if _, ok := ccc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "ChatConfig.create_time"`)}
	}
	if _, ok := ccc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "ChatConfig.update_time"`)}
	}
	if _, ok := ccc.mutation.DeviceCode(); !ok {
		return &ValidationError{Name: "device_code", err: errors.New(`ent: missing required field "ChatConfig.device_code"`)}
	}
	if v, ok := ccc.mutation.DeviceCode(); ok {
		if err := chatconfig.DeviceCodeValidator(v); err != nil {
			return &ValidationError{Name: "device_code", err: fmt.Errorf(`ent: validator failed for field "ChatConfig.device_code": %w`, err)}
		}
	}
	if v, ok := ccc.mutation.ForwardMode(); ok {
		if err := chatconfig.ForwardModeValidator(v); err != nil {
			return &ValidationError{Name: "forward_mode", err: fmt.Errorf(`ent: validator failed for field "ChatConfig.forward_mode": %w`, err)}
		}
	}
	return nil
}

func (ccc *ChatConfigCreate) sqlSave(ctx context.Context) (*ChatConfig, error) {
	_node, _spec := ccc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ccc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ccc *ChatConfigCreate) createSpec() (*ChatConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &ChatConfig{config: ccc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: chatconfig.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: chatconfig.FieldID,
			},
		}
	)
	_spec.OnConflict = ccc.conflict
	if value, ok := ccc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatconfig.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ccc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatconfig.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ccc.mutation.DeviceCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chatconfig.FieldDeviceCode,
		})
		_node.DeviceCode = value
	}
	if value, ok := ccc.mutation.DiscordUserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chatconfig.FieldDiscordUserID,
		})
		_node.DiscordUserID = &value
	}
	if value, ok := ccc.mutation.ForwardMode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: chatconfig.FieldForwardMode,
		})
		_node.ForwardMode = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ChatConfig.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ChatConfigUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (ccc *ChatConfigCreate) OnConflict(opts ...sql.ConflictOption) *ChatConfigUpsertOne {
	ccc.conflict = opts
	return &ChatConfigUpsertOne{
		create: ccc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ChatConfig.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccc *ChatConfigCreate) OnConflictColumns(columns ...string) *ChatConfigUpsertOne {
	ccc.conflict = append(ccc.conflict, sql.ConflictColumns(columns...))
	return &ChatConfigUpsertOne{
		create: ccc,
	}
}

type (
	// ChatConfigUpsertOne is the builder for "upsert"-ing
	//  one ChatConfig node.
	ChatConfigUpsertOne struct {
		create *ChatConfigCreate
	}

	// ChatConfigUpsert is the "OnConflict" setter.
	ChatConfigUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreateTime sets the "create_time" field.
func (u *ChatConfigUpsert) SetCreateTime(v time.Time) *ChatConfigUpsert {
	u.Set(chatconfig.FieldCreateTime, v)
	return u
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *ChatConfigUpsert) UpdateCreateTime() *ChatConfigUpsert {
	u.SetExcluded(chatconfig.FieldCreateTime)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *ChatConfigUpsert) SetUpdateTime(v time.Time) *ChatConfigUpsert {
	u.Set(chatconfig.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ChatConfigUpsert) UpdateUpdateTime() *ChatConfigUpsert {
	u.SetExcluded(chatconfig.FieldUpdateTime)
	return u
}

// SetDeviceCode sets the "device_code" field.
func (u *ChatConfigUpsert) SetDeviceCode(v string) *ChatConfigUpsert {
	u.Set(chatconfig.FieldDeviceCode, v)
	return u
}

// UpdateDeviceCode sets the "device_code" field to the value that was provided on create.
func (u *ChatConfigUpsert) UpdateDeviceCode() *ChatConfigUpsert {
	u.SetExcluded(chatconfig.FieldDeviceCode)
	return u
}

// SetDiscordUserID sets the "discord_user_id" field.
func (u *ChatConfigUpsert) SetDiscordUserID(v string) *ChatConfigUpsert {
	u.Set(chatconfig.FieldDiscordUserID, v)
	return u
}

// UpdateDiscordUserID sets the "discord_user_id" field to the value that was provided on create.
func (u *ChatConfigUpsert) UpdateDiscordUserID() *ChatConfigUpsert {
	u.SetExcluded(chatconfig.FieldDiscordUserID)
	return u
}

// ClearDiscordUserID clears the value of the "discord_user_id" field.
func (u *ChatConfigUpsert) ClearDiscordUserID() *ChatConfigUpsert {
	u.SetNull(chatconfig.FieldDiscordUserID)
	return u
}

// SetForwardMode sets the "forward_mode" field.
func (u *ChatConfigUpsert) SetForwardMode(v chatconfig.ForwardMode) *ChatConfigUpsert {
	u.Set(chatconfig.FieldForwardMode, v)
	return u
}

// UpdateForwardMode sets the "forward_mode" field to the value that was provided on create.
func (u *ChatConfigUpsert) UpdateForwardMode() *ChatConfigUpsert {
	u.SetExcluded(chatconfig.FieldForwardMode)
	return u
}

// ClearForwardMode clears the value of the "forward_mode" field.
func (u *ChatConfigUpsert) ClearForwardMode() *ChatConfigUpsert {
	u.SetNull(chatconfig.FieldForwardMode)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.ChatConfig.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ChatConfigUpsertOne) UpdateNewValues() *ChatConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(chatconfig.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ChatConfig.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ChatConfigUpsertOne) Ignore() *ChatConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ChatConfigUpsertOne) DoNothing() *ChatConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ChatConfigCreate.OnConflict
// documentation for more info.
func (u *ChatConfigUpsertOne) Update(set func(*ChatConfigUpsert)) *ChatConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ChatConfigUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *ChatConfigUpsertOne) SetCreateTime(v time.Time) *ChatConfigUpsertOne {
	return u.Update(func(s *ChatConfigUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *ChatConfigUpsertOne) UpdateCreateTime() *ChatConfigUpsertOne {
	return u.Update(func(s *ChatConfigUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *ChatConfigUpsertOne) SetUpdateTime(v time.Time) *ChatConfigUpsertOne {
	return u.Update(func(s *ChatConfigUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ChatConfigUpsertOne) UpdateUpdateTime() *ChatConfigUpsertOne {
	return u.Update(func(s *ChatConfigUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDeviceCode sets the "device_code" field.
func (u *ChatConfigUpsertOne) SetDeviceCode(v string) *ChatConfigUpsertOne {
	return u.Update(func(s *ChatConfigUpsert) {
		s.SetDeviceCode(v)
	})
}

// UpdateDeviceCode sets the "device_code" field to the value that was provided on create.
func (u *ChatConfigUpsertOne) UpdateDeviceCode() *ChatConfigUpsertOne {
	return u.Update(func(s *ChatConfigUpsert) {
		s.UpdateDeviceCode()
	})
}

// SetDiscordUserID sets the "discord_user_id" field.
func (u *ChatConfigUpsertOne) SetDiscordUserID(v string) *ChatConfigUpsertOne {
	return u.Update(func(s *ChatConfigUpsert) {
		s.SetDiscordUserID(v)
	})
}

// UpdateDiscordUserID sets the "discord_user_id" field to the value that was provided on create.
func (u *ChatConfigUpsertOne) UpdateDiscordUserID() *ChatConfigUpsertOne {
	return u.Update(func(s *ChatConfigUpsert) {
		s.UpdateDiscordUserID()
	})
}

// ClearDiscordUserID clears the value of the "discord_user_id" field.
func (u *ChatConfigUpsertOne) ClearDiscordUserID() *ChatConfigUpsertOne {
	return u.Update(func(s *ChatConfigUpsert) {
		s.ClearDiscordUserID()
	})
}

// SetForwardMode sets the "forward_mode" field.
func (u *ChatConfigUpsertOne) SetForwardMode(v chatconfig.ForwardMode) *ChatConfigUpsertOne {
	return u.Update(func(s *ChatConfigUpsert) {
		s.SetForwardMode(v)
	})
}

// UpdateForwardMode sets the "forward_mode" field to the value that was provided on create.
func (u *ChatConfigUpsertOne) UpdateForwardMode() *ChatConfigUpsertOne {
	return u.Update(func(s *ChatConfigUpsert) {
		s.UpdateForwardMode()
	})
}

// ClearForwardMode clears the value of the "forward_mode" field.
func (u *ChatConfigUpsertOne) ClearForwardMode() *ChatConfigUpsertOne {
	return u.Update(func(s *ChatConfigUpsert) {
		s.ClearForwardMode()
	})
}

// Exec executes the query.
func (u *ChatConfigUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ChatConfigCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ChatConfigUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ChatConfigUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ChatConfigUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ChatConfigCreateBulk is the builder for creating many ChatConfig entities in bulk.
type ChatConfigCreateBulk struct {
	config
	builders []*ChatConfigCreate
	conflict []sql.ConflictOption
}

// Save creates the ChatConfig entities in the database.
func (cccb *ChatConfigCreateBulk) Save(ctx context.Context) ([]*ChatConfig, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cccb.builders))
	nodes := make([]*ChatConfig, len(cccb.builders))
	mutators := make([]Mutator, len(cccb.builders))
	for i := range cccb.builders {
		func(i int, root context.Context) {
			builder := cccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChatConfigMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cccb *ChatConfigCreateBulk) SaveX(ctx context.Context) []*ChatConfig {
	v, err := cccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cccb *ChatConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := cccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cccb *ChatConfigCreateBulk) ExecX(ctx context.Context) {
	if err := cccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ChatConfig.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ChatConfigUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (cccb *ChatConfigCreateBulk) OnConflict(opts ...sql.ConflictOption) *ChatConfigUpsertBulk {
	cccb.conflict = opts
	return &ChatConfigUpsertBulk{
		create: cccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ChatConfig.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cccb *ChatConfigCreateBulk) OnConflictColumns(columns ...string) *ChatConfigUpsertBulk {
	cccb.conflict = append(cccb.conflict, sql.ConflictColumns(columns...))
	return &ChatConfigUpsertBulk{
		create: cccb,
	}
}

// ChatConfigUpsertBulk is the builder for "upsert"-ing
// a bulk of ChatConfig nodes.
type ChatConfigUpsertBulk struct {
	create *ChatConfigCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ChatConfig.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ChatConfigUpsertBulk) UpdateNewValues() *ChatConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(chatconfig.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ChatConfig.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ChatConfigUpsertBulk) Ignore() *ChatConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ChatConfigUpsertBulk) DoNothing() *ChatConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ChatConfigCreateBulk.OnConflict
// documentation for more info.
func (u *ChatConfigUpsertBulk) Update(set func(*ChatConfigUpsert)) *ChatConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ChatConfigUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *ChatConfigUpsertBulk) SetCreateTime(v time.Time) *ChatConfigUpsertBulk {
	return u.Update(func(s *ChatConfigUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *ChatConfigUpsertBulk) UpdateCreateTime() *ChatConfigUpsertBulk {
	return u.Update(func(s *ChatConfigUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *ChatConfigUpsertBulk) SetUpdateTime(v time.Time) *ChatConfigUpsertBulk {
	return u.Update(func(s *ChatConfigUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ChatConfigUpsertBulk) UpdateUpdateTime() *ChatConfigUpsertBulk {
	return u.Update(func(s *ChatConfigUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDeviceCode sets the "device_code" field.
func (u *ChatConfigUpsertBulk) SetDeviceCode(v string) *ChatConfigUpsertBulk {
	return u.Update(func(s *ChatConfigUpsert) {
		s.SetDeviceCode(v)
	})
}

// UpdateDeviceCode sets the "device_code" field to the value that was provided on create.
func (u *ChatConfigUpsertBulk) UpdateDeviceCode() *ChatConfigUpsertBulk {
	return u.Update(func(s *ChatConfigUpsert) {
		s.UpdateDeviceCode()
	})
}

// SetDiscordUserID sets the "discord_user_id" field.
func (u *ChatConfigUpsertBulk) SetDiscordUserID(v string) *ChatConfigUpsertBulk {
	return u.Update(func(s *ChatConfigUpsert) {
		s.SetDiscordUserID(v)
	})
}

// UpdateDiscordUserID sets the "discord_user_id" field to the value that was provided on create.
func (u *ChatConfigUpsertBulk) UpdateDiscordUserID() *ChatConfigUpsertBulk {
	return u.Update(func(s *ChatConfigUpsert) {
		s.UpdateDiscordUserID()
	})
}

// ClearDiscordUserID clears the value of the "discord_user_id" field.
func (u *ChatConfigUpsertBulk) ClearDiscordUserID() *ChatConfigUpsertBulk {
	return u.Update(func(s *ChatConfigUpsert) {
		s.ClearDiscordUserID()
	})
}

// SetForwardMode sets the "forward_mode" field.
func (u *ChatConfigUpsertBulk) SetForwardMode(v chatconfig.ForwardMode) *ChatConfigUpsertBulk {
	return u.Update(func(s *ChatConfigUpsert) {
		s.SetForwardMode(v)
	})
}

// UpdateForwardMode sets the "forward_mode" field to the value that was provided on create.
func (u *ChatConfigUpsertBulk) UpdateForwardMode() *ChatConfigUpsertBulk {
	return u.Update(func(s *ChatConfigUpsert) {
		s.UpdateForwardMode()
	})
}

// ClearForwardMode clears the value of the "forward_mode" field.
func (u *ChatConfigUpsertBulk) ClearForwardMode() *ChatConfigUpsertBulk {
	return u.Update(func(s *ChatConfigUpsert) {
		s.ClearForwardMode()
	})
}

// Exec executes the query.
func (u *ChatConfigUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ChatConfigCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ChatConfigCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ChatConfigUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

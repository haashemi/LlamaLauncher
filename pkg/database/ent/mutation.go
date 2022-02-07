// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"llamalauncher/pkg/database/ent/predicate"
	"llamalauncher/pkg/database/ent/user"
	"sync"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser = "User"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op                      Op
	typ                     string
	id                      *int
	account_id              *string
	display_name            *string
	access_token            *string
	access_token_expires_at *time.Time
	refresh_token           *string
	refres_token_expires_at *time.Time
	clearedFields           map[string]struct{}
	done                    bool
	oldValue                func(context.Context) (*User, error)
	predicates              []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetAccountID sets the "account_id" field.
func (m *UserMutation) SetAccountID(s string) {
	m.account_id = &s
}

// AccountID returns the value of the "account_id" field in the mutation.
func (m *UserMutation) AccountID() (r string, exists bool) {
	v := m.account_id
	if v == nil {
		return
	}
	return *v, true
}

// OldAccountID returns the old "account_id" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAccountID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAccountID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAccountID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAccountID: %w", err)
	}
	return oldValue.AccountID, nil
}

// ResetAccountID resets all changes to the "account_id" field.
func (m *UserMutation) ResetAccountID() {
	m.account_id = nil
}

// SetDisplayName sets the "display_name" field.
func (m *UserMutation) SetDisplayName(s string) {
	m.display_name = &s
}

// DisplayName returns the value of the "display_name" field in the mutation.
func (m *UserMutation) DisplayName() (r string, exists bool) {
	v := m.display_name
	if v == nil {
		return
	}
	return *v, true
}

// OldDisplayName returns the old "display_name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldDisplayName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDisplayName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDisplayName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDisplayName: %w", err)
	}
	return oldValue.DisplayName, nil
}

// ResetDisplayName resets all changes to the "display_name" field.
func (m *UserMutation) ResetDisplayName() {
	m.display_name = nil
}

// SetAccessToken sets the "access_token" field.
func (m *UserMutation) SetAccessToken(s string) {
	m.access_token = &s
}

// AccessToken returns the value of the "access_token" field in the mutation.
func (m *UserMutation) AccessToken() (r string, exists bool) {
	v := m.access_token
	if v == nil {
		return
	}
	return *v, true
}

// OldAccessToken returns the old "access_token" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAccessToken(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAccessToken is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAccessToken requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAccessToken: %w", err)
	}
	return oldValue.AccessToken, nil
}

// ResetAccessToken resets all changes to the "access_token" field.
func (m *UserMutation) ResetAccessToken() {
	m.access_token = nil
}

// SetAccessTokenExpiresAt sets the "access_token_expires_at" field.
func (m *UserMutation) SetAccessTokenExpiresAt(t time.Time) {
	m.access_token_expires_at = &t
}

// AccessTokenExpiresAt returns the value of the "access_token_expires_at" field in the mutation.
func (m *UserMutation) AccessTokenExpiresAt() (r time.Time, exists bool) {
	v := m.access_token_expires_at
	if v == nil {
		return
	}
	return *v, true
}

// OldAccessTokenExpiresAt returns the old "access_token_expires_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAccessTokenExpiresAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAccessTokenExpiresAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAccessTokenExpiresAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAccessTokenExpiresAt: %w", err)
	}
	return oldValue.AccessTokenExpiresAt, nil
}

// ResetAccessTokenExpiresAt resets all changes to the "access_token_expires_at" field.
func (m *UserMutation) ResetAccessTokenExpiresAt() {
	m.access_token_expires_at = nil
}

// SetRefreshToken sets the "refresh_token" field.
func (m *UserMutation) SetRefreshToken(s string) {
	m.refresh_token = &s
}

// RefreshToken returns the value of the "refresh_token" field in the mutation.
func (m *UserMutation) RefreshToken() (r string, exists bool) {
	v := m.refresh_token
	if v == nil {
		return
	}
	return *v, true
}

// OldRefreshToken returns the old "refresh_token" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldRefreshToken(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRefreshToken is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRefreshToken requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRefreshToken: %w", err)
	}
	return oldValue.RefreshToken, nil
}

// ResetRefreshToken resets all changes to the "refresh_token" field.
func (m *UserMutation) ResetRefreshToken() {
	m.refresh_token = nil
}

// SetRefresTokenExpiresAt sets the "refres_token_expires_at" field.
func (m *UserMutation) SetRefresTokenExpiresAt(t time.Time) {
	m.refres_token_expires_at = &t
}

// RefresTokenExpiresAt returns the value of the "refres_token_expires_at" field in the mutation.
func (m *UserMutation) RefresTokenExpiresAt() (r time.Time, exists bool) {
	v := m.refres_token_expires_at
	if v == nil {
		return
	}
	return *v, true
}

// OldRefresTokenExpiresAt returns the old "refres_token_expires_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldRefresTokenExpiresAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRefresTokenExpiresAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRefresTokenExpiresAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRefresTokenExpiresAt: %w", err)
	}
	return oldValue.RefresTokenExpiresAt, nil
}

// ResetRefresTokenExpiresAt resets all changes to the "refres_token_expires_at" field.
func (m *UserMutation) ResetRefresTokenExpiresAt() {
	m.refres_token_expires_at = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.account_id != nil {
		fields = append(fields, user.FieldAccountID)
	}
	if m.display_name != nil {
		fields = append(fields, user.FieldDisplayName)
	}
	if m.access_token != nil {
		fields = append(fields, user.FieldAccessToken)
	}
	if m.access_token_expires_at != nil {
		fields = append(fields, user.FieldAccessTokenExpiresAt)
	}
	if m.refresh_token != nil {
		fields = append(fields, user.FieldRefreshToken)
	}
	if m.refres_token_expires_at != nil {
		fields = append(fields, user.FieldRefresTokenExpiresAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAccountID:
		return m.AccountID()
	case user.FieldDisplayName:
		return m.DisplayName()
	case user.FieldAccessToken:
		return m.AccessToken()
	case user.FieldAccessTokenExpiresAt:
		return m.AccessTokenExpiresAt()
	case user.FieldRefreshToken:
		return m.RefreshToken()
	case user.FieldRefresTokenExpiresAt:
		return m.RefresTokenExpiresAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldAccountID:
		return m.OldAccountID(ctx)
	case user.FieldDisplayName:
		return m.OldDisplayName(ctx)
	case user.FieldAccessToken:
		return m.OldAccessToken(ctx)
	case user.FieldAccessTokenExpiresAt:
		return m.OldAccessTokenExpiresAt(ctx)
	case user.FieldRefreshToken:
		return m.OldRefreshToken(ctx)
	case user.FieldRefresTokenExpiresAt:
		return m.OldRefresTokenExpiresAt(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldAccountID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAccountID(v)
		return nil
	case user.FieldDisplayName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDisplayName(v)
		return nil
	case user.FieldAccessToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAccessToken(v)
		return nil
	case user.FieldAccessTokenExpiresAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAccessTokenExpiresAt(v)
		return nil
	case user.FieldRefreshToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRefreshToken(v)
		return nil
	case user.FieldRefresTokenExpiresAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRefresTokenExpiresAt(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldAccountID:
		m.ResetAccountID()
		return nil
	case user.FieldDisplayName:
		m.ResetDisplayName()
		return nil
	case user.FieldAccessToken:
		m.ResetAccessToken()
		return nil
	case user.FieldAccessTokenExpiresAt:
		m.ResetAccessTokenExpiresAt()
		return nil
	case user.FieldRefreshToken:
		m.ResetRefreshToken()
		return nil
	case user.FieldRefresTokenExpiresAt:
		m.ResetRefresTokenExpiresAt()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}

// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// PeopleInCharge is an object representing the database table.
type PeopleInCharge struct {
	ID     int `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID int `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	MenuID int `boil:"menu_id" json:"menu_id" toml:"menu_id" yaml:"menu_id"`

	R *peopleInChargeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L peopleInChargeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PeopleInChargeColumns = struct {
	ID     string
	UserID string
	MenuID string
}{
	ID:     "id",
	UserID: "user_id",
	MenuID: "menu_id",
}

// Generated where

var PeopleInChargeWhere = struct {
	ID     whereHelperint
	UserID whereHelperint
	MenuID whereHelperint
}{
	ID:     whereHelperint{field: `id`},
	UserID: whereHelperint{field: `user_id`},
	MenuID: whereHelperint{field: `menu_id`},
}

// PeopleInChargeRels is where relationship names are stored.
var PeopleInChargeRels = struct {
	User string
	Menu string
}{
	User: "User",
	Menu: "Menu",
}

// peopleInChargeR is where relationships are stored.
type peopleInChargeR struct {
	User *User
	Menu *Menu
}

// NewStruct creates a new relationship struct
func (*peopleInChargeR) NewStruct() *peopleInChargeR {
	return &peopleInChargeR{}
}

// peopleInChargeL is where Load methods for each relationship are stored.
type peopleInChargeL struct{}

var (
	peopleInChargeColumns               = []string{"id", "user_id", "menu_id"}
	peopleInChargeColumnsWithoutDefault = []string{"user_id", "menu_id"}
	peopleInChargeColumnsWithDefault    = []string{"id"}
	peopleInChargePrimaryKeyColumns     = []string{"id"}
)

type (
	// PeopleInChargeSlice is an alias for a slice of pointers to PeopleInCharge.
	// This should generally be used opposed to []PeopleInCharge.
	PeopleInChargeSlice []*PeopleInCharge
	// PeopleInChargeHook is the signature for custom PeopleInCharge hook methods
	PeopleInChargeHook func(context.Context, boil.ContextExecutor, *PeopleInCharge) error

	peopleInChargeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	peopleInChargeType                 = reflect.TypeOf(&PeopleInCharge{})
	peopleInChargeMapping              = queries.MakeStructMapping(peopleInChargeType)
	peopleInChargePrimaryKeyMapping, _ = queries.BindMapping(peopleInChargeType, peopleInChargeMapping, peopleInChargePrimaryKeyColumns)
	peopleInChargeInsertCacheMut       sync.RWMutex
	peopleInChargeInsertCache          = make(map[string]insertCache)
	peopleInChargeUpdateCacheMut       sync.RWMutex
	peopleInChargeUpdateCache          = make(map[string]updateCache)
	peopleInChargeUpsertCacheMut       sync.RWMutex
	peopleInChargeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var peopleInChargeBeforeInsertHooks []PeopleInChargeHook
var peopleInChargeBeforeUpdateHooks []PeopleInChargeHook
var peopleInChargeBeforeDeleteHooks []PeopleInChargeHook
var peopleInChargeBeforeUpsertHooks []PeopleInChargeHook

var peopleInChargeAfterInsertHooks []PeopleInChargeHook
var peopleInChargeAfterSelectHooks []PeopleInChargeHook
var peopleInChargeAfterUpdateHooks []PeopleInChargeHook
var peopleInChargeAfterDeleteHooks []PeopleInChargeHook
var peopleInChargeAfterUpsertHooks []PeopleInChargeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PeopleInCharge) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peopleInChargeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PeopleInCharge) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peopleInChargeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PeopleInCharge) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peopleInChargeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PeopleInCharge) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peopleInChargeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PeopleInCharge) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peopleInChargeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PeopleInCharge) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peopleInChargeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PeopleInCharge) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peopleInChargeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PeopleInCharge) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peopleInChargeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PeopleInCharge) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peopleInChargeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPeopleInChargeHook registers your hook function for all future operations.
func AddPeopleInChargeHook(hookPoint boil.HookPoint, peopleInChargeHook PeopleInChargeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		peopleInChargeBeforeInsertHooks = append(peopleInChargeBeforeInsertHooks, peopleInChargeHook)
	case boil.BeforeUpdateHook:
		peopleInChargeBeforeUpdateHooks = append(peopleInChargeBeforeUpdateHooks, peopleInChargeHook)
	case boil.BeforeDeleteHook:
		peopleInChargeBeforeDeleteHooks = append(peopleInChargeBeforeDeleteHooks, peopleInChargeHook)
	case boil.BeforeUpsertHook:
		peopleInChargeBeforeUpsertHooks = append(peopleInChargeBeforeUpsertHooks, peopleInChargeHook)
	case boil.AfterInsertHook:
		peopleInChargeAfterInsertHooks = append(peopleInChargeAfterInsertHooks, peopleInChargeHook)
	case boil.AfterSelectHook:
		peopleInChargeAfterSelectHooks = append(peopleInChargeAfterSelectHooks, peopleInChargeHook)
	case boil.AfterUpdateHook:
		peopleInChargeAfterUpdateHooks = append(peopleInChargeAfterUpdateHooks, peopleInChargeHook)
	case boil.AfterDeleteHook:
		peopleInChargeAfterDeleteHooks = append(peopleInChargeAfterDeleteHooks, peopleInChargeHook)
	case boil.AfterUpsertHook:
		peopleInChargeAfterUpsertHooks = append(peopleInChargeAfterUpsertHooks, peopleInChargeHook)
	}
}

// One returns a single peopleInCharge record from the query.
func (q peopleInChargeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PeopleInCharge, error) {
	o := &PeopleInCharge{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for people_in_charge")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PeopleInCharge records from the query.
func (q peopleInChargeQuery) All(ctx context.Context, exec boil.ContextExecutor) (PeopleInChargeSlice, error) {
	var o []*PeopleInCharge

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PeopleInCharge slice")
	}

	if len(peopleInChargeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PeopleInCharge records in the query.
func (q peopleInChargeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count people_in_charge rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q peopleInChargeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if people_in_charge exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *PeopleInCharge) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// Menu pointed to by the foreign key.
func (o *PeopleInCharge) Menu(mods ...qm.QueryMod) menuQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.MenuID),
	}

	queryMods = append(queryMods, mods...)

	query := Menus(queryMods...)
	queries.SetFrom(query.Query, "\"menus\"")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (peopleInChargeL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybePeopleInCharge interface{}, mods queries.Applicator) error {
	var slice []*PeopleInCharge
	var object *PeopleInCharge

	if singular {
		object = maybePeopleInCharge.(*PeopleInCharge)
	} else {
		slice = *maybePeopleInCharge.(*[]*PeopleInCharge)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &peopleInChargeR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &peopleInChargeR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`users`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(peopleInChargeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.PeopleInCharges = append(foreign.R.PeopleInCharges, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.PeopleInCharges = append(foreign.R.PeopleInCharges, local)
				break
			}
		}
	}

	return nil
}

// LoadMenu allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (peopleInChargeL) LoadMenu(ctx context.Context, e boil.ContextExecutor, singular bool, maybePeopleInCharge interface{}, mods queries.Applicator) error {
	var slice []*PeopleInCharge
	var object *PeopleInCharge

	if singular {
		object = maybePeopleInCharge.(*PeopleInCharge)
	} else {
		slice = *maybePeopleInCharge.(*[]*PeopleInCharge)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &peopleInChargeR{}
		}
		args = append(args, object.MenuID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &peopleInChargeR{}
			}

			for _, a := range args {
				if a == obj.MenuID {
					continue Outer
				}
			}

			args = append(args, obj.MenuID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`menus`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Menu")
	}

	var resultSlice []*Menu
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Menu")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for menus")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for menus")
	}

	if len(peopleInChargeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Menu = foreign
		if foreign.R == nil {
			foreign.R = &menuR{}
		}
		foreign.R.PeopleInCharges = append(foreign.R.PeopleInCharges, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.MenuID == foreign.ID {
				local.R.Menu = foreign
				if foreign.R == nil {
					foreign.R = &menuR{}
				}
				foreign.R.PeopleInCharges = append(foreign.R.PeopleInCharges, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the peopleInCharge to the related item.
// Sets o.R.User to related.
// Adds o to related.R.PeopleInCharges.
func (o *PeopleInCharge) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"people_in_charge\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, peopleInChargePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &peopleInChargeR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			PeopleInCharges: PeopleInChargeSlice{o},
		}
	} else {
		related.R.PeopleInCharges = append(related.R.PeopleInCharges, o)
	}

	return nil
}

// SetMenu of the peopleInCharge to the related item.
// Sets o.R.Menu to related.
// Adds o to related.R.PeopleInCharges.
func (o *PeopleInCharge) SetMenu(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Menu) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"people_in_charge\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"menu_id"}),
		strmangle.WhereClause("\"", "\"", 2, peopleInChargePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.MenuID = related.ID
	if o.R == nil {
		o.R = &peopleInChargeR{
			Menu: related,
		}
	} else {
		o.R.Menu = related
	}

	if related.R == nil {
		related.R = &menuR{
			PeopleInCharges: PeopleInChargeSlice{o},
		}
	} else {
		related.R.PeopleInCharges = append(related.R.PeopleInCharges, o)
	}

	return nil
}

// PeopleInCharges retrieves all the records using an executor.
func PeopleInCharges(mods ...qm.QueryMod) peopleInChargeQuery {
	mods = append(mods, qm.From("\"people_in_charge\""))
	return peopleInChargeQuery{NewQuery(mods...)}
}

// FindPeopleInCharge retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPeopleInCharge(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*PeopleInCharge, error) {
	peopleInChargeObj := &PeopleInCharge{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"people_in_charge\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, peopleInChargeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from people_in_charge")
	}

	return peopleInChargeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PeopleInCharge) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no people_in_charge provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(peopleInChargeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	peopleInChargeInsertCacheMut.RLock()
	cache, cached := peopleInChargeInsertCache[key]
	peopleInChargeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			peopleInChargeColumns,
			peopleInChargeColumnsWithDefault,
			peopleInChargeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(peopleInChargeType, peopleInChargeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(peopleInChargeType, peopleInChargeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"people_in_charge\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"people_in_charge\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into people_in_charge")
	}

	if !cached {
		peopleInChargeInsertCacheMut.Lock()
		peopleInChargeInsertCache[key] = cache
		peopleInChargeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PeopleInCharge.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PeopleInCharge) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	peopleInChargeUpdateCacheMut.RLock()
	cache, cached := peopleInChargeUpdateCache[key]
	peopleInChargeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			peopleInChargeColumns,
			peopleInChargePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update people_in_charge, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"people_in_charge\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, peopleInChargePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(peopleInChargeType, peopleInChargeMapping, append(wl, peopleInChargePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update people_in_charge row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for people_in_charge")
	}

	if !cached {
		peopleInChargeUpdateCacheMut.Lock()
		peopleInChargeUpdateCache[key] = cache
		peopleInChargeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q peopleInChargeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for people_in_charge")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for people_in_charge")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PeopleInChargeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), peopleInChargePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"people_in_charge\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, peopleInChargePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in peopleInCharge slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all peopleInCharge")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PeopleInCharge) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no people_in_charge provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(peopleInChargeColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	peopleInChargeUpsertCacheMut.RLock()
	cache, cached := peopleInChargeUpsertCache[key]
	peopleInChargeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			peopleInChargeColumns,
			peopleInChargeColumnsWithDefault,
			peopleInChargeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			peopleInChargeColumns,
			peopleInChargePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert people_in_charge, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(peopleInChargePrimaryKeyColumns))
			copy(conflict, peopleInChargePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"people_in_charge\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(peopleInChargeType, peopleInChargeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(peopleInChargeType, peopleInChargeMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert people_in_charge")
	}

	if !cached {
		peopleInChargeUpsertCacheMut.Lock()
		peopleInChargeUpsertCache[key] = cache
		peopleInChargeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PeopleInCharge record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PeopleInCharge) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PeopleInCharge provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), peopleInChargePrimaryKeyMapping)
	sql := "DELETE FROM \"people_in_charge\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from people_in_charge")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for people_in_charge")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q peopleInChargeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no peopleInChargeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from people_in_charge")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for people_in_charge")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PeopleInChargeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PeopleInCharge slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(peopleInChargeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), peopleInChargePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"people_in_charge\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, peopleInChargePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from peopleInCharge slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for people_in_charge")
	}

	if len(peopleInChargeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PeopleInCharge) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPeopleInCharge(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PeopleInChargeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PeopleInChargeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), peopleInChargePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"people_in_charge\".* FROM \"people_in_charge\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, peopleInChargePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PeopleInChargeSlice")
	}

	*o = slice

	return nil
}

// PeopleInChargeExists checks if the PeopleInCharge row exists.
func PeopleInChargeExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"people_in_charge\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if people_in_charge exists")
	}

	return exists, nil
}

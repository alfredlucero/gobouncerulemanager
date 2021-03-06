// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// BounceRule is an object representing the database table.
type BounceRule struct {
	ID           int16  `boil:"id" json:"id" toml:"id" yaml:"id"`
	ResponseCode int16  `boil:"response_code" json:"response_code" toml:"response_code" yaml:"response_code"`
	EnhancedCode string `boil:"enhanced_code" json:"enhanced_code" toml:"enhanced_code" yaml:"enhanced_code"`
	Regex        string `boil:"regex" json:"regex" toml:"regex" yaml:"regex"`
	Priority     int8   `boil:"priority" json:"priority" toml:"priority" yaml:"priority"`
	Description  string `boil:"description" json:"description" toml:"description" yaml:"description"`
	BounceAction string `boil:"bounce_action" json:"bounce_action" toml:"bounce_action" yaml:"bounce_action"`

	R *bounceRuleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bounceRuleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BounceRuleColumns = struct {
	ID           string
	ResponseCode string
	EnhancedCode string
	Regex        string
	Priority     string
	Description  string
	BounceAction string
}{
	ID:           "id",
	ResponseCode: "response_code",
	EnhancedCode: "enhanced_code",
	Regex:        "regex",
	Priority:     "priority",
	Description:  "description",
	BounceAction: "bounce_action",
}

var BounceRuleTableColumns = struct {
	ID           string
	ResponseCode string
	EnhancedCode string
	Regex        string
	Priority     string
	Description  string
	BounceAction string
}{
	ID:           "bounce_rule.id",
	ResponseCode: "bounce_rule.response_code",
	EnhancedCode: "bounce_rule.enhanced_code",
	Regex:        "bounce_rule.regex",
	Priority:     "bounce_rule.priority",
	Description:  "bounce_rule.description",
	BounceAction: "bounce_rule.bounce_action",
}

// Generated where

type whereHelperint16 struct{ field string }

func (w whereHelperint16) EQ(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint16) NEQ(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint16) LT(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint16) LTE(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint16) GT(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint16) GTE(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint16) IN(slice []int16) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint16) NIN(slice []int16) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint8 struct{ field string }

func (w whereHelperint8) EQ(x int8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint8) NEQ(x int8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint8) LT(x int8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint8) LTE(x int8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint8) GT(x int8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint8) GTE(x int8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint8) IN(slice []int8) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint8) NIN(slice []int8) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var BounceRuleWhere = struct {
	ID           whereHelperint16
	ResponseCode whereHelperint16
	EnhancedCode whereHelperstring
	Regex        whereHelperstring
	Priority     whereHelperint8
	Description  whereHelperstring
	BounceAction whereHelperstring
}{
	ID:           whereHelperint16{field: "`bounce_rule`.`id`"},
	ResponseCode: whereHelperint16{field: "`bounce_rule`.`response_code`"},
	EnhancedCode: whereHelperstring{field: "`bounce_rule`.`enhanced_code`"},
	Regex:        whereHelperstring{field: "`bounce_rule`.`regex`"},
	Priority:     whereHelperint8{field: "`bounce_rule`.`priority`"},
	Description:  whereHelperstring{field: "`bounce_rule`.`description`"},
	BounceAction: whereHelperstring{field: "`bounce_rule`.`bounce_action`"},
}

// BounceRuleRels is where relationship names are stored.
var BounceRuleRels = struct {
}{}

// bounceRuleR is where relationships are stored.
type bounceRuleR struct {
}

// NewStruct creates a new relationship struct
func (*bounceRuleR) NewStruct() *bounceRuleR {
	return &bounceRuleR{}
}

// bounceRuleL is where Load methods for each relationship are stored.
type bounceRuleL struct{}

var (
	bounceRuleAllColumns            = []string{"id", "response_code", "enhanced_code", "regex", "priority", "description", "bounce_action"}
	bounceRuleColumnsWithoutDefault = []string{"enhanced_code", "regex", "description", "bounce_action"}
	bounceRuleColumnsWithDefault    = []string{"id", "response_code", "priority"}
	bounceRulePrimaryKeyColumns     = []string{"id"}
)

type (
	// BounceRuleSlice is an alias for a slice of pointers to BounceRule.
	// This should almost always be used instead of []BounceRule.
	BounceRuleSlice []*BounceRule
	// BounceRuleHook is the signature for custom BounceRule hook methods
	BounceRuleHook func(context.Context, boil.ContextExecutor, *BounceRule) error

	bounceRuleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bounceRuleType                 = reflect.TypeOf(&BounceRule{})
	bounceRuleMapping              = queries.MakeStructMapping(bounceRuleType)
	bounceRulePrimaryKeyMapping, _ = queries.BindMapping(bounceRuleType, bounceRuleMapping, bounceRulePrimaryKeyColumns)
	bounceRuleInsertCacheMut       sync.RWMutex
	bounceRuleInsertCache          = make(map[string]insertCache)
	bounceRuleUpdateCacheMut       sync.RWMutex
	bounceRuleUpdateCache          = make(map[string]updateCache)
	bounceRuleUpsertCacheMut       sync.RWMutex
	bounceRuleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bounceRuleBeforeInsertHooks []BounceRuleHook
var bounceRuleBeforeUpdateHooks []BounceRuleHook
var bounceRuleBeforeDeleteHooks []BounceRuleHook
var bounceRuleBeforeUpsertHooks []BounceRuleHook

var bounceRuleAfterInsertHooks []BounceRuleHook
var bounceRuleAfterSelectHooks []BounceRuleHook
var bounceRuleAfterUpdateHooks []BounceRuleHook
var bounceRuleAfterDeleteHooks []BounceRuleHook
var bounceRuleAfterUpsertHooks []BounceRuleHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BounceRule) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bounceRuleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BounceRule) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bounceRuleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BounceRule) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bounceRuleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BounceRule) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bounceRuleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BounceRule) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bounceRuleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BounceRule) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bounceRuleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BounceRule) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bounceRuleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BounceRule) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bounceRuleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BounceRule) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bounceRuleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBounceRuleHook registers your hook function for all future operations.
func AddBounceRuleHook(hookPoint boil.HookPoint, bounceRuleHook BounceRuleHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		bounceRuleBeforeInsertHooks = append(bounceRuleBeforeInsertHooks, bounceRuleHook)
	case boil.BeforeUpdateHook:
		bounceRuleBeforeUpdateHooks = append(bounceRuleBeforeUpdateHooks, bounceRuleHook)
	case boil.BeforeDeleteHook:
		bounceRuleBeforeDeleteHooks = append(bounceRuleBeforeDeleteHooks, bounceRuleHook)
	case boil.BeforeUpsertHook:
		bounceRuleBeforeUpsertHooks = append(bounceRuleBeforeUpsertHooks, bounceRuleHook)
	case boil.AfterInsertHook:
		bounceRuleAfterInsertHooks = append(bounceRuleAfterInsertHooks, bounceRuleHook)
	case boil.AfterSelectHook:
		bounceRuleAfterSelectHooks = append(bounceRuleAfterSelectHooks, bounceRuleHook)
	case boil.AfterUpdateHook:
		bounceRuleAfterUpdateHooks = append(bounceRuleAfterUpdateHooks, bounceRuleHook)
	case boil.AfterDeleteHook:
		bounceRuleAfterDeleteHooks = append(bounceRuleAfterDeleteHooks, bounceRuleHook)
	case boil.AfterUpsertHook:
		bounceRuleAfterUpsertHooks = append(bounceRuleAfterUpsertHooks, bounceRuleHook)
	}
}

// One returns a single bounceRule record from the query.
func (q bounceRuleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BounceRule, error) {
	o := &BounceRule{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for bounce_rule")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all BounceRule records from the query.
func (q bounceRuleQuery) All(ctx context.Context, exec boil.ContextExecutor) (BounceRuleSlice, error) {
	var o []*BounceRule

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BounceRule slice")
	}

	if len(bounceRuleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all BounceRule records in the query.
func (q bounceRuleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count bounce_rule rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q bounceRuleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if bounce_rule exists")
	}

	return count > 0, nil
}

// BounceRules retrieves all the records using an executor.
func BounceRules(mods ...qm.QueryMod) bounceRuleQuery {
	mods = append(mods, qm.From("`bounce_rule`"))
	return bounceRuleQuery{NewQuery(mods...)}
}

// FindBounceRule retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBounceRule(ctx context.Context, exec boil.ContextExecutor, iD int16, selectCols ...string) (*BounceRule, error) {
	bounceRuleObj := &BounceRule{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `bounce_rule` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, bounceRuleObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from bounce_rule")
	}

	if err = bounceRuleObj.doAfterSelectHooks(ctx, exec); err != nil {
		return bounceRuleObj, err
	}

	return bounceRuleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BounceRule) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bounce_rule provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bounceRuleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bounceRuleInsertCacheMut.RLock()
	cache, cached := bounceRuleInsertCache[key]
	bounceRuleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bounceRuleAllColumns,
			bounceRuleColumnsWithDefault,
			bounceRuleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bounceRuleType, bounceRuleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bounceRuleType, bounceRuleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `bounce_rule` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `bounce_rule` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `bounce_rule` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, bounceRulePrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into bounce_rule")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int16(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == bounceRuleMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for bounce_rule")
	}

CacheNoHooks:
	if !cached {
		bounceRuleInsertCacheMut.Lock()
		bounceRuleInsertCache[key] = cache
		bounceRuleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the BounceRule.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BounceRule) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bounceRuleUpdateCacheMut.RLock()
	cache, cached := bounceRuleUpdateCache[key]
	bounceRuleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bounceRuleAllColumns,
			bounceRulePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update bounce_rule, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `bounce_rule` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, bounceRulePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bounceRuleType, bounceRuleMapping, append(wl, bounceRulePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update bounce_rule row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for bounce_rule")
	}

	if !cached {
		bounceRuleUpdateCacheMut.Lock()
		bounceRuleUpdateCache[key] = cache
		bounceRuleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q bounceRuleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for bounce_rule")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for bounce_rule")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BounceRuleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bounceRulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `bounce_rule` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bounceRulePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in bounceRule slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all bounceRule")
	}
	return rowsAff, nil
}

var mySQLBounceRuleUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BounceRule) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bounce_rule provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bounceRuleColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLBounceRuleUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	bounceRuleUpsertCacheMut.RLock()
	cache, cached := bounceRuleUpsertCache[key]
	bounceRuleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bounceRuleAllColumns,
			bounceRuleColumnsWithDefault,
			bounceRuleColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			bounceRuleAllColumns,
			bounceRulePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert bounce_rule, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`bounce_rule`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `bounce_rule` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(bounceRuleType, bounceRuleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bounceRuleType, bounceRuleMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for bounce_rule")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int16(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == bounceRuleMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(bounceRuleType, bounceRuleMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for bounce_rule")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for bounce_rule")
	}

CacheNoHooks:
	if !cached {
		bounceRuleUpsertCacheMut.Lock()
		bounceRuleUpsertCache[key] = cache
		bounceRuleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single BounceRule record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BounceRule) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BounceRule provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bounceRulePrimaryKeyMapping)
	sql := "DELETE FROM `bounce_rule` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from bounce_rule")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for bounce_rule")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bounceRuleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no bounceRuleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bounce_rule")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bounce_rule")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BounceRuleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bounceRuleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bounceRulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `bounce_rule` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bounceRulePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bounceRule slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bounce_rule")
	}

	if len(bounceRuleAfterDeleteHooks) != 0 {
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
func (o *BounceRule) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBounceRule(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BounceRuleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BounceRuleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bounceRulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `bounce_rule`.* FROM `bounce_rule` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bounceRulePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BounceRuleSlice")
	}

	*o = slice

	return nil
}

// BounceRuleExists checks if the BounceRule row exists.
func BounceRuleExists(ctx context.Context, exec boil.ContextExecutor, iD int16) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `bounce_rule` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if bounce_rule exists")
	}

	return exists, nil
}

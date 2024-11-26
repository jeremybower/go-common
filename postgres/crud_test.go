package postgres

import (
	"context"
	"strconv"
	"testing"

	"github.com/jeremybower/go-common/pagination"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type valueRow struct {
	ID    int64  `db:"id"`
	Name  string `db:"name"`
	Value string `db:"value"`
}

func TestCountT(t *testing.T) {
	t.Parallel()

	dbPool := databasePoolForTesting(t)
	defer dbPool.Close()

	var id int64
	ctx := context.Background()
	err := dbPool.QueryRow(ctx, "INSERT INTO values (name, value) VALUES ('name', 'value') RETURNING id;").Scan(&id)
	require.NoError(t, err)

	templ := MustParse("SELECT COUNT(*) FROM values WHERE id = {{ arg .ID }}")
	data := map[string]any{"ID": id}
	c, err := CountT(ctx, dbPool, templ, data)
	require.NoError(t, err)
	assert.Equal(t, int64(1), c)
}

func TestExecT(t *testing.T) {
	t.Parallel()

	dbPool := databasePoolForTesting(t)
	defer dbPool.Close()

	var id int64
	ctx := context.Background()
	err := dbPool.QueryRow(ctx, "INSERT INTO values (name, value) VALUES ('name', 'value') RETURNING id;").Scan(&id)
	require.NoError(t, err)

	templ := MustParse("DELETE FROM values WHERE id = {{ arg .ID }}")
	data := map[string]any{"ID": id}
	c, err := ExecT(ctx, dbPool, templ, data)
	require.NoError(t, err)
	assert.Equal(t, int64(1), c)

	c, err = CountT(ctx, dbPool, MustParse("SELECT COUNT(*) FROM values WHERE id = {{ arg .ID }}"), map[string]any{"ID": id})
	require.NoError(t, err)
	assert.Equal(t, int64(0), c)
}

func TestReadOneT(t *testing.T) {
	t.Parallel()

	dbPool := databasePoolForTesting(t)
	defer dbPool.Close()

	var id int64
	ctx := context.Background()
	err := dbPool.QueryRow(ctx, "INSERT INTO values (name, value) VALUES ('name', 'value') RETURNING id;").Scan(&id)
	require.NoError(t, err)

	templ := MustParse("SELECT * FROM values WHERE id = {{ arg .ID }}")
	data := map[string]any{"ID": id}
	row, err := ReadOneT[valueRow](ctx, dbPool, templ, data)
	require.NoError(t, err)
	assert.Equal(t, id, row.ID)
	assert.Equal(t, "name", row.Name)
	assert.Equal(t, "value", row.Value)
}

func TestReadManyT(t *testing.T) {
	t.Parallel()

	dbPool := databasePoolForTesting(t)
	defer dbPool.Close()

	var ids []int64
	ctx := context.Background()
	for i := 0; i < 3; i++ {
		var id int64
		err := dbPool.QueryRow(ctx, "INSERT INTO values (name, value) VALUES ('name', 'value') RETURNING id;").Scan(&id)
		require.NoError(t, err)
		ids = append(ids, id)
	}

	templ := MustParse("SELECT * FROM values WHERE id = {{ arg .ID }}")
	data := map[string]any{"ID": ids[0]}
	rows, err := ReadManyT[valueRow](ctx, dbPool, templ, data)
	require.NoError(t, err)
	require.Len(t, rows, 1)
	assert.Equal(t, ids[0], rows[0].ID)
	assert.Equal(t, "name", rows[0].Name)
	assert.Equal(t, "value", rows[0].Value)
}

func TestListT(t *testing.T) {
	t.Parallel()

	dbPool := databasePoolForTesting(t)
	defer dbPool.Close()

	var ids []int64
	ctx := context.Background()
	for i := 0; i < 3; i++ {
		var id int64
		name := "name" + strconv.Itoa(i)
		value := "value" + strconv.Itoa(i)
		err := dbPool.QueryRow(ctx, "INSERT INTO values (name, value) VALUES ($1, $2) RETURNING id;", name, value).Scan(&id)
		require.NoError(t, err)
		ids = append(ids, id)
	}

	templ := MustParse(`SELECT {{ if counting }} COUNT(*) {{ else }} * {{ end }} FROM values {{ if not counting }} LIMIT {{ pageSize }} OFFSET {{ firstItemIndex }} {{ end }};`)
	paged, err := ListT[valueRow](ctx, dbPool, templ, map[string]any{}, 0, 2, 1, 10, 100)
	require.NoError(t, err)
	assert.Equal(t, pagination.Result[*valueRow]{
		PageIndex:      0,
		PageSize:       2,
		FirstItemIndex: 0,
		TotalItems:     3,
		TotalPages:     2,
		Items: []*valueRow{
			{ID: ids[0], Name: "name0", Value: "value0"},
			{ID: ids[1], Name: "name1", Value: "value1"},
		},
	}, *paged)
}

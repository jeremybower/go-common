package postgres

import (
	"testing"

	"github.com/jeremybower/go-optional"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTemplateParse(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() { MustParse("SELECT * FROM table WHERE id = {{ arg .ID }}") })
	assert.Panics(t, func() { MustParse("SELECT * FROM table WHERE id = {{ arg .ID") })
	assert.Panics(t, func() { MustParse("SELECT * FROM table WHERE id = {{ unknown }}") })
}

func TestTemplateExecute(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		text         string
		data         map[string]any
		expectedErr  error
		expectedSQL  string
		expectedArgs []any
	}{
		{
			name:         "simple",
			text:         `SELECT * FROM table WHERE id = {{ arg .ID }}`,
			data:         map[string]any{"ID": "123"},
			expectedSQL:  `SELECT * FROM table WHERE id = $1`,
			expectedArgs: []any{"123"},
		},
		{
			name:         "join",
			text:         `SELECT * FROM {{- join "AND" -}} {{ sep }} tableA {{ sep }} tableB {{ endJoin -}} WHERE id = {{ arg .ID }}`,
			data:         map[string]any{"ID": "123"},
			expectedSQL:  `SELECT * FROM tableA AND tableB WHERE id = $1`,
			expectedArgs: []any{"123"},
		},
		{
			name:         "join not ended",
			text:         `SELECT * FROM {{- join "AND" -}} {{ sep }} tableA {{ sep }} tableB WHERE id = {{ arg .ID }}`,
			data:         map[string]any{"ID": "123"},
			expectedErr:  ErrJoinNotEnded,
			expectedSQL:  "",
			expectedArgs: nil,
		},
		{
			name:         "endJoin not available",
			text:         `SELECT * FROM tableA {{ endJoin }} WHERE id = {{ arg .ID }}`,
			data:         map[string]any{"ID": "123"},
			expectedErr:  ErrJoinNotStarted,
			expectedSQL:  "",
			expectedArgs: nil,
		},
		{
			name:         "sep not available",
			text:         `SELECT * FROM tableA {{ sep }} tableB WHERE id = {{ arg .ID }}`,
			data:         map[string]any{"ID": "123"},
			expectedErr:  ErrTemplateFuncNotAvail,
			expectedSQL:  "",
			expectedArgs: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := Parse(tt.text)
			require.NoError(t, err)

			sql, args, err := tmpl.Execute(tt.data)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedSQL, sql)
				assert.Equal(t, tt.expectedArgs, args)
			}
		})
	}
}

func TestTemplateExecuteCount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		text         string
		data         map[string]any
		expectedErr  error
		expectedSQL  string
		expectedArgs []any
	}{
		{
			name:         "simple",
			text:         `{{ if counting -}} SELECT * FROM table WHERE id = {{ arg .ID }} {{- end }}`,
			data:         map[string]any{"ID": "123"},
			expectedSQL:  `SELECT * FROM table WHERE id = $1`,
			expectedArgs: []any{"123"},
		},
		{
			name:         "firstItemIndex not available",
			text:         `SELECT * FROM table WHERE id = {{ arg .ID }} OFFSET {{ firstItemIndex }}`,
			data:         map[string]any{"ID": "123"},
			expectedErr:  ErrTemplateFuncNotAvail,
			expectedSQL:  "",
			expectedArgs: nil,
		},
		{
			name:         "pageSize not available",
			text:         `SELECT * FROM table WHERE id = {{ arg .ID }} LIMIT {{ pageSize }}`,
			data:         map[string]any{"ID": "123"},
			expectedErr:  ErrTemplateFuncNotAvail,
			expectedSQL:  "",
			expectedArgs: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := Parse(tt.text)
			require.NoError(t, err)

			sql, args, err := tmpl.ExecuteCount(tt.data)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedSQL, sql)
				assert.Equal(t, tt.expectedArgs, args)
			}
		})
	}
}

func TestTemplateExecuteList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		text         string
		data         map[string]any
		expectedErr  error
		expectedSQL  string
		expectedArgs []any
	}{
		{
			name:         "simple",
			text:         `SELECT * FROM table WHERE id = {{ arg .ID }} LIMIT {{ pageSize }} OFFSET {{ firstItemIndex }}`,
			data:         map[string]any{"ID": "123"},
			expectedSQL:  `SELECT * FROM table WHERE id = $1 LIMIT $2 OFFSET $3`,
			expectedArgs: []any{"123", int64(10), int64(30)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := Parse(tt.text)
			require.NoError(t, err)

			sql, args, err := tmpl.ExecuteList(tt.data, 30, 10)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedSQL, sql)
				assert.Equal(t, tt.expectedArgs, args)
			}
		})
	}
}

var benchTempl *Template

func BenchmarkParseTemplateSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchTempl = MustParse("SELECT * FROM table WHERE id = {{ arg .ID }}")
	}
}

func BenchmarkParseTemplateComplex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchTempl = MustParse(
			`SELECT * FROM values
			{{ if (or .Name.Valid .Value.Valid) }}
				WHERE
				{{ join "AND" }}
					{{ if .Name.Valid }} {{ sep }} name = {{ arg .Name }} {{ end }}
					{{ if .Value.Valid }} {{ sep }} value = {{ arg .Value }} {{ end }}
				{{ endJoin }}
			{{ end }};`,
		)
	}
}

var benchTemplSQL string
var benchTemplArgs []any
var benchTemplErr error

func BenchmarkExecTemplateSimple(b *testing.B) {
	benchTempl = MustParse("SELECT * FROM table WHERE id = {{ arg .ID }}")
	data := map[string]any{"ID": "123"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchTemplSQL, benchTemplArgs, benchTemplErr = benchTempl.Execute(data)
	}
}

func BenchmarkExecTemplateComplex(b *testing.B) {
	benchTempl = MustParse(
		`SELECT * FROM values
		{{ if (or .Name.Valid .Value.Valid) }}
			WHERE
			{{ join "AND" }}
				{{ if .Name.Valid }} {{ sep }} name = {{ arg .Name }} {{ end }}
				{{ if .Value.Valid }} {{ sep }} value = {{ arg .Value }} {{ end }}
			{{ endJoin }}
		{{ end }};`,
	)
	data := map[string]any{
		"Name":  optional.New("name"),
		"Value": optional.Value[string]{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchTemplSQL, benchTemplArgs, benchTemplErr = benchTempl.Execute(data)
	}
}

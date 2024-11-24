package postgres

import (
	"net/netip"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func TestRequiredArray(t *testing.T) {
	t.Parallel()

	arr := pgtype.Array[int]{
		Elements: []int{1, 2, 3},
		Valid:    true,
	}

	assert.Equal(t, []int{1, 2, 3}, RequiredArray(arr))
	assert.Panics(t, func() { RequiredArray(pgtype.Array[int]{}) })
}

func TestNilableArray(t *testing.T) {
	t.Parallel()

	assert.Nil(t, NilableArray(pgtype.Array[int]{}).Slice)
	assert.Equal(t, []int{1, 2, 3}, NilableArray(pgtype.Array[int]{
		Elements: []int{1, 2, 3},
		Valid:    true,
	}).Slice)
}

func TestRequiredFlatArray(t *testing.T) {
	t.Parallel()

	assert.Equal(t, []int{1, 2, 3}, RequiredFlatArray([]int{1, 2, 3}))
	assert.Panics(t, func() { RequiredFlatArray[int](nil) })
}

func TestNilableFlatArray(t *testing.T) {
	t.Parallel()

	assert.Nil(t, NilableFlatArray[int](nil).Slice)
	assert.Equal(t, []int{1, 2, 3}, NilableFlatArray([]int{1, 2, 3}).Slice)
}

func TestRequiredBool(t *testing.T) {
	t.Parallel()

	assert.True(t, RequiredBool(pgtype.Bool{Bool: true, Valid: true}))
	assert.Panics(t, func() { RequiredBool(pgtype.Bool{Valid: false}) })
}

func TestNilableBool(t *testing.T) {
	t.Parallel()

	assert.Nil(t, NilableBool(pgtype.Bool{Valid: false}).Value)
	assert.True(t, *NilableBool(pgtype.Bool{Bool: true, Valid: true}).Value)
}

func TestRequiredFloat4(t *testing.T) {
	t.Parallel()

	assert.Equal(t, float32(1.23), RequiredFloat4[float32](pgtype.Float4{Float32: 1.23, Valid: true}))
	assert.Panics(t, func() { RequiredFloat4[float32](pgtype.Float4{Valid: false}) })
}

func TestNilableFloat4(t *testing.T) {
	t.Parallel()

	assert.Nil(t, NilableFloat4[float32](pgtype.Float4{Valid: false}).Value)
	assert.Equal(t, float32(1.23), *NilableFloat4[float32](pgtype.Float4{Float32: 1.23, Valid: true}).Value)
}

func TestRequiredFloat8(t *testing.T) {
	t.Parallel()

	assert.Equal(t, float64(1.23), RequiredFloat8[float64](pgtype.Float8{Float64: 1.23, Valid: true}))
	assert.Panics(t, func() { RequiredFloat8[float64](pgtype.Float8{Valid: false}) })
}

func TestNilableFloat8(t *testing.T) {
	t.Parallel()

	assert.Nil(t, NilableFloat8[float64](pgtype.Float8{Valid: false}).Value)
	assert.Equal(t, float64(1.23), *NilableFloat8[float64](pgtype.Float8{Float64: 1.23, Valid: true}).Value)
}

func TestRequiredInt4(t *testing.T) {
	t.Parallel()

	assert.Equal(t, int32(123), RequiredInt4[int32](pgtype.Int4{Int32: 123, Valid: true}))
	assert.Panics(t, func() { RequiredInt4[int32](pgtype.Int4{Valid: false}) })
}

func TestNilableInt4(t *testing.T) {
	t.Parallel()

	assert.Nil(t, NilableInt4[int32](pgtype.Int4{Valid: false}).Value)
	assert.Equal(t, int32(123), *NilableInt4[int32](pgtype.Int4{Int32: 123, Valid: true}).Value)
}

func TestRequiredInt8(t *testing.T) {
	t.Parallel()

	assert.Equal(t, int64(123), RequiredInt8[int64](pgtype.Int8{Int64: 123, Valid: true}))
	assert.Panics(t, func() { RequiredInt8[int64](pgtype.Int8{Valid: false}) })
}

func TestNilableInt8(t *testing.T) {
	t.Parallel()

	assert.Nil(t, NilableInt8[int64](pgtype.Int8{Valid: false}).Value)
	assert.Equal(t, int64(123), *NilableInt8[int64](pgtype.Int8{Int64: 123, Valid: true}).Value)
}

func TestRequiredIPAddr(t *testing.T) {
	t.Parallel()

	assert.Equal(t, netip.AddrFrom4([4]byte{1, 2, 3, 4}), RequiredIPAddr(netip.AddrFrom4([4]byte{1, 2, 3, 4})))
	assert.Panics(t, func() { RequiredIPAddr(netip.Addr{}) })
}

func TestNilableIPAddr(t *testing.T) {
	t.Parallel()

	assert.Nil(t, NilableIPAddr(netip.Addr{}).Value)
	assert.Equal(t, netip.AddrFrom4([4]byte{1, 2, 3, 4}), *NilableIPAddr(netip.AddrFrom4([4]byte{1, 2, 3, 4})).Value)
}

func TestRequiredJSON(t *testing.T) {
	t.Parallel()

	assert.Equal(t, map[string]any{"a": 1}, RequiredJSON(map[string]any{"a": 1}))
	assert.Panics(t, func() { RequiredJSON(nil) })
}

func TestNilableJSON(t *testing.T) {
	t.Parallel()

	assert.Nil(t, NilableJSON(nil).Map)
	assert.Equal(t, map[string]any{"a": 1}, NilableJSON(map[string]any{"a": 1}).Map)
}

func TestRequiredPoint(t *testing.T) {
	t.Parallel()

	assert.Equal(t, pgtype.Vec2{X: 1.23, Y: 4.56}, RequiredPoint(pgtype.Point{P: pgtype.Vec2{X: 1.23, Y: 4.56}, Valid: true}, ToVec2))
	assert.Panics(t, func() { RequiredPoint(pgtype.Point{}, ToVec2) })
}

func TestNilablePoint(t *testing.T) {
	t.Parallel()

	assert.Nil(t, NilablePoint(pgtype.Point{}, ToVec2).Value)
	assert.Equal(t, pgtype.Vec2{X: 1.23, Y: 4.56}, *NilablePoint(pgtype.Point{P: pgtype.Vec2{X: 1.23, Y: 4.56}, Valid: true}, ToVec2).Value)
}

func TestRequiredText(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "abc", RequiredText[string](pgtype.Text{String: "abc", Valid: true}))
	assert.Panics(t, func() { RequiredText[string](pgtype.Text{Valid: false}) })
}

func TestNilableText(t *testing.T) {
	t.Parallel()

	assert.Nil(t, NilableText[string](pgtype.Text{Valid: false}).Value)
	assert.Equal(t, "abc", *NilableText[string](pgtype.Text{String: "abc", Valid: true}).Value)
}

func TestRequiredTimestamp(t *testing.T) {
	t.Parallel()

	assert.Equal(t, time.Date(2024, 1, 2, 3, 4, 5, 6, time.UTC), RequiredTimestamp(pgtype.Timestamp{Time: time.Date(2024, 1, 2, 3, 4, 5, 6, time.UTC), Valid: true}))
	assert.Panics(t, func() { RequiredTimestamp(pgtype.Timestamp{Valid: false}) })
}

func TestNilableTimestamp(t *testing.T) {
	t.Parallel()

	assert.Nil(t, NilableTimestamp(pgtype.Timestamp{Valid: false}).Value)
	assert.Equal(t, time.Date(2024, 1, 2, 3, 4, 5, 6, time.UTC), *NilableTimestamp(pgtype.Timestamp{Time: time.Date(2024, 1, 2, 3, 4, 5, 6, time.UTC), Valid: true}).Value)
}

func TestRequiredUUID(t *testing.T) {
	t.Parallel()

	id := uuid.New()

	assert.Equal(t, id.String(), RequiredUUID[string](pgtype.UUID{Bytes: id, Valid: true}))
	assert.Panics(t, func() { RequiredUUID[string](pgtype.UUID{Valid: false}) })
}

func TestNilableUUID(t *testing.T) {
	t.Parallel()

	id := uuid.New()

	assert.Nil(t, NilableUUID[string](pgtype.UUID{Valid: false}).Value)
	assert.Equal(t, id.String(), *NilableUUID[string](pgtype.UUID{Bytes: id, Valid: true}).Value)
}

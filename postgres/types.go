package postgres

import (
	"fmt"
	"net/netip"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jeremybower/go-common/nilable"
)

//-----------------------------------------------------------------------------
// Array
//-----------------------------------------------------------------------------

func RequiredArray[T any](a pgtype.Array[T]) []T {
	if !a.Valid {
		panic(fmt.Sprintf("invalid array of type %T", a))
	}

	return a.Elements
}

func NilableArray[T any](a pgtype.Array[T]) nilable.Slice[T] {
	if !a.Valid {
		return nilable.NilSlice[T]()
	}

	return nilable.NewSlice(a.Elements)
}

func RequiredFlatArray[T any](a pgtype.FlatArray[T]) []T {
	if a == nil {
		panic(fmt.Sprintf("invalid array of type %T", a))
	}

	return a
}

func NilableFlatArray[T any](a pgtype.FlatArray[T]) nilable.Slice[T] {
	if a == nil {
		return nilable.NilSlice[T]()
	}

	return nilable.NewSlice(a)
}

//-----------------------------------------------------------------------------
// Bool
//-----------------------------------------------------------------------------

func RequiredBool(b pgtype.Bool) bool {
	if !b.Valid {
		panic("invalid boolean")
	}

	return b.Bool
}

func NilableBool(b pgtype.Bool) nilable.Value[bool] {
	if !b.Valid {
		return nilable.NilValue[bool]()
	}

	return nilable.NewValue(&b.Bool)
}

//-----------------------------------------------------------------------------
// Float4
//-----------------------------------------------------------------------------

func RequiredFloat4[T ~float32](n pgtype.Float4) T {
	if !n.Valid {
		panic("float32 is required")
	}

	return T(n.Float32)
}

func NilableFloat4[T ~float32](n pgtype.Float4) nilable.Value[T] {
	if !n.Valid {
		return nilable.NilValue[T]()
	}

	val := T(n.Float32)
	return nilable.NewValue(&val)
}

//-----------------------------------------------------------------------------
// Float8
//-----------------------------------------------------------------------------

func RequiredFloat8[T ~float64](n pgtype.Float8) T {
	if !n.Valid {
		panic("float64 is required")
	}

	val := T(n.Float64)
	return val
}

func NilableFloat8[T ~float64](n pgtype.Float8) nilable.Value[T] {
	if !n.Valid {
		return nilable.NilValue[T]()
	}

	val := T(n.Float64)
	return nilable.NewValue(&val)
}

//-----------------------------------------------------------------------------
// Int4
//-----------------------------------------------------------------------------

func RequiredInt4[T ~int32](n pgtype.Int4) T {
	if !n.Valid {
		panic("invalid value")
	}

	return T(n.Int32)
}

func NilableInt4[T ~int32](n pgtype.Int4) nilable.Value[T] {
	if !n.Valid {
		return nilable.NilValue[T]()
	}

	val := T(n.Int32)
	return nilable.NewValue(&val)
}

//-----------------------------------------------------------------------------
// Int8
//-----------------------------------------------------------------------------

func RequiredInt8[T ~int64](n pgtype.Int8) T {
	if !n.Valid {
		panic("int64 is required")
	}

	return T(n.Int64)
}

func NilableInt8[T ~int64](n pgtype.Int8) nilable.Value[T] {
	if !n.Valid {
		return nilable.NilValue[T]()
	}

	val := T(n.Int64)
	return nilable.NewValue(&val)
}

//-----------------------------------------------------------------------------
// IP Address
//-----------------------------------------------------------------------------

func RequiredIPAddr(ip netip.Addr) netip.Addr {
	if !ip.IsValid() {
		panic("invalid ip address")
	}

	return ip
}

func NilableIPAddr(ip netip.Addr) nilable.Value[netip.Addr] {
	if !ip.IsValid() {
		return nilable.NilValue[netip.Addr]()
	}

	return nilable.NewValue(&ip)
}

//-----------------------------------------------------------------------------
// JSON
//-----------------------------------------------------------------------------

func RequiredJSON(value map[string]any) map[string]any {
	if value == nil {
		panic("invalid json")
	}

	return value
}

func NilableJSON(value map[string]any) nilable.Map[string, any] {
	if value == nil {
		return nilable.NilMap[string, any]()
	}

	return nilable.NewMap(value)
}

//-----------------------------------------------------------------------------
// Point
//-----------------------------------------------------------------------------

func RequiredPoint[T any](p pgtype.Point, fn func(pgtype.Point) T) T {
	if !p.Valid {
		panic("invalid point")
	}

	return fn(p)
}

func NilablePoint[T any](p pgtype.Point, fn func(pgtype.Point) T) nilable.Value[T] {
	if !p.Valid {
		return nilable.NilValue[T]()
	}

	val := fn(p)
	return nilable.NewValue(&val)
}

func ToVec2(p pgtype.Point) pgtype.Vec2 {
	return p.P
}

//-----------------------------------------------------------------------------
// Text
//-----------------------------------------------------------------------------

func RequiredText[T ~string](t pgtype.Text) T {
	if !t.Valid {
		panic("text is required")
	}

	return T(t.String)
}

func NilableText[T ~string](t pgtype.Text) nilable.Value[T] {
	if !t.Valid {
		return nilable.NilValue[T]()
	}

	str := T(t.String)
	return nilable.NewValue(&str)
}

//-----------------------------------------------------------------------------
// Timestamp
//-----------------------------------------------------------------------------

func RequiredTimestamp(t pgtype.Timestamp) time.Time {
	if !t.Valid {
		panic("timestamp is required")
	}

	return t.Time
}

func NilableTimestamp(t pgtype.Timestamp) nilable.Value[time.Time] {
	if !t.Valid {
		return nilable.NilValue[time.Time]()
	}

	time := t.Time
	return nilable.NewValue(&time)
}

//-----------------------------------------------------------------------------
// UUID
//-----------------------------------------------------------------------------

func RequiredUUID[T ~string](u pgtype.UUID) T {
	if !u.Valid {
		panic("uuid is required")
	}

	id, err := uuid.FromBytes(u.Bytes[:])
	if err != nil {
		// FromBytes will only panic if there aren't 16 bytes, which is
		// virtually impossible for a UUID from Postgres.
		panic(err)
	}

	return T(id.String())
}

func NilableUUID[T ~string](u pgtype.UUID) nilable.Value[T] {
	if !u.Valid {
		return nilable.NilValue[T]()
	}

	id, err := uuid.FromBytes(u.Bytes[:])
	if err != nil {
		// FromBytes will only panic if there aren't 16 bytes, which is
		// virtually impossible for a UUID from Postgres.
		panic(err)
	}

	str := T(id.String())
	return nilable.NewValue(&str)
}

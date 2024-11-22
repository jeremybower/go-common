package postgres

import (
	"net/netip"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

//-----------------------------------------------------------------------------
// Array
//-----------------------------------------------------------------------------

func RequiredArray[T any](a pgtype.Array[T]) []T {
	if !a.Valid {
		panic("array is required")
	}

	return a.Elements
}

//-----------------------------------------------------------------------------
// Bool
//-----------------------------------------------------------------------------

func OptionalBool(b pgtype.Bool) *bool {
	if !b.Valid {
		return nil
	}

	val := b.Bool
	return &val
}

func RequiredBool(b pgtype.Bool) bool {
	if !b.Valid {
		panic("bool is required")
	}

	return b.Bool
}

//-----------------------------------------------------------------------------
// IP Address
//-----------------------------------------------------------------------------

func OptionalIPAddr(ip netip.Addr) netip.Addr {
	return ip
}

func RequiredIPAddr(ip netip.Addr) netip.Addr {
	if !ip.IsValid() {
		panic("ip address is required")
	}

	return ip
}

//-----------------------------------------------------------------------------
// JSON
//-----------------------------------------------------------------------------

func OptionalJSON(value map[string]any) map[string]any {
	return value
}

func RequiredJSON(value map[string]any) map[string]any {
	if value == nil {
		panic("json is required")
	}

	return value
}

//-----------------------------------------------------------------------------
// Numeric
//-----------------------------------------------------------------------------

func OptionalInt32(n pgtype.Int4) *int32 {
	if !n.Valid {
		return nil
	}

	val := n.Int32
	return &val
}

func RequiredInt32(n pgtype.Int4) int32 {
	if !n.Valid {
		panic("int32 is required")
	}

	return n.Int32
}

func OptionalInt64(n pgtype.Int8) *int64 {
	if !n.Valid {
		return nil
	}

	val := n.Int64
	return &val
}

func RequiredInt64(n pgtype.Int8) int64 {
	if !n.Valid {
		panic("uint32 is required")
	}

	return n.Int64
}

func OptionalFloat32(n pgtype.Float4) *float32 {
	if !n.Valid {
		return nil
	}

	val := n.Float32
	return &val
}

func RequiredFloat32(n pgtype.Float4) float32 {
	if !n.Valid {
		panic("float32 is required")
	}

	return n.Float32
}

func OptionalFloat64(n pgtype.Float8) *float64 {
	if !n.Valid {
		return nil
	}

	val := n.Float64
	return &val
}

func RequiredFloat64(n pgtype.Float8) float64 {
	if !n.Valid {
		panic("float64 is required")
	}

	return n.Float64
}

//-----------------------------------------------------------------------------
// Text
//-----------------------------------------------------------------------------

func OptionalText[T ~string](t pgtype.Text) *T {
	if !t.Valid {
		return nil
	}

	str := T(t.String)
	return &str
}

func RequiredText[T ~string](t pgtype.Text) T {
	if !t.Valid {
		panic("text is required")
	}

	return T(t.String)
}

//-----------------------------------------------------------------------------
// Timestamp
//-----------------------------------------------------------------------------

func OptionalTimestamp(t pgtype.Timestamp) *time.Time {
	if !t.Valid {
		return nil
	}

	time := t.Time
	return &time
}

func RequiredTimestamp(t pgtype.Timestamp) time.Time {
	if !t.Valid {
		panic("timestamp is required")
	}

	return t.Time
}

//-----------------------------------------------------------------------------
// UUID
//-----------------------------------------------------------------------------

func OptionalUUID[T ~string](u pgtype.UUID) *T {
	if !u.Valid {
		return nil
	}

	v := RequiredUUID[T](u)
	return &v
}

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

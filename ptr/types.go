package ptr

// Bool returns a point to bool
func Bool(v bool) *bool { return &v }

// Int returns a pointer to int
func Int(v int) *int { return &v }

// Int8 returns a pointer to int8
func Int8(v int8) *int8 { return &v }

// Int16 returns a pointer to int16
func Int16(v int16) *int16 { return &v }

// Int32 returns a pointer to int32
func Int32(v int32) *int32 { return &v }

// Int64 returns a pointer to int64
func Int64(v int64) *int64 { return &v }

// Uint returns a pointer to uint
func Uint(v uint) *uint { return &v }

// Uint8 returns a pointer to uint8
func Uint8(v uint8) *uint8 { return &v }

// Uint16 returns a pointer to uint16
func Uint16(v uint16) *uint16 { return &v }

// Uint32 returns a pointer to uint32
func Uint32(v uint32) *uint32 { return &v }

// Uint64 returns a pointer to uint64
func Uint64(v uint64) *uint64 { return &v }

// Float32 returns a pointer to float32
func Float32(v float32) *float32 { return &v }

// Float64 returns a pointer to float64
func Float64(v float64) *float64 { return &v }

// String returns a pointer to string
func String(v string) *string { return &v }

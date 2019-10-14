package val

// Bool returns a bool value of p or false when p is nil
func Bool(p *bool) bool {
	if p != nil {
		return *p
	}
	return false
}

// Int returns an int value of p or 0 when p is nil
func Int(p *int) int {
	if p != nil {
		return *p
	}
	return 0
}

// Int8 returns an int8 value of p or 0 when p is nil
func Int8(p *int8) int8 {
	if p != nil {
		return *p
	}
	return 0
}

// Int16 returns an int16 value of p or 0 when p is nil
func Int16(p *int16) int16 {
	if p != nil {
		return *p
	}
	return 0
}

// Int32 returns an int32 value of p or 0 when p is nil
func Int32(p *int32) int32 {
	if p != nil {
		return *p
	}
	return 0
}

// Int64 returns an int64 value of p or 0 when p is nil
func Int64(p *int64) int64 {
	if p != nil {
		return *p
	}
	return 0
}

// Uint returns an uint value of p or 0 when p is nil
func Uint(p *uint) uint {
	if p != nil {
		return *p
	}
	return 0
}

// Uint8 returns an uint8 value of p or 0 when p is nil
func Uint8(p *uint8) uint8 {
	if p != nil {
		return *p
	}
	return 0
}

// Uint16 returns an uint16 value of p or 0 when p is nil
func Uint16(p *uint16) uint16 {
	if p != nil {
		return *p
	}
	return 0
}

// Uint32 returns an uint32 value of p or 0 when p is nil
func Uint32(p *uint32) uint32 {
	if p != nil {
		return *p
	}
	return 0
}

// Uint64 returns an uint64 value of p or 0 when p is nil
func Uint64(p *uint64) uint64 {
	if p != nil {
		return *p
	}
	return 0
}

// Float32 returns a float32 value of p or 0 when p is nil
func Float32(p *float32) float32 {
	if p != nil {
		return *p
	}
	return 0
}

// Float64 returns a float64 value of p or 0 when p is nil
func Float64(p *float64) float64 {
	if p != nil {
		return *p
	}
	return 0
}

// String returns a string value of p or "" when p is nil
func String(p *string) string {
	if p != nil {
		return *p
	}
	return ""
}

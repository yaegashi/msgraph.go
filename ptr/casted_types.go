package ptr

// CastBool returns a pointer to bool or nil when v is not bool
func CastBool(v interface{}) *bool {
	if x, ok := v.(bool); ok {
		return &x
	}
	return nil
}

// CastInt returns a pointer to int or nil when v is not int
func CastInt(v interface{}) *int {
	if x, ok := v.(int); ok {
		return &x
	}
	return nil
}

// CastInt8 returns a pointer to int8 or nil when v is not int8
func CastInt8(v interface{}) *int8 {
	if x, ok := v.(int8); ok {
		return &x
	}
	return nil
}

// CastInt16 returns a pointer to int16 or nil when v is not int16
func CastInt16(v interface{}) *int16 {
	if x, ok := v.(int16); ok {
		return &x
	}
	return nil
}

// CastInt32 returns a pointer to int32 or nil when v is not int32
func CastInt32(v interface{}) *int32 {
	if x, ok := v.(int32); ok {
		return &x
	}
	return nil
}

// CastInt64 returns a pointer to int64 or nil when v is not int64
func CastInt64(v interface{}) *int64 {
	if x, ok := v.(int64); ok {
		return &x
	}
	return nil
}

// CastUint returns a pointer to uint or nil when v is not uint
func CastUint(v interface{}) *uint {
	if x, ok := v.(uint); ok {
		return &x
	}
	return nil
}

// CastUint8 returns a pointer to uint8 or nil when v is not uint8
func CastUint8(v interface{}) *uint8 {
	if x, ok := v.(uint8); ok {
		return &x
	}
	return nil
}

// CastUint16 returns a pointer to uint16 or nil when v is not uint16
func CastUint16(v interface{}) *uint16 {
	if x, ok := v.(uint16); ok {
		return &x
	}
	return nil
}

// CastUint32 returns a pointer to uint32 or nil when v is not uint32
func CastUint32(v interface{}) *uint32 {
	if x, ok := v.(uint32); ok {
		return &x
	}
	return nil
}

// CastUint64 returns a pointer to uint64 or nil when v is not uint64
func CastUint64(v interface{}) *uint64 {
	if x, ok := v.(uint64); ok {
		return &x
	}
	return nil
}

// CastFloat32 returns a pointer to float32 or nil when v is not float32
func CastFloat32(v interface{}) *float32 {
	if x, ok := v.(float32); ok {
		return &x
	}
	return nil
}

// CastFloat64 returns a pointer to float64 or nil when v is not float64
func CastFloat64(v interface{}) *float64 {
	if x, ok := v.(float64); ok {
		return &x
	}
	return nil
}

// CastString returns a pointer to string or nil when v is not string
func CastString(v interface{}) *string {
	if x, ok := v.(string); ok {
		return &x
	}
	return nil
}

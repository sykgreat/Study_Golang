package cache

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestMemCacheOP(t *testing.T) {
	testData := []struct {
		key    string
		value  interface{}
		expire time.Duration
	}{
		{"keyfdsa1", 123321, 10},
		{"keyfdsagsdafs1", "123321", 10},
		{"kedafs1", 123321, 10},
		{"keyfddfsgfdhssa1", "value1gfdsg", 12},
		{"keyfdhga1", true, 14},
		{"keyfdsdfsafa1", []int{1, 2, 3, 4}, 16},
		{"keyfdsadfashg1", []string{"gfdsgsfds", "gfdsgsfdljkhfds", "gfdsgsfgdsfhds", "gfdsgsfdsfdsaf"}, 18},
		{"keyfdsafds1", map[string]interface{}{"adfs": 15415, "fdsafghd": "fadsfkhsfksgasdf"}, 20},
	}
	c := NewMemCache()
	c.SetMaxMemory("10MB")
	for _, v := range testData {
		c.Set(v.key, v.value, v.expire)
		value, ok := c.Get(v.key)
		if !ok {
			t.Errorf("Get key %s failed", v.key)
		}
		if !reflect.DeepEqual(value, v.value) {
			t.Errorf("Get key %s value %v not equal to %v", v.key, value, v.value)
		}
	}

	if int64(len(testData)) != c.Keys() {
		t.Errorf("Keys() = %d, want %d", c.Keys(), len(testData))
	}

	c.Del(testData[0].key)
	c.Del(testData[1].key)

	if int64(len(testData)) != c.Keys()+2 {
		t.Errorf("Keys() = %d, want %d", c.Keys(), len(testData))
	}

	t.Log(c.currentMemorySize)
	time.Sleep(16 * time.Second)
	t.Log(c.currentMemorySize)
	if c.Keys() != 0 {
		t.Errorf("Keys() = %d, want %d", c.Keys(), 0)
	}
}

func TestMemCache_currentMemorySize(t *testing.T) {
	testData := []struct {
		key    string
		value  interface{}
		expire time.Duration
	}{
		{"keyfdsa1", 123321, 10},
		{"keyfdsagsdafs1", "123321", 10},
	}
	c := NewMemCache()
	c.SetMaxMemory("10MB")
	for _, v := range testData {
		c.Set(v.key, v.value, v.expire)
		value, ok := c.Get(v.key)
		if !ok {
			t.Errorf("Get key %s failed", v.key)
		}
		if !reflect.DeepEqual(value, v.value) {
			t.Errorf("Get key %s value %v not equal to %v", v.key, value, v.value)
		}
	}

	t.Log(c.currentMemorySize)
}

func TestMemCache_Del(t *testing.T) {
	type fields struct {
		maxMemorySize                int64
		maxMemorySizeString          string
		currentMemorySize            int64
		values                       map[string]*memCacheValue
		locker                       sync.RWMutex
		clearExpiredItemTimeInterval time.Duration
	}
	type args struct {
		key string
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemCache{
				maxMemorySize:                tt.fields.maxMemorySize,
				maxMemorySizeString:          tt.fields.maxMemorySizeString,
				currentMemorySize:            tt.fields.currentMemorySize,
				values:                       tt.fields.values,
				locker:                       tt.fields.locker,
				clearExpiredItemTimeInterval: tt.fields.clearExpiredItemTimeInterval,
			}
			if got := m.Del(tt.args.key); got != tt.want {
				t.Errorf("Del() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemCache_Exists(t *testing.T) {
	type fields struct {
		maxMemorySize                int64
		maxMemorySizeString          string
		currentMemorySize            int64
		values                       map[string]*memCacheValue
		locker                       sync.RWMutex
		clearExpiredItemTimeInterval time.Duration
	}
	type args struct {
		key string
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemCache{
				maxMemorySize:                tt.fields.maxMemorySize,
				maxMemorySizeString:          tt.fields.maxMemorySizeString,
				currentMemorySize:            tt.fields.currentMemorySize,
				values:                       tt.fields.values,
				locker:                       tt.fields.locker,
				clearExpiredItemTimeInterval: tt.fields.clearExpiredItemTimeInterval,
			}
			if got := m.Exists(tt.args.key); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemCache_Flush(t *testing.T) {
	type fields struct {
		maxMemorySize                int64
		maxMemorySizeString          string
		currentMemorySize            int64
		values                       map[string]*memCacheValue
		locker                       sync.RWMutex
		clearExpiredItemTimeInterval time.Duration
	}
	var tests []struct {
		name   string
		fields fields
		want   bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemCache{
				maxMemorySize:                tt.fields.maxMemorySize,
				maxMemorySizeString:          tt.fields.maxMemorySizeString,
				currentMemorySize:            tt.fields.currentMemorySize,
				values:                       tt.fields.values,
				locker:                       tt.fields.locker,
				clearExpiredItemTimeInterval: tt.fields.clearExpiredItemTimeInterval,
			}
			if got := m.Flush(); got != tt.want {
				t.Errorf("Flush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemCache_Get(t *testing.T) {
	type fields struct {
		maxMemorySize                int64
		maxMemorySizeString          string
		currentMemorySize            int64
		values                       map[string]*memCacheValue
		locker                       sync.RWMutex
		clearExpiredItemTimeInterval time.Duration
	}
	type args struct {
		key string
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemCache{
				maxMemorySize:                tt.fields.maxMemorySize,
				maxMemorySizeString:          tt.fields.maxMemorySizeString,
				currentMemorySize:            tt.fields.currentMemorySize,
				values:                       tt.fields.values,
				locker:                       tt.fields.locker,
				clearExpiredItemTimeInterval: tt.fields.clearExpiredItemTimeInterval,
			}
			got, got1 := m.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMemCache_Keys(t *testing.T) {
	type fields struct {
		maxMemorySize                int64
		maxMemorySizeString          string
		currentMemorySize            int64
		values                       map[string]*memCacheValue
		locker                       sync.RWMutex
		clearExpiredItemTimeInterval time.Duration
	}
	var tests []struct {
		name   string
		fields fields
		want   int64
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemCache{
				maxMemorySize:                tt.fields.maxMemorySize,
				maxMemorySizeString:          tt.fields.maxMemorySizeString,
				currentMemorySize:            tt.fields.currentMemorySize,
				values:                       tt.fields.values,
				locker:                       tt.fields.locker,
				clearExpiredItemTimeInterval: tt.fields.clearExpiredItemTimeInterval,
			}
			if got := m.Keys(); got != tt.want {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemCache_Set(t *testing.T) {
	type fields struct {
		maxMemorySize                int64
		maxMemorySizeString          string
		currentMemorySize            int64
		values                       map[string]*memCacheValue
		locker                       sync.RWMutex
		clearExpiredItemTimeInterval time.Duration
	}
	type args struct {
		key    string
		value  interface{}
		expire time.Duration
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemCache{
				maxMemorySize:                tt.fields.maxMemorySize,
				maxMemorySizeString:          tt.fields.maxMemorySizeString,
				currentMemorySize:            tt.fields.currentMemorySize,
				values:                       tt.fields.values,
				locker:                       tt.fields.locker,
				clearExpiredItemTimeInterval: tt.fields.clearExpiredItemTimeInterval,
			}
			if got := m.Set(tt.args.key, tt.args.value, tt.args.expire); got != tt.want {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemCache_SetMaxMemory(t *testing.T) {
	type fields struct {
		maxMemorySize                int64
		maxMemorySizeString          string
		currentMemorySize            int64
		values                       map[string]*memCacheValue
		locker                       sync.RWMutex
		clearExpiredItemTimeInterval time.Duration
	}
	type args struct {
		size string
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemCache{
				maxMemorySize:                tt.fields.maxMemorySize,
				maxMemorySizeString:          tt.fields.maxMemorySizeString,
				currentMemorySize:            tt.fields.currentMemorySize,
				values:                       tt.fields.values,
				locker:                       tt.fields.locker,
				clearExpiredItemTimeInterval: tt.fields.clearExpiredItemTimeInterval,
			}
			if got := m.SetMaxMemory(tt.args.size); got != tt.want {
				t.Errorf("SetMaxMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemCache_add(t *testing.T) {
	type fields struct {
		maxMemorySize                int64
		maxMemorySizeString          string
		currentMemorySize            int64
		values                       map[string]*memCacheValue
		locker                       sync.RWMutex
		clearExpiredItemTimeInterval time.Duration
	}
	type args struct {
		key   string
		value *memCacheValue
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemCache{
				maxMemorySize:                tt.fields.maxMemorySize,
				maxMemorySizeString:          tt.fields.maxMemorySizeString,
				currentMemorySize:            tt.fields.currentMemorySize,
				values:                       tt.fields.values,
				locker:                       tt.fields.locker,
				clearExpiredItemTimeInterval: tt.fields.clearExpiredItemTimeInterval,
			}
			m.add(tt.args.key, tt.args.value)
		})
	}
}

func TestMemCache_clearExpiredTime(t *testing.T) {
	type fields struct {
		maxMemorySize                int64
		maxMemorySizeString          string
		currentMemorySize            int64
		values                       map[string]*memCacheValue
		locker                       sync.RWMutex
		clearExpiredItemTimeInterval time.Duration
	}
	var tests []struct {
		name   string
		fields fields
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemCache{
				maxMemorySize:                tt.fields.maxMemorySize,
				maxMemorySizeString:          tt.fields.maxMemorySizeString,
				currentMemorySize:            tt.fields.currentMemorySize,
				values:                       tt.fields.values,
				locker:                       tt.fields.locker,
				clearExpiredItemTimeInterval: tt.fields.clearExpiredItemTimeInterval,
			}
			m.clearExpiredTime()
		})
	}
}

func TestMemCache_del(t *testing.T) {
	type fields struct {
		maxMemorySize                int64
		maxMemorySizeString          string
		currentMemorySize            int64
		values                       map[string]*memCacheValue
		locker                       sync.RWMutex
		clearExpiredItemTimeInterval time.Duration
	}
	type args struct {
		key string
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemCache{
				maxMemorySize:                tt.fields.maxMemorySize,
				maxMemorySizeString:          tt.fields.maxMemorySizeString,
				currentMemorySize:            tt.fields.currentMemorySize,
				values:                       tt.fields.values,
				locker:                       tt.fields.locker,
				clearExpiredItemTimeInterval: tt.fields.clearExpiredItemTimeInterval,
			}
			m.del(tt.args.key)
		})
	}
}

func TestMemCache_get(t *testing.T) {
	type fields struct {
		maxMemorySize                int64
		maxMemorySizeString          string
		currentMemorySize            int64
		values                       map[string]*memCacheValue
		locker                       sync.RWMutex
		clearExpiredItemTimeInterval time.Duration
	}
	type args struct {
		key string
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   *memCacheValue
		want1  bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemCache{
				maxMemorySize:                tt.fields.maxMemorySize,
				maxMemorySizeString:          tt.fields.maxMemorySizeString,
				currentMemorySize:            tt.fields.currentMemorySize,
				values:                       tt.fields.values,
				locker:                       tt.fields.locker,
				clearExpiredItemTimeInterval: tt.fields.clearExpiredItemTimeInterval,
			}
			got, got1 := m.get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewMemCache(t *testing.T) {
	var tests []struct {
		name string
		want *MemCache
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

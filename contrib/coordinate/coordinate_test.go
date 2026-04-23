package coordinate

import (
	"fmt"
	"math"
	"testing"
)

func TestBaidu2AMap(t *testing.T) {
	tests := []struct {
		name    string
		lng     float64
		lat     float64
		wantLng float64
		wantLat float64
	}{
		{
			name:    "天安门广场-百度坐标",
			lng:     116.404,
			lat:     39.915,
			wantLng: 116.397627,
			wantLat: 39.909006,
		},
		{
			name:    "上海东方明珠-百度坐标",
			lng:     121.499706,
			lat:     31.239652,
			wantLng: 121.493421,
			wantLat: 31.234053,
		},
		{
			name:    "上海同普大厦",
			lng:     121.38,
			lat:     31.24,
			wantLng: 121.493421,
			wantLat: 31.234053,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLng, gotLat := Baidu2AMap(tt.lng, tt.lat)
			// 允许一定误差
			if math.Abs(gotLng-tt.wantLng) > 0.0001 {
				t.Errorf("Baidu2AMap() lng = %v, want %v", gotLng, tt.wantLng)
			}
			if math.Abs(gotLat-tt.wantLat) > 0.0001 {
				t.Errorf("Baidu2AMap() lat = %v, want %v", gotLat, tt.wantLat)
			}
		})
	}
}

func TestAMap2Baidu(t *testing.T) {
	tests := []struct {
		name    string
		lng     float64
		lat     float64
		wantLng float64
		wantLat float64
	}{
		{
			name:    "天安门广场-高德坐标",
			lng:     116.397627,
			lat:     39.909006,
			wantLng: 116.404,
			wantLat: 39.915,
		},
		{
			name:    "上海东方明珠-高德坐标",
			lng:     121.493421,
			lat:     31.234053,
			wantLng: 121.499706,
			wantLat: 31.239652,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLng, gotLat := AMap2Baidu(tt.lng, tt.lat)
			// 允许一定误差
			if math.Abs(gotLng-tt.wantLng) > 0.0001 {
				t.Errorf("AMap2Baidu() lng = %v, want %v", gotLng, tt.wantLng)
			}
			if math.Abs(gotLat-tt.wantLat) > 0.0001 {
				t.Errorf("AMap2Baidu() lat = %v, want %v", gotLat, tt.wantLat)
			}
		})
	}
}

func TestWGS842GCJ02(t *testing.T) {
	tests := []struct {
		name    string
		lng     float64
		lat     float64
		wantLng float64
		wantLat float64
	}{
		{
			name:    "天安门广场-GPS坐标",
			lng:     116.391281,
			lat:     39.907084,
			wantLng: 116.397627,
			wantLat: 39.909006,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLng, gotLat := WGS842GCJ02(tt.lng, tt.lat)
			// 允许一定误差
			if math.Abs(gotLng-tt.wantLng) > 0.0001 {
				t.Errorf("WGS842GCJ02() lng = %v, want %v", gotLng, tt.wantLng)
			}
			if math.Abs(gotLat-tt.wantLat) > 0.0001 {
				t.Errorf("WGS842GCJ02() lat = %v, want %v", gotLat, tt.wantLat)
			}
		})
	}
}

func TestGCJ022WGS84(t *testing.T) {
	tests := []struct {
		name    string
		lng     float64
		lat     float64
		wantLng float64
		wantLat float64
	}{
		{
			name:    "天安门广场-高德坐标",
			lng:     116.397627,
			lat:     39.909006,
			wantLng: 116.391281,
			wantLat: 39.907084,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLng, gotLat := GCJ022WGS84(tt.lng, tt.lat)
			// 允许一定误差
			if math.Abs(gotLng-tt.wantLng) > 0.0001 {
				t.Errorf("GCJ022WGS84() lng = %v, want %v", gotLng, tt.wantLng)
			}
			if math.Abs(gotLat-tt.wantLat) > 0.0001 {
				t.Errorf("GCJ022WGS84() lat = %v, want %v", gotLat, tt.wantLat)
			}
		})
	}
}

func TestBaidu2WGS84(t *testing.T) {
	tests := []struct {
		name    string
		lng     float64
		lat     float64
		wantLng float64
		wantLat float64
	}{
		{
			name:    "天安门广场-百度坐标",
			lng:     116.404,
			lat:     39.915,
			wantLng: 116.391281,
			wantLat: 39.907084,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLng, gotLat := Baidu2WGS84(tt.lng, tt.lat)
			// 允许一定误差
			if math.Abs(gotLng-tt.wantLng) > 0.0001 {
				t.Errorf("Baidu2WGS84() lng = %v, want %v", gotLng, tt.wantLng)
			}
			if math.Abs(gotLat-tt.wantLat) > 0.0001 {
				t.Errorf("Baidu2WGS84() lat = %v, want %v", gotLat, tt.wantLat)
			}
		})
	}
}

func TestWGS842Baidu(t *testing.T) {
	tests := []struct {
		name    string
		lng     float64
		lat     float64
		wantLng float64
		wantLat float64
	}{
		{
			name:    "天安门广场-GPS坐标",
			lng:     116.391281,
			lat:     39.907084,
			wantLng: 116.404,
			wantLat: 39.915,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLng, gotLat := WGS842Baidu(tt.lng, tt.lat)
			// 允许一定误差
			if math.Abs(gotLng-tt.wantLng) > 0.0001 {
				t.Errorf("WGS842Baidu() lng = %v, want %v", gotLng, tt.wantLng)
			}
			if math.Abs(gotLat-tt.wantLat) > 0.0001 {
				t.Errorf("WGS842Baidu() lat = %v, want %v", gotLat, tt.wantLat)
			}
		})
	}
}

// TestCoordinateReversibility 测试坐标转换的可逆性
func TestCoordinateReversibility(t *testing.T) {
	type testCase struct {
		name          string
		convertFunc1  func(float64, float64) (float64, float64)
		convertFunc2  func(float64, float64) (float64, float64)
		originalLng   float64
		originalLat   float64
		shouldFailOut bool // 是否预期境外坐标测试失败
	}

	tests := []testCase{
		{
			name:         "BD09 <-> GCJ02 可逆性",
			convertFunc1: Baidu2AMap,
			convertFunc2: AMap2Baidu,
			originalLng:  116.404,
			originalLat:  39.915,
		},
		{
			name:         "WGS84 <-> GCJ02 可逆性",
			convertFunc1: WGS842GCJ02,
			convertFunc2: GCJ022WGS84,
			originalLng:  116.391281,
			originalLat:  39.907084,
		},
		{
			name:         "BD09 <-> WGS84 可逆性",
			convertFunc1: Baidu2WGS84,
			convertFunc2: WGS842Baidu,
			originalLng:  116.404,
			originalLat:  39.915,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 第一步转换
			lng1, lat1 := tt.convertFunc1(tt.originalLng, tt.originalLat)
			// 第二步反向转换
			lng2, lat2 := tt.convertFunc2(lng1, lat1)

			// 验证是否恢复到原始坐标（允许小误差）
			if math.Abs(lng2-tt.originalLng) > 0.00001 {
				t.Errorf("%s: 经度恢复失败: got %v, want %v", tt.name, lng2, tt.originalLng)
			}
			if math.Abs(lat2-tt.originalLat) > 0.00001 {
				t.Errorf("%s: 纬度恢复失败: got %v, want %v", tt.name, lat2, tt.originalLat)
			}
		})
	}
}

// TestOutOfChina 测试境外坐标判断
func TestOutOfChina(t *testing.T) {
	tests := []struct {
		name     string
		lng      float64
		lat      float64
		expected bool
	}{
		{
			name:     "天安门广场-国内",
			lng:      116.404,
			lat:      39.915,
			expected: false,
		},
		{
			name:     "美国纽约-境外",
			lng:      -74.006,
			lat:      40.7128,
			expected: true,
		},
		{
			name:     "日本东京-境外",
			lng:      139.6917,
			lat:      35.6895,
			expected: true,
		},
		{
			name:     "边界值-西边界",
			lng:      72.004,
			lat:      40,
			expected: false,
		},
		{
			name:     "边界值-东边界",
			lng:      137.8347,
			lat:      40,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := outOfChina(tt.lng, tt.lat)
			if result != tt.expected {
				t.Errorf("outOfChina(%v, %v) = %v, want %v", tt.lng, tt.lat, result, tt.expected)
			}
		})
	}
}

// ExampleBaidu2AMap 示例函数
func ExampleBaidu2AMap() {
	// 将百度坐标转换为高德坐标
	lng, lat := Baidu2AMap(116.404, 39.915)
	fmt.Printf("高德坐标: 经度=%.6f, 纬度=%.6f\n", lng, lat)
	// Output: 高德坐标: 经度=116.397627, 纬度=39.909006
}

package coordinate

import (
	"math"
)

// 坐标系常量
const (
	xPI = 3.14159265358979324 * 3000.0 / 180.0
	PI  = 3.1415926535897932384626
	a   = 6378245.0
	ee  = 0.00669342162296594323
)

// Baidu2AMap BD09坐标系转换为GCJ02坐标系
// 百度地图 -> 高德地图、腾讯地图等
func Baidu2AMap(lng, lat float64) (float64, float64) {
	if outOfChina(lng, lat) {
		return lng, lat
	}

	x := lng - 0.0065
	y := lat - 0.006
	z := math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*xPI)
	theta := math.Atan2(y, x) - 0.000003*math.Cos(x*xPI)

	gcjLng := z * math.Cos(theta)
	gcjLat := z * math.Sin(theta)

	return gcjLng, gcjLat
}

// AMap2Baidu GCJ02坐标系转换为BD09坐标系
// 高德地图、腾讯地图等 -> 百度地图
func AMap2Baidu(lng, lat float64) (float64, float64) {
	if outOfChina(lng, lat) {
		return lng, lat
	}

	z := math.Sqrt(lng*lng+lat*lat) + 0.00002*math.Sin(lat*xPI)
	theta := math.Atan2(lat, lng) + 0.000003*math.Cos(lng*xPI)

	bdLng := z*math.Cos(theta) + 0.0065
	bdLat := z*math.Sin(theta) + 0.006

	return bdLng, bdLat
}

// WGS842GCJ02 WGS84坐标系转换为GCJ02坐标系
// GPS坐标 -> 高德地图、腾讯地图等
func WGS842GCJ02(lng, lat float64) (float64, float64) {
	if outOfChina(lng, lat) {
		return lng, lat
	}

	dLng, dLat := delta(lng, lat)
	return lng + dLng, lat + dLat
}

// GCJ022WGS84 GCJ02坐标系转换为WGS84坐标系
// 高德地图、腾讯地图等 -> GPS坐标
func GCJ022WGS84(lng, lat float64) (float64, float64) {
	if outOfChina(lng, lat) {
		return lng, lat
	}

	dLng, dLat := delta(lng, lat)
	return lng - dLng, lat - dLat
}

// delta 计算 WGS84 与 GCJ02 之间的偏差值
func delta(lng, lat float64) (float64, float64) {
	dLat := transformLat(lng-105.0, lat-35.0)
	dLng := transformLng(lng-105.0, lat-35.0)
	radLat := lat / 180.0 * PI
	magic := math.Sin(radLat)
	magic = 1 - ee*magic*magic
	sqrtMagic := math.Sqrt(magic)

	dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * PI)
	dLng = (dLng * 180.0) / (a / sqrtMagic * math.Cos(radLat) * PI)
	return dLng, dLat
}

// Baidu2WGS84 BD09坐标系转换为WGS84坐标系
// 百度地图 -> GPS坐标
func Baidu2WGS84(lng, lat float64) (float64, float64) {
	// 先转换为GCJ02
	gcjLng, gcjLat := Baidu2AMap(lng, lat)
	// 再转换为WGS84
	return GCJ022WGS84(gcjLng, gcjLat)
}

// WGS842Baidu WGS84坐标系转换为BD09坐标系
// GPS坐标 -> 百度地图
func WGS842Baidu(lng, lat float64) (float64, float64) {
	// 先转换为GCJ02
	gcjLng, gcjLat := WGS842GCJ02(lng, lat)
	// 再转换为BD09
	return AMap2Baidu(gcjLng, gcjLat)
}

// transformLat 纬度转换
func transformLat(lng, lat float64) float64 {
	ret := -100.0 + 2.0*lng + 3.0*lat + 0.2*lat*lat + 0.1*lng*lat + 0.2*math.Sqrt(math.Abs(lng))
	ret += (20.0*math.Sin(6.0*lng*PI) + 20.0*math.Sin(2.0*lng*PI)) * 2.0 / 3.0
	ret += (20.0*math.Sin(lat*PI) + 40.0*math.Sin(lat/3.0*PI)) * 2.0 / 3.0
	ret += (160.0*math.Sin(lat/12.0*PI) + 320.0*math.Sin(lat*PI/30.0)) * 2.0 / 3.0
	return ret
}

// transformLng 经度转换
func transformLng(lng, lat float64) float64 {
	ret := 300.0 + lng + 2.0*lat + 0.1*lng*lng + 0.1*lng*lat + 0.1*math.Sqrt(math.Abs(lng))
	ret += (20.0*math.Sin(6.0*lng*PI) + 20.0*math.Sin(2.0*lng*PI)) * 2.0 / 3.0
	ret += (20.0*math.Sin(lng*PI) + 40.0*math.Sin(lng/3.0*PI)) * 2.0 / 3.0
	ret += (150.0*math.Sin(lng/12.0*PI) + 300.0*math.Sin(lng/30.0*PI)) * 2.0 / 3.0
	return ret
}

// outOfChina 判断坐标是否在中国境外
func outOfChina(lng, lat float64) bool {
	if lng < 72.004 || lng > 137.8347 {
		return true
	}
	if lat < 0.8293 || lat > 55.8271 {
		return true
	}
	return false
}

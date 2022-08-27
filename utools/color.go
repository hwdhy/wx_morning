package utools

import "math/rand"

// GetColor 生成随机颜色
func GetColor() string {
	colors := []byte("0123456789ABCDEF")

	color := "#"
	for i := 0; i < 6; i++ {
		color += string(colors[rand.Intn(16)])
	}
	return color
}

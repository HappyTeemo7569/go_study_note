package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

//画笔
type Brush struct {
	gray *image.Gray
}

//初始画笔
func initBrush() *Brush {
	// 图片大小
	const size = 300
	// 根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))

	// 遍历每个像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			// 填充为白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	return &Brush{gray: pic}
}

//画
func (b *Brush) draw(x, y int) {
	b.gray.SetGray(x, y, color.Gray{0})
}

//保存
func (b *Brush) save() {
	// 创建文件
	file, err := os.Create("./move.png")

	if err != nil {
		log.Fatal(err)
	}
	// 使用png格式将数据写入文件
	png.Encode(file, b.gray) //将image信息写入文件中

	// 关闭文件
	file.Close()
}

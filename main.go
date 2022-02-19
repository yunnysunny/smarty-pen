package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/smarty-pen/drawer/config"
	"github.com/smarty-pen/drawer/impl"
	"io/ioutil"
)

func main() {
	var (
		filename        string
		output          string
		background      string
		disablePressure int
		drawer          impl.BaseDrawer
	)

	flag.StringVar(&filename, "input", "", "数据文件的完整路径")
	flag.StringVar(&background, "background", "", "背景图片的完整路径")
	flag.StringVar(&output, "output", "result.png", "生成图片的完整路径")
	flag.IntVar(&disablePressure, "disable-pressure", 0, "是否禁用压感，默认为启用，设置为1代表禁用")
	flag.Parse()
	fmt.Println("开始处理数据文件", filename)
	dots, err := readData(filename)
	if err != nil {
		fmt.Println(filename, "读取文件失败", err)
		return
	}
	if disablePressure == 1 {
		drawer = &impl.SimpleDrawer{}
	} else {
		drawer = &impl.PressureDrawer{}
	}
	err = drawer.DoDraw(dots, background, output)
	if err != nil {
		fmt.Println(filename, "绘制图片失败", err)
		return
	}
	fmt.Println(filename, "生成图片完成")
}

func readData(filename string) (*[]config.Info, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	dots := make([]config.Info, 0)
	err = json.Unmarshal(body, &dots)
	if err != nil {
		return nil, err
	}
	return &dots, nil
}

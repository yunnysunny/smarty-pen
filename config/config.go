package config

const b5_width = 182; // B5 纸的码点宽度
const b5_height = 256; // B5纸的码点高度
const x_codepoint_size = 1.524; // 横坐标码点的大小
const y_codepoint_size = 1.524; // 纵坐标码点的大小
const MAX_X float64 = float64( b5_width ) / x_codepoint_size
const MAX_Y float64 = float64( b5_height ) / y_codepoint_size
const IMAGE_WIDTH  = 600
const IMAGE_HEIGHT  = 844 
const TRANSFORM_SCALE = 120.00

type Info struct {
	D   int64   `json:"d"`
	Dot InfoDot `json:"dot"`
	S   int     `json:"s"`
	U   int     `json:"u"`
	W   float64 `json:"w"`
	X   float64 `json:"x"`
	Y   float64 `json:"y"`
}

type InfoDot struct {
	BookID    int     `json:"BookID"`
	Counter   int     `json:"Counter"`
	OwnerID   int     `json:"OwnerID"`
	PageID    int     `json:"PageID"`
	SectionID int     `json:"SectionID"`
	AbX       float64 `json:"ab_x"`
	AbY       float64 `json:"ab_y"`
	Angle     int     `json:"angle"`
	Color     int     `json:"color"`
	Force     int     `json:"force"`
	Fx        int     `json:"fx"`
	Fy        int     `json:"fy"`
	Timelong  int64   `json:"timelong"`
	Type      string  `json:"type"`
	X         int     `json:"x"`
	Y         int     `json:"y"`
}

// type PointPair struct {
// 	x1 float64
// 	y1 float64
// 	width1 float64
// 	x2 float64
// 	y2 float64
// 	width2 float64
// }
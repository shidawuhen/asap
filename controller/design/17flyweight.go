/**
 * @Author: Jason Pang
 * @Description: 棋子类，有文案、颜色、规则，这三种不变属性
 */
package design

import "fmt"

type Piece struct {
	text  string
	color string
	rule  string
}

/**
 * @Author: Jason Pang
 * @Description: 棋子信息说明
 * @receiver p
 * @return string
 */
func (p *Piece) String() string {
	return fmt.Sprintf("%s,颜色为%s,规则为%s", p.text, p.color, p.rule)
}

/**
 * @Author: Jason Pang
 * @Description: 棋子在棋盘位置
 */
type Pos struct {
	x int64
	y int64
}

/**
 * @Author: Jason Pang
 * @Description: 游戏中的棋子
 */
type GamePiece struct {
	piece   *Piece //棋子指针
	pos     Pos    //棋子位置
	ownerId int64  //玩家ID
	roomId  int64  //房间ID
}

/**
 * @Author: Jason Pang
 * @Description: 游戏中的棋子说明
 * @receiver g
 * @return string
 */
func (g *GamePiece) String() string {
	return fmt.Sprintf("%s位置为(%d,%d)", g.piece, g.pos.x, g.pos.y)
}

/**
 * @Author: Jason Pang
 * @Description: 棋子工厂，包含32颗棋子信息
 */
type PieceFactory struct {
	pieces []*Piece
}

/**
 * @Author: Jason Pang
 * @Description: 创建棋子。棋子的信息都是不变的
 * @receiver f
 */
func (f *PieceFactory) CreatePieces() {
	f.pieces = make([]*Piece, 32)
	f.pieces[0] = &Piece{
		text:  "兵",
		color: "红",
		rule:  "过河前只能一步一步前进，过河后只能一步一步前进或者左右移",
	}
	f.pieces[1] = &Piece{
		text:  "兵",
		color: "黑",
		rule:  "过河前只能一步一步前进，过河后只能一步一步前进或者左右移",
	}
	//todo 创建其它棋子。此处可以使用配置文件创建，能方便一些。系统中可以设置一个规则引擎，控制棋子运动。
}

/**
 * @Author: Jason Pang
 * @Description: 获取棋子信息
 * @receiver f
 * @param id
 * @return *Piece
 */
func (f *PieceFactory) GetPiece(id int64) *Piece {
	return f.pieces[id]
}

/**
 * @Author: Jason Pang
 * @Description: 初始化棋盘
 * @param roomId
 * @param u1
 * @param u2
 */
func InitBoard(roomId int64, u1 int64, u2 int64, factory *PieceFactory) {
	fmt.Printf("创建房间%d,玩家为%d和%d \n", roomId, u1, u2)
	fmt.Println("初始化棋盘")

	fmt.Printf("玩家%d的棋子为 \n", u1)
	piece := &GamePiece{
		piece:   factory.GetPiece(0),
		pos:     Pos{1, 1},
		roomId:  roomId,
		ownerId: u1,
	}
	fmt.Println(piece)

	fmt.Printf("玩家%d的棋子为 \n", u2)
	piece2 := &GamePiece{
		piece:   factory.GetPiece(1),
		pos:     Pos{16, 1},
		roomId:  roomId,
		ownerId: u2,
	}
	fmt.Println(piece2)
}
func flyweightMain() {
	factory := &PieceFactory{}
	factory.CreatePieces()
	InitBoard(1, 66, 88, factory)
}

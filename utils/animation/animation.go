package animation

import "github.com/nsf/termbox-go"

var Characters = [][][]rune{
	{
		{' ', 'O', ' '},
		{'/', '|', '\\'},
		{'/', ' ', '\\'},
	},
	{
		{' ', 'O', ' '},
		{'/', '|', '\\'},
		{' ', '/', '\\'},
	},
	{
		{' ', 'O', ' '},
		{'/', '|', '/'},
		{'/', ' ', ' '},
	},
	{
		{' ', 'O', ' '},
		{' ', '|', '\\'},
		{'/', ' ', '\\'},
	},
	{
		{' ', 'O', ' '},
		{'/', '|', ' '},
		{'/', ' ', '\\'},
	},
}

var Wave = [][]rune{
	{'~', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', '~', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', '~', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', '~', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', '~', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', '~', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', '~', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', '~', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '~', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '~'},
}

// DrawCharacter отрисовывает танцующего персонажа на экране.
func DrawCharacter(character [][]rune, posX, posY int) {
	for y, row := range character {
		for x, ch := range row {
			termbox.SetCell(posX+x, posY+y, ch, termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

// DrawWave отрисовывает анимацию волны на экране.
func DrawWave(wave [][]rune, step int, width, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			char := wave[(y+step)%len(wave)][x%len(wave[0])]
			termbox.SetCell(x, y, char, termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

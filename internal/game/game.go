package game

import (
	"errors"
	"fmt"
)

type Map struct {
	height           int
	width            int
	imageHeight      int
	imageWidth       int
	image            [][]rune
	xRune            rune
	oRune            rune
	winCond          int
	LastPlayerSymbol rune
}

func InitMap(size, winCond int) (*Map, error) {
	if size < 3 {
		return nil, errors.New("map width and height should be greater 2")
	}

	gameMap := Map{}
	gameMap.height = size
	gameMap.width = size
	gameMap.imageHeight = size
	gameMap.imageWidth = size + size - 1
	gameMap.xRune = 'x'
	gameMap.oRune = 'o'
	gameMap.winCond = winCond

	gameMap.image = make([][]rune, gameMap.imageHeight, gameMap.height)

	for i := 0; i < cap(gameMap.image); i++ {
		gameMap.image[i] = make([]rune, gameMap.imageWidth, gameMap.imageWidth)
		for j := 0; j < cap(gameMap.image[i]); j++ {
			if j%2 == 0 {
				gameMap.image[i][j] = '_'
			} else {
				gameMap.image[i][j] = '|'
			}
		}
	}
	return &gameMap, nil
}

func (m *Map) PrintMap() {
	for i := 0; i < cap(m.image); i++ {
		for j := 0; j < cap(m.image[i]); j++ {
			fmt.Print(string(m.image[i][j]))
		}
		fmt.Println()
	}
}

func (m *Map) GetMapForResponse() []string {
	result := make([]string, m.imageHeight, m.imageHeight)
	for i := 0; i < m.imageHeight; i++ {
		for j := 0; j < m.imageWidth; j++ {
			result[i] += string(m.image[i][j])
		}
	}
	return result
}

func (m *Map) Move(yCord int, xCord int, symb rune) (bool, error) {
	if err := m.isValidMove(yCord, xCord); err != nil {
		return false, err
	}
	m.image[getImageYCord(yCord)][getImageXCord(xCord)] = symb
	m.LastPlayerSymbol = symb
	return m.isWin(yCord, xCord, symb), nil
}

func (m *Map) isWin(yCord int, xCord int, symb rune) bool {
	//check colum
	counter := 0
	for i := 0; i < m.height; i++ {
		if m.isRuneEquals(i, xCord, symb) {
			counter++
		}
	}
	if counter == m.winCond {
		return true
	}

	//check row
	counter = 0
	for i := 0; i < m.width; i++ {
		if m.isRuneEquals(yCord, i, symb) {
			counter++
		}
	}
	if counter == m.winCond {
		return true
	}

	//todo: check diagonal
	counter = 0
	return false
}

func (m *Map) isRuneEquals(yCord, xCord int, symb rune) bool {
	return m.image[getImageYCord(yCord)][getImageXCord(xCord)] == symb
}

func (m *Map) isValidMove(yCord, xCord int) error {
	if yCord < 0 || xCord < 0 || yCord > m.height || xCord > m.width {
		return errors.New("coordinates in out of map bounds")
	}
	if m.image[getImageYCord(yCord)][getImageXCord(xCord)] != '_' {
		return errors.New("occupied cell")
	}
	return nil
}

func getImageYCord(yCord int) int {
	return yCord
}

func getImageXCord(xCord int) int {
	return xCord * 2
}

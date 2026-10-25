package game

import "fmt"

type InvalidBoardSizeError error

func newInvalidBoardSizeError(m string) InvalidBoardSizeError {
	return fmt.Errorf("invalid board size: %s", m)
}

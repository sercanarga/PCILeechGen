package flash

import (
	"errors"
	"fmt"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/board"
)

var (
	ErrMissingBitstream  = errors.New("flash: missing bitstream")
	ErrUnsupportedBoard  = errors.New("flash: unsupported board")
	ErrUnsupportedTarget = errors.New("flash: unsupported target")
	ErrUnsupportedTool   = errors.New("flash: unsupported tool")
)

type Request struct {
	Board     *board.Board
	Bitstream string
}

type Command struct {
	Args []string
}

func (c Command) String() string {
	return strings.Join(c.Args, " ")
}

func OpenFPGALoaderCommand(req Request) (Command, error) {
	if strings.TrimSpace(req.Bitstream) == "" {
		return Command{}, ErrMissingBitstream
	}
	if req.Board == nil || req.Board.Flash == nil {
		return Command{}, fmt.Errorf("board %s has no flash metadata: %w", boardName(req.Board), ErrUnsupportedBoard)
	}
	meta := req.Board.Flash
	if meta.Tool != "openFPGALoader" {
		return Command{}, fmt.Errorf("board %s flash tool %q: %w", req.Board.Name, meta.Tool, ErrUnsupportedTool)
	}
	if meta.OpenFPGALoaderBoard == "" {
		return Command{}, fmt.Errorf("board %s missing openFPGALoader board id: %w", req.Board.Name, ErrUnsupportedBoard)
	}

	args := []string{"openFPGALoader", "-b", meta.OpenFPGALoaderBoard}
	if meta.Cable != "" {
		args = append(args, "-c", meta.Cable)
	}
	switch meta.Target {
	case "", "bitstream":
		args = append(args, req.Bitstream)
	case "flash":
		args = append(args, "-f", req.Bitstream)
	default:
		return Command{}, fmt.Errorf("board %s flash target %q: %w", req.Board.Name, meta.Target, ErrUnsupportedTarget)
	}

	return Command{Args: args}, nil
}

func boardName(b *board.Board) string {
	if b == nil || b.Name == "" {
		return "<unknown>"
	}
	return b.Name
}

package history

import t "github.com/abroudoux/dk/internal/types"

func HistoryMode(ctx t.Context, cli t.Client) error {
	err := readHistoryFile()
	if err != nil {
		return err
	}

	return nil
}

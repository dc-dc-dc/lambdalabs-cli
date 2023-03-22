package cmds

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func prettyJson(data []byte) error {
	var raw bytes.Buffer
	if err := json.Indent(&raw, data, "", "\t"); err != nil {
		return err
	}
	fmt.Println(raw.String())
	return nil
}

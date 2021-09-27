package direct

import (
	"encoding/json"

	"github.com/trevor403/gostream/pkg/input"
)

func Handle(data []byte) {
	raw := input.RawEvent{}
	_ = json.Unmarshal(data, &raw)

	switch raw.Type {
	case input.KeyEventType:
		ev := input.KeyEvent{}
		json.Unmarshal(data, &ev)
		HandleKey(ev)
	case input.MouseEventType:
		ev := input.MouseEvent{}
		json.Unmarshal(data, &ev)
		HandlePtr(ev)
	}
}

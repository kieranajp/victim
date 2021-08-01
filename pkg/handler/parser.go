package handler

import (
	"encoding/json"
	"fmt"

	"github.com/slack-go/slack"
)

func dumpMap(space string, m map[string]interface{}) {
	for k, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			fmt.Printf("{ \"%v\": \n", k)
			dumpMap(space+"\t", mv)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v %v : %v\n", space, k, v)
		}
	}
}

func InteractionCallbackParser(payload []byte) (s slack.InteractionCallback, err error) {
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(payload, &jsonMap)
	if err != nil {
		return s, err
	}

	dumpMap("", jsonMap)

	return s, nil
}

package messenger

import "encoding/json"

func Serialize(message Message) ([]byte, error) {
	return json.Marshal(message)
}

func Deserialize(data []byte) (Message, error) {
	var message Message
	err := json.Unmarshal(data, &message)
	return message, err
}

package messenger

func Publish[T any](sender chan<- []byte, body T) error {
	message := NewMessage(body)
	data, err := Serialize(message)
	if err != nil {
		return err
	}
	sender <- data
	return nil
}

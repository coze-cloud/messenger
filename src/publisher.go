package messenger

func Publish(sender chan<- []byte, message Message) error {
	data, err := Serialize(message)
	if err != nil {
		return err
	}
	sender <- data
	return nil
}

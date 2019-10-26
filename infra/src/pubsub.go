package main

func pubsub(subj string, data string) error {
	go subscribe(subj)
	go publish(subj, data)
	/*reader := bufio.NewReader(os.Stdin)
	for {
		data, _ := reader.ReadBytes('\n')
		publish(subj, bytes.Trim(data, "\n"))
	}*/
	return nil
}

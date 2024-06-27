package deck

func AddFromInteractivePrompt() error {
	println("adding deck")
	return nil
}

func AddFromFlags(name string) error {
	println("adding deck name" + name)
	return nil
}

func List() error {
	println("Listing decks")
	return nil
}

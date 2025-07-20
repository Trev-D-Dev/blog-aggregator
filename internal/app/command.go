package app

func CreateCommand(name string, args []string) command {
	c := command{
		name: name,
		args: args,
	}

	return c
}

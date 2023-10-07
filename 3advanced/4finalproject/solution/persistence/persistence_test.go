package persistence

func setup() {
	DB = newDatabase()
}

func cleanup() {
	closeDatabase(DB)
}

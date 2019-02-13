package internal

type Storage struct {
	DataDir string
}

func NewStorage(dataDir string) *Storage {
	return &Storage{
		DataDir: dataDir,
	}
}

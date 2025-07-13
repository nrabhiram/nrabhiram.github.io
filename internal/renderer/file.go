package renderer

type File struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Path     string `json:"path"`
}

func MakeFile(name string, location string, path string) *File {
	file := File{
		name,
		location,
		path,
	}
	return &file
}

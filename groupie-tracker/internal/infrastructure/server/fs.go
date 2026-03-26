package server

type FileServer struct {
	Path string
	Url string
}

func NewFileServer(path string, url string) *FileServer {
	return &FileServer{
		Path: path,
		Url: url,
	}
}
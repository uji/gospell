package gospell

func FilePerser(path string) {
}

func FilesPerser(paths []string) {
	for _, path := range paths {
		FilePerser(path)
	}
}

func DirPerser(path string) {
}

func DirsPerser(paths []string) {
	for _, path := range paths {
		DirPerser(path)
	}
}

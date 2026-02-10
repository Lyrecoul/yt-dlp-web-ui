package sys

import (
	"os"
	"path/filepath"

	"github.com/marcopiovanello/yt-dlp-web-ui/v3/server/config"
	"github.com/marcopiovanello/yt-dlp-web-ui/v3/server/internal"
)

// Build a directory tree started from the specified path using DFS.
// Then return the flattened tree represented as a list.
func DirectoryTree() (*[]string, error) {
	type Node struct {
		path     string
		children []Node
	}

	var (
		rootPath = config.Instance().DownloadPath

		stack     = internal.NewStack[Node]()
		flattened = make([]string, 0)
	)

	stack.Push(Node{path: rootPath})
	flattened = append(flattened, rootPath)

	for stack.IsNotEmpty() {
		current := stack.Pop().Value

		children, err := os.ReadDir(current.path)
		if err != nil {
			return nil, err
		}
		for _, entry := range children {
			if entry.IsDir() {
				childPath := filepath.Join(current.path, entry.Name())
				stack.Push(Node{path: childPath})
				flattened = append(flattened, childPath)
			}
		}
	}
	return &flattened, nil
}

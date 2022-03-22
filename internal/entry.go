package internal

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

// Entry is what we looking for: a project dir with a readme and stuff
type Entry struct {
	Dir     string
	Excerpt string
	Git     string
}

type EntryFunc func(e *Entry, err error)

func Walk(fsys fs.FS, depth int, fn EntryFunc) {
	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// if we got an error, lets skip the dir
			// log.Fatal(err)
			fn(nil, fmt.Errorf("cannot explore directory %s: %w", path, err))
			return fs.SkipDir
		}
		if d.IsDir() {
			// Limit depth perhaps
			if -1 < depth && strings.Count(path, string(os.PathSeparator)) > depth {
				return fs.SkipDir
			}
			// Skip typical vendor dirs
			if d.Name() == "node_modules" || d.Name() == "_archive" {
				return fs.SkipDir
			}

			return nil
		}

		// so we got a file...

		if readmeRex.Match([]byte(d.Name())) {
			b, err := fs.ReadFile(fsys, path)
			if err != nil {
				return err
			}
			e := handleReadme(path, b)
			e.Dir = filepath.Dir(path)
			fn(e, nil)
			return fs.SkipDir
		}

		if filepath.Ext(d.Name()) == ".tar" {
			b, err := fs.ReadFile(fsys, path)
			if err != nil {
				return err
			}
			tr := tar.NewReader(bytes.NewBuffer(b))
			for {
				hdr, err := tr.Next()
				if err == io.EOF {
					break // End of archive
				}
				if err != nil {
					return err
				}
				if readmeRex.Match([]byte(filepath.Base(hdr.Name))) {
					b2 := bytes.NewBuffer([]byte{})
					if _, err := io.Copy(b2, tr); err != nil {
						return err
					}
					e := handleReadme(hdr.Name, b2.Bytes())
					e.Dir = path
					fn(e, nil)
					return fs.SkipDir
				}
			}
		}

		return nil
	})
}

var readmeRex = regexp.MustCompile("^(?i)readme")

func handleReadme(path string, content []byte) *Entry {
	e := &Entry{}
	if strings.HasSuffix(path, ".md") {
		ast := goldmark.DefaultParser().Parse(text.NewReader(content))

		for n := ast.FirstChild(); n != nil; n = n.NextSibling() {
			if n.Kind().String() == "Paragraph" {
				e.Excerpt = fmt.Sprintln(string(n.Text(content)))
				break
			}
		}
	} else {
		lines := strings.SplitAfterN(string(content), "\n\n", 2)
		e.Excerpt = fmt.Sprintln(strings.Join(lines[:len(lines)-1], "\n"))
	}

	e.Excerpt = strings.TrimSpace(e.Excerpt)

	return e
}

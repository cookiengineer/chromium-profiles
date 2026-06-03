package extensions

import "io"
import "io/fs"
import "os"
import "path/filepath"
import "strings"

func CopyExtension(name string, target_root string) error {

	source_root := strings.TrimSpace(strings.ToLower(name))

	return fs.WalkDir(FS, source_root, func(source_path string, entry fs.DirEntry, err error) error {

		if err == nil {

			relative_path, err1 := filepath.Rel(source_root, source_path)

			if err1 == nil {

				target_path := filepath.Join(target_root, relative_path)

				if entry.IsDir() {

					return os.MkdirAll(target_path, 0755)

				} else {

					source, err2 := FS.Open(source_path)

					if err2 == nil {

						defer source.Close()

						target, err3 := os.Create(target_path)

						if err3 == nil {

							defer target.Close()

							_, err4 := io.Copy(target, source)

							if err4 == nil {
								return nil
							} else {
								return err4
							}

						} else {
							return err3
						}

					} else {
						return err2
					}

				}

			} else {
				return err1
			}

		} else {
			return err
		}

	})

}

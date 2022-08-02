package backup

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Archiver interface {
	DestFmt() func(int64) string
	Archive(src, dest string) error
}

type zipper struct{}

// ZIPはファイルの圧縮とその解除にZIP形式を利用するArchiver
var ZIP Archiver = (*zipper)(nil)

func (z *zipper) DestFmt() func(int64) string {
	return func(i int64) string {
		return fmt.Sprintf("%d.zip", i)
	}
}

func (z *zipper) Archive(src, dest string) error {
	// 保存先のディレクトリが存在することを確認
	if err := os.MkdirAll(filepath.Dir(dest), 0777); err != nil {
		return err
	}
	// destで指定されたパスにファイルを新規作成
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	w := zip.NewWriter(out)
	defer w.Close()
	// filepath.Waklは再起的に処理を行い、フォルダー構造が深くてもすべてのファイルを検出できる
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil // skip
		}
		if err != nil {
			return err
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		// 圧縮ファイルを新規作成
		f, err := w.Create(path)
		if err != nil {
			return err
		}
		io.Copy(f, in)
		return nil
	})
}

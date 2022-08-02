package backup

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func DirHash(path string) (string, error) {
	hash := md5.New()
	// path以下のファイルとフォルダに対して処理を行う
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		io.WriteString(hash, path)
		// 書式付き出力(%vはデフォルトの書式)
		fmt.Fprintf(hash, "%v", info.IsDir())
		fmt.Fprintf(hash, "%v", info.ModTime())
		fmt.Fprintf(hash, "%v", info.Mode())
		fmt.Fprintf(hash, "%v", info.Name())
		fmt.Fprintf(hash, "%v", info.Size())
		return nil
	})
	if err != nil {
		return "", err
	}
	// hash.HashのSumメソッドは、現時点までに書き込まれたデータのハッシュ値を計算
	// %xは16進数(アルファベットは小文字)で出力する書式
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

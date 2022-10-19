package util

import (
	"os"
	"os/exec"
	"path"
)

// 蛇形转驼峰
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func PutGoLang(name, content string) error {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	if err != nil {
		if os.IsNotExist(err) {
			if _, err = os.Stat(path.Dir(name)); err != nil {
				if os.IsNotExist(err) {
					_ = os.MkdirAll(path.Dir(name), 0755)
				}
			}
			f, err = os.Create(name)
		}

	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	cmd := exec.Command("gofmt", "-w", f.Name())
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err = cmd.Run(); err != nil {
		return err
	}
	return nil
}

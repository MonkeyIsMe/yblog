package tool

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

// CreateDateDir 根据当前日期来创建文件夹
func CreateDateDir(Path string) (string, error) {
	folderName := time.Now().Format("20060102")
	folderPath := filepath.Join(Path, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		err := os.Mkdir(folderPath, 0777)
		if err != nil {
			log.Fatalf("Mkdir Err: [%+v]", err)
			return "", err
		} //0777也可以os.ModePerm
		err = os.Chmod(folderPath, 0777)
		if err != nil {
			log.Fatalf("Chmod Err: [%+v]", err)
			return "", err
		}
	}
	return folderPath, nil
}

// PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		log.Fatalf("PathExists Stat Err: [%+v]", err)
		return false, err
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return true, nil
}

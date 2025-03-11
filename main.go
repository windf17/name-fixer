package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 净化文件名，去除无关字符
func cleanFileName(fileName string) string {
	// 移除类似 [xxx-xxx.xxx] 的内容
	cleanName := regexp.MustCompile(`\[.*?\]`).ReplaceAllString(fileName, "")
	// 移除可能剩余的前导和尾随空格
	cleanName = strings.TrimSpace(cleanName)
	return cleanName
}

func main() {
	// 检查参数数量
	if len(os.Args) < 2 {
		fmt.Println("请提供至少一个参数：文件名片段。")
		return
	}

	fileNameFragment := os.Args[1] // 文件名片段

	// 检查文件名片段长度
	if len(fileNameFragment) < 1 {
		fmt.Println("文件名片段长度不能少于一个字符！")
		return
	}

	// 获取工作目录
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取当前工作目录失败: %v\n", err)
		return
	}

	// 解析文件名片段中的路径
	dir, fragment := filepath.Split(fileNameFragment)
	if dir != "" {
		// 如果文件名片段包含路径，则更新工作目录
		workingDir = filepath.Join(workingDir, dir)
	}

	// 创建目标目录
	targetDir := filepath.Join(workingDir, fragment)
	err = os.MkdirAll(targetDir, 0755)
	if err != nil {
		fmt.Printf("创建目标目录失败: %v\n", err)
		return
	}

	// 遍历工作目录及其子目录，查找匹配的文件
	var matches []string
	err = filepath.Walk(workingDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 忽略目录
		if info.IsDir() {
			return nil
		}

		// 检查文件名是否包含指定片段
		baseName := filepath.Base(path)
		if strings.Contains(strings.ToLower(baseName), strings.ToLower(fragment)) {
			matches = append(matches, path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("遍历目录失败: %v\n", err)
		return
	}

	if len(matches) == 0 {
		fmt.Println("未找到匹配的文件。")
		return
	}

	// 移动匹配的文件到目标目录
	for _, match := range matches {
		// 获取文件名并净化
		baseName := filepath.Base(match)
		cleanName := cleanFileName(baseName)

		// 构造目标路径
		targetPath := filepath.Join(targetDir, cleanName)

		// 移动文件
		err := os.Rename(match, targetPath)
		if err != nil {
			fmt.Printf("移动文件 %s 失败: %v\n", match, err)
			continue
		}
		fmt.Printf("已移动并净化: %s -> %s\n", match, targetPath)
	}

	// 删除以文件名片段开头的空目录
	err = removeEmptyDirs(workingDir, fragment)
	if err != nil {
		fmt.Printf("删除空目录失败: %v\n", err)
		return
	}

	fmt.Println("操作完成。")
}

// 递归删除以指定片段开头的空目录
func removeEmptyDirs(dir string, prefix string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	// 如果是空目录且包含指定片段，则删除
	if len(entries) == 0 && strings.Contains(strings.ToLower(filepath.Base(dir)), strings.ToLower(prefix)) {
		fmt.Printf("删除空目录: %s\n", dir)
		return os.Remove(dir)
	}

	// 递归检查子目录
	for _, entry := range entries {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			err := removeEmptyDirs(subDir, prefix)
			if err != nil {
				return err
			}
		}
	}

	// 再次检查当前目录是否为空且包含指定片段（可能子目录被删除了）
	entries, err = os.ReadDir(dir)
	if err != nil {
		return err
	}
	if len(entries) == 0 && strings.Contains(strings.ToLower(filepath.Base(dir)), strings.ToLower(prefix)) {
		fmt.Printf("删除空目录: %s\n", dir)
		return os.Remove(dir)
	}

	return nil
}

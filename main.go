package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 净化文件名，提取数字并按指定格式重组
func cleanFileName(fileName string, seriesName string, removePatterns []string) string {
	// 分离文件名和扩展名
	ext := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, ext)

	// 移除指定的字符串
	for _, pattern := range removePatterns {
		baseName = strings.ReplaceAll(baseName, pattern, "")
	}

	// 从后向前查找最后一个数字部分
	numberRegex := regexp.MustCompile(`(\d+)[^\d]*$`)
	// 找出匹配的数字
	match := numberRegex.FindStringSubmatch(baseName)
	// 使用找到的数字
	if len(match) > 1 {
		// 如果找到数字，则使用新格式
		return fmt.Sprintf("%s-%s%s", seriesName, match[1], ext)
	}

	// 如果没有找到数字，返回清理后的文件名
	return strings.TrimSpace(baseName) + ext
}

// 获取系统语言环境，判断是否为中文
func isChineseLocale() bool {
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = os.Getenv("LANGUAGE")
	}
	if lang == "" {
		lang = os.Getenv("LC_ALL")
	}
	return strings.HasPrefix(strings.ToLower(lang), "zh")
}

func showHelp() {
	if isChineseLocale() {
		fmt.Print(`名称修复工具 (Name Fixer)

功能：
  批量处理文件名，提取集数并按指定格式重命名

用法：
  name-fixer <剧集名> [要移除的字符串...]
  name-fixer --help    显示帮助信息

参数：
  <剧集名>            必需参数，用于匹配和重命名文件
  [要移除的字符串...]  可选参数，用于移除可能影响提取集数的干扰字符

处理规则：
  1. 移除指定的字符串（如果提供）
  2. 从文件名中提取最后出现的数字作为集数
  3. 生成新文件名格式："剧集名-集数.扩展名"
  4. 如果文件名中没有数字，则保持原文件名不变

示例：
  # 基本用法：处理包含"进击的巨人"的文件
  name-fixer 进击的巨人

  # 高级用法：指定要移除的字符串
  name-fixer 进击的巨人 "[字幕组]" "1080P"
`)
	} else {
		fmt.Print(`Name Fixer

Features:
  Batch process filenames by extracting episode numbers and renaming in specified format

Usage:
  name-fixer <series-name> [strings-to-remove...]
  name-fixer --help    Show help information

Parameters:
  <series-name>         Required, used for matching and renaming files
  [strings-to-remove]   Optional, used to remove strings that may interfere with episode number extraction

Processing Rules:
  1. Remove specified strings (if provided)
  2. Extract the last occurring number from the filename as the episode number
  3. Generate new filename format: "series-name-episode-number.extension"
  4. If no number is found in the filename, keep the original filename unchanged

Examples:
  # Basic usage: Process files containing "Attack on Titan"
  name-fixer "Attack on Titan"

  # Advanced usage: Specify strings to remove
  name-fixer "Attack on Titan" "[SubGroup]" "1080P"
`)
	}
}

func main() {
	// 检查参数数量
	if len(os.Args) < 2 {
		fmt.Println("请提供至少一个参数：剧集名。使用 --help 查看帮助信息。")
		return
	}

	// 检查是否为帮助命令
	if os.Args[1] == "--help" {
		showHelp()
		return
	}

	seriesName := os.Args[1] // 剧集名
	// 获取要删除的字符串列表（如果有的话）
	removePatterns := os.Args[2:]

	// 检查剧集名长度
	if len(seriesName) < 1 {
		fmt.Println("剧集名长度不能少于一个字符！")
		return
	}

	// 获取工作目录
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取当前工作目录失败: %v\n", err)
		return
	}

	// 创建目标目录
	targetDir := filepath.Join(workingDir, seriesName)
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
		if strings.Contains(strings.ToLower(baseName), strings.ToLower(seriesName)) {
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
		cleanName := cleanFileName(baseName, seriesName, removePatterns)

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

	// 删除以剧集名开头的空目录
	err = removeEmptyDirs(workingDir, seriesName)
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

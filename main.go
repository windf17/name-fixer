package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	// 用于记录源文件所在的目录
	var sourceDirs = make(map[string]bool)

	// 检查参数数量
	if len(os.Args) < 2 {
		if isChineseLocale() {
			fmt.Println("请提供至少一个参数：剧集名。使用 --help 查看帮助信息。")
		} else {
			fmt.Println("Please provide at least one argument: series name. Use --help for help information.")
		}
		return
	}

	// 检查是否为帮助命令
	args := os.Args[1]
	// args转换成大写字符串
	argsUpper := strings.ToLower(args)
	if argsUpper == "--help" || argsUpper == "-help" || argsUpper == "-h" || argsUpper == "--h" {
		showHelp()
		return
	}
	seriesName := os.Args[1] // 剧集名
	// 获取要删除的字符串列表（如果有的话）
	removePatterns := os.Args[2:]

	// 检查剧集名长度
	if len(seriesName) < 1 {
		if isChineseLocale() {
			fmt.Println("剧集名长度不能少于一个字符！")
		} else {
			fmt.Println("Series name length cannot be less than one character!")
		}
		return
	}

	// 获取工作目录
	workingDir, err := os.Getwd()
	if err != nil {
		if isChineseLocale() {
			fmt.Printf("获取当前工作目录失败: %v\n", err)
		} else {
			fmt.Printf("Failed to get current working directory: %v\n", err)
		}
		return
	}

	newSeriesName := seriesName
	// 检查待替换的字符串列表中是否包含剧集名
	for _, pattern := range removePatterns {
		// 如果包含，再检查是否包含"=>"
		parts := strings.Split(pattern, "=>")
		if len(parts) == 2 {
			if parts[0] == seriesName {
				newSeriesName = parts[1]
				break
			}
		}
	}

	// 创建目标目录
	targetDir := filepath.Join(workingDir, newSeriesName)
	// 创建目标目录前，检查是否存在同名文件，并处理
	if _, err := os.Stat(targetDir); err == nil {
		// 路径存在，检查是否是目录
		if isDir, _ := isDirectory(targetDir); !isDir {
			if isChineseLocale() {
				fmt.Printf("错误：%s 已存在且不是目录。\n", targetDir)
			} else {
				fmt.Printf("Error: %s already exists and is not a directory.\n", targetDir)
			}
			return
		}
	} else {
		// 目录不存在，创建
		err = os.MkdirAll(targetDir, 0755)
		if err != nil {
			if isChineseLocale() {
				fmt.Printf("创建目标目录失败: %v\n", err)
			} else {
				fmt.Printf("Failed to create target directory: %v\n", err)
			}
			return
		}
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
		if isChineseLocale() {
			fmt.Printf("遍历目录失败: %v\n", err)
		} else {
			fmt.Printf("Failed to traverse directory: %v\n", err)
		}
		return
	}

	if len(matches) == 0 {
		if isChineseLocale() {
			fmt.Println("未找到匹配的文件。")
		} else {
			fmt.Println("No matching files found.")
		}
		return
	}

	// 移动匹配的文件到目标目录
	for _, match := range matches {
		// 获取文件名并净化
		baseName := filepath.Base(match)
		cleanName := cleanFileName(baseName, newSeriesName, removePatterns)

		// 构造目标路径
		targetPath := filepath.Join(targetDir, cleanName)

		// 记录源文件所在的目录
		sourceDir := filepath.Dir(match)
		sourceDirs[sourceDir] = true

		// 移动文件
		err := os.Rename(match, targetPath)
		if err != nil {
			if isChineseLocale() {
				fmt.Printf("移动文件 %s 失败: %v\n", match, err)
			} else {
				fmt.Printf("Failed to move file %s: %v\n", match, err)
			}
			continue
		}
		if isChineseLocale() {
			fmt.Printf("已移动并净化: %s -> %s\n", match, targetPath)
		} else {
			fmt.Printf("Moved and cleaned: %s -> %s\n", match, targetPath)
		}
	}

	// 删除之前包含剧集文件的空目录
	err = removeEmptyDirs(workingDir, sourceDirs)
	if err != nil {
		if isChineseLocale() {
			fmt.Printf("删除空目录失败: %v\n", err)
		} else {
			fmt.Printf("Failed to remove empty directory: %v\n", err)
		}
		return
	}

	if isChineseLocale() {
		fmt.Println("操作完成。")
	} else {
		fmt.Println("Operation completed.")
	}
}

// 递归删除指定的空目录
func removeEmptyDirs(dir string, sourceDirs map[string]bool) error {
	// 如果当前目录不在sourceDirs中，直接返回
	if !sourceDirs[dir] {
		return nil
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	// 递归检查子目录
	for _, entry := range entries {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			err := removeEmptyDirs(subDir, sourceDirs)
			if err != nil {
				return err
			}
		}
	}

	// 重新读取目录内容（因为子目录可能已被删除）
	entries, err = os.ReadDir(dir)
	if err != nil {
		return err
	}

	// 检查当前目录是否为空
	if len(entries) == 0 {
		if isChineseLocale() {
			fmt.Printf("删除空目录: %s\n", dir)
		} else {
			fmt.Printf("Removing empty directory: %s\n", dir)
		}
		return os.Remove(dir)
	}

	return nil
}

// 净化文件名，提取数字并按指定格式重组
func cleanFileName(fileName string, seriesName string, removePatterns []string) string {
	// 分离文件名和扩展名
	ext := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, ext)

	if len(removePatterns) == 0 {
		// 从后向前查找最后一个数字部分
		numberRegex := regexp.MustCompile(`(\d+)[^\d]*$`)
		// 找出匹配的数字
		match := numberRegex.FindStringSubmatch(baseName)
		// 使用找到的数字
		if len(match) > 1 {
			// 如果找到数字，则使用新格式
			return fmt.Sprintf("%s-%s%s", seriesName, match[1], ext)
		}
		// 如果没有找到数字，返回去除空格后的文件名
		return fmt.Sprintf("%s%s", seriesName, ext)
	}
	// 移除指定的字符串
	for _, pattern := range removePatterns {
		// 如果pattern包含"=>",则按"=>"分割，前面的部分是要处理的字符串，后面的部分是要替换的字符串
		parts := strings.Split(pattern, "=>")
		if len(parts) == 2 {
			baseName = strings.ReplaceAll(baseName, parts[0], parts[1])
		} else {
			baseName = strings.ReplaceAll(baseName, pattern, "")
		}
	}
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
  nf <剧集名> [要移除的字符串...]
  nf --help 或 nf -h 显示帮助信息

参数：
  <剧集名>           必需参数，用于匹配和重命名文件
  [自定义处理规则...] 可选参数，自定义处理规则(可有多条)。若包含"=>"即它之前的字符串将被它之后的字符串替换，若不包含则表示移除字符串。
                    注意：在Linux系统下，不同的处理规则都要用引号包裹。

处理规则：
  1. 若没定义处理规则，则从文件名中提取最后出现的数字作为集数，生成新文件名格式："剧集名-集数.扩展名"
  2. 若定义了处理规则，则按照处理规则进行处理。
  3. 若处理规则中包含"=>"，则它之前的字符串将被它之后的字符串替换，若不包含则表示移除字符串。

示例：
  # 基本用法：处理包含"进击的巨人"的文件
  nf 进击的巨人

  # 高级用法：指定要移除的字符串
  nf "进击的巨人" "[字幕组]" "1080P"

  # 使用字符串替换功能（注意：在Linux系统下需要用引号包裹参数）
  nf "进击的巨人" "第=>Episode " "季=>Season "
`)
	} else {
		fmt.Print(`Name Fixer

Features:
  Batch process filenames by extracting episode numbers and renaming in specified format

Usage:
  nf <series-name> [custom-rules...]
  nf --help or nf -h    Show help information

Parameters:
  <series-name>      Required, used for matching and renaming files
  [custom-rules...]  Optional, multiple custom rules for processing. If a rule contains "=>", the string before it will be replaced by the string after it. If not, the string will be removed.
                     Note: On Linux systems, each rule must be quoted.

Processing Rules:
  1. If no custom rules are defined, extract the last occurring number as episode number and generate new filename format: "series-name-episode-number.extension"
  2. If custom rules are defined, process according to the rules.
  3. If a rule contains "=>", the string before it will be replaced by the string after it. If not, the string will be removed.

Examples:
  # Basic usage: Process files containing "Attack on Titan"
  nf "Attack on Titan"

  # Advanced usage: Specify strings to remove
  nf "Attack on Titan" "[SubGroup]" "1080P"

  # Using string replacement feature (Note: On Linux, parameters with special characters must be quoted)
  nf "Attack on Titan" "Chapter=>Episode " "Season=>Season "
`)
	}
}

// 辅助函数检查是否为目录
func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

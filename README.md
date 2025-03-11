# Name Fixer 文件名处理工具

[English](#english) | [中文](#chinese)

<a name="chinese"></a>
## 中文说明

### 项目简介
Name Fixer 是一个命令行工具，用于批量处理文件名，可以移除文件名中的特定标记（如 `[xxx-xxx.xxx]`），并将匹配的文件整理到指定目录中。

### 功能特点
- 支持按文件名片段匹配文件
- 自动提取文件名中的数字编号
- 支持自定义要移除的字符串列表
- 智能识别文件编号位置
- 将匹配的文件移动到新目录
- 自动清理空目录
- 支持递归搜索子目录

### 安装方法
1. 确保已安装 Go 1.23.5 或更高版本
2. 克隆或下载本项目
3. 在项目目录下执行：
   ```bash
   go build
   ```

### 使用方法
```bash
# Windows
name-fixer.exe <剧集名> [要移除的字符串...]

# Linux/macOS
./name-fixer <剧集名> [要移除的字符串...]
```

### 使用示例
```bash
# 基本用法：匹配并处理包含"动漫名"的文件
name-fixer.exe 动漫名

# 高级用法：指定要移除的字符串
name-fixer.exe 动漫名 "[字幕组]" "1080P" "第话"
```

### 参数说明
- `<剧集名>`: 必需参数，用于匹配和重命名文件
- `[要移除的字符串...]`: 可选参数，用于移除可能影响提取集数的干扰字符
  - 如果文件名中的字符不会影响正确提取集数，则无需添加到移除列表中
  - 例如：对于文件"动画第一集第2话.mp4"，即使不指定移除"第"、"集"、"话"等字符，程序也能正确提取出集数"2"

### 处理规则
1. 程序会搜索当前目录及其子目录中包含指定片段的文件
2. 创建以搜索片段命名的新目录
3. 移动匹配的文件到新目录，同时清理文件名
4. 自动删除剩余的空目录

### 文件名处理规则
1. 移除指定的字符串（如果提供）
2. 从文件名中提取最后出现的数字作为集数
3. 生成新文件名格式："剧集名-集数.扩展名"
4. 如果文件名中没有数字，则保持原文件名不变

---

<a name="english"></a>
## English

### Introduction
Name Fixer is a command-line tool for batch processing filenames. It can extract episode numbers, remove specific markers, and organize matching files into a designated directory.

### Features
- Match files by filename fragment
- Automatically extract episode numbers from filenames
- Support custom string removal patterns
- Smart episode number detection
- Move matching files to a new directory
- Clean up empty directories
- Support recursive subdirectory search

### Installation
1. Ensure Go 1.23.5 or higher is installed
2. Clone or download this project
3. In the project directory, run:
   ```bash
   go build
   ```

### Usage
```bash
# Windows
name-fixer.exe <series-name> [strings-to-remove...]

# Linux/macOS
./name-fixer <series-name> [strings-to-remove...]
```

### Examples
```bash
# Basic usage: Match and process files containing "SeriesName"
name-fixer.exe SeriesName

# Advanced usage: Specify strings to remove
name-fixer.exe SeriesName "[Sub-Group]" "1080P" "EP"
```

### Processing Rules
1. The program searches for files containing the specified fragment in the current directory and its subdirectories
2. Creates a new directory named after the search fragment
3. Moves matching files to the new directory while cleaning their filenames
4. Automatically removes remaining empty directories

### Filename Processing Rules
1. Remove specified strings (if provided)
2. Extract the last occurring number from the filename as the episode number
3. Generate new filename format: "series-name-episode-number.extension"
4. If no number is found in the filename, keep the original filename unchanged

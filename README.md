# Name Fixer 文件名处理工具

[English](#english) | [中文](#chinese)

<a name="chinese"></a>
## 中文说明

### 项目简介
Name Fixer 是一个命令行工具，用于批量处理文件名，可以移除文件名中的特定标记（如 `[xxx-xxx.xxx]`），并将匹配的文件整理到指定目录中。

### 功能特点
- 支持按文件名片段匹配文件
- 自动清理文件名中的方括号标记
- 移除文件名前后的空格
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
name-fixer.exe <文件名片段>

# Linux/macOS
./name-fixer <文件名片段>
```

### 使用示例
```bash
# 匹配并处理包含"example"的文件
name-fixer.exe example

# 处理指定子目录中的文件
name-fixer.exe path/to/dir/example
```

### 处理规则
1. 程序会搜索当前目录及其子目录中包含指定片段的文件
2. 创建以搜索片段命名的新目录
3. 移动匹配的文件到新目录，同时清理文件名
4. 自动删除剩余的空目录

---

<a name="english"></a>
## English

### Introduction
Name Fixer is a command-line tool for batch processing filenames. It can remove specific markers (like `[xxx-xxx.xxx]`) from filenames and organize matching files into a designated directory.

### Features
- Match files by filename fragment
- Automatically clean bracket markers from filenames
- Remove leading and trailing spaces
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
name-fixer.exe <filename-fragment>

# Linux/macOS
./name-fixer <filename-fragment>
```

### Examples
```bash
# Match and process files containing "example"
name-fixer.exe example

# Process files in a specific subdirectory
name-fixer.exe path/to/dir/example
```

### Processing Rules
1. The program searches for files containing the specified fragment in the current directory and its subdirectories
2. Creates a new directory named after the search fragment
3. Moves matching files to the new directory while cleaning their filenames
4. Automatically removes remaining empty directories
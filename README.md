# Name Fixer 文件名处理工具

[English](#english) | [中文](#chinese)

<a name="chinese"></a>

## 中文说明

### 项目简介

Name Fixer 是一个命令行工具，用于批量处理文件名。它可以提取文件名中的集数编号，移除指定的字符串（如 `[xxx-xxx.xxx]`），并按照统一格式重命名文件，最后将处理后的文件整理到指定目录中。

### 功能特点

-   支持按剧集名称匹配文件
-   自动提取文件名中的集数编号
-   支持移除或替换指定字符串
-   支持自定义文件名格式
-   将匹配的文件移动到新目录
-   自动清理空目录
-   支持递归搜索子目录

### 安装方法

1. 确保已安装 Go 1.23.5 或更高版本
2. 克隆或下载本项目
3. 在项目目录下执行：
    ```bash
    go build -o nf
    ```

### 使用方法

```bash
nf <剧集名> [要移除的字符串...]
nf --help    显示帮助信息
```

### 使用示例

```bash
# 基本用法：处理包含"进击的巨人"的文件
nf 进击的巨人

# 高级用法：指定要移除的字符串
nf "进击的巨人" "[字幕组]" "1080P"

# 使用字符串替换功能
nf "进击的巨人" "第=>Episode " "季=>Season "
```

### 处理规则

1. 移除指定的字符串（如果提供）
2. 从文件名中提取最后出现的数字作为集数
3. 生成新文件名格式："剧集名-集数.扩展名"
4. 如果文件名中没有数字，则保持原文件名不变
5. 将处理后的文件移动到以剧集名命名的新目录
6. 自动删除剩余的空目录

---

<a name="english"></a>

## English

### Introduction

Name Fixer is a command-line tool for batch processing filenames. It can extract episode numbers from filenames, remove specific strings (like `[xxx-xxx.xxx]`), rename files in a unified format, and organize them into a designated directory.

### Features

-   Match files by series name
-   Automatically extract episode numbers from filenames
-   Support removing or replacing specific strings
-   Support custom filename format
-   Move matching files to a new directory
-   Clean up empty directories
-   Support recursive subdirectory search

### Installation

1. Ensure Go 1.23.5 or higher is installed
2. Clone or download this project
3. In the project directory, run:
    ```bash
    go build -o nf
    ```

### Usage

```bash
nf <series-name> [strings-to-remove...]
nf --help    Show help information
```

### Examples

```bash
# Basic usage: Process files containing "Attack on Titan"
nf "Attack on Titan"

# Advanced usage: Specify strings to remove
nf "Attack on Titan" "[SubGroup]" "1080P"

# Using string replacement feature
nf "Attack on Titan" "Ep.=>Episode " "S.=>Season "
```

### Processing Rules

1. Remove specified strings (if provided)
2. Extract the last occurring number from the filename as the episode number
3. Generate new filename format: "series-name-episode-number.extension"
4. If no number is found in the filename, keep the original filename unchanged
5. Move processed files to a new directory named after the series name
6. Automatically remove remaining empty directories

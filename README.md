# Name Fixer 文件名处理工具

[English](#english) | [中文](#chinese)

<a name="chinese"></a>

## 中文说明

### 项目简介

<<<<<<< HEAD
Name Fixer 是一个命令行工具，用于批量处理文件名，可以移除文件名中的特定标记（如 `[xxx-xxx.xxx]`），并将匹配的文件整理到指定目录中。

### 功能特点

-   支持按文件名片段匹配文件
-   自动清理文件名中的方括号标记
-   移除文件名前后的空格
=======
Name Fixer 是一个命令行工具，用于批量处理文件名。它可以提取文件名中的集数编号，移除指定的字符串（如 `[xxx-xxx.xxx]`），并按照统一格式重命名文件，最后将处理后的文件整理到指定目录中。

### 功能特点

-   支持按剧集名称匹配文件
-   自动提取文件名中的集数编号
-   支持移除或替换指定字符串
-   支持自定义文件名格式
>>>>>>> 62358f9c76a5157d5582740c71f87a28ecea7103
-   将匹配的文件移动到新目录
-   自动清理空目录
-   支持递归搜索子目录

### 安装方法

1. 确保已安装 Go 1.23.5 或更高版本
2. 克隆或下载本项目
3. 在项目目录下执行：
    ```bash
<<<<<<< HEAD
    go build
    ```
4. 全局安装（可选）：
    - Windows 系统：
        1. 以管理员身份运行命令提示符
        2. 将生成的 nf.exe 复制到 C:\Windows 目录：
            ```cmd
            copy nf.exe C:\Windows\
            ```
    - Linux 系统：
        1. 为 nf.sh 添加执行权限：
            ```bash
            chmod +x ./nf.sh
            ```
        2. 将 nf.sh 移动到/usr/bin 目录：
            ```bash
            sudo mv ./nf.sh /usr/bin/nf
            ```
=======
    go build -o nf
    ```
>>>>>>> 62358f9c76a5157d5582740c71f87a28ecea7103

### 使用方法

```bash
<<<<<<< HEAD
# Windows
nf.exe <文件名片段>

# Linux/macOS
./nf <文件名片段>
```

### 使用示例

#### 基本使用方式

```bash
# 创建测试目录和文件
❯ mkdir -p 1 && touch 1/刺青1.mp4
❯ mkdir -p 2 && touch 2/刺青2.mp4
❯ mkdir -p 3 && touch 3/刺青3.mp4
❯ tree
.
├── 1
│   └── 刺青1.mp4
├── 2
│   └── 刺青2.mp4
└── 3
    └── 刺青3.mp4

4 directories, 3 files

# 执行文件名处理
❯ nf 刺青
Moved and cleaned: /home/windf/test/2/刺青2.mp4 -> /home/windf/test/刺青/刺青-2.mp4
Moved and cleaned: /home/windf/test/3/刺青3.mp4 -> /home/windf/test/刺青/刺青-3.mp4
Moved and cleaned: /home/windf/test/1/刺青1.mp4 -> /home/windf/test/刺青/刺青-1.mp4
Operation completed.
❯ tree
.
└── 刺青
    ├── 刺青-1.mp4
    ├── 刺青-2.mp4
    └── 刺青-3.mp4
```

#### 进阶使用方式

```bash
# 创建测试目录和文件
❯ mkdir -p 1 && touch 1/刺青abc1.mp4
❯ mkdir -p 2 && touch 2/刺青abc2.mp4
❯ touch 刺青abc3.mp4
❯ tree
.
├── 1
│   └── 刺青abc1.mp4
├── 2
│   └── 刺青abc2.mp4
└── 刺青abc3.mp4

3 directories, 3 files

# 使用字符串替换功能
❯ nf 刺青 "刺青=>刺" "abc"
Moved and cleaned: /home/windf/test/2/刺青abc2.mp4 -> /home/windf/test/刺/刺2.mp4
Moved and cleaned: /home/windf/test/刺青abc3.mp4 -> /home/windf/test/刺/刺3.mp4
Moved and cleaned: /home/windf/test/1/刺青abc1.mp4 -> /home/windf/test/刺/刺1.mp4
Operation completed.
❯ tree
.
└── 刺
    ├── 刺1.mp4
    ├── 刺2.mp4
    └── 刺3.mp4
```

#### 错误使用案例

注意：在使用字符串替换功能时，参数必须使用引号，否则会导致错误的结果：

```bash
# 创建测试目录和文件
❯ mkdir -p 1 && touch 1/刺青abc1.mp4
❯ mkdir -p 2 && touch 2/刺青abc2.mp4
❯ touch 刺青abc3.mp4
❯ tree
.
├── 1
│   └── 刺青abc1.mp4
├── 2
│   └── 刺青abc2.mp4
└── 刺青abc3.mp4

3 directories, 3 files

# 错误：参数未使用引号
❯ nf 刺青 刺青=>刺 abc
❯ tree
.
├── 刺
└── 刺青
    ├── 刺青1.mp4
    ├── 刺青2.mp4
    └── 刺青3.mp4
```

### 处理规则

1. 程序会搜索当前目录及其子目录中包含指定片段的文件
2. 创建以搜索片段命名的新目录
3. 移动匹配的文件到新目录，同时清理文件名
4. 自动删除剩余的空目录

=======
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

>>>>>>> 62358f9c76a5157d5582740c71f87a28ecea7103
---

<a name="english"></a>

## English

### Introduction

<<<<<<< HEAD
Name Fixer is a command-line tool for batch processing filenames. It can remove specific markers (like `[xxx-xxx.xxx]`) from filenames and organize matching files into a designated directory.

### Features

-   Match files by filename fragment
-   Automatically clean bracket markers from filenames
-   Remove leading and trailing spaces
=======
Name Fixer is a command-line tool for batch processing filenames. It can extract episode numbers from filenames, remove specific strings (like `[xxx-xxx.xxx]`), rename files in a unified format, and organize them into a designated directory.

### Features

-   Match files by series name
-   Automatically extract episode numbers from filenames
-   Support removing or replacing specific strings
-   Support custom filename format
>>>>>>> 62358f9c76a5157d5582740c71f87a28ecea7103
-   Move matching files to a new directory
-   Clean up empty directories
-   Support recursive subdirectory search

### Installation

1. Ensure Go 1.23.5 or higher is installed
2. Clone or download this project
3. In the project directory, run:
    ```bash
<<<<<<< HEAD
    go build
    ```

### Usage

```bash
# Windows
nf.exe <filename-fragment>

# Linux/macOS
./nf <filename-fragment>
```

### Examples

#### Basic Usage

```bash
# Create test directories and files
❯ mkdir -p 1 && touch 1/tattoo1.mp4
❯ mkdir -p 2 && touch 2/tattoo2.mp4
❯ mkdir -p 3 && touch 3/tattoo3.mp4
❯ tree
.
├── 1
│   └── tattoo1.mp4
├── 2
│   └── tattoo2.mp4
└── 3
    └── tattoo3.mp4

4 directories, 3 files

# Process filenames
❯ nf tattoo
Moved and cleaned: /home/windf/test/2/tattoo2.mp4 -> /home/windf/test/tattoo/tattoo-2.mp4
Moved and cleaned: /home/windf/test/3/tattoo3.mp4 -> /home/windf/test/tattoo/tattoo-3.mp4
Moved and cleaned: /home/windf/test/1/tattoo1.mp4 -> /home/windf/test/tattoo/tattoo-1.mp4
Operation completed.
❯ tree
.
└── tattoo
    ├── tattoo-1.mp4
    ├── tattoo-2.mp4
    └── tattoo-3.mp4
```

#### Advanced Usage

```bash
# Create test directories and files
❯ mkdir -p 1 && touch 1/tattooabc1.mp4
❯ mkdir -p 2 && touch 2/tattooabc2.mp4
❯ touch tattooabc3.mp4
❯ tree
.
├── 1
│   └── tattooabc1.mp4
├── 2
│   └── tattooabc2.mp4
└── tattooabc3.mp4

3 directories, 3 files

# Using string replacement feature
❯ nf tattoo "tattoo=>tat" "abc"
Moved and cleaned: /home/windf/test/2/tattooabc2.mp4 -> /home/windf/test/tat/tat2.mp4
Moved and cleaned: /home/windf/test/tattooabc3.mp4 -> /home/windf/test/tat/tat3.mp4
Moved and cleaned: /home/windf/test/1/tattooabc1.mp4 -> /home/windf/test/tat/tat1.mp4
Operation completed.
❯ tree
.
└── tat
    ├── tat1.mp4
    ├── tat2.mp4
    └── tat3.mp4
```

#### Error Case Example

Note: When using the string replacement feature, parameters must be quoted to avoid incorrect results:

```bash
# Create test directories and files
❯ mkdir -p 1 && touch 1/tattooabc1.mp4
❯ mkdir -p 2 && touch 2/tattooabc2.mp4
❯ touch tattooabc3.mp4
❯ tree
.
├── 1
│   └── tattooabc1.mp4
├── 2
│   └── tattooabc2.mp4
└── tattooabc3.mp4

3 directories, 3 files

# Error: Parameters not quoted
❯ nf tattoo tattoo=>tat abc
❯ tree
.
├── tat
└── tattoo
    ├── tattoo1.mp4
    ├── tattoo2.mp4
    └── tattoo3.mp4
=======
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
>>>>>>> 62358f9c76a5157d5582740c71f87a28ecea7103
```

### Processing Rules

<<<<<<< HEAD
1. The program searches for files containing the specified fragment in the current directory and its subdirectories
2. Creates a new directory named after the search fragment
3. Moves matching files to the new directory while cleaning their filenames
4. Automatically removes remaining empty directories
=======
1. Remove specified strings (if provided)
2. Extract the last occurring number from the filename as the episode number
3. Generate new filename format: "series-name-episode-number.extension"
4. If no number is found in the filename, keep the original filename unchanged
5. Move processed files to a new directory named after the series name
6. Automatically remove remaining empty directories
>>>>>>> 62358f9c76a5157d5582740c71f87a28ecea7103

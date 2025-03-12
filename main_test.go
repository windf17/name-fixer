package main

import (
	"testing"
)

func TestCleanFileName(t *testing.T) {
	tests := []struct {
		name           string
		fileName       string
		seriesName     string
		removePatterns []string
		want           string
	}{
		{
			name:           "正常文件名带数字",
			fileName:       "[abc-xyz.com]剧集01.mp4",
			seriesName:     "剧集",
			removePatterns: []string{"[abc-xyz.com]"},
			want:           "剧集-01.mp4",
		},
		{
			name:           "文件名中间有数字但以其他数字结尾",
			fileName:       "动漫第01话 第02集.mp4",
			seriesName:     "动漫",
			removePatterns: []string{"话", "集"},
			want:           "动漫-02.mp4",
		},
		{
			name:           "文件名没有数字",
			fileName:       "无数字文件.mp4",
			seriesName:     "测试",
			removePatterns: nil,
			want:           "无数字文件.mp4",
		},
		{
			name:           "文件名有多个数字和特殊字符",
			fileName:       "[Sub-Group]_Series_01_[1080p]_[123ABC].mkv",
			seriesName:     "Series",
			removePatterns: []string{ "[1080p]", "[123ABC]"},
			want:           "Series-01.mkv",
		},
		{
			name:           "文件名以数字开头",
			fileName:       "01.测试文件.avi",
			seriesName:     "测试",
			removePatterns: nil,
			want:           "测试-01.avi",
		},
		{
			name:           "文件名包含中文数字",
			fileName:       "动画第一集第2话.mp4",
			seriesName:     "动画",
			removePatterns: []string{"第", "集", "话"},
			want:           "动画-2.mp4",
		},
		{
			name:           "文件名包含中文数字，removePatterns为空",
			fileName:       "动画第一集第2话.mp4",
			seriesName:     "动画",
			removePatterns: []string{},
			want:           "动画-2.mp4",
		},
		{
			name:           "文件名包含多组数字和特殊字符",
			fileName:       "[字幕组]动漫名EP01.第1话.1080P.mp4",
			seriesName:     "动漫名",
			removePatterns: []string{"[字幕组]", "EP", "第1话", "1080P"},
			want:           "动漫名-01.mp4",
		},
		{
			name:           "文件名包含下划线和空格",
			fileName:       "_Series Name_ _01_ .mkv",
			seriesName:     "Series Name",
			removePatterns: []string{"_"},
			want:           "Series Name-01.mkv",
		},
		{
			name:           "文件名不包含数字但需要移除指定字符串",
			fileName:       "[字幕组]测试文件[1080P].mp4",
			seriesName:     "测试",
			removePatterns: []string{ "[1080P]"},
			want:           "[字幕组]测试文件.mp4",
		},
		{
			name:           "文件名不包含数字且需要移除多个重复字符串",
			fileName:       "[Sub]测试[Sub]文件[Sub].mp4",
			seriesName:     "测试",
			removePatterns: []string{"[Sub]"},
			want:           "测试文件.mp4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanFileName(tt.fileName, tt.seriesName, tt.removePatterns); got != tt.want {
				t.Errorf("cleanFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
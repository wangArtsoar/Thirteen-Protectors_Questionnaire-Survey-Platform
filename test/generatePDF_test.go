package main

import (
	"github.com/go-pdf/fpdf"
	"os"
	"path/filepath"
	"testing"
)

// TestGeneratePDF generate a simple pdf for test
func TestGeneratePDF(t *testing.T) {
	// 当前执行文件目录
	exePath, _ := os.Getwd()
	// 拼接要保存的PDF目录
	pdfDir := filepath.Join(exePath, "../pdfs")
	// 如果目录不存在,递归创建目录
	if _, err := os.Stat(pdfDir); os.IsNotExist(err) {
		os.MkdirAll(pdfDir, 0755)
	}
	// 创建一个新的PDF文档
	pdf := fpdf.New("P", "mm", "A4", "")
	// 添加新页面
	pdf.AddPage()
	// 设置字体
	pdf.SetFont("Arial", "B", 16)
	// 输出文本
	pdf.Cell(40, 10, "Hello World")
	// 保存文件
	pdf.OutputFileAndClose(pdfDir + "/hello.pdf")
}

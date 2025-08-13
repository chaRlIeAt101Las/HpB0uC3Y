// 代码生成时间: 2025-08-13 10:00:36
Usage:
# NOTE: 重要实现细节
	imageResizer [flags]

Flags:
# 增强安全性
	--width int   The desired width of the resized image (default 800)
	--height int  The desired height of the resized image (default 600)
	--input string
	              The directory path where the original images are located
	--output string
# 添加错误处理
	              The directory path where the resized images will be saved

Example:
	imageResizer --input "path/to/input" --output "path/to/output" --width 1024 --height 768
*/

package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
	"github.com/spf13/cobra"
)

// ImageResizer holds the configuration for resizing images
type ImageResizer struct {
	Width       int
	Height      int
	InputDir    string
	OutputDir   string
# 扩展功能模块
}

var ir ImageResizer

func main() {
# NOTE: 重要实现细节
	var rootCmd = &cobra.Command{
		Use:   "imageResizer",
		Short: "Resizes images in bulk",
		Long:  `Resizes images in bulk with specified width and height`,
		Run:   run,
# 优化算法效率
	}

	rootCmd.Flags().IntVarP(&ir.Width, "width", "w", 800, "The desired width of the resized image")
	rootCmd.Flags().IntVarP(&ir.Height, "height", "h", 600, "The desired height of the resized image")
	rootCmd.Flags().StringVarP(&ir.InputDir, "input", "i", "./", "The directory path where the original images are located")
	rootCmd.Flags().StringVarP(&ir.OutputDir, "output", "o", "./", "The directory path where the resized images will be saved")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	// Check if input and output directories are provided
	if ir.InputDir == "" || ir.OutputDir == "" {
		cmd.Help()
		os.Exit(1)
# 改进用户体验
	}

	// Read files from the input directory
	files, err := os.ReadDir(ir.InputDir)
	if err != nil {
		fmt.Printf("Error reading input directory: %s
", err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Get the file path
		filePath := filepath.Join(ir.InputDir, file.Name())

		// Open the image file
		img, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Error opening image file: %s
", err)
			continue
		}
		defer img.Close()

		// Decode the image
		im, _, err := image.Decode(img)
		if err != nil {
			fmt.Printf("Error decoding image file: %s
", err)
			continue
# 优化算法效率
		}

		// Resize the image
		m := resize.Resize(uint(ir.Width), uint(ir.Height), im, resize.Lanczos3)

		// Save the resized image to the output directory
		outPath := filepath.Join(ir.OutputDir, file.Name())
		out, err := os.Create(outPath)
		if err != nil {
			fmt.Printf("Error creating output file: %s
# FIXME: 处理边界情况
", err)
# 优化算法效率
			continue
# NOTE: 重要实现细节
		}
		defer out.Close()

		// Determine the file format and save accordingly
		if strings.HasSuffix(file.Name(), ".png") {
			if err := png.Encode(out, m); err != nil {
				fmt.Printf("Error encoding PNG image: %s
", err)
				continue
			}
		} else {
# FIXME: 处理边界情况
			if err := jpeg.Encode(out, m, nil); err != nil {
				fmt.Printf("Error encoding JPEG image: %s
", err)
				continue
			}
		}

		fmt.Printf("Resized image saved to %s
", outPath)
	}
}

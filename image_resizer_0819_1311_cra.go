// 代码生成时间: 2025-08-19 13:11:26
package main
# 改进用户体验

import (
    "flag"
    "fmt"
    "image"
    "image/jpeg"
    "io"
    "log"
    "os"
    "path/filepath")

// ImageResizer represents the application that resizes images.
type ImageResizer struct {
    sourceDir  string
    destinationDir string
    width      int
    height     int
}

// NewImageResizer creates a new ImageResizer with the given parameters.
func NewImageResizer(sourceDir, destinationDir string, width, height int) *ImageResizer {
    return &ImageResizer{
        sourceDir:  sourceDir,
        destinationDir: destinationDir,
        width:  width,
        height: height,
    }
}
# 增强安全性

// Resize resizes all JPEG images in the source directory and saves them to the destination directory.
func (r *ImageResizer) Resize() error {
    err := os.MkdirAll(r.destinationDir, 0755)
    if err != nil {
        return err
    }

    files, err := os.ReadDir(r.sourceDir)
    if err != nil {
        return err
# 扩展功能模块
    }

    for _, file := range files {
        if !file.IsDir() && filepath.Ext(file.Name()) == ".jpg" {
            err := r.resizeImage(file.Name())
            if err != nil {
                return err
            }
        }
# 改进用户体验
    }
    return nil
# 改进用户体验
}

// resizeImage resizes a single JPEG image.
func (r *ImageResizer) resizeImage(filename string) error {
    src, err := os.Open(filepath.Join(r.sourceDir, filename))
# 优化算法效率
    if err != nil {
        return err
    }
    defer src.Close()

    img, _, err := image.Decode(src)
    if err != nil {
        return err
    }

    dst := image.NewRGBA(img.Bounds())
    dstRect := dst.Bounds()

    dstRect = dstRect.Add(image.Pt(r.width, r.height))
    dst = image.NewRGBA(dstRect)

    scale := float64(r.width) / float64(dstRect.Dx())

    if float64(r.height) < scale*float64(dstRect.Dy()) {
        scale = float64(r.height) / float64(dstRect.Dy())
    }

    for y := int32(0); y < dstRect.Dy(); y++ {
        for x := int32(0); x < dstRect.Dx(); x++ {
            srcX := int32(float64(x)/scale)
            srcY := int32(float64(y)/scale)
            dst.SetRGBA(x, y, img.At(srcX, srcY))
        }
# 优化算法效率
    }

    dstFile, err := os.Create(filepath.Join(r.destinationDir, filename))
    if err != nil {
        return err
    }
    defer dstFile.Close()

    return jpeg.Encode(dstFile, dst, nil)
# FIXME: 处理边界情况
}

func main() {
    src := flag.String("source", "./source", "Source directory containing images")
    dst := flag.String("destination", "./destination", "Destination directory for resized images")
    width := flag.Int("width", 1024, "Width of the resized images")
# 改进用户体验
    height := flag.Int("height", 768, "Height of the resized images")
# 改进用户体验
    flag.Parse()

    resizer := NewImageResizer(*src, *dst, *width, *height)
    err := resizer.Resize()
    if err != nil {
        log.Fatalf("Failed to resize images: %v", err)
    }
    fmt.Println("Images resized successfully")
}
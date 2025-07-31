// 代码生成时间: 2025-07-31 18:05:54
package main

import (
    "fmt"
    "image"
    "image/jpeg"
    "os"
    "path/filepath"
    "strings"
    "log"
    "github.com/spf13/cobra"
)

// ImageResizer represents the application
type ImageResizer struct {
    Width, Height int
}

// NewImageResizer creates a new ImageResizer instance
func NewImageResizer() *ImageResizer {
    return &ImageResizer{}
}

// Resize resizes the image to the specified width and height
func (r *ImageResizer) Resize(srcPath, destPath string) error {
    file, err := os.Open(srcPath)
    if err != nil {
        return fmt.Errorf("failed to open source image: %w", err)
    }
    defer file.Close()

    img, _, err := image.Decode(file)
    if err != nil {
        return fmt.Errorf("failed to decode image: %w", err)
    }

    // Create a new image with the desired dimensions
    imgRect := img.Bounds()
    newImg := image.NewRGBA(image.Rect(0, 0, r.Width, r.Height))

    // Calculate the scale factor to maintain aspect ratio
    scaleX := float64(r.Width) / float64(imgRect.Dx())
    scaleY := float64(r.Height) / float64(imgRect.Dy())
    scale := func(dx, dy, sx, sy float64) {
        if sx > sy {
            dx = dy * (sx / sy)
        } else {
            dy = dx * (sy / sx)
        }
        return dx, dy
    }

    width, height := scale(float64(r.Width), float64(r.Height), scaleX, scaleY)
    newImgRect := image.Rect(0, 0, int(width), int(height))

    // Resize the image
    if err := resize(img, newImg, imgRect, newImgRect); err != nil {
        return fmt.Errorf("failed to resize image: %w", err)
    }

    // Save the resized image
    outFile, err := os.Create(destPath)
    if err != nil {
        return fmt.Errorf("failed to create destination image: %w", err)
    }
    defer outFile.Close()

    if err := jpeg.Encode(outFile, newImg, nil); err != nil {
        return fmt.Errorf("failed to encode resized image: %w", err)
    }

    fmt.Printf("Resized image saved to %s
", destPath)
    return nil
}

// resize is a helper function to perform the image resizing
func resize(src image.Image, dst *image.RGBA, srcRect, dstRect image.Rectangle) error {
    // Implement the image resizing logic here
    // For simplicity, this example uses a naive approach
    for y := dstRect.Min.Y; y < dstRect.Max.Y; y++ {
        for x := dstRect.Min.X; x < dstRect.Max.X; x++ {
            srcX := int(float64(x) * float64(srcRect.Dx()) / float64(dstRect.Dx()))
            srcY := int(float64(y) * float64(srcRect.Dy()) / float64(dstRect.Dy()))
            dst.Set(x, y, src.At(srcX, srcY))
        }
    }
    return nil
}

// Command represents the command line arguments
type Command struct {
    Width, Height int
    SrcDir, DestDir string
}

// Execute runs the command
func (cmd *Command) Execute() error {
    resizer := NewImageResizer()
    resizer.Width = cmd.Width
    resizer.Height = cmd.Height

    err := filepath.WalkDir(cmd.SrcDir, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if !d.IsDir() && strings.HasSuffix(path, ".jpg") {
            destPath := filepath.Join(cmd.DestDir, d.Name())
            if err := resizer.Resize(path, destPath); err != nil {
                log.Printf("Error resizing image %s: %v", path, err)
            }
        }
        return nil
    })
    if err != nil {
        return fmt.Errorf("failed to process images: %w", err)
    }
    return nil
}

func main() {
    cmd := &Command{}

    var rootCmd = &cobra.Command{
        Use:   "imageresizer",
        Short: "Resizes images in a directory",
        Long:  `Resizes all JPEG images in a directory to a specified size.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            return cmd.Execute()
        },
    }

    rootCmd.Flags().IntVar(&cmd.Width, "width", 800, "Width of the resized images")
    rootCmd.Flags().IntVar(&cmd.Height, "height", 600, "Height of the resized images")
    rootCmd.Flags().StringVar(&cmd.SrcDir, "src", "./src", "Source directory containing images")
    rootCmd.Flags().StringVar(&cmd.DestDir, "dest", "./dest", "Destination directory for resized images")

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

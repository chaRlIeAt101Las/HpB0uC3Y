// 代码生成时间: 2025-09-18 22:23:23
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/spf13/cobra"
)

// 初始化Cobra命令
var rootCmd = &cobra.Command{
    Use:   "api-server",
    Short: "A simple RESTful API server",
    Long:  `A simple RESTful API server created with Cobra framework`,
}

// 定义一个结构体来模拟用户数据
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

// 用户列表
var users = []User{{ID: 1, Name: "John"}, {ID: 2, Name: "Jane"}}

// 定义一个处理GET请求的函数来获取用户列表
func getUserList(w http.ResponseWriter, r *http.Request) {
    // 确保只处理GET请求
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // 将用户列表编码为JSON并写入响应
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

// 定义一个处理POST请求的函数来创建新用户
func createUser(w http.ResponseWriter, r *http.Request) {
    // 确保只处理POST请求
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // 从请求中解码用户数据
    var newUser User
    if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // 将新用户添加到用户列表
    newUser.ID = len(users) + 1 // 简单的ID分配策略
    users = append(users, newUser)

    // 将新用户编码为JSON并写入响应
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newUser)
}

// 主函数设置路由并启动服务器
func main() {
    http.HandleFunc("/users", getUserList)
    http.HandleFunc("/users/create", createUser)

    // 设置Cobra命令的Run函数
    rootCmd.Run = func(cmd *cobra.Command, args []string) {
        fmt.Println("API server started...")
        if err := http.ListenAndServe(":8080", nil); err != nil {
            fmt.Println("Error starting server: ", err)
        }
    }

    // 执行Cobra命令
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
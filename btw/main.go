package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
    "encoding/json"
    "sort"
)

// 菜单项结构
type MenuItem struct {
    ID       string      `json:"id"`
    Title    string      `json:"title"`
    Path     string      `json:"-"`
    Children []*MenuItem  `json:"children,omitempty"`
    bookCollapseSection bool
    dir string
    cover bool
}

func main() {
    // 基础配置
    contentDir := "../content"
    outputDir := "./miniprogram/docs"

    // 创建输出目录
    os.MkdirAll(outputDir, 0755)

    // 生成菜单结构
    menu := processDirectory(contentDir)



    // 生成内容文件
    generateContentFiles(contentDir, outputDir, menu)    
    
    // 生成菜单JSON文件
    menuJSON, _ := json.MarshalIndent(menu, "", "  ")
    ioutil.WriteFile(filepath.Join(outputDir, "menu.js"), 
        []byte("export const menuConfig = "+string(menuJSON)), 0644)
}

func processDirectory(dir string) []*MenuItem {
    var items []*MenuItem
    
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Printf("Error reading directory: %v\n", err)
        return items
    }

    for _, file := range files {
        if file.IsDir() {
            // 处理目录
            children := processDirectory(filepath.Join(dir, file.Name()))
            item := &MenuItem{
                ID:       generateID(file.Name()),
                Title:    formatTitle(file.Name()),
                Children: children,
                dir: file.Name(),
            }
            items = append(items, item)
        } else if strings.HasSuffix(file.Name(), ".md") {
            // 处理markdown文件
            item := &MenuItem{
                ID:    generateID(file.Name()),
                Title: formatTitle(strings.TrimSuffix(file.Name(), ".md")),
                Path:  filepath.Join(dir, file.Name()),
            }
            items = append(items, item)
        }
    }

    return items
}


func generateContentFiles(contentDir, outputDir string, menu []*MenuItem) {
    for _, item := range menu {
        if item.Path != "" {
            // 读取markdown文件
            content, err := ioutil.ReadFile(item.Path)
            if err != nil {
                fmt.Printf("Error reading file %s: %v\n", item.Path, err)
                continue
            }

            contentStr := string(content)
            

            // 处理frontmatter
            if strings.HasPrefix(contentStr, "---") {
                endIndex := strings.Index(contentStr[3:], "---")
                if endIndex != -1 {
                    frontmatter := contentStr[3:endIndex+3]
                    // 解析frontmatter
                    for _, line := range strings.Split(frontmatter, "\n") {
                        line = strings.TrimSpace(line)
                        if strings.HasPrefix(line, "title:") {
                            title := strings.TrimPrefix(line, "title:")
                            item.Title = strings.Trim(title, "\"' ")
                        } 
                        if strings.HasPrefix(line, "bookCollapseSection:") {
                            item.bookCollapseSection = strings.Trim(strings.TrimPrefix(line, "bookCollapseSection:"), "\"' ") == "true"
                        }
                        if strings.HasPrefix(line, "cover:") {
                            item.cover = strings.Trim(strings.TrimPrefix(line, "cover:"), "\"' ") == "true" 
                        }
                    }
                    // 移除frontmatter部分
                    contentStr = contentStr[endIndex+6:]
                }
            }

            // 如果不是section，则生成内容文件
            if !item.bookCollapseSection {
                jsContent := fmt.Sprintf("module.exports = `%s`;", 
                    strings.ReplaceAll(contentStr, "`", "\\`"))
                
                // 保持原有目录结构
                relPath := strings.TrimPrefix(item.Path, contentDir)
                jsPath := strings.TrimSuffix(relPath, ".md") + ".js"
                outputPath := filepath.Join(outputDir, "content", jsPath)
                
                os.MkdirAll(filepath.Dir(outputPath), 0755)
                ioutil.WriteFile(outputPath, []byte(jsContent), 0644)
            }
        }

        // 递归处理子项
        if len(item.Children) > 0 {
            generateContentFiles(contentDir, outputDir, item.Children)
        }
    }

    // 生成内容索引文件
    indexContent := "module.exports = {\n"
    var generateIndex func(items []*MenuItem)
    generateIndex = func(items []*MenuItem) {
        for _, item := range items {
            if item.ID == "_index" {
                continue
            }
            if item.cover {
                relPath, _ := filepath.Rel(contentDir, item.Path)
                item.ID = strings.ReplaceAll(relPath, "/", "_")
                item.ID = strings.TrimSuffix(item.ID, ".md")
            }
            if item.Path != "" {
                relPath, _ := filepath.Rel(contentDir, item.Path)
                jsPath := strings.TrimSuffix(relPath, ".md")
                indexContent += fmt.Sprintf("  '%s': require('./content/%s.js'),\n", 
                    item.ID, jsPath)
            }
            generateIndex(item.Children)
            if cover := findCover(item.ID, item.Children); cover != nil {
                item.Title = ""
                relPath, _ := filepath.Rel(contentDir, cover.Path)
                jsPath := strings.TrimSuffix(relPath, ".md")
                indexContent += fmt.Sprintf("  '%s': require('./content/%s.js'),\n", 
                    item.ID, jsPath)
            }
            if book := findBookCollapseSection(item.ID, item.Children); book != nil {
                item.Title = book.Title
            }
            sort.Slice(item.Children, func(i, j int) bool {
                if item.Children[i].cover {
                    return true
                }
                return false
            })
        }
    }
    generateIndex(menu)
    indexContent += "};"
    
    ioutil.WriteFile(filepath.Join(outputDir, "contents.js"), 
        []byte(indexContent), 0644)
}

func generateID(name string) string {
    // 移除扩展名
    name = strings.TrimSuffix(name, filepath.Ext(name))
    // 转换为小写
    name = strings.ToLower(name)
    // 替换特殊字符为下划线
    name = strings.Map(func(r rune) rune {
        if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
            return r
        }
        return '_'
    }, name)
    return name
}

func formatTitle(name string) string {
    // 移除扩展名
    name = strings.TrimSuffix(name, filepath.Ext(name))
    // 替换下划线和连字符为空格
    name = strings.NewReplacer("_", " ", "-", " ").Replace(name)
    return strings.TrimSpace(name)
}

func findCover(id string, children []*MenuItem) *MenuItem {
    if len(id) <= 0 || len(children) <= 0 {
        return nil
    }
    for _, item := range children {
        if item == nil {
            continue
        }
        if item.cover {
            return item
        }
    }
    return nil
}

func findBookCollapseSection(id string, children []*MenuItem) *MenuItem {
    if len(id) <= 0 || len(children) <= 0 {
        return nil
    }
    for _, item := range children {
        if item == nil {
            continue
        }
        if item.bookCollapseSection {
            return item
        }
    }
    return nil
}
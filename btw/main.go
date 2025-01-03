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

    bookCollapseSection bool // 是否折叠一级目录
    dir string // 如果是目录，存目录名
    cover bool // 是否是封面，封面优先排
    dirLevel int // 目录层级
}

func main() {
    // 基础配置
    contentDir := "../content"
    outputDir := "./"

    // 创建输出目录
    os.MkdirAll(outputDir, 0755)

    // 生成菜单结构
    menu := processDirectory(contentDir, 1)

    // 生成内容文件
    generateContentFiles(contentDir, outputDir, menu)    
}

func removeIndex(items []*MenuItem) []*MenuItem {
    var children []*MenuItem
    if len(items) <= 0 {
        return children
    }
    for _, item := range items {
        if item == nil {
            continue
        }
        if len(item.Children) <= 0 && item.bookCollapseSection {
            continue
        }
        item.Children = removeIndex(item.Children)
        children = append(children, item)
    }
    return children
}

func processDirectory(dir string, dirLevel int) []*MenuItem {
    var items []*MenuItem
    
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Printf("Error reading directory: %v\n", err)
        return items
    }

    for _, file := range files {
        if file.IsDir() {
            // 处理目录
            children := processDirectory(filepath.Join(dir, file.Name()), dirLevel + 1)
            item := &MenuItem{
                ID:       generateID(file.Name()),
                Title:    formatTitle(file.Name()),
                Children: children,
                dir: file.Name(),
                dirLevel: dirLevel,
            }
            items = append(items, item)
        } else if file.Name() == "search.md" {
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
                    strings.ReplaceAll(strings.ReplaceAll(contentStr, "`\\`", "`\\\\`"), "`", "\\`"))
                
                // 保持原有目录结构
                relPath := strings.TrimPrefix(item.Path, contentDir)
                jsPath := strings.TrimSuffix(relPath, ".md") + ".js"
                outputPath := filepath.Join(outputDir, jsPath)
                firstPath := strings.Split(outputPath, "/")[0]
                outputPath = filepath.Join(firstPath, outputPath)
                
                os.MkdirAll(filepath.Dir(outputPath), 0755)
                ioutil.WriteFile(outputPath, []byte(jsContent), 0644)
            }
        }

        // 递归处理子项
        if len(item.Children) > 0 {
            generateContentFiles(contentDir, outputDir, item.Children)
        }
    }
    // 生成每个第一级目录的menu.js和contents.js
    for _, item := range menu {
        if item.dirLevel != 1 {
            continue
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
                    indexContent += fmt.Sprintf("  '%s': require('./%s.js'),\n", 
                        item.ID, jsPath)
                }
                generateIndex(item.Children)
                if cover := findCover(item.ID, item.Children); cover != nil {
                    item.Title = ""
                    relPath, _ := filepath.Rel(contentDir, cover.Path)
                    jsPath := strings.TrimSuffix(relPath, ".md")
                    indexContent += fmt.Sprintf("  '%s': require('./%s.js'),\n", 
                        item.ID, jsPath)
                }
                if book := findBookCollapseSection(item.ID, item.Children); book != nil {
                    item.Title = book.Title
                }
            }
        }
        generateIndex([]*MenuItem{item})
        indexContent += "};"
        
        ioutil.WriteFile(filepath.Join(outputDir, item.dir, "contents.js"), 
            []byte(indexContent), 0644)
        sort.Slice(item.Children, func(i, j int) bool {
            if item.Children[i].cover && !item.Children[j].cover {
                return true
            }
            if !item.Children[i].cover && item.Children[j].cover {
                return false
            }
            return item.Children[i].ID < item.Children[j].ID
        })
        // 生成菜单JSON文件
        item.Children = removeIndex(item.Children)
        menuJSON, _ := json.Marshal([]*MenuItem{item})
        ioutil.WriteFile(filepath.Join(outputDir, item.dir, "menu.js"), 
            []byte("export const menuConfig = "+string(menuJSON)), 0644)
    }

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
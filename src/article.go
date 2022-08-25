package main

import (
    "os"
    "path/filepath"
)

import (
    "flag"
    "fmt"
    "io/ioutil"
    "strings"
    "time"
)

var (
    cgs          []string
    cg           string
    title        string
    documentRoot string
)

func init() {
    pwd, _ := os.Getwd()

    documentRoot = filepath.Dir(pwd)

    scanLastDocs()
}

func existsLastDocs(name string) bool {
    for _, n := range cgs {
        if n == name {
            return true
        }
    }

    return false
}

func main() {
    parseArgs()

    if cg == "" || !existsLastDocs(cg) {
        flag.Usage()
        return
    }

    if title == "" {
        title = time.Now().Format("150405")
    }

    archive := time.Now().Format("2006-01")

    docsPath := docsPath()

    archivePath := fmt.Sprintf("%s/%s/articles/%s", docsPath, cg, archive)

    _, err := os.Stat(archivePath)

    if err != nil && os.IsNotExist(err) {
        os.Mkdir(archivePath, os.ModePerm)
    }

    os.Create(fmt.Sprintf("%s/%s.md", archivePath, title))
}

func parseArgs() {
    opts := strings.Join(cgs, "|")

    flag.StringVar(&cg, "c", "", fmt.Sprintf("usage (%s)", opts))

    flag.StringVar(&title, "t", "", fmt.Sprintf("usage title default by h:i:s"))

    flag.Parse()
}

func docsPath() string {
    return fmt.Sprintf("%s/docs", documentRoot)
}

func scanLastDocs() {
    docsPath := docsPath()

    dir, err := ioutil.ReadDir(docsPath)
    if err != nil {
        panic(err)
    }

    for _, fs := range dir {
        if fs.IsDir() {
            cgs = append(cgs, fs.Name())
        }
    }
}

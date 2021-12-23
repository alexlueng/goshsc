package myhttp

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/alexlueng/goshsc/internal/myhtml"
	"github.com/alexlueng/goshsc/internal/mylog"
)

type directory struct {
	Path    string
	Content []item
}

type item struct {
	URI  string
	Name string
}

type FileServer struct {
	Port    int
	Webroot string
}

func (fs *FileServer) router() {
	http.Handle("/", fs)
}

func (fs *FileServer) Start() {
	fs.router()
	log.Printf("Serving HTTP on 0.0.0.0 port %s\n", fs.Port)
	addr := fmt.Sprintf(":%+v", fs.Port)
	log.Panic(http.ListenAndServe(addr, nil))
}

func (fs *FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			http.Error(w, fmt.Sprintf("%+v", err), http.StatusInternalServerError)
		}
	}()
	fs.handler(w, r)
}

func (fs *FileServer) handler(w http.ResponseWriter, r *http.Request) {

	upath := r.URL.Path

	if upath == "/favicon.ico" {
		return
	}

	open := fs.Webroot + path.Clean(upath)
	fmt.Println(open)

	file, err := os.Open(open)
	if os.IsNotExist(err) {
		// log.Printf("Error: cannot not read file or folder: %+v", err)
		fs.handle404(w, r)
		return
	}

	if os.IsPermission(err) {
		// log.Printf("Error: cannot not read file or folder: %+v", err)
		fs.handle500(w, r)
		return
	}

	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	mylog.LogRequest(r.RemoteAddr, r.Method, r.URL.Path, r.Proto, "200")

	stat, _ := file.Stat()
	if stat.IsDir() {
		fs.processDir(w, r, file, upath)
	} else {
		fs.sendFile(w, file)
	}
}

func (fs *FileServer) processDir(w http.ResponseWriter, r *http.Request, file *os.File, relpath string) {

	fis, err := file.ReadDir(-1)
	if err != nil {
		fs.handle404(w, r)
		return
	}
	items := make([]item, 0, len(fis))
	for _, fi := range fis {
		itemname := fi.Name()
		itemuri := url.PathEscape(path.Join(relpath, itemname))
		if fi.IsDir() {
			itemname += "/"
		}
		it := item{
			Name: itemname,
			URI:  itemuri,
		}
		items = append(items, it)
	}
	sort.Slice(items, func(i, j int) bool {
		return strings.ToLower(items[i].Name) < strings.ToLower(items[j].Name)
	})

	t := template.New("index")
	t.Parse(myhtml.GetTemplate("display"))
	d := &directory{
		Path:    relpath,
		Content: items,
	}
	t.Execute(w, d)
}

func (fs *FileServer) sendFile(w http.ResponseWriter, file *os.File) {
	io.Copy(w, file)
}

func (fs *FileServer) handle404(w http.ResponseWriter, r *http.Request) {
	mylog.LogRequest(r.RemoteAddr, r.Method, r.URL.Path, r.Proto, "404")
	mylog.LogMessage("404: File not found")
	t := template.New("404")
	t.Parse(myhtml.GetTemplate("404"))
	t.Execute(w, nil)
}

func (fs *FileServer) handle500(w http.ResponseWriter, r *http.Request) {
	mylog.LogRequest(r.RemoteAddr, r.Method, r.URL.Path, r.Proto, "500")
	mylog.LogMessage("500: File not found")
	t := template.New("500")
	t.Parse(myhtml.GetTemplate("500"))
	t.Execute(w, nil)
}

func (fs *FileServer) upload() {

}

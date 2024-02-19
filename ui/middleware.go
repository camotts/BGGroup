package ui

import (
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
)

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/v1") {
			logrus.Debugf("%v -> api handler", r.URL.Path)
			handler.ServeHTTP(w, r)
			return
		}

		logrus.Debugf("%v -> static handler", r.URL.Path)

		staticPath := "out"
		indexPath := "index.html"

		pathStr := path.Join(staticPath, r.URL.Path)

		_, err := FS.Open(pathStr)
		if os.IsNotExist(err) {
			logrus.Errorf("Unable to find %v: %v", pathStr, err)
			index, err := FS.ReadFile(path.Join(staticPath, indexPath))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusAccepted)
			_, _ = w.Write(index)
			return

		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if statics, err := fs.Sub(FS, staticPath); err == nil {
			http.FileServer(http.FS(statics)).ServeHTTP(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

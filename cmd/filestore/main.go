package main

import (
	"html/template"
	"log"
	"net/http"

	"gitlab.login.no/tekkom/web/beehive/admin-api/images"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	store := images.NewFileStore("./testimages")

	http.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		file, headers, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		name := r.FormValue("name")
		if name == "" {
			name = headers.Filename
		}

		err = images.CheckImage(file, headers.Size, 3, 2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = store.UploadImage("", name, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("."))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		files, err := store.GetImages("")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		template.Must(template.New("index").
			Parse(`<!DOCTYPE html>
<html>		
	<head>
		<title>File Upload</title>
	</head>
	<body>
		<form action="http://localhost:8181/v1/images/events/banner" method="post" enctype="multipart/form-data">
			<div>
				<input type="text" name="name" />
			</div>
			<div>
				<input type="file" name="file" />
			</div>
			<input type="submit">
		</form>

		{{ range .Files }}
			<p>{{ .Name }} - {{ .Size }}</p>
			<img src="/files/{{ .Path }}" alt="{{ .Name }}" style="max-width: 100px; max-height: 70px;" />
		{{ end }}
	</body>
</html>`)).
			Execute(w, map[string]any{
				"Files": files,
			})
	})

	log.Println("Listening on :6969")
	http.ListenAndServe(":6969", nil)
}

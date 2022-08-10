package handler

import (
	"html/template"
	"julpin_1/entity"
	"net/http"
	"strconv"
)

var data = []entity.Task{}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	tmpl, err := template.ParseFiles("pages/home.html", "pages/mainlayout.html")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	if req.Method == "POST" {
		if req.FormValue("task") == "" {
			tmpl.Execute(res, data)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		task := entity.Task{}
		task.ID = len(data) + 1
		task.Task = req.FormValue("task")
		task.Description = req.FormValue("description")
		task.Deadline = req.FormValue("deadline")
		task.PIC = req.FormValue("pic")
		task.IsDone = req.PostFormValue("isDone") == "on"

		data = append(data, task)
	}

	tmpl.Execute(res, data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateHandler(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	id2, _ := strconv.Atoi(id)

	data[id2].Task = req.FormValue("task")
	data[id2].Description = req.FormValue("description")
	data[id2].PIC = req.FormValue("pic")
	data[id2].IsDone = req.PostFormValue("isDone") == "on"
	data[id2].Deadline = req.FormValue("deadline")

	http.Redirect(res, req, "/", http.StatusFound)
}

func DeleteHandler(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	id2, _ := strconv.Atoi(id)

	data = append(data[:id2], data[id2+1:]...)

	http.Redirect(res, req, "/", http.StatusFound)
	// fmt.Println(data)
}

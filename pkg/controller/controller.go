package controller

import (
	"fmt"
	"go_shapes/pkg/utils"
	"net/http"
	"os/exec"
	"strconv"
	"text/template"

	svg "github.com/ajstarks/svgo"
)

type Command struct {
	Output string
}

func HandleForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./static/form.html")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	command := r.FormValue("command")
	cmd := exec.Command("sh", "-c", command)
	stdout, err := cmd.Output()
	c := Command{}
	if err != nil {
		c.Output = err.Error()
	} else {
		c.Output = string(stdout)
	}
	tmpl.Execute(w, c)
}

func GetShape(w http.ResponseWriter, r *http.Request) {
	var number string = r.URL.Query().Get("num")
	shape, err := utils.DefineShape(number)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(600, 600)
	if shape.Shape == 0 {
		s.Circle(shape.H, shape.H, (shape.H)/2, fmt.Sprintf("fill:#%s;stroke:black", shape.Color))
	} else {
		s.Rect(100, 100, shape.W, shape.H, fmt.Sprintf("fill:#%s;stroke:black", shape.Color))
	}
	s.End()
}

func GetSortedArray(w http.ResponseWriter, r *http.Request) {
	var strArr string = r.URL.Query().Get("str")
	arr, err := utils.StringToArray(strArr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	utils.QSort(arr, 0, len(arr)-1)
	res := []byte{}
	for i := 0; i < len(arr); i++ {
		res = append(res, []byte(strconv.Itoa(arr[i]))...)
		res = append(res, 32)
	}
	w.Write(res)
}

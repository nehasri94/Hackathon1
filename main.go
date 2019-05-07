package main

import (
	"fmt"
	"html/template"
	"mux"
	"net/http"
)

type employee struct {
	EName    string
	EId      string
	EMail    string
	EContact string
	ESkills  string
}
type student struct {
	SName       string
	SId         string
	SMail       string
	SDepartment string
	SContact    string
}

type enode struct {
	edata employee
	en    *enode
}
type snode struct {
	sdata student
	sn    *snode
}


func screate() *snode {

	sn1 := snode{}

	return &sn1
}
func (sn1 *snode) sinsert(e student) {
	sn2 := screate()

	for sn1.sn != nil {
		sn1 = sn1.sn
	}
	sn2.sdata = e
	sn2.sn = nil
	sn1.sn = sn2

}
func (sn1 *snode) sdisplay() *snode{
	for sn1.sn != nil {

		sn1 = sn1.sn


	}
	return sn1
}

func ecreate() *enode {

	en1 := enode{}

	return &en1
}
func (en1 *enode) einsert(e employee) {
	en2 := ecreate()

	for en1.en != nil {
		en1 = en1.en
	}
	en2.edata = e
	en2.en = nil
	en1.en = en2

}
func (en1 *enode) edisplay() *enode{
	for en1.en != nil {

		en1 = en1.en


	}
	return en1
}

func ins1(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("form_emp.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := employee{
		EName:    r.FormValue("EName"),
		EId:      r.FormValue("EId"),
		ESkills:  r.FormValue("ESkills"),
		EContact: r.FormValue("EContact"),
		EMail:    r.FormValue("EMail"),
	}


	lis:=ecreate()
	lis.einsert(details)
	t1:=lis.edisplay()

	t2:="<html><body><h3>Your data has been recorded</h3></br>Name: "+t1.edata.EName+"</br>Employee Id: "+t1.edata.EId+"</br>Skills: "+t1.edata.ESkills+"</br>Contact Number: "+t1.edata.EContact+"</br>Email Id: "+t1.edata.EMail+"</body></html>"
	fmt.Fprintln(w,t2)

	tmpl.Execute(w, struct{ Success bool }{true})
}

func ins2(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("form_student.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := student{
		SName:    r.FormValue("SName"),
		SId:      r.FormValue("SId"),
		SDepartment:  r.FormValue("SDepartment"),
		SContact: r.FormValue("SContact"),
		SMail:    r.FormValue("SMail"),

	}


	lis:=screate()
	lis.sinsert(details)
	t1:=lis.sdisplay()

	t2:="<html><body><h3>Your data has been recorded</h3></br>Name: "+t1.sdata.SName+"</br>Student Id: "+t1.sdata.SId+"</br>Department: "+t1.sdata.SDepartment+"</br>Contact Number: "+t1.sdata.SContact+"</br>Email Id: "+t1.sdata.SMail+"</body></html>"
	fmt.Fprintln(w,t2)

	tmpl.Execute(w, struct{ Success bool }{true})
}


func main() {

	m := mux.NewRouter()


	m.HandleFunc("/employee",ins1)
	m.HandleFunc("/student",ins2)


	http.ListenAndServe(":8080", m)

}

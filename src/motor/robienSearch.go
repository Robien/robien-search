package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Page struct {
	Question string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("question.html"))
var id int


func questionHandler(w http.ResponseWriter, r *http.Request) {

	var needNewQuestion = false
	var c *Concept
	c = motor.GetCurrentConcept(id)
	if c==nil{
		needNewQuestion = true
	}
	
	if (r.FormValue("reset") != ""){
		motor.Reset(id)
		needNewQuestion = true
		
		fmt.Fprintf(w, "<h1>il y a %v éléments, %v concepts et %v liens</h1>", len(motor.elements), len(motor.concepts), motor.nbLink)
	}
	if (r.FormValue("answer") != ""){
		f, _ := strconv.ParseFloat(r.FormValue("answer"), 64)
		f = f/100
		motor.Propage(id, f)
		needNewQuestion = true
	}
	if (r.FormValue("addQuestion") != ""){
		c = motor.ConceptFactory("NONAME", r.FormValue("other"))
		motor.ForceQuestion(id, c)
	}
	if (r.FormValue("found") != ""){
	var idFound int
		idFound, _ = strconv.Atoi(r.FormValue("found"))
		if idFound == -1{
			newElem := motor.ElementFactory(r.FormValue("other"))
			idFound = newElem.Id
		}
		motor.Found(id, idFound)
		motor.Reset(id)
		needNewQuestion = true
		
		fmt.Fprintf(w, "<h1>il y a %v éléments, %v concepts et %v liens</h1>", len(motor.elements), len(motor.concepts), motor.nbLink)
	}
	
	propositions := motor.GetListElement(id)
	
	
	if needNewQuestion {
		c = motor.GetNextConcept(id)
	}

	
	if c!=nil{
		q := c.GetQuestion()
		var p *Page = new (Page)
		p.Question = q
		renderTemplate(w, "question", p)
	}
	


	fmt.Fprintf(w, "<br /><br />")
	for _, n := range propositions{
		fmt.Fprintf(w, "<form action=\"/\" method=\"POST\">%v (%g) <input type=\"hidden\" name=\"found\" value=\"%v\"><input type=\"submit\" value=\"Selectionner\"></form>", n.name, n.proba, n.id)
	}
	fmt.Fprintf(w, "<form action=\"/\" method=\"POST\">Rien de tout cela :  <input type=\"text\" name=\"other\" value=\"\"><input type=\"hidden\" name=\"found\" value=\"-1\"><input type=\"submit\" value=\"Ajouter et selectionner\"></form>")
	fmt.Fprintf(w, "<br /><br />")
	fmt.Fprintf(w, "<form action=\"/\" method=\"POST\">J'aurais aimé qu'on me demande :  <input type=\"text\" name=\"other\" value=\"\"><input type=\"hidden\" name=\"addQuestion\" value=\"-1\"><input type=\"submit\" value=\"Ajouter\"></form>")
	fmt.Fprintf(w, "<form action=\"/\" method=\"POST\">Recommencer une recherche<input type=\"hidden\" name=\"reset\" value=\"-2\"><input type=\"submit\" value=\"valider\"></form>")
	
}





func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fn(w, r)
	}
}
func test(){
	fruit := motor.ConceptFactory("fruit", "Est-ce un fruit ?")
	legume := motor.ConceptFactory("legume", "Est-ce un legume ?")
	mangeable := motor.ConceptFactory("mangeable", "Est-ce que ça se mange ?")
	outil := motor.ConceptFactory("outil", "Est-ce un outil ?")
	animal := motor.ConceptFactory("animal", "Est-ce un animal ?")
	rond := motor.ConceptFactory("rond", "Est-ce rond ?")
	long := motor.ConceptFactory("long", "Est-ce long ?")
	
	pomme := motor.ElementFactory("pomme")
	banane := motor.ElementFactory("banane")
	poireau := motor.ElementFactory("poireau")
	tournevis := motor.ElementFactory("tournevis")
	vache := motor.ElementFactory("vache")
	chat := motor.ElementFactory("chat")
	chien := motor.ElementFactory("chien")
	marteau := motor.ElementFactory("marteau")

	motor.LinkFactory(pomme, fruit, 1)
	motor.LinkFactory(pomme, mangeable, 1)
	motor.LinkFactory(pomme, rond, 0.9)

	motor.LinkFactory(banane, fruit, 1)
	motor.LinkFactory(banane, mangeable, 1)
	motor.LinkFactory(banane, long, 0.8)

	motor.LinkFactory(poireau, legume, 1)
	motor.LinkFactory(poireau, mangeable, 1)
	motor.LinkFactory(poireau, long, 0.9)

	motor.LinkFactory(tournevis, outil, 1)
	motor.LinkFactory(tournevis, long, 0.9)

	motor.LinkFactory(vache, animal, 1)
	motor.LinkFactory(vache, mangeable, 0.7)

	motor.LinkFactory(chat, animal, 1)
	motor.LinkFactory(chien, animal, 1)

	motor.LinkFactory(marteau, outil, 1)
	
	id = motor.CreateNewResearch()
}

func main() {
test()
	http.HandleFunc("/", makeHandler(questionHandler))
	http.ListenAndServe(":8080", nil)
}

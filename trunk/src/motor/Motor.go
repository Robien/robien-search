package main

//import "fmt"
//import "strconv"
import 	"sort"

type Motor struct {

	elements []*Element
	concepts []*Concept
	research []*Research
	links []*Link
	nbLink int
	
}

var motor Motor

func (m *Motor) LinkFactory(e *Element, c *Concept, proba float64)  {
	l := new(Link)
	l.Create(c, e)
	l.value = proba
	l.weight = 1
	m.nbLink++
	m.links = append(m.links, l)
}

func (m *Motor) LinkFactoryWithWeight(e *Element, c *Concept, proba float64, weight float64)  {
	l := new(Link)
	l.Create(c, e)
	l.value = proba
	l.weight = weight
	m.nbLink++
	m.links = append(m.links, l)
}

func (m *Motor) ElementFactory(name string)  *Element{
	e := new(Element)
	e.Name = name
	m.AddAnElement(e)
	return e
}
func (m *Motor) ConceptFactory(name string, question string)  *Concept{
	c := new(Concept)
	c.Name = name
	c.Question = question
	m.AddAConcept(c)
	return c
}

func (m *Motor) AddAnElement(e *Element)  {
	e.Id = len(m.elements)
	m.elements = append(m.elements, e)
}
func (m *Motor) AddAConcept(c *Concept)  {
	c.Id = len(m.concepts)
	m.concepts = append(m.concepts, c)
}
func (m *Motor) GetAConcept(id int) *Concept {
	return m.concepts[id]
}
func (m *Motor) GetAElement(id int) *Element {
	return m.elements[id]
}
func (m *Motor) CreateNewResearch() int {
	r := new(Research)
	m.research = append(m.research, r)
	r.Prepare(m.elements, m.concepts)
	return len(m.research)-1
}
func (m *Motor) Reset(id int) {
	m.research[id] = new(Research)
	m.research[id].Prepare(m.elements, m.concepts)
}

func (m *Motor) GetNextConcept(id int) *Concept {
	return m.research[id].GetNextConcept()
}
func (m *Motor) GetCurrentConcept(id int) *Concept {
	return m.research[id].currentConcept
}
func (m *Motor) ForceQuestion(id int, c *Concept)  {
	m.research[id].currentConcept = c
}

func (m *Motor) Propage(id int, note float64) {
	m.research[id].Propage(note)
}

func (m *Motor) Found(id int, idElement int) {
	m.research[id].Found(m.elements[idElement])
}

type Proposition struct {
	name string
	proba float64
	id int
	
}
type ByProba []*Proposition

func (a ByProba) Len() int           { return len(a) }
func (a ByProba) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByProba) Less(i, j int) bool { return a[i].proba > a[j].proba }

func (m *Motor) GetListElement(id int) []*Proposition {
	p := make([]*Proposition, len(m.elements))

	for i, e := range m.elements{
		p[i] = new(Proposition)
		p[i].name = e.Name
		p[i].proba = m.research[id].proba[e.Id]
		p[i].id = i
	}
	
	sort.Sort(ByProba(p))
		
	m.research[id].lastPropositions = &p
	
	return p
}
func (m *Motor) save(){

	saveElements(m.elements)
	saveConcepts(m.concepts)
	saveLinks(m.links)
	
}
func (m *Motor) load(){

	loadElements()
	loadConcepts()
	loadLinks()
	
}


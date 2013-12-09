package main

import "fmt"
import "math"
import "sort"

type Research struct {
	proba []float64
	nbModified []int
	currentConcept *Concept
	answers []*Answer
	sortedConcepts []*Concept
	lastPropositions *[]*Proposition
}



func (l *Research) Prepare(elements []*Element, concepts []*Concept)  {
	for _, e := range elements{
		if e.Id != len(l.proba){
			fmt.Println("sp'a normal")
		}
		l.proba = append(l.proba, 0.0)
		l.nbModified = append(l.nbModified, 0)
	}
	for _, c := range concepts{
		l.sortedConcepts = append(l.sortedConcepts, c)
	}
	l.lastPropositions = nil
	l.currentConcept = nil
}
func (r *Research) GetNextConcept() *Concept{
	if len(r.sortedConcepts) == 0{
		r.sortedConcepts = nil
		return nil
	}
	t := r.sort()
	fmt.Println("----------------------------")
	for _,e := range r.sortedConcepts{
		fmt.Println(e.Question)
		fmt.Println(e.GetUtilitysortedElement(&t))
	}
	fmt.Println("----------------------------")
	r.currentConcept = r.sortedConcepts[0]
	r.sortedConcepts = r.sortedConcepts[1:len(r.sortedConcepts)]
	return r.currentConcept

//	if r.nbCurentConcept >= len(c)-1{
//		return nil
	//}
	//r.nbCurentConcept++
	//r.currentConcept = c[r.nbCurentConcept]
	//return r.currentConcept
}
func (r *Research) Propage(note float64) {

	for _, l := range r.currentConcept.links{
		//r.proba[l.element] = ((r.proba[l.element]*float64(r.nbModified[l.element]))+(1 - 2 * (math.Abs(l.value - note))))/(float64(r.nbModified[l.element])+1)
		r.proba[l.element] = ((r.proba[l.element])+(1 - 2 * (math.Abs(l.value - note))))
		r.nbModified[l.element] += 1
	}
	a := new(Answer)
	a.concept = r.currentConcept
	a.note = note
	r.answers = append(r.answers, a)
}

func (r *Research) Found(e *Element) {
	for _, a := range r.answers{
		a.Found(e)
	}
	r.lastPropositions = nil
	r.currentConcept = nil
	r.answers = make([]*Answer, 0)

}


type ByUtility []*Concept

func (a ByUtility) Len() int           { return len(a) }
func (a ByUtility) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUtility) Less(i, j int) bool { return a[i].GetUtility() > a[j].GetUtility() }

type conceptsToSort struct {
	concept *Concept
	sortedIdElement *[]int
}

type ByUtilityElement []*conceptsToSort

func (a ByUtilityElement) Len() int           { return len(a) }
func (a ByUtilityElement) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUtilityElement) Less(i, j int) bool { return a[i].concept.GetUtilitysortedElement(a[i].sortedIdElement) > a[j].concept.GetUtilitysortedElement(a[i].sortedIdElement) }



func (r *Research) sort() []int {

var t []int

	if r.lastPropositions!= nil{
		t = make([]int, len(*r.lastPropositions))
		for i, e := range *r.lastPropositions{
			t[e.id] = i
		}
		fmt.Println("id1: ",  motor.elements[t[0]].Name)
		fmt.Println("id1: ", motor.elements[t[1]].Name)
	}


	p := make([]*conceptsToSort, len(r.sortedConcepts))

	
	for i, e := range r.sortedConcepts{
		p[i] = new(conceptsToSort)
		p[i].concept = e
		if r.lastPropositions!= nil{
			p[i].sortedIdElement = &t
		}else{
			p[i].sortedIdElement = nil
		}
		
	}

	sort.Sort(ByUtilityElement(p))
	
	for i, e := range p{
		r.sortedConcepts[i] = e.concept
	}
	
	return t
	
}
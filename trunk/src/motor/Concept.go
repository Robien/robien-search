package main

import "fmt"
import "math"

type Concept struct {
	Name string
	Question string
	Id int
	links []*Link
}


func (c *Concept) Print()  {
	fmt.Println(c.Name)
}

func (c *Concept) PrintDebug()  {
	fmt.Println(c.Name + " (" , c.Id , ")")
}
func (c *Concept) PrintQuestion()  {
	fmt.Println(c.Question)
}
func (c *Concept) GetQuestion()  string{
	return c.Question
}
func (c *Concept) AddALink(l *Link)  {
	c.links = append(c.links, l)
}
func (c *Concept) GetLinkWithAnElement(e *Element) *Link {
	i := 0
	for i<len(c.links) && e.Id != c.links[i].element{
		i++
	}
	if i==len(c.links){
		return nil
	}else{
		return c.links[i]
	}
}
func (c *Concept) GetUtility()  float64 {
	score := 0.0
	for _,l := range c.links{
		score += math.Abs(l.value - 0.5)
	}
	return score
}

func (c *Concept) GetUtilitysortedElement(sortedElementId *[]int, proba *[]float64)  float64 {
	score := 0.0
	first := -1.0
	second := -1.0
	bonus := 0
	
	for _,l := range c.links{
		if sortedElementId != nil{
			if ((*sortedElementId)[l.element] == 0){
				first = l.value
			} else if ((*sortedElementId)[l.element] == 1){
				second = l.value
			}
			if (*proba)[l.element] < 0{
				bonus--
			} else if (*proba)[l.element] > 0{
				bonus++
			}
		}
		
		score += math.Abs(l.value - 0.5)
	}
	score = score / (math.Pow(2.0,float64(bonus)))
	if first != -1 && second != -1 {
		score *=  (4*math.Abs(first - second)*math.Abs(first - second))
		//fmt.Println(c.Name, " score :", score)
	}
	return score
}
package main

type Answer struct {
concept *Concept
note float64

}

func (a *Answer) Found(e *Element) {
	l := a.concept.GetLinkWithAnElement(e)
	if l==nil{
		motor.LinkFactory(e, a.concept,a.note)
	}else{
		value := l.value
		w := l.weight
		l.value = (value * w + a.note) / (w + 1)
		l.weight = w+1
	}

}
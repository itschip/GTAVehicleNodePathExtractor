package main

type Scene struct {
	Objects Objects `xml:"objects"`
}

type Objects struct {
	ObjectRoot []ObjectRoot `xml:"object"`
}

type ObjectRoot struct {
	Class     string    `xml:"class,attr"`
	Transform Transform `xml:"transform"`
}

type Transform struct {
	Object Object `xml:"object"`
}

type Object struct {
	Position Position `xml:"position"`
}

type Position struct {
	X string `xml:"x,attr"`
	Y string `xml:"y,attr"`
}

package janus

type graph struct {
	graphName   string  `xml:"graphName"`
	datasetPath string  `xml:"datasetPath"`
	jobID       string  `xml:"jobID"`
	nodes       *[]node `xml:"nodes"`
}

type node struct {
	nodeID     string  `xml:"nodeID"`
	language   string  `xml:"language"`
	scriptFile string  `xml:"scriptFile"`
	className  string  `xml:"className"`
	args       *[]arg  `xml:"args"`
	edges      *[]edge `xml:"edges"`
	antiEdges  *[]edge `xml:"antiEdges"`
}

type arg struct {
	source string `xml:"source"`
	value  string `xml:"value"`
}

type edge struct {
	string `xml:"edge"`
}

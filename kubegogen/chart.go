package kubegogen

type Chart interface {
	Construct
	Namespace() string
}

type ChartProps struct {
	Namespace string
}

type chart struct {
	*construct
	props ChartProps
}

func (c *chart) Namespace() string {
	return c.props.Namespace
}

func newChart(id string, props ChartProps, parentNode *node[Construct]) Chart {
	// fmt.Println("AddChart", c.node.FullID(), id)
	chart := &chart{construct: &construct{}, props: props}
	chart.construct.node = parentNode.AddChildNode(id, chart)
	if props.Namespace != "" {
		chart.SetContext("namespace", props.Namespace)
	}
	// parentNode.AddChildNode(id, chart)
	return chart
}

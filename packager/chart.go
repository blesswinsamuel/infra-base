package packager

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

func (c *construct) Chart(id string, props ChartProps) Chart {
	// fmt.Println("AddChart", c.node.FullID(), id)
	chart := &chart{construct: &construct{}, props: props}
	chart.construct.node = c.node.AddChildNode(id, chart)
	if props.Namespace != "" {
		chart.SetContext("namespace", props.Namespace)
	}
	c.node.AddChildNode(id, chart)
	return chart
}

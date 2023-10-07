package renderer

var R Renderer

type Renderer struct {
	n  Navigator
}

func init() {
	R = Renderer{
		n: n,
	}
}

func (r Renderer) Render() error {
	for {
		err := r.n.show()
		if err != nil {
			return err
		}

		err = r.n.navigate()
		if err != nil {
			return err
		}
	}
}
package react

// Define reactor, cell and canceler types here.
// These types will implement the Reactor, Cell and Canceler interfaces, respectively.
type cell struct {
	value int
}

// inputCell and computeCell both embed cell type hence also qualifying for implementing Cell interface
type inputCell struct {
	cell
}

type computeCell struct {
	cell

	// the cells on which the value of this cell depends
	deps []Cell

	callbacks map[int](func(int))
}

// cell implements Cell
func (c *cell) Value() int {
	return c.value
}

func (c *inputCell) SetValue(newValue int) {
	c.value = newValue

}

func (c *computeCell) AddCallback(f func(int)) Canceler {
	index := len(c.callbacks)
	c.callbacks[index] = f

	return &canceler{
		f: func() {
			delete(c.callbacks, index)
		},
	}
}

type canceler struct {
	f func()
}

// canceler now implements Canceler
func (c *canceler) Cancel() {
	// just call the function
	c.f()
}

type reactor struct {
	adjList map[Cell][]Cell
}

func (*reactor) Update(cell Cell) {

}

func New() Reactor {
	panic("not implemented")
}

func (r *reactor) CreateInput(initial int) InputCell {
	panic("not implemented")
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	panic("not implemented")
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	panic("not implemented")
}

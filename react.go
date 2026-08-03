package react

// Define reactor, cell and canceler types here.
// These types will implement the Reactor, Cell and Canceler interfaces, respectively.
type cell struct {
	value     int
	isCompute bool
	_reactor  *reactor
	compute1  func(int) int      // defines how the value of compute cell with a change to an input cell
	compute2  func(int, int) int // same as above but for case 2
	callbacks map[int](func(int))
}

// cell implements Cell
func (c *cell) Value() int {
	return c.value
}

// *cell now implements InputCell
func (c *cell) SetValue(newValue int) {
	c.value = newValue
	if c.isCompute == false {
		// if compute, propagate change to its dependencies
		for inputCell, computeCells := range c._reactor.dependencies {
			if inputCell == c {
				for _, cell_t := range computeCells {
					// assert that the object implementing Cell is cell struct
					computeCell := (cell_t).(*cell)
					computeCell.value = newValue

					// execute this computeCell's callbacks
					for _, fCallback := range computeCell.callbacks {
						fCallback(newValue)
					}
				}
			}
		}
	}
}

// *cell now implements ComputeCell
func (c *cell) AddCallback(fCallback func(int)) Canceler {
	callbackId := len(c.callbacks)
	c.callbacks[callbackId] = fCallback

	return &canceler{
		f: func() {
			delete(c.callbacks, callbackId)
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
	dependencies map[Cell]([]Cell)
}

func New() Reactor {
	return &reactor{}
}

func (r *reactor) CreateInput(initial int) InputCell {
	return &cell{
		_reactor:  r,
		value:     initial,
		isCompute: false,
	}
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {

	newComputeCell := &cell{
		_reactor:  r,
		compute1:  compute,
		isCompute: true,
	}

	// inputCell := dep.(InputCell)

	r.dependencies[dep] = append(r.dependencies[dep], newComputeCell)

	return newComputeCell // return a zero initailzed cell
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {

	newComputeCell := &cell{
		_reactor:  r,
		compute2:  compute,
		isCompute: true,
	}

	r.dependencies[dep1] = append(r.dependencies[dep2], newComputeCell)
	r.dependencies[dep2] = append(r.dependencies[dep2], newComputeCell)

	return newComputeCell // return a zero initailzed cell
}

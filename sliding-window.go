package main

// Window : Sliding Window data type
type Window struct {
	length int   // length of the stack
	mirror []int // unordered list of values
	size   int   // size of the sliding window
	stack  []int // ordered list of values
}

// AddDelay : adds a delay value to the stack
func (w *Window) AddDelay(delay int) {
	// check if stack of values is full
	if w.full() {
		w.cut()
	}

	// no values have been stored yet or the value goes to the end of the stack
	if w.length < 1 || w.stack[w.length-1] <= delay {
		w.stack = append(w.stack, delay)
		w.copy(delay)
		return
	}

	// check if the value goes to the beginning of the stack
	if w.stack[0] > delay {
		w.stack = append([]int{delay}, w.stack...)
		w.copy(delay)
		return
	}

	// find the offset position and insert the value in the middle
	offset := w.offset(delay)
	w.stack = append(w.stack[:offset], append([]int{delay}, w.stack[offset:]...)...)
	w.copy(delay)
}

// mirror : mirror the stack values in their order of arrival
func (w *Window) copy(val int) {
	w.mirror = append(w.mirror, val)
	w.length++
}

// cut : cut the oldest value from the stack
func (w *Window) cut() {
	target := w.mirror[0]
	w.mirror = w.mirror[1:]

	for key, val := range w.stack {
		if val == target {
			w.stack = append(w.stack[:key], w.stack[key+1:]...)
			break
		}
	}

	w.length--
}

// full : check if the window is full
func (w *Window) full() bool {
	if w.length < w.size {
		return false
	}

	return true
}

// Median : calculate the median value from the stack of values
func (w *Window) Median() int {
	// check if there is only 1 or less values in the stack
	if w.length < 2 {
		return -1
	}

	// calculate the median from an odd number of values
	if w.length%2 == 1 {
		l := w.length + 1
		idx := l/2 - 1
		return w.stack[idx]
	}

	// calculate the median from the even number of values
	idx := w.length/2 - 1
	val := w.stack[idx]
	val += w.stack[idx+1]
	return val / 2
}

// offset : find the offset for a delay value in the stack
func (w *Window) offset(val int) int {
	for key, v := range w.stack {
		if v > val || key == w.length-1 {
			return key
		}
	}

	return -1
}

// Size : set the size of the sliding window
func (w *Window) Size(val int) {
	w.size = val
}

// NewSlidingWindow : Create a new Sliding Window object
func NewSlidingWindow() *Window {
	win := &Window{}
	win.length = 0
	return win
}

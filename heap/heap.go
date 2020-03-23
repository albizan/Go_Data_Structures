package heap

import "fmt";

/*
 * HEAP
 * Un HEAP è una struttura dati rappresentabile con un Albero Binario completo
 * Proprietà: Per ogni nodo che non sia la radice, il valore del nodo è sempre minore del valore del nodo padre
 * Utile per implementare Code a priorità (Priority Queque)
 * Essendo un albero binario completo, una heap è facile da implementare come array
 * Esempio: heap:= []int -> dato heap[i] il figlio sx è heap[2i+1] mentre il figlio destro è heap[2i+2]
*/

// Heap è una struct che contiene uno slice di interi, la dimensione corrente e la capacità totale dello slice
type Heap struct {
	items []int;
	currentSize int;
	capacity int;
}

// MakeHeap crea una nuova Heap
func MakeHeap(capacity int) *Heap {
	heap := &Heap{
		items: []int{},
		currentSize: 0,
		capacity: capacity,
	}

	return heap;
}

// Verify if an element is a leaf of the tree
func (heap *Heap) isLeaf(index int) bool {
	// Get index of the left child, if left child's index is greater than heap's size, node has no child so node is a leaf
	maxIndex := heap.currentSize-1;
	leftChildIndex := index*2 +1;
	if leftChildIndex > maxIndex {
		return true
	}

	return false;
}

// Return the parent (in this case the index of the parent)
func (heap *Heap) getParentIndex(index int) int {
	if index == 0 {
		return index
	}
	parentIndex := (index-1) / 2;
	return parentIndex;
}

// Return left child (in this case the index of the left child)
func (heap *Heap) getLeftChildIndex(index int) int {
	// Child does not exist
	if heap.isLeaf(index) {
		return -1
	}

	return index*2 +1;
}

// Return right child (in this case the index of the right child)
func (heap *Heap) getRightChildIndex(index int) int {
	// Child does not exist
	if heap.isLeaf(index) {
		return -1
	}
	maxIndex := heap.currentSize-1;
	rigtChildIndex := index*2 +2;
	if rigtChildIndex > maxIndex {
		return -1
	}
	return rigtChildIndex;
}

func (heap *Heap) swap(index1, index2 int) {
	tmp := heap.items[index1];
	heap.items[index1] = heap.items[index2];
	heap.items[index2] = tmp;
}

// Insert is used to insert new items to the heap
// A new item gets added at the end of the underlying slice, thus breaking the structure of the heap
// The new item gets compared with his parent and swapped untill his vale is less than his parent's value
func (heap *Heap) Insert(newItem int) error {
	if(heap.capacity == heap.currentSize) {
		return fmt.Errorf("ERRORE: Heap piena")
	}
	heap.items = append(heap.items, newItem)
	fmt.Println(newItem, "inserted")
	heap.currentSize++;
	index := heap.currentSize -1;
	if heap.items[index] <= heap.items[heap.getParentIndex(index)] {
		fmt.Println("Inserted value is smaller than his parent, no need to swap")
		fmt.Println()
		return nil
	}
	for heap.items[index] > heap.items[heap.getParentIndex(index)] {
		fmt.Println("Item:", heap.items[index])
		fmt.Println("Parent:", heap.items[heap.getParentIndex(index)])
		fmt.Println("Swapping", heap.items[index], "with", heap.items[heap.getParentIndex(index)])
		fmt.Println()
		heap.swap(index, heap.getParentIndex(index))
		index = heap.getParentIndex(index)
	}
	return nil
}

// GetItems returns the underlying slice of heap
func (heap Heap) GetItems() []int {
	items :=heap.items
	return items
}

// Heapify is a function that restores the structure of a heap considering Root the element with given index
func (heap *Heap) Heapify(index int) {
	if heap.isLeaf(index) {
		return
	}
	largest := index
  leftChildIndex := heap.getLeftChildIndex(index)
	rightChildIndex := heap.getRightChildIndex(index)
	if leftChildIndex != -1  {
		if heap.items[leftChildIndex] > heap.items[largest] {
			largest = leftChildIndex
		}
	}

	if rightChildIndex != -1  {
		if heap.items[rightChildIndex] > heap.items[largest] {
			largest = rightChildIndex
		}
	}
	if largest != index {
		heap.swap(index, largest)
		heap.Heapify(largest)
	}
	return
}

// Pop returns the root node and delete root item from the heap, proceeds to regenerting heap properties with Heapify on new "root"
func (heap *Heap) Pop() int {
	root := heap.items[0]
	heap.items[0] = heap.items[heap.currentSize-1]
	// set items to new smaller slice
	heap.items = heap.items[:heap.currentSize-1]
	heap.currentSize--
	// To regenerate a heap, heapify from new root
	heap.Heapify(0)
	return root;
}

// MakeHeapFromArray generates a new heap based on given array
func MakeHeapFromArray(array []int) *Heap {
	h := MakeHeap(len(array))

	for i := 0; i < len(array); i++ {
		error := h.Insert(array[i])
	  if(error != nil) {
		  fmt.Println(error)
	  }
	}

	return h
}

// Heapsort sorts the array given as input 
func Heapsort(inputArray []int) []int {
	h := MakeHeapFromArray(inputArray)
	sortedArray := make([]int, h.currentSize)

	for h.currentSize > 0 {
		// Pop decrease h.currentsize by one
		sortedArray[h.currentSize] = h.Pop()
	}

	return sortedArray
}
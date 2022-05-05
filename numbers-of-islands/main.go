package main

// https://leetcode.com/problems/number-of-islands/

type Point struct {
	x, y int
}

type PointQueue struct {
	queue []*Point
}

func NewPointQueue() *PointQueue {
	return &PointQueue{
		queue: make([]*Point, 0),
	}
}

func (q *PointQueue) enqueue(element *Point) {
	if q.queue == nil {
		q.queue = make([]*Point, 0)
	}
	q.queue = append(q.queue, element)
}

func (q *PointQueue) dequeue() *Point {
	if len(q.queue) == 0 {
		return nil
	}
	result := q.queue[0]
	q.queue = q.queue[1:]
	return result
}

func (q *PointQueue) peek() *Point {
	if len(q.queue) == 0 {
		return nil
	}
	result := q.queue[0]
	return result
}

func (q *PointQueue) isEmpty() bool {
	return len(q.queue) == 0
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	count := 0
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	var bfs func(x, y int)
	bfs = func(x, y int) {
		q := NewPointQueue()
		q.enqueue(&Point{x, y})
		visited[x][y] = true

		for !q.isEmpty() {
			p := q.dequeue()
			for _, t := range directions {
				offsetX, offsetY := t[0], t[1]
				r, c := p.x+offsetX, p.y+offsetY
				if r >= 0 &&
					c >= 0 &&
					r < rows &&
					c < cols &&
					!visited[r][c] &&
					grid[r][c]&1 == 1 {
					visited[r][c] = true
					q.enqueue(&Point{r, c})
				}
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j]&1 == 1 && !visited[i][j] {
				bfs(i, j)
				count++
			}
		}
	}

	return count
}

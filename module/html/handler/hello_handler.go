package handler

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo"
)

// Response ...
type Response struct {
	Message string
}

func sleepyGopher(id int, ch chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	log.Println(fmt.Sprintf("... %d snore ...", id))
	ch <- id
}

func select1() {
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go sleepyGopher(i, ch)
	}

	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-ch:
			log.Println(fmt.Sprintf("gopher %d はスリープを終えました", gopherID))

		case <-timeout:
			log.Println("忍耐の限度に達した！")
		}
	}
}

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	// downstream <- ""
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	// for {
	// 	item, ok := <-upstream
	// 	if !ok {
	// 		close(downstream)
	// 		return
	// 	} else if !strings.Contains(item, "bad") {
	// 		downstream <- item
	// 	}
	// }

	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	// for {
	// 	v := <-upstream
	// 	if v == "" {
	// 		return
	// 	}
	// 	log.Println(v)
	// }
	for v := range upstream {
		log.Println(v)
	}
}

func stream() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}

// Visited ...
type Visited struct {
	mu      sync.Mutex
	visited map[string]int
}

// VisitURL ...
func (v *Visited) VisitURL(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

func (v *Visited) crawl(url string) {
	// time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	time.Sleep(3000 * time.Millisecond)
	v.VisitURL(url)
	log.Println(fmt.Sprintf("url: %s %d", url, v.visited[url]))
}

func visit() {
	v := &Visited{sync.Mutex{}, map[string]int{}}
	// urls := []string{"a", "a", "a", "a", "a", "b", "b", "b", "c", "c", "c", "c", "c", "d", "d", "d", "d", "d", "d", "d", "d", "d"}
	urls := []string{"a", "a", "a", "a", "a", "a", "a", "a"}
	for _, url := range urls {
		go v.crawl(url)
	}
	time.Sleep(4 * time.Second)
	log.Println("passed")
}

// RoverDriver ...
type RoverDriver struct {
	commandc chan command
}

// NewRoverDriver ...
func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

type command int

const (
	right = command(0)
	left  = command(1)
)

func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	done := time.After(10 * time.Second)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			}
			log.Println(fmt.Sprintf("new direction %v", direction))
		case <-nextMove:
			pos = pos.Add(direction)
			log.Println(fmt.Sprintf("moved to %v", pos))
			nextMove = time.After(updateInterval)
		case <-done:
			return
		}
	}
}

func worker() {
	// pos := image.Point{X: 10, Y: 10}
	// direction := image.Point{X: 1, Y: 0}
	// next := time.After(time.Second)
	// done := time.After(10 * time.Second)
	// for {
	// 	select {
	// 	case <-next:
	// 		pos = pos.Add(direction)
	// 		log.Println(fmt.Sprintf("current position is %v", pos))
	// 		next = time.After(time.Second)
	// 	case <-done:
	// 		return
	// 	}
	// }
	NewRoverDriver()
}

// HelloHandler ...
func HelloHandler(c echo.Context) error {
	message := "Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!"

	// select1()
	// stream()
	// visit()
	// go worker()

	data := map[string]interface{}{"Message": message, "Number": 2364821976}
	return renderer.HTML(c, http.StatusOK, "example.html", data)
}

// ParrotHandler ...
func ParrotHandler(c echo.Context) error {
	message := c.Param("message")

	res := &Response{
		Message: message,
	}

	return c.JSON(http.StatusOK, res)
}

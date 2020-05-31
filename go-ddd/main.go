package main

import (
	"go-ddd/cmd"
	"go-ddd/domain"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	r domain.Repository = domain.NewRepository()
	s domain.Service    = domain.NewService(r)
)

func init() {

}

func main() {
	cmd.Slices()
}

func test() {
	log.Println("Running a test")

	wg := sync.WaitGroup{}

	for i := 0; i < 10000; i++ {
		go func() {
			wg.Add(1)
			s.Create(&domain.Entity{Name: "Test " + strconv.Itoa(i)})
			delay := time.Duration(rand.Intn(10))
			log.Printf("waiting for %d", delay)
			time.Sleep(delay * time.Second)
			wg.Done()
		}()
	}

	wg.Wait()

	//time.Sleep(3 * time.Second)

	for _, v := range s.RetrieveAll() {
		log.Println(v)
	}

	log.Printf("%d item(s) found in the repository", r.Size())
}

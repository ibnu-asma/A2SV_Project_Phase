package main

import (
	"fmt"
	"library_management/controllers"
	"library_management/models"
	"library_management/services"
	"sync"
	"time"
)

func main() {
	library := services.NewLibrary()
	controller := controllers.NewLibraryController(library)

	book := models.Book{
		ID: 101, Title: "Concurrency in Go",
		Status: models.StatusAvailable,
	}
	library.AddBook(book)
	fmt.Println("Adding test members...")
	for i := 1; i <= 10; i++ {
		member := models.Member{
			ID:   i,
			Name: fmt.Sprintf("Member %d", i),
		}
		library.AddMember(member)
	}

	
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(memberID int) {
			defer wg.Done()

			err := library.ReserveBook(101, memberID)
			if err != nil {
				fmt.Printf("Member %d: %v\n", memberID, err)
				return
			}

			fmt.Printf("Member %d RESERVED book 101!\n", memberID)
			time.Sleep(2 * time.Second)

			borrowErr := library.BorrowBook(101, memberID)
			if borrowErr != nil {
				fmt.Printf("Member %d failed to borrow: %v\n", memberID, borrowErr)
			} else {
				fmt.Printf("Member %d BORROWED book 101!\n", memberID)
			}
		}(i)
	}
	wg.Wait()
	controller.Start()

}

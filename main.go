package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"
	"github.com/jamescyeh/gobuffalo/models"
)

func main() {
	eager()
	// noteager()

}

func eager() {
	c, err := pop.Connect("development")
	if err != nil {
		fmt.Println("could not connect to db")
		return
	}
	// [...]
	u1, _ := uuid.NewV4()
	u2, _ := uuid.NewV4()

	fmt.Println(u1)
	fmt.Println(u2)

	job_one := &models.Job{
		ID: u2,
	}
	group := &models.Group{
		ID:   u1,
		Jobs: []models.Job{*job_one},
	}

	err = c.Eager().Create(group)
	spew.Dump(err)

	outputGroup := models.Group{}
	err = c.Eager().Find(&outputGroup, u1)
	spew.Dump(err)
	spew.Dump(outputGroup)
}

func noteager() {
	c, err := pop.Connect("development")
	if err != nil {
		fmt.Println("could not connect to db")
		return
	}

	u1, _ := uuid.NewV4()
	u2, _ := uuid.NewV4()

	fmt.Println(u1)
	fmt.Println(u2)

	group := &models.Group{
		ID: u1,
	}

	job_one := &models.Job{
		ID:      u2,
		GroupID: u1,
	}

	err = c.Create(group)
	spew.Dump(err)
	err = c.Create(job_one)
	spew.Dump(err)

	outputGroup := models.Group{}
	err = c.Eager().Find(&outputGroup, u1)
	spew.Dump(err)
	spew.Dump(outputGroup)
}

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/IvanRodriguez09/desafio-goweb-IvanRodriguez/cmd/server/handler"
	"github.com/IvanRodriguez09/desafio-goweb-IvanRodriguez/internal/domain"
	"github.com/IvanRodriguez09/desafio-goweb-IvanRodriguez/internal/tickets"
	"github.com/gin-gonic/gin"
)

func main() {

	// Cargo csv.
	list, err := LoadTicketsFromFile("../../tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}

	repository := tickets.NewRepository(list)
	service := tickets.NewService(repository)
	handler := handler.NewService(service)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	// Rutas a desarollar:
	// GET - “/ticket/getByCountry/:dest”
	// GET - “/ticket/getAverage/:dest”
	r.GET("/ticket/getByCountry/:dest", handler.GetTicketsByCountry())
	r.GET("/ticket/getAverage/:dest", handler.AverageDestination())

	if err := r.Run(); err != nil {
		panic(err)
	}

}

func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}

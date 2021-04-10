package product

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Product struct {
	ID          string  `json: "id"`
	Name        string  `json: "name"`
	Description string  `json: "description"`
	Price       float32 `json: "price"`
}

func InitProduct(m *kafka.Message) *Product {
	var product *Product
	fmt.Println("Eaeae")
	if err := json.Unmarshal([]byte(m.Value), &product); err == nil {
		fmt.Println(product)
		return product
	}
	return &Product{}
}

func (p *Product) Buy() error {
	if p.ID == "" {
		return errors.New("Product not found")
	}
	fmt.Printf("Novo produto comprado com sucesso %v(%v) e custou %v", p.Name, p.ID, p.Price)
	return nil
}

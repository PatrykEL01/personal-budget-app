package services

import (
	"context"
	"log"
	"github.com/jackc/pgx/v4"
)


func LoadData(ctx context.Context, conn *pgx.Conn) error {
	query := `
		INSERT INTO personal_budget (name, amount) VALUES 
		('Marketing Budget', 5000.0),
		('Development Budget', 12000.0),
		('Operations Budget', 7500.0),
		('Employee Benefits', 3000.0),
		('Research and Development', 15000.0),
		('Sales Incentives', 6000.0),
		('Training Programs', 2500.0),
		('Customer Support', 4000.0),
		('Office Supplies', 1500.0),
		('IT Infrastructure', 8000.0),
		('Travel Expenses', 3500.0),
		('Event Sponsorship', 4500.0),
		('Marketing Campaign 1', 5500.0),
		('Marketing Campaign 2', 5200.0),
		('Miscellaneous Expenses', 1200.0);
	`
	_, err := conn.Exec(ctx, query)
	if err != nil {
		return err
	}
	log.Println("Test data seeded successfully!")
	return nil
}



func CleanTestData(ctx context.Context, conn *pgx.Conn) error {
    _, err := conn.Exec(ctx, "TRUNCATE TABLE personal_budget RESTART IDENTITY")
    if err != nil {
        return err
    }
    log.Println("Test data cleaned successfully!")
    return nil
}

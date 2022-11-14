package inventory

import (
	"context"
	"log"

	"database/sql"

	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/models"
	_ "github.com/lib/pq"
)

// Create an interface to prevent unwanted use of these methods

type InventoryRepo struct {
	db *sql.DB
}

// Receives the db instance as argument and sets it in the struct before returning the struct itself
func NewInventoryRepo(db *sql.DB) *InventoryRepo {
	return &InventoryRepo{
		db: db,
	}
}

func (m InventoryRepo) GetAllInventories(ctx context.Context, limit, offset int32) (items []*models.Skus, err error) {

	d, cerr := m.db.Conn(ctx)
	if cerr != nil {
		err = cerr
	}
	defer d.Close()

	rows, qErr := d.QueryContext(ctx, "SELECT * FROM skus")
	if qErr != nil {
		log.Fatalf("eror querying : %+v", qErr)
	}
	defer rows.Close()
	items = []*models.Skus{}
	for rows.Next() {
		i := models.Skus{}
		err = rows.Scan(&i.ID, &i.SkuID, &i.Name, &i.Price, &i.Quantity, &i.Type, &i.Description, &i.ImageURL, &i.CreatedAt, &i.UpdatedAt, &i.DeletedAt)
		if err != nil {
			log.Fatalf("errorr scanning : %+v\n", err)
		}
		items = append(items, &i)
	}
	return
}

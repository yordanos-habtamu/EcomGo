package product

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/yordanos-habtamu/EcomGo.git/types"
)

type Store struct {
	db *sql.DB
}

// CreateProduct implements types.ProductStore.
func (s *Store) CreateProduct(payload types.Product) ( error) {
	_,err := s.db.Exec("INSERT INTO products (name,description,price,stock_quantity,category,image_url,is_active) VALUES (?,?,?,?,?,?,?)",payload.Name,payload.Description,payload.Price,payload.Stock,payload.Catagory,payload.ImgUrl,payload.IsActive)
     if err != nil {
		return err
	 }
	 return nil
}

// DeleteProduct implements types.ProductStore.
func (s *Store) DeleteProduct(id uint) error {
    // Prepare the SQL query to delete the product by ID
    result, err := s.db.Exec("DELETE FROM products WHERE id = ?", id)
    if err != nil {
        return fmt.Errorf("failed to delete product: %w", err)
    }

    // Check how many rows were affected
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to retrieve affected rows: %w", err)
    }

    // If no rows were affected, the product with the given ID was not found
    if rowsAffected == 0 {
        return fmt.Errorf("product with ID %d not found", id)
    }

    return nil
}

// GetAllProducts implements types.ProductStore.
func (s *Store) GetAllProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
    if err != nil {
        return nil, err
    }
    defer rows.Close() // Ensure rows are closed

    // Create a slice to hold the products
    products := []types.Product{}

    // Iterate through the rows
    for rows.Next() {
        product, err := scanRowsIntoProducts(rows)
        if err != nil {
            return nil, err
        }
        products = append(products, *product)
    }

    // Check for errors during iteration
    if err = rows.Err(); err != nil {
        return nil, err
    }

    return products, nil
}




// GetProductsByCategory implements types.ProductStore.
func (s *Store) GetProductsByCategory(category string) ([]types.Product, error) {
    // Execute the query
    rows, err := s.db.Query("SELECT * FROM products WHERE category = ?", category)
    if err != nil {
        return nil, err
    }
    defer rows.Close() // Ensure rows are closed

    // Create a slice to hold the products
    products := []types.Product{}

    // Iterate through the rows
    for rows.Next() {
        product, err := scanRowsIntoProducts(rows)
        if err != nil {
            return nil, err
        }
        products = append(products, *product)
    }

    // Check for errors during iteration
    if err = rows.Err(); err != nil {
        return nil, err
    }

    // If no products are found, return an error
    if len(products) == 0 {
        return nil, fmt.Errorf("no products found for category: %s", category)
    }

    return products, nil
}


// UpdateProduct implements types.ProductStore.
func (s *Store) UpdateProduct(id uint, payload types.RegisterProductPayload) (*types.Product, error) {
    // Prepare the SQL query to update the product by ID
    query := `
        UPDATE products 
        SET name = ?, category = ?, price = ?, stock = ?
        WHERE id = ?
    `

    // Execute the query
    result, err := s.db.Exec(query, payload.Name, payload.Catagory, payload.Price, payload.Stock, id)
    if err != nil {
        return &types.Product{}, fmt.Errorf("failed to update product: %w", err)
    }

    // Check how many rows were affected
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return &types.Product{}, fmt.Errorf("failed to retrieve affected rows: %w", err)
    }

    // If no rows were affected, the product with the given ID was not found
    if rowsAffected == 0 {
        return &types.Product{}, fmt.Errorf("product with ID %d not found", id)
    }

    // Fetch the updated product to return it
    updatedProduct, err := s.GetProductById(id)
    if err != nil {
        return &types.Product{}, fmt.Errorf("failed to retrieve updated product: %w", err)
    }

    return updatedProduct, nil
}


func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProductByName(name string) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	u := new(types.Product)
	for rows.Next() {
		u, err = scanRowsIntoProducts(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func scanRowsIntoProducts(rows *sql.Rows) (*types.Product, error) {
	product:= new(types.Product)
	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.Catagory,
		&product.ImgUrl,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.IsActive,
	)
	if err != nil {
		log.Fatal(err)
	}
	return product, nil
}

func (s *Store) GetProductById(id uint) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	u := new(types.Product)
	for rows.Next() {
		u, err = scanRowsIntoProducts(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("product not found")
	}
	return u, nil
}


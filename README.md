# EcomGo

EcomGo is an eCommerce API built using Golang. It provides core functionalities for managing users and products, supporting essential eCommerce workflows.

## Features

- 🧑‍💻 User Registration
- ➕ Add, ❌ Remove, ✏️ Edit Products
- 🛒 Checkout Process

## Technologies Used

- **Language:** 🐹 Golang
- **Frameworks/Tools:** 📦 Mux, 🔐 JWT
- **Database:** 🛢️ MySQL
- **Migrations:** 🗂️ Migrate

## Installation

### Prerequisites

- ✅ Go installed on your system
- ✅ MySQL database set up

### Steps

1. Clone the repository:

   ```bash
   git clone <repository_url>
   cd EcomGo
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Configure the database connection in your project files (ensure the database name is `EcomGo`).

4. Apply migrations:

   ```bash
   migrate -path ./migrations -database "mysql://user:password@tcp(localhost:3306)/EcomGo" up
   ```

5. Run the project:

   ```bash
   go run main.go
   ```

## Usage

Use tools like Postman or cURL to interact with the API. Below are some example endpoints:

- **Register User:** `/api/register`
- **Add Product:** `/api/product/add`
- **Edit Product:** `/api/product/edit`
- **Remove Product:** `/api/product/remove`
- **Checkout:** `/api/checkout`

## License

📜 This project is licensed under [LICENSE_NAME].

## Contributors

- 👩‍💻 [Yordanos Habtamu](https://github.com/yordanos-habtamu)

## Acknowledgments

💡 Thanks to the open-source community for providing tools and inspiration for this project.


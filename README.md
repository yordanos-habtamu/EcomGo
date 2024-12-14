# EcomGo

Welcome to **EcomGo**! 🚀 This is an eCommerce API built with the power of **Golang** to simplify and enhance the way eCommerce applications are built and managed. Whether you're a developer looking to explore backend APIs or building your next online store, EcomGo has you covered. 🎉

## ✨ Features

- 🧑‍💻 **User Registration:** Seamless user account creation.
- ➕ **Add Products:** Easily expand your product catalog.
- ❌ **Remove Products:** Remove outdated or unwanted items.
- ✏️ **Edit Products:** Keep your product details up-to-date.
- 🛒 **Checkout Process:** Smooth checkout workflows for users.

## 🔧 Technologies Used

- 🐹 **Golang:** The heart of the application.
- 📦 **Mux:** For efficient and fast routing.
- 🔐 **JWT:** Secure user authentication.
- 🛢️ **MySQL:** Reliable database for all your data needs.
- 🗂️ **Migrate:** Simple and effective database migrations.

## 🛠️ Installation

### Prerequisites

- ✅ Ensure **Go** is installed on your system.
- ✅ Set up a **MySQL** database instance.

### Steps to Get Started

1. Clone the repository:

   ```bash
   git clone <repository_url>
   cd EcomGo
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Configure your database connection in the project configuration files (make sure to name your database `EcomGo`).

4. Apply migrations:

   ```bash
   migrate -path ./migrations -database "mysql://user:password@tcp(localhost:3306)/EcomGo" up
   ```

5. Run the project:

   ```bash
   go run main.go
   ```

## 🚀 Usage

Interact with the API using tools like **Postman** or **cURL**. Here are some example endpoints to get you started:

- 📝 **Register User:** `/api/register`
- ➕ **Add Product:** `/api/product/add`
- ✏️ **Edit Product:** `/api/product/edit`
- ❌ **Remove Product:** `/api/product/remove`
- 🛒 **Checkout:** `/api/checkout`

## 📜 License

This project is licensed under [LICENSE_NAME].

## 👥 Contributors

- 👩‍💻 [Yordanos Habtamu](https://github.com/yordanos-habtamu)  

## 💡 Acknowledgments

A huge thanks to the open-source community for providing tools, inspiration, and guidance. Together, we build better software! 🌟


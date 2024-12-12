package routes

import (
    "net/http"

    "github.com/CVWO/sample-go-app/internal/handlers/posts"
    "github.com/go-chi/chi/v5"
)

// PostsRoutes sets up the routes for post-related operations.
func PostsRoutes() http.Handler {
    r := chi.NewRouter()

    // Define post-related routes
    r.Get("/", posts.HandleList)          // List all posts
    r.Get("/{id}", posts.HandleGet)       // Get a specific post by ID
    r.Post("/", posts.HandleCreate)       // Create a new post
    r.Put("/{id}", posts.HandleUpdate)    // Update a specific post by ID
    r.Delete("/{id}", posts.HandleDelete) // Delete a specific post by ID

    return r
}
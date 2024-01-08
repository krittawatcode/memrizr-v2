package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/memrizr-v2/account/model"
)

// Handler struct holds required services for handler to function
type Handler struct {
	UserService  model.UserService
	TokenService model.TokenService
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R            *gin.Engine
	UserService  model.UserService
	TokenService model.TokenService
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{
		UserService:  c.UserService,
		TokenService: c.TokenService,
	} // currently has no properties

	// Create an account group
	g := c.R.Group(os.Getenv("ACCOUNT_API_URL"))

	g.GET("/me", h.Me)
	g.POST("/signUp", h.SignUp)
	g.POST("/signIn", h.SignIn)
	g.POST("/signOut", h.SignOut)
	g.POST("/tokens", h.Tokens)
	g.POST("/image", h.Image)
	g.DELETE("/image", h.DeleteImage)
	g.PUT("/details", h.Details)
}

// SignIn handler
func (h *Handler) SignIn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's sign in",
	})
}

// SignOut handler
func (h *Handler) SignOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's sign out",
	})
}

// Tokens handler
func (h *Handler) Tokens(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's tokens",
	})
}

// Image handler
func (h *Handler) Image(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's image",
	})
}

// DeleteImage handler
func (h *Handler) DeleteImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's deleteImage",
	})
}

// Details handler
func (h *Handler) Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's details",
	})
}

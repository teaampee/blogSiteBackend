package main

import (
	"database/sql"
	"log"
	"os"

	"example.com/blog/v2/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port in not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL in not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("can't connect to the db", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	r := gin.Default()

	//api endpoints

	// user
	// create user {json:: name=?,email=?,password=?}
	r.POST("/users", apiCfg.handlerCreateUser)
	//get user
	r.GET("users/:userID", apiCfg.handlerGetUser)
	//update user {json:: name=?,email=?,password=?} update the fields as needed can be left null
	r.PATCH("/users", apiCfg.UserAuth, apiCfg.handlerUpdateUser)
	//login {json:: email=?,password=?}
	r.POST("/login", apiCfg.handlerLogin)
	//validate user
	r.GET("/validate", apiCfg.UserAuth, apiCfg.handlerValidate)

	// get user blog
	r.GET("users/:userID/blog", apiCfg.handlerGetUserBlog)

	// blogs

	//post blog {json:: title=?, description=?}
	r.POST("/blogs", apiCfg.UserAuth, apiCfg.handlerCreateBlog)
	//Get blogs "/blogs?limit=?&offset=?"
	r.GET("/blogs", apiCfg.handlerGetBlogs)
	//get active blogs "/blogs/active?limit=?&offset=?"
	r.GET("/blogs/active", apiCfg.handlerGetActiveBlogs)
	//get blog
	r.GET("blogs/:blogID", apiCfg.handlerGetBlog)
	//update blog {jsonLL title=?, description=?}
	r.PATCH("/blogs/:blogID", apiCfg.UserAuth, apiCfg.handlerUpdateUserBlog)
	//delete blog
	r.DELETE("/blogs/:blogID", apiCfg.UserAuth, apiCfg.handlerDeleteUserBlog)

	// posts

	//create post "/posts?blog_id=?"
	r.POST("/posts", apiCfg.UserAuth, apiCfg.handlerCreatePost)
	//get blogs post "/posts?blog_id=?"
	r.GET("/posts", apiCfg.handlerGetBlogPosts)
	//get post
	r.GET("/posts/:postID", apiCfg.handlerGetPost)
	//update user post
	r.PATCH("/posts/:postID", apiCfg.UserAuth, apiCfg.handlerUpdateUserPost)
	//delete user post
	r.DELETE("/posts/:postID", apiCfg.UserAuth, apiCfg.handlerDeleteUserPost)

	//get post comments
	r.GET("/posts/:postID/comments", apiCfg.handlerGetPostComments)

	//comments

	//create comment "/comments?post_id=?"
	r.POST("/comments", apiCfg.UserAuth, apiCfg.handlerCreateComment)
	//get comment
	r.GET("/comments/:commentID", apiCfg.handlerGetComment)
	//update comment {json:: content=?}
	r.PATCH("/comments/:commentID", apiCfg.UserAuth, apiCfg.handlerUpdateComment)
	//delete comment
	r.DELETE("/comments/:commentID", apiCfg.UserAuth, apiCfg.handlerDeleteUserComment)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")}
	r.Run()

}

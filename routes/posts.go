package routes

import (
	"net/http"
	"strconv"

	"example.com/blog-rest-api/models"
	"github.com/gin-gonic/gin"
)

func getPosts(context *gin.Context) {
	posts, err := models.GetAllPosts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch posts. Try again later."})
		return
	}
	context.JSON(http.StatusOK, posts)
}

func getPost(context *gin.Context) {
	postId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse post id."})
		return
	}

	post, err := models.GetPostByID(postId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch post."})
		return
	}

	context.JSON(http.StatusOK, post)
}

func createPost(context *gin.Context) {
	var post models.Post
	err := context.ShouldBindJSON(&post)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	userId := context.GetInt64("userId")
	post.UserID = userId

	err = post.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create post. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Post created!", "post": post})
}

func updatePost(context *gin.Context) {
	postId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse post id."})
		return
	}

	userId := context.GetInt64("userId")
	post, err := models.GetPostByID(postId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch post."})
		return
	}

	if post.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update post."})
		return
	}

	var updatedPost models.Post
	err = context.ShouldBindJSON(&updatedPost)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	updatedPost.ID = postId
	err = updatedPost.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update post."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Post updated successfully!"})
}

func deletePost(context *gin.Context) {
	postId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse post id."})
		return
	}

	userId := context.GetInt64("userId")
	post, err := models.GetPostByID(postId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch post."})
		return
	}

	if post.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete post."})
		return
	}

	err = post.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete post."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully!"})
}

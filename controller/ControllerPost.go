package controller

import (
	"WHisperHArbor-backend/model"
	"WHisperHArbor-backend/service"
	"WHisperHArbor-backend/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func UserPost(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	claim, _ := utils.ParseToken(auth)
	PostContext := model.Post{}
	if err := c.ShouldBindJSON(&PostContext); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "发布失败!" + err.Error(),
		})
		return
	} else {
		if user, err := service.GetUser(*claim); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    400,
				"message": "发布失败!" + err.Error(),
			})
			return
		} else {
			if PostContext.Encrypted {
				content, err := utils.EncryptPost(PostContext.Content, user.AES)
				if err != nil {
					c.JSON(http.StatusOK, gin.H{
						"code":    400,
						"message": "发布失败!" + err.Error(),
					})
					return
				}
				PostContext.Content = content
			}
			PostContext.UserId = user.ID
			PostContext.Time = time.Unix(time.Now().Unix(), 0)
			if err := HandlePost(&PostContext); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code":    400,
					"message": "发布失败!" + err.Error(),
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code":    200,
					"message": "发布成功!",
				})
			}
		}
	}
}

func HandlePost(post *model.Post) error {
	if err := model.DB.Create(post).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func PublicGetPost(c *gin.Context) {
	limist := &model.Pagination{}
	if c.ShouldBind(limist) != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "query参数错误",
		})
		return
	} else {
		if limist.Limit < 0 || limist.Limit > 100 {
			limist.Limit = 10
		}
		if limist.Offset < 1 {
			limist.Offset = 1
		}
	}
	if posts, err := service.PublicPost(*limist); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "获取列表失败",
		})
		return
	} else {
		json := make([]struct {
			Content  string    `json:"content"`
			Nickname string    `json:"nickname"`
			Likes    int64     `json:"likes"`
			Time     time.Time `json:"time"`
		}, len(posts))
		for i := range posts {
			log.Println(posts[i].User)
			json[i].Nickname = posts[i].User.Nickname
			json[i].Time = posts[i].Time
			json[i].Likes = posts[i].Likes
			json[i].Content = posts[i].Content
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": json,
		})

	}
}

func UserGetPost(c *gin.Context) {
	limist := &model.Pagination{}
	if c.ShouldBind(limist) != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "query参数错误",
		})
		return
	} else {
		if limist.Limit < 0 || limist.Limit > 100 {
			limist.Limit = 10
		}
		if limist.Offset < 1 {
			limist.Offset = 1
		}
	}
	auth := c.Request.Header.Get("Authorization")
	claim, _ := utils.ParseToken(auth)
	if user, err := service.GetUser(*claim); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "获取列表失败",
		})
		return
	} else {
		if post, err := service.UserPost(user, *limist); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    400,
				"message": "获取列表失败",
			})
			return
		} else {
			temp_post := post
			for i, _ := range temp_post {
				if post[i].Encrypted == true {
					if content, err := utils.DecryptPost(post[i].Content, user.AES); err != nil {
						c.JSON(http.StatusOK, gin.H{
							"code":    400,
							"message": "获取列表失败",
						})
						return
					} else {
						temp_post[i].Content = content
					}
				}
				temp_post[i].User = model.User{}
			}
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": temp_post,
			})
		}
	}

}

func UserLikePost(c *gin.Context) {
	Posts := &model.AddLikes{}
	if err := c.ShouldBindJSON(Posts); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "点赞失败" + string(err.Error()),
		})
	} else {
		if err := service.IncrLike(*Posts); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    400,
				"message": "点赞失败" + string(err.Error()),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "点赞成功",
			})
		}
	}
}

func UserFavoritePost(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	claim, _ := utils.ParseToken(auth)
	Posts := &model.AddLikes{}
	if err := c.ShouldBindJSON(Posts); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "收藏失败1" + string(err.Error()),
		})
		return
	} else {
		if user, err := service.GetUser(*claim); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    400,
				"message": "收藏失败2" + string(err.Error()),
			})
			return
		} else {
			if err := service.FavoritePost(user, *Posts); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code":    400,
					"message": "收藏失败3" + string(err.Error()),
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code":    200,
					"message": "收藏成功",
				})
			}

		}

	}
}

func GetUserFavorites(c *gin.Context) {
	limist := &model.Pagination{}
	if c.ShouldBind(limist) != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "query参数错误",
		})
		return
	} else {
		if limist.Limit < 0 || limist.Limit > 100 {
			limist.Limit = 10
		}
		if limist.Offset < 1 {
			limist.Offset = 1
		}
	}
	auth := c.Request.Header.Get("Authorization")
	claim, _ := utils.ParseToken(auth)
	if user, err := service.GetUser(*claim); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "获取失败" + string(err.Error()),
		})
		return
	} else {
		if postList, err := service.GetFavoritePost(user, *limist); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    400,
				"message": "获取失败" + string(err.Error()),
			})
			return
		} else {
			postAns := make([]model.Post, len(postList))
			for i := range postList {
				postAns[i], _ = service.GetPost(postList[i].PostID)
			}
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": postAns,
			})
		}
	}
}

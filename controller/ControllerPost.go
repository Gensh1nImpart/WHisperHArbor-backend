package controller

import (
	"WHisperHArbor-backend/model"
	"WHisperHArbor-backend/service"
	"WHisperHArbor-backend/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

/*
@ 发表帖子
@ method: Post
@ Headers: Authorization
@ Parmas: Json{}: Post
*/
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

/*
@ 获取所有用户公开的帖子
@ method: GET
@ Headers: Authorization
@ Parmas: limist
*/
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
			json[i].Nickname = posts[i].User.Nickname
			json[i].Time = posts[i].Time
			json[i].Likes = posts[i].Likes
			json[i].Content = posts[i].Content
		}
		//sort.Slice(json, func(i, j int) bool {
		//	return json[i].Time.After(json[j].Time)
		//})
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": json,
		})

	}
}

/*
@ 获取用户的帖子
@ method: GET
@ Headers: Authorization
@ Parmas: limist
*/
func UserGetPost(c *gin.Context) {
	limist := &model.Pagination{}
	limist.Limit, _ = strconv.Atoi(c.DefaultQuery("limit", "10"))
	limist.Offset, _ = strconv.Atoi(c.DefaultQuery("offset", "0"))
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
			json := make([]struct {
				Content  string    `json:"content"`
				Nickname string    `json:"nickname"`
				Likes    int64     `json:"likes"`
				Time     time.Time `json:"time"`
			}, len(post))
			for i, _ := range post {
				if post[i].Encrypted == true {
					if content, err := utils.DecryptPost(post[i].Content, user.AES); err != nil {
						c.JSON(http.StatusOK, gin.H{
							"code":    400,
							"message": "获取列表失败",
						})
						return
					} else {
						json[i].Content = content
						json[i].Nickname = post[i].User.Nickname
						json[i].Likes = post[i].Likes
						json[i].Time = post[i].Time
					}
				}
				json[i].Content = post[i].Content
				json[i].Nickname = post[i].User.Nickname
				json[i].Likes = post[i].Likes
				json[i].Time = post[i].Time
			}
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": json,
			})
		}
	}

}

/*
@ 点赞帖子
@ method: POST
@ Headers: Authorization
@ Parmas: Json{}:PostId
*/
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

/*
@ 收藏帖子
@ method: Post
@ Headers: Authorization
@ Parmas: Json{}: PostID
*/
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

/*
@ 获取用户喜爱的帖子
@ method: GET
@ Headers: Authorization
@ Parmas: limist
*/
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

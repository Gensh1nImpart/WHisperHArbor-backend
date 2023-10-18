package service

import (
	"WHisperHArbor-backend/model"
	"errors"
)

//func checkLogin(user model.LoginUser) (bool, error) {
//
//}

func GetUser(claim model.MyClaims) (model.User, error) {
	user := model.User{}
	if err := model.DB.Where("account = ?", claim.User.Account).First(&user).Error; err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func UserPost(claim model.User, limit model.Pagination) ([]model.Post, error) {
	var posts []model.Post
	if err := model.DB.Where("user_id = ?", claim.ID).Limit(limit.Limit).Offset(limit.Offset).Find(&posts).Error; err != nil {
		return posts, err
	} else {
		return posts, nil
	}
}

func PublicPost(limit model.Pagination) ([]model.Post, error) {
	var posts []model.Post
	if err := model.DB.Preload("User").Where("encrypted = ?", false).Limit(limit.Limit).Offset(limit.Offset).Find(&posts).Error; err != nil {
		return posts, err
	} else {
		return posts, nil
	}
}

func IncrLike(likepost model.AddLikes) error {
	post := &model.Post{}
	if err := model.DB.Where("id = ?", likepost.PostId).Find(&post).Error; err != nil {
		return err
	}
	if err := model.DB.Model(&post).Update("likes", post.Likes+1).Error; err != nil {
		return err
	}
	return nil
}

func FavoritePost(claim model.User, post model.AddLikes) error {
	var existingFavorite []model.Favorites
	if err := model.DB.Where("user_id = ? AND post_id = ?", claim.ID, post.PostId).Find(&existingFavorite).Error; err != nil {
		return err
	} else {
		if len(existingFavorite) > 0 {
			return errors.New("该帖子已经被收藏过")
		}
		newFavotire := model.Favorites{
			PostID: post.PostId,
			UserID: claim.ID,
		}
		err := model.DB.Create(&newFavotire).Error
		return err
	}
}

func GetFavoritePost(user model.User, limit model.Pagination) ([]model.Favorites, error) {
	var posts []model.Favorites
	err := model.DB.Where("user_id = ?", user.ID).Limit(limit.Limit).Offset(limit.Offset).Find(&posts).Error
	return posts, err
}

func GetPost(id uint) (model.Post, error) {
	var post model.Post
	err := model.DB.Where("id = ?", id).Find(&post)
	return post, err.Error
}

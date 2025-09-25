package gormtest

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Posts     []Post
	PostCount int `gorm:"default:0"` // 用户的文章数量统计
}

type Post struct {
	gorm.Model
	Title         string `gorm:"type:varchar(200);not null"`
	Content       string `gorm:"type:text;not null"`
	UserID        uint   `gorm:"not null"`
	User          User
	Comments      []Comment
	CommentCount  int    `gorm:"default:0"` // 文章的评论数量统计
	CommentStatus string `gorm:"type:varchar(20);default:'无评论'"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	UserID  uint   `gorm:"not null"`
	PostID  uint   `gorm:"not null"`
	User    User
	Post    Post
}

func RunCreate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	users := []User{
		{Name: "Alice", Email: "alice@example.com"},
		{Name: "Bob", Email: "bob@example.com"},
		{Name: "Charlie", Email: "charlie@example.com"},
	}
	db.Create(&users)
	posts := []Post{
		{Title: "Post 1", Content: "This is the first post.", UserID: 1},
		{Title: "Post 2", Content: "This is the second post.", UserID: 2},
		{Title: "Post 3", Content: "This is the third post.", UserID: 3},
	}
	db.Create(&posts)
	comments := []Comment{
		{Content: "Comment 1", UserID: 1, PostID: 1},
		{Content: "Comment 2", UserID: 2, PostID: 2},
		{Content: "Comment 3", UserID: 3, PostID: 3},
	}
	db.Create(&comments)
	fmt.Println("数据创建成功")
}

func RunQueryByUserName(db *gorm.DB, name string) {
	var posts []Post
	var user User
	result := db.Debug().Where("name = ?", name).Find(&user)
	if result.Error != nil {
		fmt.Println("查询失败:", result.Error)
		return
	}
	err := db.Debug().
		Model(&Post{}).
		Preload("User").
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Order("comments.created_at DESC")
		}).
		Preload("Comments.User").
		Where("user_id = ?", user.ID).
		Order("posts.created_at DESC").
		Find(&posts).Error

	if err != nil {
		fmt.Printf("查询文章失败: %v", err)
	}

	fmt.Printf("查询结果为: %v\n", result.RowsAffected)
	for _, post := range posts {
		fmt.Printf("文章ID:%d,文章标题:%s,文章内容:%s,作者ID:%d,作者名称:%s\n", post.ID, post.Title, post.Content, post.UserID, post.User.Name)
		for _, comment := range post.Comments {
			fmt.Printf(" - 评论: %s (用户: %s)\n", comment.Content, comment.User.Name)
		}
	}
}

func GetMostCommentedPost(db *gorm.DB) (Post, error) {
	var post Post

	err := db.Debug().
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).
		Preload("Comments").
		First(&post).Error

	return post, err
}

// Comment模型的BeforeCreate钩子：创建评论时更新文章评论数量
func (c *Comment) BeforeCreate(db *gorm.DB) error {
	// 更新文章的评论数量
	err := db.Model(&Post{}).Where("id = ?", c.PostID).
		Update("comment_count", gorm.Expr("comment_count + 1")).Error
	if err != nil {
		return err
	}

	// 更新文章的评论状态
	err = db.Model(&Post{}).Where("id = ?", c.PostID).
		Update("comment_status", "有评论").Error
	return err
}

// Comment模型的AfterDelete钩子：删除评论时检查文章评论数量
func (c *Comment) AfterDelete(db *gorm.DB) error {
	// 获取当前文章的评论数量
	var commentCount int64
	err := db.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error
	if err != nil {
		return err
	}

	// 如果评论数量为0，更新评论状态
	if commentCount == 0 {
		err = db.Model(&Post{}).Where("id = ?", c.PostID).
			Update("comment_status", "无评论").Error
	} else {
		// 更新评论数量
		err = db.Model(&Post{}).Where("id = ?", c.PostID).
			Update("comment_count", commentCount).Error
	}
	return err
}

// Post模型的BeforeCreate钩子：创建文章时更新用户的文章数量
func (p *Post) BeforeCreate(db *gorm.DB) error {
	// 更新用户的文章数量
	err := db.Model(&User{}).Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + 1")).Error
	if err != nil {
		return err
	}
	return nil
}

// Post模型的AfterDelete钩子：删除文章时更新用户的文章数量
func (p *Post) AfterDelete(db *gorm.DB) error {
	// 更新用户的文章数量（减少）
	err := db.Model(&User{}).Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count - 1")).Error
	return err
}

func DataCreate(db *gorm.DB) {
	// 创建测试数据
	user := User{
		Name:  "testuser",
		Email: "test@example.com",
	}

	db.Create(&user)

	post := Post{
		Title:   "我的第一篇文章",
		Content: "这是文章内容...",
		UserID:  user.ID,
	}

	db.Create(&post)

	comment := Comment{
		Content: "这是一条评论",
		UserID:  user.ID,
		PostID:  post.ID,
	}

	db.Create(&comment)

	mostCommentedPost, _ := GetMostCommentedPost(db)
	fmt.Printf("评论最多的文章: %s\n", mostCommentedPost.Title)
}

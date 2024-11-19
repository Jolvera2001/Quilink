package blogComponent

import (
	"fmt"
	"log"
	m "quilink/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogService struct {
	db *gorm.DB
}

func NewBlogService(db *gorm.DB) *BlogService {
	return &BlogService{
		db: db,
	}
}

func (s *BlogService) GetBlog(id uuid.UUID) (m.Blog, error) {
	var blog m.Blog

	if err := s.db.First(&blog, "id = ?", id).Error; err != nil {
		log.Printf("[BlogService.GetBlog][blogId=%s] error finding blog with id %s: %v", id, id, err)
		if err == gorm.ErrRecordNotFound {
			return m.Blog{}, fmt.Errorf("blog not found with id %s", id)
		}
		return m.Blog{}, fmt.Errorf("error fetching blog from database: %w", err)
	}

	return blog, nil
}

func (s *BlogService) GetTotalCount(profileId uuid.UUID) (int64, error) {
	var count int64

	if err := s.db.Model(&m.Blog{}).
		Where("user_id = ?", profileId).
		Count(&count).Error; err != nil {
		log.Printf("[BlogService.GetTotalCount][userId=%s] error counting blogs for user %s: %v", profileId, profileId, err)
		return 0, fmt.Errorf("failed to count blogs: %w", err)
	}

	return count, nil
}

func (s *BlogService) GetPostBySlug(slug string) (m.Blog, error) {
	var blog m.Blog

	if err := s.db.First(blog, "slug = ?", slug).Error; err != nil {
		log.Printf("[BlogService.GetPostBySlug][slug=%s] error getting blog with slug %s: %v", slug, slug, err)

		if err == gorm.ErrRecordNotFound {
			return m.Blog{}, fmt.Errorf("blog not found with slug: %w", err)
		}

		return m.Blog{}, fmt.Errorf("failed to find blog by slug: %w", err)
	}

	return blog, nil
}

func (s *BlogService) GetBlogs(profileId uuid.UUID, page, size int) ([]m.Blog, error) {
	var blogs []m.Blog
	offset := (page - 1) * size

	result := s.db.
		Where("user_id = ?", profileId).
		Order("created_at DESC").
		Offset(offset).
		Limit(size).
		Find(&blogs)

	if result.Error != nil {
		log.Printf("[BlogService.GetBlogs][userId=%s, page=%d, size=%d] error getting blogs with id %s: %v", profileId, page, size, profileId, result.Error)
		return nil, fmt.Errorf("error getting blogs with id %s: %v", profileId, result.Error)
	}

	return blogs, nil
}

func (s *BlogService) GetPublishedBlogs(profileId uuid.UUID, page, size int) ([]m.Blog, error) {
	var blogs []m.Blog
	offset := (page - 1) * size

	result := s.db.
		Where("user_id = ?, published = true", profileId).
		Order("created_at DESC").
		Offset(offset).
		Limit(size).
		Find(&blogs)

	if result.Error != nil {
		log.Printf("[BlogService.GetPublishedBlogs][userId=%s, page=%d, size=%d] error getting published blogs with id %s: %v", profileId, page, size, profileId, result.Error)
		return nil, fmt.Errorf("error getting blogs with id %s: %v", profileId, result.Error)
	}

	return blogs, nil
}

func (s *BlogService) CreateBlog(dto m.BlogDto) (m.Blog, error) {
	blog := m.Blog{
		Title:     dto.Title,
		Content:   dto.Content,
		Slug:      dto.Slug,
		Published: dto.Published,
		ProfileId: dto.ProfileId,
	}

	result := s.db.Create(&blog)

	if result.Error != nil {
		log.Printf("[BlogService.CreateBlog] error creating blog for user %s: %v", dto.ProfileId, result.Error)
		return m.Blog{}, fmt.Errorf("failed to create blog %w", result.Error)
	}

	return blog, nil
}

func (s *BlogService) UpdateBlog(id uuid.UUID, dto m.BlogDto) (m.Blog, error) {
	var blog m.Blog

	if err := s.db.First(&blog, "id = ?", id.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("[BlogService.UpdateBlog] blog not found: id=%s", id)
			return m.Blog{}, fmt.Errorf("blog not found: %s", id)
		}
		log.Printf("[BlogService.UpdateBlog] error fetching blog: id=%s error=%v", id, err)
		return m.Blog{}, fmt.Errorf("failed to fetch blog: %w", err)
	}

	blog.Title = dto.Title
	blog.Content = dto.Content
	blog.Slug = dto.Slug
	blog.Published = dto.Published

	if err := s.db.Save(&blog).Error; err != nil {
		log.Printf("[BlogService.UpdateBlog] error updating blog: id=%s error=%v", id, err)
		return m.Blog{}, fmt.Errorf("failed to update blog: %w", err)
	}

	return blog, nil
}

func (s *BlogService) TogglePublishStatus(id uuid.UUID) (bool, error) {
	var blog m.Blog

	if err := s.db.First(&blog, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("[BlogService.TogglePublishedStatus] blog not found: id=%s", id)
			return false, fmt.Errorf("blog not found: %s", id)
		}
		log.Printf("[BlogService.TogglePublishedStatus] error fetching blog: id=%s error=%v", id, err)
		return false, fmt.Errorf("failed to fetch blog: %w", err)
	}

	blog.Published = !blog.Published

	if err := s.db.Save(&blog).Error; err != nil {
		log.Printf("[BlogService.TogglePublishedStatus] error toggling published status: id=%s error=%v", id, err)
		return false, fmt.Errorf("failed to toggle published status: %w", err)
	}

	return true, nil
}

func (s *BlogService) DeleteBlog(id uuid.UUID) error {
	if err := s.db.Delete(&m.Blog{}, id).Error; err != nil {
		log.Printf("[BlogService.DeleteBlog] error deleting blog: id=%s error=%v", id, err)
		return fmt.Errorf("failed to delete blog: %w", err)
	}

	return nil
}

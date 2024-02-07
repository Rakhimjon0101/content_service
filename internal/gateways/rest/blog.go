package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"projects/content_service/internal/domains"
	"projects/content_service/internal/errs"
	"projects/content_service/internal/gateways/rest/views"
)

type reqCreateBlog struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// blog godoc
// @Router /blog [POST]
// @Summary blog
// @Description create blog
// @Tags blog
// @Accept json
// @Produce json
// @Param request body reqCreateBlog true "Request to Create a Blog"
// @Success 200 {object} R{data=views.Blog}
// @Failure 422 {object} R
// @Failure 500 {object} R
func (s *Server) create(c *gin.Context) {
	var req reqCreateBlog
	if err := c.ShouldBindJSON(&req); err != nil {
		err := errs.Errf(errs.ErrUnauthorized, "request validation failed")
		s.handleErrorResponse(c, http.StatusUnprocessableEntity, InternalError, err)

		return
	}

	d := domains.Blog{
		ID:      uuid.NewString(),
		Title:   req.Title,
		Content: req.Content,
	}

	if err := s.blog.CreateBlog(c.Request.Context(), &d); err != nil {
		s.handleErrorResponse(c, http.StatusInternalServerError, InternalError, err)

		return
	}

	s.handleSuccessResponse(c, views.Blog{
		ID:        d.ID,
		Title:     d.Title,
		Content:   d.Content,
		CreatedAt: d.CreatedAt,
		UpdatedAT: d.UpdatedAT,
	})
}

// blog godoc
// @Router /blog/{id} [GET]
// @Summary blog
// @Description get blog by it`s id
// @Tags blog
// @Accept json
// @Produce json
// @Param id path string true "Blog ID"
// @Success 200 {object} R{data=views.Blog}
// @Failure 422 {object} R
// @Failure 500 {object} R
func (s *Server) getByID(c *gin.Context) {
	blogID := c.Param("id")

	b, err := s.blog.GetBlogByID(c.Request.Context(), blogID)
	if err != nil {
		s.handleErrorResponse(c, http.StatusInternalServerError, InternalError, err)

		return
	}

	s.handleSuccessResponse(c, views.Blog{
		ID:        blogID,
		Title:     b.Title,
		Content:   b.Content,
		CreatedAt: b.CreatedAt,
		UpdatedAT: b.UpdatedAT,
	})
}

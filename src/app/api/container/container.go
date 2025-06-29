package container

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/misalima/nano-link-backend/src/core/ports"
	"github.com/misalima/nano-link-backend/src/core/services"
	"github.com/misalima/nano-link-backend/src/infra/postgres"
)

type Container struct {
	db *pgxpool.Pool

	urlRepo      ports.URLRepository
	tagRepo      ports.TagRepository
	urlTagRepo   ports.URLTagRepository
	urlVisitRepo ports.URLVisitRepository

	urlService    ports.URLService
	tagService    ports.TagService
	urlTagService ports.URLTagService
}

func New(db *pgxpool.Pool) *Container {
	container := &Container{
		db: db,
	}

	container.initRepositories()

	container.initServices()

	return container
}

func (c *Container) initRepositories() {
	c.urlRepo = postgres.NewURLRepository(c.db)
	c.tagRepo = postgres.NewTagRepository(c.db)
	c.urlTagRepo = postgres.NewURLTagRepository(c.db)
	c.urlVisitRepo = postgres.NewURLVisitRepository(c.db)
}

func (c *Container) initServices() {
	c.urlService = services.NewURLService(c.urlRepo, c.urlVisitRepo)
	c.tagService = services.NewTagService(c.tagRepo)
	c.urlTagService = services.NewURLTagService(c.urlTagRepo, c.urlRepo, c.tagRepo)
}

func (c *Container) URLService() ports.URLService {
	return c.urlService
}

func (c *Container) TagService() ports.TagService {
	return c.tagService
}

func (c *Container) URLTagService() ports.URLTagService {
	return c.urlTagService
}

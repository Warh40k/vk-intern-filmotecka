package service

import (
	"errors"
	"github.com/Warh40k/vk-intern-filmotecka/internal/api/repository"
	"github.com/Warh40k/vk-intern-filmotecka/internal/domain"
	"log/slog"
)

//go:generate mockery --all --dry-run=false

var (
	ErrNotFound     = errors.New("not found")
	ErrBadRequest   = errors.New("bad request")
	ErrInternal     = errors.New("internal errors")
	ErrUnauthorized = errors.New("not authorized")
)

type Service struct {
	Authorization
	Actor
	Film
}

type Authorization interface {
	SignUp(user domain.User) error
	SignIn(username, password string) (string, error)
	GetUserById(id int) (domain.User, error)
}

type Actor interface {
	CreateActor(actor domain.Actor) (int, error)
	DeleteActor(id int) error
	UpdateActor(actor domain.Actor) error
	PatchActor(actor domain.ActorInput) (domain.Actor, error)
	ListActors(filmId int) ([]domain.Actor, error)
}

type Film interface {
	CreateFilm(film domain.Film, actorIds []int) (int, error)
	DeleteFilm(id int) error
	UpdateFilm(film domain.Film, actorIds []int) error
	PatchFilm(input domain.NullableFilm, actorIds []int) (domain.Film, error)
	ListFilms(sortBy, sortDir string, actorId int) ([]domain.Film, error)
	SearchFilm(query string) ([]domain.Film, error)
}

func NewService(repos *repository.Repository, log *slog.Logger) *Service {
	return &Service{
		Authorization: NewAuthService(repos, log),
		Actor:         NewActorService(repos, log),
		Film:          NewFilmService(repos, log),
	}
}

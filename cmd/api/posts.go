package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/atomicmeganerd/rcd-gopher-social/internal/store"
	"github.com/go-chi/chi"
)

type CreatePostPayload struct {
	Title string `json:"title"`

	// TODO: Add validation rules as we have set this to text type in the database. We need
	// to set a maximum length for the content. The instructor for this course did this
	// deliberately to show how to add validation rules and to highlight the vulnerability.
	Content string `json:"content"`

	Tags []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {

	var payload CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		_ = writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	userId := 1

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		// TODO: We'll replace this with the actual user ID once we have authentication
		UserID: int64(userId),
	}

	ctx := r.Context()
	if err := app.store.Posts.Create(ctx, post); err != nil {
		_ = writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		_ = writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		_ = writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ctx := r.Context()
	post, err := app.store.Posts.GetByID(ctx, int64(id))
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			_ = writeJSONError(w, http.StatusNotFound, err.Error())
		default:
			_ = writeJSONError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err = writeJSON(w, http.StatusOK, post); err != nil {
		_ = writeJSONError(w, http.StatusInternalServerError, err.Error())
	}
}

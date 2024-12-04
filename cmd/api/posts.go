package main

import (
	"net/http"

	"github.com/atomicmeganerd/rcd-gopher-social/internal/store"
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

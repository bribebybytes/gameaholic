package actions

import (

  "fmt"
  "net/http"
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/pop"
  "bribebybytes.in/gameaholic/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Level)
// DB Table: Plural (levels)
// Resource: Plural (Levels)
// Path: Plural (/levels)
// View Template Folder: Plural (/templates/levels/)

// LevelsResource is the resource for the Level model
type LevelsResource struct{
  buffalo.Resource
}

// List gets all Levels. This function is mapped to the path
// GET /levels
func (v LevelsResource) List(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  levels := &models.Levels{}

  // Paginate results. Params "page" and "per_page" control pagination.
  // Default values are "page=1" and "per_page=20".
  q := tx.PaginateFromParams(c.Params())

  // Retrieve all Levels from the DB
  if err := q.All(levels); err != nil {
    return err
  }

  // Add the paginator to the context so it can be used in the template.
  c.Set("pagination", q.Paginator)

  return c.Render(http.StatusOK, r.Auto(c, levels))
}

// Show gets the data for one Level. This function is mapped to
// the path GET /levels/{level_id}
func (v LevelsResource) Show(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Level
  level := &models.Level{}

  // To find the Level the parameter level_id is used.
  if err := tx.Find(level, c.Param("level_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  return c.Render(http.StatusOK, r.Auto(c, level))
}

// New renders the form for creating a new Level.
// This function is mapped to the path GET /levels/new
func (v LevelsResource) New(c buffalo.Context) error {
  return c.Render(http.StatusOK, r.Auto(c, &models.Level{}))
}
// Create adds a Level to the DB. This function is mapped to the
// path POST /levels
func (v LevelsResource) Create(c buffalo.Context) error {
  // Allocate an empty Level
  level := &models.Level{}

  // Bind level to the html form elements
  if err := c.Bind(level); err != nil {
    return err
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(level)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    // Make the errors available inside the html template
    c.Set("errors", verrs)

    // Render again the new.html template that the user can
    // correct the input.
    return c.Render(http.StatusUnprocessableEntity, r.Auto(c, level))
  }

  // If there are no errors set a success message
  c.Flash().Add("success", T.Translate(c, "level.created.success"))
  // and redirect to the levels index page
  return c.Render(http.StatusCreated, r.Auto(c, level))
}

// Edit renders a edit form for a Level. This function is
// mapped to the path GET /levels/{level_id}/edit
func (v LevelsResource) Edit(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Level
  level := &models.Level{}

  if err := tx.Find(level, c.Param("level_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  return c.Render(http.StatusOK, r.Auto(c, level))
}
// Update changes a Level in the DB. This function is mapped to
// the path PUT /levels/{level_id}
func (v LevelsResource) Update(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Level
  level := &models.Level{}

  if err := tx.Find(level, c.Param("level_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  // Bind Level to the html form elements
  if err := c.Bind(level); err != nil {
    return err
  }

  verrs, err := tx.ValidateAndUpdate(level)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    // Make the errors available inside the html template
    c.Set("errors", verrs)

    // Render again the edit.html template that the user can
    // correct the input.
    return c.Render(http.StatusUnprocessableEntity, r.Auto(c, level))
  }

  // If there are no errors set a success message
  c.Flash().Add("success", T.Translate(c, "level.updated.success"))
  // and redirect to the levels index page
  return c.Render(http.StatusOK, r.Auto(c, level))
}

// Destroy deletes a Level from the DB. This function is mapped
// to the path DELETE /levels/{level_id}
func (v LevelsResource) Destroy(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Level
  level := &models.Level{}

  // To find the Level the parameter level_id is used.
  if err := tx.Find(level, c.Param("level_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  if err := tx.Destroy(level); err != nil {
    return err
  }

  // If there are no errors set a flash message
  c.Flash().Add("success", T.Translate(c, "level.destroyed.success"))
  // Redirect to the levels index page
  return c.Render(http.StatusOK, r.Auto(c, level))
}

package models

import (
    "errors"
    "strconv"
    "time"
)

const MaxPostLength = 2000

type Post struct {
    ID         int       `json:"id"`
    UserID     int       `json:"user_id"`
    CategoryID int       `json:"category_id"`
    Title      string    `json:"title"`
    Content    string    `json:"content"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}

func (post *Post) Validate() error {
    if len(post.Content) > MaxPostLength {
        return errors.New("post content exceeds maximum length of " + strconv.Itoa(MaxPostLength) + " characters")
    }
    return nil
}

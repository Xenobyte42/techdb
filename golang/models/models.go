package models

import (
	_ "encoding/json"
)

type ForumModel struct {
	Posts   int64  `json:"posts,readonly"`
	Slug    string `json:"slug"`
	Threads int32  `json:"threads,readonly"`
	Title   string `json:"title"`
	User    string `json:"user"`
}

type ForumGetRequest struct {
	Limit int64  `json:"limit"`
	Since string `json:"since"`
	Desc  bool   `json:"desc"`
}

type ThreadModel struct {
	Author  string `json:"author"`
	Created string `json:"created,omitempty"`
	Forum   string `json:"forum,readonly"`
	Id      int32  `json:"id,readonly"`
	Message string `json:"message"`
	Slug    string `json:"slug,omitempty"`
	Title   string `json:"title"`
	Votes   int32  `json:"votes,readonly"`
}

type ThreadUpdate struct {
	Message string `json:"message"`
	Title   string `json:"title"`
}

type ThreadGetRequest struct {
	Limit int64  `json:"limit"`
	Since int64  `json:"since"`
	Sort  string `json:"sort"`
	Desc  bool   `json:"desc"`
}

type UserModel struct {
	About    string `json:"about"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Nickname string `json:"nickname,readonly"`
}

type UserUpdate struct {
	About    string `json:"about"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

type DatabaseStatus struct {
	Forum  int32 `json:"forum"`
	Post   int64 `json:"post"`
	Thread int32 `json:"thread"`
	User   int32 `json:"user"`
}

type PostModel struct {
	Author   string `json:"author"`
	Created  string `json:"created,omitempty,readonly"`
	Forum    string `json:"forum,readonly"`
	Id       int64  `json:"id,readonly"`
	IsEdited bool   `json:"isEdited,readonly"`
	Message  string `json:"message"`
	Parent   int64  `json:"parent"`
	Thread   int32  `json:"thread,readonly"`
}

type PostUpdate struct {
	Message string `json:"message"`
}

type PostGetRequest struct {
	Related []string `json:"related"`
}

type PostFull struct {
	Author UserModel   `json:"author"`
	Forum  ForumModel  `json:"forum"`
	Post   PostModel   `json:"post"`
	Thread ThreadModel `json:"thread"`
}

type VoteModel struct {
	Nickname string `json:"nickname"`
	Voice    int32  `json:"voice"`
}

type ErrorResponse struct {
	Message string `json:"message,readonly"`
}

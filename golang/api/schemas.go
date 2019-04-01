package api

const (
	userSchema = `
	CREATE TABLE IF NOT EXISTS Users (
	    id BIGSERIAL PRIMARY KEY,
	    about TEXT NOT NULL,
	    email TEXT NOT NULL UNIQUE,
	    fullname TEXT NOT NULL,
	    nickname TEXT NOT NULL UNIQUE 
	)
	`

	forumSchema = `
	CREATE TABLE IF NOT EXISTS Forums (
		id BIGSERIAL PRIMARY KEY,
		slug TEXT NOT NULL UNIQUE,
		title TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES Users (id)
	)
	`

	threadSchema = `
	CREATE TABLE IF NOT EXISTS Threads (
		id BIGSERIAL PRIMARY KEY,
		author_id INTEGER NOT NULL,
		created TIMESTAMP NOT NULL,
		forum_id INTEGER NOT NULL,
		message TEXT NOT NULL,
		slug TEXT NOT NULL,
		title TEXT NOT NULL,
		FOREIGN KEY (author_id) REFERENCES Users (id),
		FOREIGN KEY (forum_id) REFERENCES Forums (id)
	)
	`

	voteSchema = `
	CREATE TABLE IF NOT EXISTS Votes (
		user_id INTEGER NOT NULL,
		thread_id INTEGER NOT NULL,
		voice INTEGER CHECK (voice = -1 OR voice = 1),
		FOREIGN KEY (user_id) REFERENCES Users (id),
		FOREIGN KEY (thread_id) REFERENCES Threads (id)
	)
	`

	postSchema = `
	CREATE TABLE IF NOT EXISTS Posts (
		id BIGSERIAL PRIMARY KEY,
		author_id INTEGER NOT NULL,
		created TIMESTAMP NOT NULL,
		is_edited BOOLEAN NOT NULL DEFAULT FALSE,
		message TEXT NOT NULL,
		parent INTEGER,
		thread_id INTEGER NOT NULL,
		FOREIGN KEY (author_id) REFERENCES Users (id),
		FOREIGN KEY (thread_id) REFERENCES Threads (id)
	)
	`
)

var (
	tables  = [...]string{"Users", "Forums", "Threads", "Votes", "Posts"}
	schemas = [...]string{userSchema, forumSchema, threadSchema, voteSchema, postSchema}
)

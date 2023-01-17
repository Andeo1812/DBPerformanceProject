package pkg

type GetSlugThreadsParams struct {
	Limit uint32
	Since string
	Desc  bool
}

type GetSlugUsersParams struct {
	Limit uint32
	Since string
	Desc  bool
}

type GetPostsParams struct {
	Limit uint32
	Since string
	Desc  bool
	Sort  string
}

type VoteParams struct {
	Nickname string
	Voice    int32
}

type PostDetails struct {
	Related []string
}

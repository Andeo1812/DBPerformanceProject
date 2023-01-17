package pkg

type GetThreadsParams struct {
	Limit uint32
	Since string
	Desc  bool
}

type GetUsersParams struct {
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

type PostDetailsParams struct {
	Related []string
}

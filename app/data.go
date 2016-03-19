package app

type User struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Nickname    string  `json:"nickname"`
	AvatarURL   string  `json:"avatar_url"`
	Age         int     `json:"age"`
	IsSingle    bool    `json:"is_single"`
	BestFriends []*User `json:"best_friends"`
}
type Post struct {
	ID         int     `json:"id"`
	Title      string  `json:"title"`
	Author     *User   `json:"author"`
	Liked      []*User `json:"liked"`
	LikedCount int     `json:"likedCount"`
	ViewCount  int     `json:"viewCount"`
}

var Users = []*User{}
var Posts = []*Post{}

func init() {

	user1 := &User{
		ID:          1,
		Name:        "Hafiz",
		Nickname:    "sogko",
		AvatarURL:   "/assets/avatar/1.jpg",
		Age:         27,
		IsSingle:    false,
		BestFriends: []*User{},
	}
	user2 := &User{
		ID:          2,
		Name:        "Emily",
		Nickname:    "M",
		AvatarURL:   "/assets/avatar/2.jpg",
		Age:         27,
		IsSingle:    true,
		BestFriends: []*User{},
	}
	user3 := &User{
		ID:          3,
		Name:        "Sam",
		Nickname:    "Starbucks",
		AvatarURL:   "/assets/avatar/3.jpg",
		Age:         24,
		IsSingle:    false,
		BestFriends: []*User{user1, user2},
	}
	post1 := &Post{
		ID:        1,
		Title:     "Hello World",
		Author:    user1,
		Liked:     []*User{user1, user2, user3},
		ViewCount: 10023,
	}
	post2 := &Post{
		ID:        2,
		Title:     "Good morning, John",
		Author:    user2,
		Liked:     []*User{user1, user2},
		ViewCount: 123,
	}
	post3 := &Post{
		ID:        3,
		Title:     "Hey, how about me?",
		Author:    user3,
		Liked:     []*User{},
		ViewCount: 2,
	}
	post4 := &Post{
		ID:        4,
		Title:     "Hey, how about meeeee?",
		Author:    user3,
		Liked:     []*User{},
		ViewCount: 2,
	}
	post5 := &Post{
		ID:        5,
		Title:     "HMU boo boo",
		Author:    user1,
		Liked:     []*User{user3},
		ViewCount: 234,
	}
	post6 := &Post{
		ID:        6,
		Title:     "Sent you a PM",
		Author:    user3,
		Liked:     []*User{},
		ViewCount: 234,
	}
	post7 := &Post{
		ID:        7,
		Title:     "Lorem ipsum",
		Author:    user3,
		Liked:     []*User{},
		ViewCount: 234,
	}
	post8 := &Post{
		ID:        8,
		Title:     "dolor sit amet",
		Author:    user3,
		Liked:     []*User{},
		ViewCount: 234,
	}
	post9 := &Post{
		ID:        9,
		Title:     "consectetur adipiscing elit",
		Author:    user3,
		Liked:     []*User{},
		ViewCount: 234,
	}
	post10 := &Post{
		ID:        10,
		Title:     "sed do eiusmod tempor incididunt",
		Liked:     []*User{},
		ViewCount: 234,
	}

	Posts = []*Post{
		post1,
		post2,
		post3,
		post4,
		post5,
		post6,
		post7,
		post8,
		post9,
		post10,
	}

	Users = []*User{
		user1,
		user2,
		user3,
	}
}

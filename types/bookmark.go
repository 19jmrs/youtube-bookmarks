package types

type Tag struct{
	Name string
}

type Bookmark struct{
	ID string
	Note string
	Tag Tag
	User User	
}

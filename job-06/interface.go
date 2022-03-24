package main

type CircleOfFriends interface {
	PostState(pi *PersonalInformation) error
	DeleteState(id string) error
	GetState() ([]*CircleOfFriendsInfo, error)
	DeleteAll() error
}

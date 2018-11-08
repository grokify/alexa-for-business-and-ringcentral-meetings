package a4boffice365

type AlexaForBusinessOffice365Data struct {
	AllMailboxes []bool `short:"m" long:"allmailboxes" description:"All Mailboxes"`
	Username     string `short:"u" long:"username" description:"Username" required:"true"`
	FirstName    string `short:"f" long:"firstname" description:"First name" required:"true"`
	LastName     string `short:"l" long:"lastname" description:"Last name" required:"true"`
	RoomName     string `short:"r" long:"roomname" description:"A room name" required:"true"`
}

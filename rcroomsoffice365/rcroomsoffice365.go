package rcroomsoffice365

type RingCentralRoomsOffice365Data struct {
	China    []bool `short:"c" long:"china" description:"China"`
	RoomName string `short:"n" long:"name" description:"A room name" required:"true"`
}

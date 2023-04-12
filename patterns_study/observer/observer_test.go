package observer

import "testing"

func TestStreamingService(t *testing.T) {
	service := &StreamingService{}
	user1, user2 := &User{}, &User{}
	service.AddObserver(user1)
	service.AddObserver(user2)
	service.Notify("12345")
	if user1.data != user2.data {
		t.Fatalf("User datas not equal! Values: %v, %v", user1.data, user2.data)
	}
}

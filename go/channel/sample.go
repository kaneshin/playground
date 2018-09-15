package main

import "fmt"

type User struct {
	ID   int
	Name string
}

type Image struct {
	ID     int
	UserID int
	URL    string
}

func GetUsersByIDs(ids []int) ([]User, error) {
	return nil, nil
}

func GetImagesByIDs(ids []int) ([]Images, error) {
	return nil, nil
}

func run() error {
	usersCh := make(chan map[int]User)
	imagesCh := make(chan map[int]Image)
	errCh := make(chan error)

	targetUserIDs := []int{1, 2, 3}
	targetImageIDs := []int{10, 20, 30}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		// defer func() { wg.Done() }()

		for {
			select {
			case err := <-errCh:
				wg.Done()
				return
			}
		}
	}()

	go func(chan map[int]User, chan error) {
		users, err := GetUsersByIDs(targetUserIDs)
		if err != nil {
			errCh <- err
			return
		}

		um := map[int]User{}
		for _, u := range users {
			um[u.ID] = u
		}
		usersCh <- um
	}(usersCh, errCh)

	go func(chan map[int]Image, chan error) {
		images, err := GetImagesByIDs(targetImageIDs)
		if err != nil {
			errCh <- err
			return
		}

		im := map[int]Image{}
		for _, i := range images {
			im[i.UserID] = i
		}
		imagesCh <- im
	}(imagesCh, errCh)

	wg.Wait()

	return 0
}

func main() {
	fmt.Println(run())
}

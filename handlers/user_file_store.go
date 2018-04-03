package handlers

type FileUserInfoStore struct {
}

func (s *FileUserInfoStore) Stop(userID, videoID string) error {
	file, err := LockFile(userID)
	if err != nil {
		return err
	}
	defer file.Unlock()

	err = UpdateUserInfo(file.Path(), "stop", videoID)
	if err != nil {
		return err
	}

	return nil
}

package utils

import (
	"errors"
	"github.com/gofrs/flock"
	"github.com/spf13/viper"
	"time"
)

func UidAcquireLock() (*flock.Flock, error) {
	lockCount := 0

reAcquireLock:
	fileLock := flock.New(viper.GetString("ldap_adm.uid_lock_file"))

	locked, err := fileLock.TryLock()
	if err != nil {
		return nil, err
	}
	if !locked {
		lockCount++
		time.Sleep(100 * time.Millisecond)
		if lockCount > 100 {
			return nil, errors.New("failed to acquire lock, Please wait for 1 minute and try again")
		}
		goto reAcquireLock
	}
	return fileLock, nil
}

func GidAcquireLock() (*flock.Flock, error) {
	lockCount := 0

reAcquireLock:
	fileLock := flock.New(viper.GetString("ldap_adm.gid_lock_file"))

	locked, err := fileLock.TryLock()
	if err != nil {
		return nil, err
	}
	if !locked {
		lockCount++
		time.Sleep(100 * time.Millisecond)
		if lockCount > 100 {
			return nil, errors.New("failed to acquire lock, Please wait for 1 minute and try again")
		}
		goto reAcquireLock
	}
	return fileLock, nil
}

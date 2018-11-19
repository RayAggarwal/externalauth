package util

import (
	"os"
	"sync"
	"time"
)

const hoursRotateLog uint8 = 24

type RotateWriter struct {
	lock sync.Mutex
	filename string
	file *os.File
	previousRotate time.Time
}

func (w *RotateWriter) Write(output []byte) (int, error) {
	w.lock.Lock()
	defer w.lock.Unlock()
	var currentTime = time.Now()
	if currentTime.Sub(w.previousRotate).Hours() >= hoursRotateLog {
		w.rotate()
	}
	return w.file.Write(output)
}

func (w *RotateWriter) rotate() (err, err) {
	if w.file != nil {
		err := w.file.Close()
		if err != nil {
			panic("Can't close log file")
			return
		}
	}
	_, err = os.Stat(w.filename)
	if err == nil {
		err = os.Rename(w.filename, w.filename+"."+time.Now().Format(time.RFC3339))
		if err != nil {
			panic("Error renaming log file")
			return
		}
	}

	w.file, err = os.Create(w.filename)
}
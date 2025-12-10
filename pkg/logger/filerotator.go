// SPDX-FileCopyrightText: 2025 UnionTech Software Technology Co., Ltd.
// SPDX-License-Identifier: MIT
package logger

import (
        "os"
        "time"
)

type FileRotator struct {
        basePath  string
        maxSize   int64
        maxAge    time.Duration
        current   *os.File
        size      int64
        startTime time.Time
}

func (fr *FileRotator) Write(p []byte) (n int, err error) {
        err = fr.setupCurrent()
        if err != nil {
                return 0, err
        }
        n, err = fr.current.Write(p)
        if err != nil {
                return n, err
        }
        fr.size += int64(n)
        return n, nil
}

func (fr *FileRotator) setupCurrent() error {
        if fr.current == nil {
                fileinfo, err := os.Stat(fr.basePath)
                if err == nil {
                        fr.current, err = os.OpenFile(fr.basePath, os.O_APPEND|os.O_WRONLY, 0600)
                        if err != nil {
                                return err
                        }
                        fr.size = fileinfo.Size()
                        fr.startTime = fileinfo.ModTime()
                } else if os.IsNotExist(err) {
                        fr.current, err = os.Create(fr.basePath)
                        if err != nil {
                                return err
                        }
                        fr.startTime = time.Now()
                } else {
                        return err
                }
        }
        return nil
}

package filestore

import (
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type couponStore struct {
	paths []string
	index map[string]int
	mu    sync.RWMutex
}

func newCouponStore(paths []string) *couponStore {
	return &couponStore{
		paths: paths,
		index: make(map[string]int),
	}
}

func (c *couponStore) loadCoupons() error {
	if len(c.paths) == 0 {
		return errors.New("no coupon files configured")
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(c.paths))

	for _, path := range c.paths {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()

			file, err := os.Open(path)
			if err != nil {
				errCh <- fmt.Errorf("failed to open %s: %w", path, err)
				return
			}
			defer file.Close()

			gzr, err := gzip.NewReader(file)
			if err != nil {
				errCh <- fmt.Errorf("failed to create gzip reader for %s: %w", path, err)
				return
			}
			defer gzr.Close()

			seen := make(map[string]struct{})
			reader := bufio.NewReader(gzr)

			for {
				line, err := reader.ReadString('\n')
				if len(line) > 0 {
					code := strings.ToUpper(strings.TrimSpace(line))
					if code == "" {
						continue
					}
					if _, exists := seen[code]; !exists {
						seen[code] = struct{}{}

						c.mu.Lock()
						c.index[code]++
						c.mu.Unlock()
					}
				}
				if err == io.EOF {
					break
				}
				if err != nil {
					errCh <- fmt.Errorf("read error in %s: %w", path, err)
					return
				}
			}
		}(path)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *couponStore) validate(code string) error {
	code = strings.TrimSpace(strings.ToUpper(code))
	if l := len(code); l < 8 || l > 10 {
		return errors.New("invalid promo: code length must be between 8 and 10")
	}

	c.mu.RLock()
	count := c.index[code]
	c.mu.RUnlock()

	if count >= 2 {
		return nil
	}
	return errors.New("invalid promo: not found in enough coupon files")
}

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

// loadCoupons builds the in-memory index from all files.
func (c *couponStore) loadCoupons() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.paths) == 0 {
		return errors.New("no coupon files configured")
	}

	for _, path := range c.paths {
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open %s: %w", path, err)
		}
		defer file.Close()

		gzr, err := gzip.NewReader(file)
		if err != nil {
			return fmt.Errorf("failed to create gzip reader for %s: %w", path, err)
		}
		defer gzr.Close()

		seen := map[string]struct{}{}
		reader := bufio.NewReader(gzr)

		for {
			line, err := reader.ReadString('\n')
			if len(line) > 0 {
				code := strings.ToUpper(strings.TrimSpace(line))
				if code == "" {
					continue
				}
				if _, exists := seen[code]; !exists {
					c.index[code]++
					seen[code] = struct{}{}
				}
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				return fmt.Errorf("read error in %s: %w", path, err)
			}
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
	defer c.mu.RUnlock()

	count := c.index[code]
	if count >= 2 {
		return nil
	}
	return errors.New("invalid promo: not found in enough coupon files")
}

package replacer

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/no-src/log"
	"github.com/no-src/nsgo/jsonutil"
	"github.com/no-src/nsgo/yamlutil"
	"github.com/no-src/replacer/conf"
)

const skipValue = "TODO"

// Replacer a configuration-based file replace tool
type Replacer struct {
	confFile string
	conf     conf.Conf
}

// NewReplacer returns the replacer instance
func NewReplacer(confFile string) (*Replacer, error) {
	r := &Replacer{
		confFile: confFile,
	}
	if err := r.parse(); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Replacer) parse() error {
	data, err := os.ReadFile(r.confFile)
	if err != nil {
		return err
	}
	err = yamlutil.Unmarshal(data, &r.conf)
	if err != nil {
		return err
	}
	return nil
}

// Replace start execute replace in the root workspace
func (r *Replacer) Replace(root, tag string) error {
	return r.replaceOrRevert(root, tag, false)
}

// Revert try to revert the replace operation
func (r *Replacer) Revert(root, tag string) error {
	return r.replaceOrRevert(root, tag, true)
}

func (r *Replacer) replaceOrRevert(root, tag string, reverse bool) error {
	for _, item := range r.conf.Items {
		if item.Disabled {
			log.Info("[replacer skip] [%s] is disabled", item.Name)
			continue
		}
		for _, path := range item.Paths {
			if err := r.replaceRecur(item.Rules, tag, path, root, reverse); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *Replacer) replaceRecur(rules []conf.Rule, tag string, matchPath string, root string, reverse bool) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		ok, err := filepath.Match(filepath.ToSlash(matchPath), filepath.ToSlash(path))
		if err != nil {
			return err
		}
		if ok {
			return r.replace(rules, tag, path, reverse)
		}
		return nil
	})
}

func (r *Replacer) replace(rules []conf.Rule, tag string, path string, reverse bool) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(data)
	newContent := content
	var records []record
	for _, rule := range rules {
		for _, old := range rule.Old {
			newStr, ok := rule.New[tag]
			isTodo := newStr == skipValue
			if ok && !isTodo {
				temp := newContent
				if reverse {
					old, newStr = newStr, old
				}
				newContent = strings.ReplaceAll(newContent, old, newStr)
				if temp != newContent {
					records = append(records, newRecord(old, newStr))
				}
			}
		}
	}
	if content == newContent {
		return nil
	}
	recordsData, _ := jsonutil.Marshal(records)
	log.Info("replace success, matched %d rules, %s => %s", len(records), path, recordsData)
	return os.WriteFile(path, []byte(newContent), os.ModePerm)
}

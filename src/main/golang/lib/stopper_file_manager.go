package lib

import (
	"bytes"
	"crypto/sha1"
	"os"
	"strconv"
	"strings"

	"github.com/starter-go/afs"
	"github.com/starter-go/application"
	"github.com/starter-go/base/lang"
)

const (
	stopperFileNamePrefix = "starter."
	stopperFileNameSuffix = ".stopper"
)

////////////////////////////////////////////////////////////////////////////////

type stopperContext struct {
	ac application.Context
	fs afs.FS
}

func (inst *stopperContext) locateAppContextDir() afs.Path {

	// compute sum
	mm := inst.ac.GetMainModule()
	sumBuilder := &bytes.Buffer{}
	sumBuilder.WriteString(mm.Name())
	sum := sha1.Sum(sumBuilder.Bytes())
	hex := lang.HexFromBytes(sum[0:10])

	// find base path
	dir := inst.fs.NewPath("/etc/starter/stopper/")
	home, err := os.UserHomeDir()
	if err == nil {
		dir = inst.fs.NewPath(home)
	} else {
		thisExe := os.Args[0]
		exe := inst.fs.NewPath(thisExe)
		dir = exe.GetParent()
	}

	return dir.GetChild(".starter/stopper/apps/" + hex.String())
}

////////////////////////////////////////////////////////////////////////////////

func getStopperFileManager(ctx *stopperContext) *stopperFileManager {
	man := &stopperFileManager{}
	man.context = *ctx
	man.appContextDir = ctx.locateAppContextDir()
	return man
}

////////////////////////////////////////////////////////////////////////////////

type stopperFileManager struct {
	context       stopperContext
	appContextDir afs.Path
}

func (inst *stopperFileManager) listAll() []*stopperFile {
	const (
		prefix = stopperFileNamePrefix
		suffix = stopperFileNameSuffix
	)
	src := inst.appContextDir.ListChildren()
	dst := make([]*stopperFile, 0)
	for _, item := range src {
		name := item.GetName()
		if strings.HasPrefix(name, prefix) && strings.HasSuffix(name, suffix) {
			file := inst.loadFile(item)
			dst = append(dst, file)
		}
	}
	return dst
}

func (inst *stopperFileManager) loadFile(path afs.Path) *stopperFile {
	info := path.GetInfo()
	t1 := info.CreatedAt()
	file := &stopperFile{
		manager:   inst,
		file:      path,
		timestamp: lang.NewTime(t1),
	}
	return file
}

func (inst *stopperFileManager) getNewer() *stopperFile {
	all := inst.listAll()
	var have *stopperFile
	for _, item := range all {
		if have != nil {
			if item.timestamp > have.timestamp {
				have = item
			}
		} else {
			have = item
		}
	}
	return have
}

func (inst *stopperFileManager) getOlder() *stopperFile {
	all := inst.listAll()
	var have *stopperFile
	for _, item := range all {
		if have != nil {
			if item.timestamp < have.timestamp {
				have = item
			}
		} else {
			have = item
		}
	}
	return have
}

func (inst *stopperFileManager) new() *stopperFile {
	const (
		prefix = stopperFileNamePrefix
		suffix = stopperFileNameSuffix
	)
	now := lang.Now()
	name := &strings.Builder{}
	name.WriteString(prefix)
	name.WriteString(strconv.FormatInt(int64(now), 10))
	name.WriteString(suffix)
	dir := inst.appContextDir
	file := dir.GetChild(name.String())
	dst := &stopperFile{
		timestamp: now,
		manager:   inst,
		file:      file,
	}
	return dst
}

////////////////////////////////////////////////////////////////////////////////

type stopperFile struct {
	timestamp lang.Time
	manager   *stopperFileManager
	file      afs.Path
}

func (inst *stopperFile) create() error {
	file := inst.file
	if file.Exists() {
		return nil
	}

	dir := file.GetParent()
	if !dir.Exists() {
		opt := afs.ToMakeDir()
		dir.Mkdirs(opt)
	}

	content := "" // 内容必须为空
	opt := afs.Todo().Create(true).Write(true).Options()
	return file.GetIO().WriteText(content, opt)
}

func (inst *stopperFile) remove() error {
	file := inst.file
	if file == nil {
		return nil
	}
	if !file.IsFile() {
		return nil
	}
	size := file.GetInfo().Length()
	if size > 0 {
		return nil // 必须是空文件
	}
	return file.Delete()
}

func (inst *stopperFile) exists() bool {
	return inst.file.Exists()
}

////////////////////////////////////////////////////////////////////////////////

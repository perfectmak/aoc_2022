package day7

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Cmd string

const Cd Cmd = "cd"
const Ls Cmd = "ls"

type Instruction struct {
	Cmd    Cmd
	Args   []string
	Output []string
}

type Path string

type Dir struct {
	Path Path
	Size uint64
}

func (d *Dir) ParentPath() Path {
	return Path(filepath.Dir(string(d.Path)))
}

func (d *Dir) ChildPath(path string) Path {
	return Path(filepath.Join(string(d.Path), path))
}

type Fs struct {
	dirs map[Path]*Dir
	cwd  *Dir
}

func NewFs() *Fs {
	return &Fs{
		dirs: map[Path]*Dir{},
	}
}

func (f *Fs) ProcessInstruction(instruction Instruction) {
	switch instruction.Cmd {
	case Cd:
		switch instruction.Args[0] {
		case "/":
			f.cwd = f.getOrCreateDir("/")
		case "..":
			if f.cwd == nil {
				panic("cwd is nil") // maybe set this to / instead of panicing
			}
			f.cwd = f.getOrCreateDir(f.cwd.ParentPath())
		default:
			f.cwd = f.getOrCreateDir(f.cwd.ChildPath(instruction.Args[0]))
		}

	case Ls:
		for _, output := range instruction.Output {
			parts := strings.Split(output, " ")
			switch parts[0] {
			case "dir":
				f.getOrCreateDir(f.cwd.ChildPath(parts[1]))
			default:
				fileSize, err := strconv.ParseUint(parts[0], 10, 64)
				if err != nil {
					panic(err)
				}
				f.addFileSizeToCwd(fileSize)
			}
		}
	}
}

func (f *Fs) getOrCreateDir(path Path) *Dir {
	if dir, exists := f.dirs[path]; exists {
		return dir
	}
	f.dirs[path] = &Dir{Path: path, Size: 0}
	return f.dirs[path]
}

func (f *Fs) addFileSizeToCwd(size uint64) {
	cwd := f.cwd
	for cwd.Path != "/" {
		cwd.Size += size
		cwd = f.getOrCreateDir(cwd.ParentPath())
	}

	// set parent directory size
	cwd.Size += size
}

func (f *Fs) TotalDirSize100K() uint64 {
	total := uint64(0)
	for _, dir := range f.dirs {
		//fmt.Printf("Path: %s, Size: %d\n", dir.Path, dir.Size)
		if dir.Size <= 100_000 {
			total += dir.Size
		}
	}

	return total
}

func (f *Fs) SmallestDeletableDir() uint64 {
	unusedSpace := 70_000_000 - f.dirs[Path("/")].Size
	smallestSize := uint64(70_000_000)
	for _, dir := range f.dirs {
		if unusedSpace+dir.Size >= 30_000_000 && dir.Size < smallestSize {
			smallestSize = dir.Size
		}
	}

	return smallestSize
}

func Solve() (uint64, error) {
	return Solve2()
}

func Solve1() (uint64, error) {
	//inputFile, err := os.Open("day7/sample.txt")
	inputFile, err := os.Open("day7/input.txt")
	defer inputFile.Close()
	if err != nil {
		return 0, err
	}

	inputReader := bufio.NewScanner(inputFile)

	cmd := ""
	args := []string{}
	output := []string{}
	fs := NewFs()
	for inputReader.Scan() {
		line := inputReader.Text()
		if line[0] == '$' {
			if cmd != "" {
				fs.ProcessInstruction(Instruction{
					Cmd:    Cmd(cmd),
					Args:   args,
					Output: output,
				})
			}

			parts := strings.Split(line, " ")
			cmd = parts[1]
			args = parts[2:]
			output = []string{}
		} else {
			output = append(output, line)
		}
	}

	if inputReader.Err() != nil {
		return 0, inputReader.Err()
	}

	// process final cmd/output
	fs.ProcessInstruction(Instruction{
		Cmd:    Cmd(cmd),
		Args:   args,
		Output: output,
	})

	return fs.TotalDirSize100K(), nil
}

func Solve2() (uint64, error) {
	//inputFile, err := os.Open("day7/sample.txt")
	inputFile, err := os.Open("day7/input.txt")
	defer inputFile.Close()
	if err != nil {
		return 0, err
	}

	inputReader := bufio.NewScanner(inputFile)

	cmd := ""
	args := []string{}
	output := []string{}
	fs := NewFs()
	for inputReader.Scan() {
		line := inputReader.Text()
		if line[0] == '$' {
			if cmd != "" {
				fs.ProcessInstruction(Instruction{
					Cmd:    Cmd(cmd),
					Args:   args,
					Output: output,
				})
			}

			parts := strings.Split(line, " ")
			cmd = parts[1]
			args = parts[2:]
			output = []string{}
		} else {
			output = append(output, line)
		}
	}

	if inputReader.Err() != nil {
		return 0, inputReader.Err()
	}

	// process final cmd/output
	fs.ProcessInstruction(Instruction{
		Cmd:    Cmd(cmd),
		Args:   args,
		Output: output,
	})

	return fs.SmallestDeletableDir(), nil
}

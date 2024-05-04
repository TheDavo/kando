package kando

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	// "time"
)

type StatusDetailed string
type Status int
type Statuses map[string][]string

const (
	Todo Status = iota
	InProgress
	Done
)

type Task struct {
	Id             int
	Status         Status
	StatusDetailed StatusDetailed
	Description    string
	// DueDate        time.Time
}

type Project struct {
	Name        string
	Description string
	LatestId    int `json:"latest-id"`
	Statuses    Statuses
	Tasks       map[int]Task
}

type Meta struct {
	Projects []string
	Filepath string
}

type Kando struct {
	Meta     Meta
	Projects map[string]*Project
}

var testJson = `

{
  "meta": {
    "projects": [
      "projectname"
	]
  },
  "projects": 
    {
      "projectname": {
        "name": "name",
        "description": "desc",
        "latest-id": 3,
        "statuses": {
          "todo": [
            "todo"
          ],
          "in-progress": [
            "in progress",
            "waiting"
          ],
          "done": [
            "done"
          ]
        },
        "tasks": {
          "1":{
            "id": 1,
            "status": 0,
            "description": "hello"
          },
          "2":{
            "status": 1,
            "status-detailed": "waiting",
            "id": 2,
            "description": "my"
          },
          "3":{
            "id": 3,
            "status": 2,
            "description": "world"
          }
        }
      }
    }
}`

func notMain() {
	var test Kando

	err := json.Unmarshal([]byte(testJson), &test)
	if err != nil {
		panic(err)
	}

	proj := test.Projects["projectname"]
	taskToAdd := Task{
		Status:         Todo,
		StatusDetailed: "lazy",
		Description:    "lazy",
	}

	proj.AddTask(taskToAdd)
	err = proj.RemoveTask(0)
	if err != nil {
		fmt.Println(err)
	}

	proj.RemoveTask(1)

	for _, val := range proj.Tasks {
		fmt.Printf("\nTask \tid:%d\n\tDescription: %s\n",
			val.Id, val.Description)
	}

	test.AddProject("testAddProject")
	fmt.Println(test)

}

func KandoFileExists(makeIfNot bool) (string, bool, error) {
	existed := true
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	pathToKando := filepath.Join(homeDir, "kando", "kando.json")

	_, err = os.Stat(pathToKando)

	// if the file does not exist and a new one has to be made
	if err != nil && makeIfNot {
		os.Mkdir(filepath.Join(homeDir, "kando"), os.ModeDir)
		file, err := os.Create(pathToKando)
		if err != nil {
			return "", false, err
		}
		file.Close()
		existed = false
	} else if err != nil && !makeIfNot {
		return "", false, err
	}

	return pathToKando, existed, nil
}

func Open() *Kando {
	p, _, err := KandoFileExists(false)

	if err != nil {
		panic(err)
	}

	k := FromFilePath(p)

	return k
}

func NewProject(name string) *Project {
	p := &Project{
		Name:        name,
		Description: "",
		LatestId:    0,
		Statuses: Statuses{
			"todo":        []string{"todo"},
			"in-progress": []string{"in-progress"},
			"done":        []string{"done"},
		},
		Tasks: make(map[int]Task),
	}

	return p
}

func FromFilePath(fp string) *Kando {
	var k *Kando

	bits, err := os.ReadFile(fp)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bits, &k)
	if err != nil {
		panic(err)
	}
	k.Meta.Filepath = fp

	return k
}

func NewKando(firstProjectName, fp string) *Kando {
	newMeta := Meta{
		Projects: []string{firstProjectName},
		Filepath: fp,
	}
	newProj := NewProject(firstProjectName)
	projs := map[string]*Project{
		firstProjectName: newProj,
	}
	k := &Kando{
		Meta:     newMeta,
		Projects: projs,
	}
	return k
}

func (p *Project) AddTask(task Task) {
	p.LatestId++
	task.Id = p.LatestId
	p.Tasks[p.LatestId] = task
}

func (p *Project) RemoveTask(id int) error {

	_, exists := p.Tasks[id]

	if exists {
		delete(p.Tasks, id)
		return nil
	}

	return errors.New(
		fmt.Sprintf("Task with id %d does not exist!", id))
}

func (k *Kando) Save() error {
	file, err := os.OpenFile(k.Meta.Filepath, os.O_RDWR, 0644)

	// Clear the contents of the file before writing the updated Kando content
	file.Truncate(0)
	if err != nil {
		fmt.Println("File does not exist!")
		panic("Error opening file")
	}

	bitties, _ := json.MarshalIndent(k, "", "\t")
	_, err = file.Write(bitties)
	if err != nil {
		return err
	}
	return nil
}
func (k *Kando) AddProject(projName string) error {
	k.Meta.Projects = append(k.Meta.Projects, projName)
	k.Projects[projName] = NewProject(projName)

	return nil
}

func (k *Kando) RemoveProject(projName string) error {
	_, exists := k.Projects[projName]

	if exists {
		delete(k.Projects, projName)

		// Find index of the project in the Meta field and modify the slice
		// to remove the Project
		var idx int
		for i, v := range k.Meta.Projects {
			if v == projName {
				idx = i
			}
		}
		if idx == len(k.Meta.Projects) {
			k.Meta.Projects = k.Meta.Projects[:idx]
		} else {
			k.Meta.Projects = append(k.Meta.Projects[:idx],
				k.Meta.Projects[idx+1:]...)
		}

		return nil
	}

	return errors.New(
		fmt.Sprintf("Project \"%s\" not found, nothing to delete",
			projName))

}

package rover

import(
	"fmt"
	"math"
        "errors"
        "strings"
)

type Coordinates struct {
        X int
	Y int
}

// TODO: Obstacle has to be filled and handle, including other existing rovers
type Obstacle struct {
        Coordinates
}

type Map struct {
        Height int
        Width int
        Rovers map[string]*rover
        Obstacles []Obstacle
}

// TODO: Wrap rover movements to the map height and width
func (m *Map) NewRover(name string, p Position) (*rover) {

	r := &rover{Name: name}
	r.Path = append(r.Path, p)

	if m.Rovers == nil {
		m.Rovers = make(map[string]*rover)
	}

	m.Rovers[r.Name] = r
	return r
}

type Position struct {
	Facing int16 // In Degrees
        Coordinates
}

func (p *Position) Move(amount int) {
	p.X = p.X + amount*int(math.Cos(float64(p.Facing)*math.Pi/180.0))
	p.Y = p.Y + amount*int(math.Sin(float64(p.Facing)*math.Pi/180.0))
}

func (p *Position) Turn(degree int16) {
	p.Facing = (p.Facing + degree) % 360
}

type rover struct {
        Name string
        Path []Position
}

func (r *rover) RunCommands(cmds string) error {

        commands := strings.Split(cmds, "")
        for _, c := range commands {
                switch c {
                        case "f":
                                r.moveForward()
                        case "b":
                                r.moveBackward()
                        case "r":
                                r.turnRight()
                        case "l":
                                r.turnLeft()
			default:
				// Halt upon unknown command
                                return errors.New(fmt.Sprintf("Unknown command: %s", c))
		}
		fmt.Println(r.Representation())
        }

        return nil
}

func (r *rover) Representation() string {
	p := r.Position()
	return fmt.Sprintf("%s is at X:%v Y:%v Facing:%v°", r.Name, p.X, p.Y, p.Facing)
}

func (r *rover) Position() Position{
	return r.Path[len(r.Path)-1]
}

func (r *rover) moveForward() {
	fmt.Println("Moving forward")
	p := r.Position()
	p.Move(1)
	r.Path = append(r.Path, p)
}

func (r *rover) moveBackward() {
	fmt.Println("Moving backward")
	p := r.Position()
	p.Move(-1)
	r.Path = append(r.Path, p)
}

func (r *rover) turnRight() {
	fmt.Println("Turning right")
	p := r.Position()
	p.Turn(90)
	r.Path = append(r.Path, p)
}

func (r *rover) turnLeft() {
	fmt.Println("Turning left")
	p := r.Position()
	p.Turn(-90)
	r.Path = append(r.Path, p)
}
	



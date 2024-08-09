package logic

type Action int

const (
	Move Action = iota + 1
	Attack
	Defend
	Exit
)

var Actions = []string{"Move", "Attack", "Defend", "Exit"}

func (a Action) String() string {
	return Actions[a-1]
}

func (a Action) IsValid() bool {
	return a >= Move && a <= Exit
}

func (a Action) IsExit() bool {
	return a == Exit
}

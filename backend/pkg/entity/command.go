package entity

import "strings"

type BashCommandsBuilder struct {
	subCommands []string
	numCommands int
}

func NewBashCommandsBuilder() *BashCommandsBuilder {
	return &BashCommandsBuilder{numCommands: 0}
}

func (b *BashCommandsBuilder) NewCommand(command string, args ...string) *BashCommandsBuilder {
	if b.numCommands > 0 {
		b.subCommands = append(b.subCommands, "&&")
	}
	b.subCommands = append(b.subCommands, command)
	b.subCommands = append(b.subCommands, args...)
	b.numCommands++
	return b
}

func (b *BashCommandsBuilder) Build() []string {
	commands := []string{"bash", "-c"}
	commands = append(commands, strings.Join(b.subCommands, " "))
	return commands
}

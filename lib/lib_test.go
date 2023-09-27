package lib

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestSuite(t *testing.T) {
	should.Run(&Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Suite struct {
	*should.T
}

func (this *Suite) Setup() {}

func (this *Suite) TestIsMatch() {
	input := "w" + "a" + "n" + "d" + "s" + "t"
	this.So(IsMatch(input, "wand"), should.BeTrue)
	this.So(IsMatch(input, "wat"), should.BeTrue)
	this.So(IsMatch(input, "nope"), should.BeFalse)
	this.So(IsMatch(input, "waand"), should.BeFalse)
}
func (this *Suite) TestFindMatches() {
	input := "w" + "a" + "n" + "d" + "s" + "t"
	dict := []string{
		"a",    // too short
		"an",   // yes
		"and",  // yes
		"am",   // no
		"wand", // yes
		"sand", // yes
		"tan",  // yes
	}
	this.So(FindMatches(input, dict...), should.Equal, []string{
		"and",  // yes
		"wand", // yes
		"sand", // yes
		"tan",  // yes
	})
}

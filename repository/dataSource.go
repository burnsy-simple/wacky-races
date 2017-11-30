package repository

// We initialise our repository source data here.
// In reality this would be coming from e.g. a DB but we hard-code it
// so we can at least demonstrate something
import (
	"time"

	"github.com/burnsy/wacky-races/models"
	"gopkg.in/mgo.v2/bson"
)

// We'll use hard-coded data rather than talk to a data source.

var competitorsSrc []*models.Competitor
var racesSrc models.Races
var meetsSrc []*models.Meet

func init() {
	competitorsSrc = []*models.Competitor{
		newThoroughbred("Shergar"),
		newThoroughbred("Johnny White Sox"),
		newThoroughbred("Blue Jade"),
		newThoroughbred("Black Caviar"),
		newThoroughbred("Phar Lap"),
		newThoroughbred("Going Going Go"),
		newThoroughbred("Carbine"),
		newThoroughbred("Rain Lover"),
		newThoroughbred("Red Rum"),
		newThoroughbred("Saintly"),
		newGreyhound("Flying Amy"),
		newGreyhound("Rapid Journey"),
		newGreyhound("Macareena"),
		newGreyhound("Black Top"),
		newGreyhound("Leg Spinner"),
		newGreyhound("Fastly"),
		newHarness("Beautide"),
		newHarness("Mr Feelgood"),
		newHarness("Blacks A Fake"),
		newHarness("Weona Warrior"),
		newHarness("Jack Morris"),
		newHarness("Bill Bixby"),
	}

	doombenMeet := newMeet("Doomben", models.Thoroughbred)
	sunshineMeet := newMeet("DoomSunshine Coastben", models.Greyhound)
	goldMeet := newMeet("Gold Coast", models.Harness)
	meetsSrc = []*models.Meet{
		doombenMeet,
		sunshineMeet,
		goldMeet,
	}

	nextThoroughbredStartTime := time.Now().Add(-2 * time.Minute)
	nextHarnessStartTime := time.Now().Add(-4 * time.Minute)
	nextGreyhoundStartTime := time.Now().Add(-4 * time.Minute)
	racesSrc = models.Races{
		newThoroughbredRace(doombenMeet.ID, "Jake Brown Handicap", &nextThoroughbredStartTime),
		newThoroughbredRace(doombenMeet.ID, "Gold Handicap", &nextThoroughbredStartTime),
		newThoroughbredRace(doombenMeet.ID, "Smalls Handicap", &nextThoroughbredStartTime),
		newThoroughbredRace(doombenMeet.ID, "Bloom Handicap", &nextThoroughbredStartTime),
		newThoroughbredRace(doombenMeet.ID, "Ulysees Handicap", &nextThoroughbredStartTime),
		newHarnessRace(goldMeet.ID, "Bladerunner Memorial", &nextHarnessStartTime),
		newHarnessRace(goldMeet.ID, "Doolittle Do", &nextHarnessStartTime),
		newHarnessRace(goldMeet.ID, "Race 3", &nextHarnessStartTime),
		newHarnessRace(goldMeet.ID, "Race 4", &nextHarnessStartTime),
		newGreyhoundRace(sunshineMeet.ID, "Race 1", &nextGreyhoundStartTime),
		newGreyhoundRace(sunshineMeet.ID, "Race 2", &nextGreyhoundStartTime),
		newGreyhoundRace(sunshineMeet.ID, "Race 3", &nextGreyhoundStartTime),
		newGreyhoundRace(sunshineMeet.ID, "Race 4", &nextGreyhoundStartTime),
	}
}

func newMeet(location string, category models.RaceCategory) *models.Meet {
	return models.NewMeet(bson.NewObjectId().String(), location, category)
}

func newThoroughbred(name string) *models.Competitor {
	return models.NewThoroughbred(bson.NewObjectId().String(), name)
}
func newGreyhound(name string) *models.Competitor {
	return models.NewGreyhound(bson.NewObjectId().String(), name)
}
func newHarness(name string) *models.Competitor {
	return models.NewHarness(bson.NewObjectId().String(), name)
}

func newThoroughbredRace(meetID string, name string, start *time.Time) *models.Race {
	*start = start.Add(2 * time.Minute)
	return models.NewThoroughbredRace(bson.NewObjectId().String(), meetID, name, *start, start.Add(-5*time.Second))
}
func newGreyhoundRace(meetID string, name string, start *time.Time) *models.Race {
	*start = start.Add(3 * time.Minute)
	return models.NewGreyhoundRace(bson.NewObjectId().String(), meetID, name, *start, start.Add(-5*time.Second))
}
func newHarnessRace(meetID string, name string, start *time.Time) *models.Race {
	*start = start.Add(5 * time.Minute)
	return models.NewHarnessRace(bson.NewObjectId().String(), meetID, name, *start, start.Add(-5*time.Second))
}

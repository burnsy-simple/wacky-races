package repository

// We initialise our repository source data here.
// In reality this would be coming from e.g. a DB but we hard-code it
// so we can at least demonstrate something
import (
	"math/rand"
	"sort"
	"time"

	"github.com/burnsy/wacky-races/models"
	"gopkg.in/mgo.v2/bson"
)

// We'll use hard-coded data rather than talk to a data source.

var thoroughbredCompetitors []*models.Competitor
var harnessCompetitors []*models.Competitor
var greyhoundCompetitors []*models.Competitor
var thoroughbredRaces models.Races
var greyhoundRaces models.Races
var harnessRaces models.Races
var allRaces models.Races
var racesByID map[string]*models.RaceDetails
var meets []*models.Meet

func init() {
	rand.Seed(42)

	thoroughbredCompetitors = []*models.Competitor{
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
	}

	greyhoundCompetitors = []*models.Competitor{
		newGreyhound("Flying Amy"),
		newGreyhound("Rapid Journey"),
		newGreyhound("Macareena"),
		newGreyhound("Black Top"),
		newGreyhound("Leg Spinner"),
		newGreyhound("Fastly"),
	}

	harnessCompetitors = []*models.Competitor{
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
	meets = []*models.Meet{
		doombenMeet,
		sunshineMeet,
		goldMeet,
	}

	nextThoroughbredStartTime := time.Now().UTC().Add(-2 * time.Minute)
	nextHarnessStartTime := time.Now().UTC().Add(-4 * time.Minute)
	nextGreyhoundStartTime := time.Now().UTC().Add(-4 * time.Minute)
	racesByID = make(map[string]*models.RaceDetails)

	competitorIndex := 0
	thoroughbredRaces = models.Races{
		*newThoroughbredRace(doombenMeet.ID, "Jake Brown Handicap", &nextThoroughbredStartTime, &competitorIndex),
		*newThoroughbredRace(doombenMeet.ID, "Gold Handicap", &nextThoroughbredStartTime, &competitorIndex),
		*newThoroughbredRace(doombenMeet.ID, "Smalls Handicap", &nextThoroughbredStartTime, &competitorIndex),
		*newThoroughbredRace(doombenMeet.ID, "Bloom Handicap", &nextThoroughbredStartTime, &competitorIndex),
		*newThoroughbredRace(doombenMeet.ID, "Ulysees Handicap", &nextThoroughbredStartTime, &competitorIndex),
	}

	competitorIndex = 0
	harnessRaces = models.Races{
		*newHarnessRace(goldMeet.ID, "Bladerunner Memorial", &nextHarnessStartTime, &competitorIndex),
		*newHarnessRace(goldMeet.ID, "Doolittle Do", &nextHarnessStartTime, &competitorIndex),
		*newHarnessRace(goldMeet.ID, "Race 3", &nextHarnessStartTime, &competitorIndex),
		*newHarnessRace(goldMeet.ID, "Race 4", &nextHarnessStartTime, &competitorIndex),
	}

	competitorIndex = 0
	greyhoundRaces = models.Races{
		*newGreyhoundRace(sunshineMeet.ID, "Race 1", &nextGreyhoundStartTime, &competitorIndex),
		*newGreyhoundRace(sunshineMeet.ID, "Race 2", &nextGreyhoundStartTime, &competitorIndex),
		*newGreyhoundRace(sunshineMeet.ID, "Race 3", &nextGreyhoundStartTime, &competitorIndex),
		*newGreyhoundRace(sunshineMeet.ID, "Race 4", &nextGreyhoundStartTime, &competitorIndex),
	}

	allRaces = make(models.Races, 0, len(thoroughbredRaces)+len(harnessRaces)+len(greyhoundRaces))
	allRaces = append(allRaces, thoroughbredRaces...)
	allRaces = append(allRaces, harnessRaces...)
	allRaces = append(allRaces, greyhoundRaces...)

	sort.Sort(allRaces)
}

func newMeet(location string, category models.RaceCategory) *models.Meet {
	return models.NewMeet(bson.NewObjectId().Hex(), location, category)
}

func newThoroughbred(name string) *models.Competitor {
	return models.NewThoroughbred(bson.NewObjectId().Hex(), name)
}
func newGreyhound(name string) *models.Competitor {
	return models.NewGreyhound(bson.NewObjectId().Hex(), name)
}
func newHarness(name string) *models.Competitor {
	return models.NewHarness(bson.NewObjectId().Hex(), name)
}

func newThoroughbredRace(meetID string, name string, start *time.Time, index *int) *models.Race {
	race := models.NewThoroughbredRace(bson.NewObjectId().Hex(), meetID, name, *start, start.Add(-5*time.Second))
	*start = start.Add(2 * time.Minute)
	raceDetails := models.RaceDetails{
		Race: race,
	}
	addCompetitors(&raceDetails, thoroughbredCompetitors, index)
	return race
}

func newGreyhoundRace(meetID string, name string, start *time.Time, index *int) *models.Race {
	race := models.NewGreyhoundRace(bson.NewObjectId().Hex(), meetID, name, *start, start.Add(-5*time.Second))
	*start = start.Add(3 * time.Minute)
	raceDetails := models.RaceDetails{
		Race: race,
	}
	addCompetitors(&raceDetails, greyhoundCompetitors, index)
	return race
}

func newHarnessRace(meetID string, name string, start *time.Time, index *int) *models.Race {
	race := models.NewHarnessRace(bson.NewObjectId().Hex(), meetID, name, *start, start.Add(-5*time.Second))
	*start = start.Add(5 * time.Minute)
	raceDetails := models.RaceDetails{
		Race: race,
	}
	addCompetitors(&raceDetails, harnessCompetitors, index)
	return race
}

func addCompetitors(raceDetails *models.RaceDetails, competitors []*models.Competitor, index *int) {
	numCompetitors := 4 + rand.Intn(3)
	raceDetails.Competitors = make([]models.Competitor, 0, numCompetitors)
	for i := 0; i < numCompetitors; i++ {
		competitor := *competitors[*index]
		competitor.Position = int8(i + 1)
		raceDetails.Competitors = append(raceDetails.Competitors, competitor)
		*index = (*index + 1) % len(competitors)
	}
	racesByID[raceDetails.Race.ID] = raceDetails
}

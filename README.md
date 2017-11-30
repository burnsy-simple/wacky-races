# raceService

By default this service will retrieve the next 5 races and provide details for each of the
races on demand.

#Assumptions:
A competitor's position number is a number between 1 and N where N < 128
No sign-in/credential checking is required
A competitor competes in only one type/category of race
i.e. Thoroughbreds do not compete in Harness races

Possible improvements:
 - unit tests
 - use web sockets to push race updates to the index page. Too late now..
 
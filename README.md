# Rocket-Elevators-Golang-Controller
## Description

This controller's whole purpose is to handle a personalized amount of elevators with a personalized amount of floors in a personalized amount of columns in a battery. 

It can be controller from any floor from the outside of the elevators, but mainly in the lobby. 

When used from the lobby, the battery choose the column that serves the floor selected by the user and then sends the best elevator possible for that spec floor and direction. Then, when used from the inside of the elevator that was selected by the column, the elevator is moved to the to the user's destination.

Otherwise, when called from another floor than the lobby, the column selects again the best elevator possible 

Elevator selection is based on the elevator's status, current floor, direction and floor request list and on the user's floor and direction.

## Dependencies

First you need to install Golang on your computer.

With golang installed on your computer, all you need to do is initialize the module:

`go mod init Rocket-Elevators-Commercial-Controller`

The code to run the scenarios is included, and can be executed with:

`go run . <SCENARIO-NUMBER>`

### Scenario 1

For the column B(2)
- Elevator A is on the 20th floor going down to the 5th floor.
- Elevator B is on the 3rd floor going up to the 15th floor.
- Elevator C is on the 13th floor going down to the 1st floor.
- Elevator D is on the 15th floor going down to the 2nd floor.
- Elevator E is on the 6th floor going down to the 2nd floor.

- User is on the 1st floor and wants to go to the 20th floor. Elevator E should be sent.

### Scenario 2

For the column C(3)
- Elevator A is stopped on the 1st floor but going up to the 21st floor.
- Elevator B is on the 23rd floor going up to the 28th floor.
- Elevator C is on the 33rd floor going down to the 1st floor.
- Elevator D is on the 40th floor going down to the 24th floor.
- Elevator E is on the 39th floor going down to the 1st floor.

- User is on the 1st floor and wants to go to the 36th floor. Elevator A should be sent.

### Scenario 3

For the column D(4)
- Elevator A is on the 58th floor going down to the 1st floor.
- Elevator B is on the 50th floor going up to the 60th floor.
- Elevator C is on the 46th floor going up to the 58th floor.
- Elevator D is on the 1st floor going up to the 54th floor.
- Elevator E is on the 60th floor going down to the 1st floor.

- User is on the 54th floor and wants to go back to the 1st floor. Elevator A should be sent.

### Scenario 4

For the column A(1)
- Elevator A is idle on the B4th floor.
- Elevator B is idle on the 1st floor
- Elevator C is on the B3rd floor going down to the B5th floor.
- Elevator D is on the B6st floor going up to the 1st floor.
- Elevator E is on the B1st floor going down to the B6th floor.

- User is on the 3rd floor and wants to go to the 1st floor. Elevator D should be sent.

### Running the tests

To launch the tests:

`go test`

With a fully completed project, you should get an output like:

![Screenshot from 2021-06-15 15-25-10](https://user-images.githubusercontent.com/28630658/122111573-e6ea7380-cded-11eb-95e3-95e0096a1b3a.png)

You can also get more details about each test by adding the `-v` flag: 

`go test -v` 

which should give something like: 

![Screenshot from 2021-06-15 15-25-51](https://user-images.githubusercontent.com/28630658/122111659-fd90ca80-cded-11eb-991b-d9f6fe1d317b.png)

The test and scenarios files can be left in your final project. The grader will run tests similar to the ones provided.

Of course, make sure to edit this Readme file to describe your own project!

# Optiopay Challenge - Find closest common Manager

To build run the following from the parent most directory of the project

    go get github.com/manifoldco/promptui
    go build optiopay

To run use the following command:

    ./optiopay -cmd

## Assumptions and limitations

* All the employee names are unique
* I am writing Go code for the first time!
* A manager can have n number of subordinates
* All the negative cases have been omitted - like during search it is assumed that you enter a name that you have already fed as an employee
* Test cases are there, but limitted

## Note
* Design has been given more importance
* It has been coded to make it adaptable - you can have different storage types later, different APIs later, that the employee model could have more functionalities etc
* There are files lying around without any code. They are for future implementation/improvement
* I have added comments in few places where I felt I could improve or implement in different way
* I started off coding in Python to get a feel of the problem statement and it has been put in `reference` directory. It is not part of solution.
* LCA has been implemented in the easiest possible way. It may not be the perfect one.

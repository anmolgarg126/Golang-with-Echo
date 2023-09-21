package service

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"echo_framework/database"
	"echo_framework/model"
	"echo_framework/utility"
	// "github.com/pkg/errors"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

var userMapping = make(map[string]model.User)

func SaveUser(c echo.Context) error {
	// User ID from path `users/:id`

	fmt.Println("Inside SaveUser")

	u := new(model.User)

	// Echo.Validator = &model.UserValidator{validator: Validator}

	if err := c.Bind(u); err != nil {
		return err
	}

	// if err := c.Validate(u); err != nil {
	// 	return err
	// }

	if err := utility.ValidateStruct(u); err != nil {
		return err
	}

	// if _, ok := userMapping[u.ID]; ok {
	// 	return c.String(http.StatusConflict, "User already exists")
	// }
	// userMapping[u.ID] = *u

	if err := database.InsertOneUsingStruct(*u); err != nil {
		// fmt.Println(errors.Wrap(err, "User Creation Failed"))
		// fmt.Println(errors.Cause(err))
		return c.JSON(http.StatusConflict, "User already exists")
	}
	return c.JSON(http.StatusOK, "user created")
}

// func saveUser(c echo.Context) error {
// 	// User ID from path `users/:id`

// 	u := new(User)
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}

// 	if err := v.Struct(u); err != nil {
// 		return err
// 	}

// 	if _, ok := userMapping[u.Id]; ok {
// 		return c.String(http.StatusConflict, "User already exists")
// 	}
// 	userMapping[u.Id] = *u
// 	return c.JSON(http.StatusOK, "user created")
// }

func GetUserById(c echo.Context) error {
	// User ID from path `users/:id`

	fmt.Println("Inside GetUserById")

	id := c.QueryParam("id")
	userById, err := database.FindOneByID(id)

	// time.Sleep(100 * time.Second)
	// if user, ok := userMapping[id]; ok {
	// 	return c.JSON(http.StatusOK, user)
	// }
	if err != nil {
		return c.String(http.StatusBadRequest, "User not found for id :: "+id)
	}
	return c.JSON(http.StatusOK, userById)
}

func GetUserByName(c echo.Context) error {
	// User ID from path `users/:id`

	fmt.Println("Inside GetUserByName")

	name := c.QueryParam("name")
	userByName, err := database.FindOneByName(name)

	if err != nil {
		return c.String(http.StatusBadRequest, "User not found for name :: "+name)
	}
	return c.JSON(http.StatusOK, userByName)
}

func FindByNumberOfJobsAUserHad(c echo.Context) error {
	fmt.Println("Inside FindByNumberOfJobsAUserHad")

	number, err := strconv.Atoi(c.QueryParam("jobsNumber"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Provide proper integer value")
	}

	userByName, err := database.FindByNumberOfJobs(number)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "No User found having min number of jobs :: "+strconv.Itoa(number))
	}
	return c.JSON(http.StatusOK, userByName)
}

func FindUserByExperienceAndSalary(c echo.Context) error {
	fmt.Println("Inside FindUserByExperienceAndSalary")

	experience, err := strconv.Atoi(c.Param("experience"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Provide proper experience value")
	}

	salary, err := strconv.ParseFloat(c.Param("salary"), 32)
	if err != nil {
		return c.String(http.StatusBadRequest, "Provide proper salary value")
	}

	userByName, err := database.FindByExperienceAndSalary(salary, experience)

	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("No User found for given min salary :: %f and experience :: %d", salary, experience))
	}
	return c.JSON(http.StatusOK, userByName)
}

func FindUserWorkedInGivenOrganisation(c echo.Context) error {
	fmt.Println("Inside FindUserWorkedInGivenOrganisation")

	organisation := c.QueryParam("organisation")

	userByName, err := database.FindUsersWorkedInGivenOrganisation(organisation)

	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("No User found for given organisation :: %s", organisation))
	}
	return c.JSON(http.StatusOK, userByName)
}

func FindUserHavingNameField(c echo.Context) error {
	fmt.Println("Inside FindUserHavingNameField")

	userByName, err := database.FindUsersWhoseNameExists()

	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("No User whose name does not exists"))
	}
	return c.JSON(http.StatusOK, userByName)
}

func DeleteUser(c echo.Context) error {
	// User ID from path `users/:id`

	fmt.Println("Inside DeleteUser")
	id := c.QueryParam("id")

	// if user, ok := userMapping[id]; ok {
	// 	delete(userMapping, id)
	// 	return c.JSON(http.StatusOK, user)
	// }
	if err := database.DeleteById(id); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, id)
}

func GetAllUsers(c echo.Context) error {
	// User ID from path `users/:id`
	time.Sleep(2 * time.Second)
	fmt.Println("inside GetAllUsers")
	result := []model.User{}

	for _, val := range userMapping {
		result = append(result, val)
	}

	fmt.Println("Result ", result)
	return c.JSON(http.StatusOK, result)
}

package main

import (
	"database/sql"
	"fmt" // package that contain fromattings/strings/inputs/outputs
	"log"
	"os" // package that contain functions to intract with os

	_ "github.com/lib/pq"
)

type Task struct { // declare struct  colletions tor store task data
    TaskId int
	TaskName string 
    TaskTime  string 
    TaskDate  string 
}

var tasks  []Task //declare variable name tasks that slice of Task object


const (
    host     = "localhost"
    port     = 5432
    user     = "pasan"
    password = "12345"
    dbname   = "my-crud"
)


func initializeDB() (*sql.DB, error) {
    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil,err
    }
    if err := db.Ping(); err != nil {
        db.Close()
        return nil, err
    }

    fmt.Println("Connected to the database successfully!")
    return db, nil
}

func saveTaskToDB(db *sql.DB, task Task) {
    _, err := db.Exec("INSERT INTO tasks (taskid, taskname, tasktime, taskdate) VALUES ($1, $2, $3, $4)",task.TaskId, task.TaskName, task.TaskTime, task.TaskDate)
    if err != nil {
        log.Fatal(err)
    }
}


func loadTasksFromDB(db *sql.DB) []Task {
    rows, err := db.Query("SELECT * FROM tasks ORDER BY taskid LIMIT 1")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var loadedTasks []Task
    for rows.Next() {
        var task Task
        if err := rows.Scan(&task.TaskId,&task.TaskName, &task.TaskTime, &task.TaskDate); err != nil {
            log.Fatal(err)
        }
        loadedTasks = append(loadedTasks, task)
    }

    return loadedTasks
}


func readTask(tasks []Task){
 if tasks == nil {
	fmt.Printf("Tasks are empty...!")
 }else{
	for _, tasks := range tasks { 
		fmt.Printf("Task ID:%d | Task Name:%s | Task Time:%s | Task Date:%s\n", tasks.TaskId, tasks.TaskName, tasks.TaskTime,tasks.TaskDate)	
	}	//%s this is an string fromat specifer that says where shoud include values
 }
}

func showTaskById(id int, tasks []Task){
	i := id-1
	if tasks != nil {
		for _, tasklist := range tasks { 
			//go through the task collection ignoreing indexes 
			//by getting values to the tasklist without specifing the type of the variable

			if tasks[i] == tasklist {
			 fmt.Printf("Task ID:%d | Task Name:%s | Task Time:%s| Task Date:%s\n\n", tasklist.TaskId, tasklist.TaskName,tasklist.TaskTime,tasklist.TaskDate)
			}
		 }
	}else{
		fmt.Printf("Tasks are empty...!\n")
	 }
   }

func createTask(tasks *[]Task, db *sql.DB){
	var taski int
	var taskN string
	var taskD string
	var taskT string
	fmt.Println()


	fmt.Print("Enter Task ID : ")
	fmt.Scan(&taski)

	fmt.Print("Enter Task Name : ")
	fmt.Scan(&taskN)
	//& used to pass the memory address of taskN to the fmt using scan.
	// Calling Scan() function for
    // scanning and reading the input
    // texts given in standard input

	fmt.Print("Enter Task Time : ")
	fmt.Scan(&taskT)

	fmt.Print("Enter Task Date : ")
	fmt.Scan(&taskD)

	Task_ := Task{TaskId:taski, TaskName:taskN, TaskDate:taskD, TaskTime:taskT}
	//declare new variable and save Task struct to data by assigning new values

	*tasks = append(*tasks, Task_) 
	//says uses to append new slices to *tasks
	//*uses to access the tasks values 

	saveTaskToDB(db, Task_)

	fmt.Println("Task successfully Created..!")
	fmt.Printf("\n")

}

func deleteTaskById(id int, tasks *[]Task){
	//id is index to be deleted
	i := id -1
   fmt.Println("Task number that need to be delted is:", id)
   if i >=0 && i< len(*tasks) {
	//remove an element from a slice at index i
	*tasks = append((*tasks)[:i],(*tasks)[i+1:]...)
 
	//get portion of the struct till i: and i+1: 
	//then concat both parts of the struct and equal it to the *tasks
	for _, tasks := range *tasks {
		fmt.Printf("Task ID:%d | Task Name:%s | Task Time:%s | Task Date:%s\n", tasks.TaskId,tasks.TaskName,tasks.TaskTime,tasks.TaskDate)
		fmt.Println("Task successfully Deleted..!")	
		}
   } else {
	  fmt.Println("The given index is not in the list.")
   }

}

func updateTask(id int, tasks []Task){
		//id is index to be update
	i := id -1
	var taski int
	var taskN string
	var taskD string
	var taskT string
	if tasks != nil {
		fmt.Println()

		fmt.Print("Enter Task Name : ")
		fmt.Scan(&taski)

		fmt.Print("Enter Task Name : ")
		fmt.Scan(&taskN)
		//& used to pass the memory address of taskN to the fmt using scan.

		fmt.Print("Enter Task Time : ")
		fmt.Scan(&taskT)

		fmt.Print("Enter Task Date : ")
		fmt.Scan(&taskD)
		
		//go through the task collection ignoreing indexes 
		//by getting values to the tasklist without specifing the type of the variable
		//check wheather user's desire thats matches with the tasks indexes
		for _, tasklist := range tasks {
			if tasks[i] == tasklist {
		//when goes throug the each index of task struct check if user index matches the tasks indexes
		//if it does update the values of that belongs to that exact index
				tasks[i] = Task{TaskId:taski, TaskName:taskN, TaskTime:taskT, TaskDate:taskD}
				fmt.Println("Task successfully Updated..!")
			}
		}
	}else{
		fmt.Printf("Tasks are empty...!\n")
	 }
}

func listOfOptions() {
	fmt.Println("HI This is an TODO APP")
	fmt.Println("1. Show List of Tasks")
	fmt.Println("2. Show Task by Task Number")
	fmt.Println("3. Add a New Task")
	fmt.Println("4. Update a Task ")
	fmt.Println("5. Delete a Task")
	fmt.Println("0. Exit")
	fmt.Print("Enter your option number: ")
}


func reload(db *sql.DB){
	var input string
	fmt.Print("Type 'Enter' go back to menu: ")
	fmt.Scan(&input)
	if input == "Enter"{
		menu(db)
	}else{
		fmt.Printf("Invalid Input...!\n")
		reload(db)
	}
	//reload the menu function to be shown
}

func menu(db *sql.DB) {
	var option int
	listOfOptions()
	for {
		fmt.Scan(&option)
	if option == 0 {
		fmt.Println("BYE BYE.!")
		os.Exit(0)
	}else{
		switch option{
		case 1:
			readTask(tasks)
			fmt.Printf("\n")
			listOfOptions()
		case 2:
			var id int
			fmt.Println()
			fmt.Print("Enter Task ID that you want to get :")
			fmt.Scan(&id)
			showTaskById(id, tasks)
			reload(db)
		case 3:
			createTask(&tasks , db)
			reload(db)
		case 4:
			var id int
			fmt.Println()
			fmt.Print("Enter Task ID that you want to Update: ")
			fmt.Scan(&id)
			updateTask(id, tasks)
			fmt.Printf("\n")
			reload(db)
		case 5:
			var id int
			fmt.Println()
			fmt.Print("Enter Task ID that you want to Delete: ")
			fmt.Scan(&id)
			deleteTaskById(id, &tasks)
			fmt.Printf("\n")
			reload(db)
		default:
			fmt.Println("Re-enter your choice!")
			menu(db)
		}
	}

	}
}

func main() {
	db, err := initializeDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a 'tasks' table if it doesn't exist
	_, createTableErr := db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		taskid int PRIMARY KEY,
		taksname varchar,
		tasktime varchar,
		taskdate varchar
	)`)
	if createTableErr != nil {
		log.Fatal(createTableErr)
	}

	tasks = loadTasksFromDB(db)
	menu(db)
}

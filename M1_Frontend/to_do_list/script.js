var taskCollection = []; 
var taskInput = document.getElementById("taskInput");
var taskContainer = document.getElementById("taskContainer");
var recycledKeys = [];
var activeFilter = "allTasks";

function createTask() {
    console.log("Task creation initiated");
    var taskDescription = taskInput.value;
    if (taskDescription.trim() !== "") {
        var taskElement = document.createElement("li");
        taskElement.classList.add("list-group-item", "d-flex", "justify-content-between", "align-items-center", "bg-primary-subtle", "rounded", "mb-1");
        taskElement.innerHTML = `<div><input type="checkbox" onclick="toggleTaskStatus(this)" class="me-1">` +
                                `<span class="task-text">${taskDescription}</span></div>` +
                                `<button class="badge text-bg-danger rounded-pill border-0" style="width:1.5rem;height:1.5rem;" ` +
                                `onclick="removeTask(this)">X</button>`;
        taskContainer.appendChild(taskElement);
        taskInput.value = "";

        if (activeFilter === "completedTasks") {
            taskElement.classList.add("d-none");
        }
    }
}

function toggleTaskStatus(checkbox) {
    var taskText = checkbox.parentElement.querySelector(".task-text");
    taskText.style.textDecoration = checkbox.checked ? "line-through" : "none";

    if (activeFilter !== "allTasks") {
        var taskItem = checkbox.closest("li");
        if ((activeFilter === "currentTasks" && checkbox.checked) || 
            (activeFilter === "completedTasks" && !checkbox.checked)) {
            taskItem.classList.add("d-none");
        } else {
            taskItem.classList.remove("d-none");
        }
    }
}

function removeTask(deleteButton) {
    var taskToRemove = deleteButton.closest("li");
    taskToRemove.remove();
}

function filterTasks(filterOption) {
    console.log("Switching filter from", activeFilter, "to", filterOption);

    if (activeFilter !== filterOption) {
        document.getElementById(activeFilter).classList.remove("bg-primary-subtle", "border-bottom", "border-primary");
        document.getElementById(filterOption).classList.add("bg-primary-subtle", "border-bottom", "border-primary");

        var tasks = taskContainer.children;
        if (filterOption === "allTasks") {
            for (let task of tasks) {
                task.classList.remove("d-none");
            }
        } else if (filterOption === "currentTasks") {
            for (let task of tasks) {
                var isChecked = task.querySelector("input[type='checkbox']").checked;
                task.classList.toggle("d-none", isChecked);
            }
        } else if (filterOption === "completedTasks") {
            for (let task of tasks) {
                var isChecked = task.querySelector("input[type='checkbox']").checked;
                task.classList.toggle("d-none", !isChecked);
            }
        }

        activeFilter = filterOption;
    }
}

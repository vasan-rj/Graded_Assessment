// Quiz Questions Array
const questions = [
    { 
        question: "Who was the first Prime Minister of India?", 
        options: ["Mahatma Gandhi", "Jawaharlal Nehru", "Subhas Chandra Bose", "Sardar Vallabhbhai Patel"], 
        correct: "Jawaharlal Nehru", 
        selected: null 
    },
    { 
        question: "What is the national flower of India?", 
        options: ["Rose", "Lotus", "Sunflower", "Marigold"], 
        correct: "Lotus", 
        selected: null 
    },
    { 
        question: "Which Indian state is known as the 'Land of Five Rivers'?", 
        options: ["Punjab", "Haryana", "Uttar Pradesh", "Rajasthan"], 
        correct: "Punjab", 
        selected: null 
    },
    { 
        question: "Who is known as the 'Missile Man of India'?", 
        options: ["APJ Abdul Kalam", "Homi Bhabha", "Vikram Sarabhai", "Rakesh Sharma"], 
        correct: "APJ Abdul Kalam", 
        selected: null 
    },
    { 
        question: "Which Indian city is famous for its Charminar?", 
        options: ["Mumbai", "Delhi", "Hyderabad", "Jaipur"], 
        correct: "Hyderabad", 
        selected: null 
    }
];

// DOM Elements
const currentQuestionElement = document.getElementById("currentQuestion");
const totalQuestionsElement = document.getElementById("totalQuestions");
const quizQuestionElement = document.getElementById("quizQuestion");
const optionGroupElement = document.getElementById("optionGroup");

let currentIndex = 0; // Current question index
totalQuestionsElement.textContent = questions.length;

// Load Question
function loadQuestion(index) {
    const questionData = questions[index];
    quizQuestionElement.textContent = questionData.question;
    currentQuestionElement.textContent = index + 1;

    // Clear and populate options
    optionGroupElement.innerHTML = "";
    questionData.options.forEach((option, i) => {
        const optionDiv = document.createElement("div");
        optionDiv.className = "form-check";
        optionDiv.innerHTML = `
            <input class="form-check-input" type="radio" name="options" id="option${i}" value="${option}" ${option === questionData.selected ? "checked" : ""}>
            <label class="form-check-label" for="option${i}">${option}</label>
        `;
        optionGroupElement.appendChild(optionDiv);
    });
}

// Store Answer and Proceed to Next Question
function storeAndProceed() {
    const selectedOption = document.querySelector('input[name="options"]:checked');
    if (selectedOption) {
        questions[currentIndex].selected = selectedOption.value;
    }

    if (currentIndex === questions.length - 1) {
        const score = calculateScore();
        alert(`Quiz Completed! Your score: ${score}/${questions.length}`);
    } else {
        currentIndex++;
        loadQuestion(currentIndex);
    }
}

// Navigate Back to Previous Question
function navigateBack() {
    if (currentIndex > 0) {
        currentIndex--;
        loadQuestion(currentIndex);
    }
}

// Navigate Forward to Next Question
function navigateForward() {
    if (currentIndex < questions.length - 1) {
        currentIndex++;
        loadQuestion(currentIndex);
    }
}

// Reset Current Question's Selection
function resetSelection() {
    questions[currentIndex].selected = null;
    loadQuestion(currentIndex);
}

// Calculate Score
function calculateScore() {
    return questions.reduce((score, question) => score + (question.selected === question.correct ? 1 : 0), 0);
}

// Initial Load
loadQuestion(currentIndex);

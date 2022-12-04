import { getInputLines } from "../../input.js";

const lines = await getInputLines();
const games = lines.map((line) => line.split(' '));

// Rock  Paper  Scissors
// A     B      C
// X     Y      Z

const selectionScoreMap = {
    'X': 1,
    'Y': 2,
    'Z': 3,
};

const lose = 0;
const draw = 3;
const win = 6;

const rules = {
    'A': {
        'X': draw,
        'Y': win,
        'Z': lose,
    },
    'B': {
        'X': lose,
        'Y': draw,
        'Z': win,
    },
    'C': {
        'X': win,
        'Y': lose,
        'Z': draw,
    }
};

const scores = games.map((game) => selectionScoreMap[game[1]] + rules[game[0]][game[1]]);
const total = scores.reduce((accumulator, value) => accumulator += value, 0);

console.log('answer1', total); //12586

// Lose Draw Win
// X    Y    Z

const outcomeScoreMap = {
    'X': lose,
    'Y': draw,
    'Z': win,
};

function findSelection(opponent, expectation) {
    for (const [selection, outcome] of Object.entries(rules[opponent])) {
        if (outcome === outcomeScoreMap[expectation]) return selection;
    }
}

const updatedScores = games.map((game) => outcomeScoreMap[game[1]] + selectionScoreMap[findSelection(game[0], game[1])]);
const updatedTotal = updatedScores.reduce((accumulator, value) => accumulator += value, 0);

console.log('answer2', updatedTotal); //13193
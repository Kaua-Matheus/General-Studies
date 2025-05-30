"use strict";
// Declare algumas variáveis com uma série de tipos.
// String
let var01 = "nome";
// Number
let var02 = 19;
// Any
let var03 = "arroz";
var03 = true;
// Boolean
let var04 = false;
// array Number
let var05 = [1, 2, 5];
console.log(var05[2]);
// Tuple
let var06 = [
    ["arroz", "macarrão"],
    ["batata", "cebola"],
];
console.log(var06[0][1]);
// enum
var Direction;
(function (Direction) {
    Direction[Direction["Up"] = 1] = "Up";
    Direction["Down"] = "Baixo";
    Direction[Direction["Left"] = 3] = "Left";
    Direction["Right"] = "Direita";
})(Direction || (Direction = {}));
console.log(Direction.Down, Direction.Right);
function function1() {
    return 9 + 11;
}
function function2() {
    let value = 10;
    value = value + 2;
    value += 2;
    return value;
}
var executioner;
(function (executioner) {
    executioner[executioner["A"] = function1()] = "A";
    executioner[executioner["B"] = function2()] = "B";
})(executioner || (executioner = {}));
console.log(executioner.A, executioner.B);

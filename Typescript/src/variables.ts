// Declare algumas variáveis com uma série de tipos.

// String
let var01: string = "nome";

// Number
let var02: number = 19;

// Any
let var03: any = "arroz";
var03 = true;

// Boolean
let var04: boolean = false;

// array Number
let var05: number[] = [1, 2, 5];
console.log(var05[2]);

// Tuple
let var06: string[][] = [
  ["arroz", "macarrão"],
  ["batata", "cebola"],
];
console.log(var06[0][1]);

// enum
enum Direction {
    Up = 1,
    Down = 'Baixo',
    Left = 3,
    Right = 'Direita',
}
console.log(Direction.Down, Direction.Right)

function function1(): number{
    return 9 + 11;

}

function function2(): number{
    let value: number = 10;
    value = value + 2;
    value += 2;
    return value;
}


enum executioner {
    A = function1(),
    B = function2(),
}
console.log(executioner.A, executioner.B)
/**
 * @author Ethan White
 * @version 1.0.0
 * @date 25-11-19
 * @fileoverview This program will take four lines and output them is reverse order
 */

//making variables
let line1: string;
let line2: string;
let line3: string;
let line4: string;

//inputs
line1 = prompt("Type in a line") || ("You didnt type anything")
line2 = prompt("Type in a line") || ("You didnt type anything")
line3 = prompt("Type in a line") || ("You didnt type anything")
line4 = prompt("Type in a line") || ("You didnt type anything")

//output

console.log("\n\n" + line4)
console.log(line3)
console.log(line2)
console.log(line1)

console.log("\nDone.")
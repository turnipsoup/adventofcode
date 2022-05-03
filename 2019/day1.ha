use fmt;
use io;
use math;
use strconv;
use strings;
use os;

export fn main() void = {

	// Day 1
	let inputs = strings::split(openFileAsString("input/day1.txt"), "\n");
	let fuelSum = 0;
	let allFuelSum = 0;
##
	for (let i = 0z; i<len(inputs); i += 1) {
		if (len(inputs[i]) > 0) {
			let fuel = getFuelCost(strconv::stoi(inputs[i])!);
			fuelSum += fuel;
			allFuelSum += fuel + calcFuelForFuel(fuel);
		};
	};

	fmt::printfln("Part 1: {}", fuelSum)!;
	fmt::printfln("Part 2: {}", allFuelSum)!;

};

// Calculate the fuel cost given a mass
fn getFuelCost(mass: int) int = {
	let fmass = mass / 3;
	return fmass - 2;
};

// Opens file as a string
fn openFileAsString(filename: str) str = {
	const source = os::open(filename)!;
	const source = io::drain(source)!;
	const source = strings::fromutf8(source);

	return source;
};

// Calculates all of the fuel used by the fuel for the fuel....
fn calcFuelForFuel(fuel: int) int = {
	let newfuel = 0;
	let remainder = getFuelCost(fuel);

	if (remainder >= 0) {
		newfuel += remainder;
	};

	for (true) {
		remainder = getFuelCost(remainder);

		if (remainder >= 0) {
			newfuel += remainder;
		} else {
			break;
		};
	};

	return newfuel;
};




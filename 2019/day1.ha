use fmt;
use io;
use math;
use strconv;
use strings;
use os;

export fn main() void = {

	let inputs = strings::split(openFileAsString("input/day1.txt"), "\n");
	let fuelSum = 0;

	for (let i = 0z; i<len(inputs); i += 1) {
		if (len(inputs[i]) > 0) {
			fuelSum += getFuelCost(inputs[i]);
		};
	};

	fmt::printfln("Day 1: {}", fuelSum)!;
};

fn getFuelCost(mass: str) int = {
	let fmass = strconv::stoi(mass);
	let fmass = fmass / 3;
	return fmass - 2;
};

fn openFileAsString(filename: str) str = {
	const source = os::open(filename)!;
	const source = io::drain(source)!;
	const source = strings::fromutf8(source);

	return source;
};

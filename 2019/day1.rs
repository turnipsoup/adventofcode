use std::fs;

fn main() {
    let finput = fs::read_to_string("./input/day1.txt").expect("Something went wrong!");
    let filist = finput.split_terminator("\n");

    let mut fuel_sum = 0;
    let mut all_fuel_sum = 0;

    for i in filist {
        if i.len() > 0 {
            let fuel = fuel_cost(i.parse().unwrap());
            fuel_sum = fuel_sum + fuel;
            all_fuel_sum = all_fuel_sum + fuel + calc_fuel_for_fuel(fuel);
        };
    };
    
    println!("Day 1: {}", fuel_sum);
    print!("Day 2: {}", all_fuel_sum);
}

fn fuel_cost(mass: i32) -> i32 {
    let fmass = mass / 3;
    fmass - 2
}

fn calc_fuel_for_fuel(fuel: i32) -> i32 {
    let mut newfuel = 0;
    let mut remainder = fuel_cost(fuel);

    if remainder >= 0 {
        newfuel = newfuel + remainder;
    };


    loop {
        remainder = fuel_cost(remainder);

        if remainder >= 0 {
            newfuel = newfuel + remainder;
        } else {
            break;
        };
    };

    newfuel
}